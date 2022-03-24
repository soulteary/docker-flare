package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/anaskhan96/soup"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

//go:embed VERSION
var version string

var (
	app      = kingpin.New("importer", "A command line tool to convert Chrome bookmarks to flare bookmarks.")
	filename = kingpin.Flag("filename", "filename of Chrome's bookmarks").Short('f').Required().String()
	output   = kingpin.Flag("output", "output filename for flare bookmarks, print to stdout if not set").Short('o').String()
	smooth   = kingpin.Flag("smooth", "smooth the output").Default("true").Bool()
)

// Record 是单条书签
type Record struct {
	Title string
	Url   string
}

// Bookmarks 是一个嵌套的书签列表
type Bookmarks struct {
	Name           string       `json:"name"`
	Records        []*Record    `json:"records"`
	EmbedBookmarks []*Bookmarks `json:"embedBookmarks"`
}

func main() {
	app.HelpFlag.Short('h')
	kingpin.Version(version)
	kingpin.Parse()

	root := Bookmarks{
		Name: "root",
	}
	// 解析书签
	ParseChromeBookmarks(*filename, &root)

	// 将书签转换为flare要的格式
	context, err := ConvertToFlareBookmarks(&root, *smooth)
	if err != nil {
		log.Fatalln("convert to flare bookmarks failed,", err)
	}

	// 将结果写文件
	if *output != "" {
		err = os.WriteFile(*output, context, 0644)
		if err != nil {
			log.Fatalln("write output failed, ", err)
		}
	} else {
		fmt.Println(string(context))
	}
}

func ParseChromeBookmarks(filename string, root *Bookmarks) {
	// 加载chrome书签
	res, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	doc := soup.HTMLParse(string(res))

	// chrome书签是固定格式的，以第一个DL为根节点
	body := doc.Find("body")
	dls := findAllChild(body, "dl")

	rootTag := dls[0]
	var f func(n *soup.Root, b *Bookmarks)
	f = func(n *soup.Root, b *Bookmarks) {
		for _, child := range n.Children() {
			if isDir(&child) {
				newDir := Bookmarks{
					Name: child.Children()[0].Text(),
				}
				dl := findBookmarksEntrypoint(&child)
				if dl == nil {
					fmt.Println("找不到dl")
					continue
				}

				f(dl, &newDir)
				b.EmbedBookmarks = append(b.EmbedBookmarks, &newDir)
			}
			if isLink(&child) {
				record := Record{
					Title: child.Children()[0].Text(),
					Url:   child.Children()[0].Attrs()["href"],
				}
				b.Records = append(b.Records, &record)
			}
		}
	}

	f(&rootTag, root)
}

func findAllChild(root soup.Root, tag string) []soup.Root {
	var all []soup.Root
	for _, child := range root.Children() {
		if child.NodeValue == tag {
			all = append(all, child)
		}
	}
	return all
}

// 以DT+H3，匹配的为目录
func isDir(n *soup.Root) bool {
	return isMatch(n, "dt", "h3")
}

// 以DT+A,匹配的为书签
func isLink(n *soup.Root) bool {
	return isMatch(n, "dt", "a")
}

func isMatch(n *soup.Root, parentTag, firstChildTag string) bool {
	if n.NodeValue != parentTag {
		return false
	}
	childs := n.Children()
	if len(childs) == 0 {
		return false
	}
	return childs[0].NodeValue == firstChildTag
}

func findBookmarksEntrypoint(n *soup.Root) *soup.Root {
	for _, child := range n.Children() {
		if child.NodeValue == "dl" {
			return &child
		}
	}
	return nil
}

type Categories struct {
	Id    int    `yaml:"id"` // TODO：flare支持id为string, 未来这里将中文的书签目录转换为拼音？？
	Title string `yaml:"title"`
}
type Links struct {
	Name     string `yaml:"name"`
	Link     string `yaml:"link"`
	Icon     string `yaml:"icon"`
	Category int    `yaml:"category"`
}

// flare bookmarks config
type FlareBookmarks struct {
	Categories []*Categories `yaml:"categories"`
	Links      []*Links      `yaml:"links"`
}

// ConvertToFlareBookmarks 将chrome书签转换为flare书签
// 格式见这里： https://github.com/soulteary/docker-flare/blob/main/app/bookmarks.yml
func ConvertToFlareBookmarks(root *Bookmarks, smooth bool) ([]byte, error) {
	// TODO 如果不smooth那有些嵌套文件夹怎么处理, 这里默认就平铺了
	bms := &FlatBookmarks{}
	flatBookmarks(root, bms)

	flareBms := &FlareBookmarks{}
	bms.Range(func(i int, b *Bookmarks) {
		// 就以索引下标+1作为categories的id
		flareBms.Categories = append(flareBms.Categories, &Categories{
			Id:    i + 1,
			Title: b.Name,
		})

		// 将书签转换为links
		for _, record := range b.Records {
			flareBms.Links = append(flareBms.Links, &Links{
				Name:     record.Title,
				Link:     record.Url,
				Icon:     "",
				Category: i + 1,
			})
		}
	})

	return yaml.Marshal(flareBms)
}

// flatBookmarks 将Bookmarks中的内容平铺返回，因为flare只支持一级目录
type FlatBookmarks []*Bookmarks

func (fb *FlatBookmarks) Append(b *Bookmarks) {
	*fb = append(*fb, b)
}
func (fb *FlatBookmarks) Range(f func(i int, b *Bookmarks)) {
	for i, b := range *fb {
		f(i, b)
	}
}

func flatBookmarks(root *Bookmarks, bms *FlatBookmarks) {
	embeds := root.EmbedBookmarks
	if len(root.Records) > 0 {
		// 有记录，把父目录保存起来
		root.EmbedBookmarks = nil
		bms.Append(root)
	}
	for _, b := range embeds {
		flatBookmarks(b, bms)
	}
}
