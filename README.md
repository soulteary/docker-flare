# Flare ✨

如果你觉得这个项目有帮到你，欢迎点赞✨（star）给予鼓励；如果你希望收到这个项目的更新推送，可以点击关注 👀（watch）并选择适合自己的关注模式（推荐 release）。

---

轻量、快速、美观的个人导航页面，适用于 HomeLab 或其他注重私密的场景。

无任何数据库依赖，应用数据完全开放透明，100% 属于用户自己。

支持 x86 以及常见的 ARM 设备，应用资源消耗非常低：

- CPU: < 1%
- MEM: < 30M
- Docker Image: < 10M

![Flare Docker Image Size](./screenshots/docker-image-size.png)

## 快速上手

快速上手 Flare，需要两步：**下载**包含示例的代码、**启动**程序访问浏览器。

### 下载包含示例的代码

你可以使用 `git clone` 或者选择使用 “Download ZIP” 的方式，下载包含了基础的配置示例（书签和应用）的代码：

```bash
git clone https://github.com/soulteary/docker-flare.git
cd docker-flare
```

`app/*yml` 目录中包含了你的书签和应用数据，你可以根据你的需求对其进行调整。如果目录中没有配置文件，应用将在首次运行的时候，进行自动创建。

### 启动程序访问浏览器

启动应用非常简单，如果你习惯使用 Docker，可以执行：

```bash
# 可以使用最新镜像
docker run --rm -it -p 5005:5005 -v `pwd`/app:/app soulteary/flare
# 也可以追求明确，使用固定版本
docker run --rm -it -p 5005:5005 -v `pwd`/app:/app soulteary/flare:0.2.9
```

如果你习惯使用 docker-compose，只需要执行：

```bash
docker-compose up -d
```

如果你是 Traefik 用户，可以参考 `docker-compose.traefik.yml` 配置文件来使用。

不论是哪一种方式，在命令执行完毕之后，默认情况下，我们访问浏览器的 `5005` 端口，就能看到下面的界面啦：

![Flare Web UI](./screenshots/ui.png)

### 程序使用向导

为了方便你的使用，我制作了一个简单的向导程序，在 flare 启动之后，你可以随时访问 `/guide`，获取 flare 基础界面功能的介绍。

![Flare Guide](./screenshots/flare-guide.png)


## 程序在线编辑页面

为了满足随时随地编辑的需求，程序新增了“在线编辑”的页面。

![Flare Editor](./screenshots/editor-beta.png)

工具页面地址：`/editor`

## 程序帮助页面

为了减少不必要的地址记忆负担，程序新增了一个“帮助页面”，默认展示所有的程序内的工具页面。

![Flare Help](./screenshots/flare-help.png)

工具页面地址：`/help`

## 程序性能

“快”作为 Flare 对主要优势而言，自然是需要“满分”来加持。

![Flare Lighthouse Scores](./screenshots/lighthouse.png)

如果你好奇这是如何实现的，可以阅读这篇文章：[《Flare 制作记录：应用前后端性能优化》](https://soulteary.com/2022/01/19/flare-production-record-application-frontend-and-backend-performance-optimization.html)。

## 进阶文档

- [自定义启动参数](./docs/advanced-startup.md)
- [关闭免登陆模式后，如何设置用户账号](./docs/application-account.md)
- [如何挑选和使用图标](./docs/material-design-icons.md)
- [如何和 Traefik 一起使用](https://github.com/soulteary/traefik-example)

## 相比较 Flame

- 服务资源消耗极低，可以跑在任何规格的机器上，甚至是一台搭载2015年S805芯片的ARM盒子。
- 程序页面性能非常好，渲染速度更快，支持同时渲染大量（数千）书签，而不必担心风扇起飞。
- 使用声明的配置来进行导航内容的管理，无需担心数据迁移问题。
- 简化了天气数据的获取方式，不再需要申请天气网站的 `API_KEY` ，避免了不必要的成本开销。
- 简化了 Flame 中的K8S、Docker 集成等不必要的功能。
- 内置了大量风格统一、高质量的矢量图标，减少选择困难症，确保界面长期“耐看”。
- 默认使用免登陆模式，避免 HomeLab、本地使用的用户有额外的登陆操作。

## 关于内置图标

程序内置了目前 [materialdesignicons.com](https://materialdesignicons.com/) 中所有的 Material Design Icons，你可以让你的每一个书签都拥有风格统一、高质量的矢量图标。

![](./screenshots/icon-cheat-sheets.png)

更多信息，可以参考 [如何挑选和使用图标](./docs/material-design-icons.md)。

## TODO

- [ ] 持续完善程序定制化功能
- [ ] 支持使用 API 进行内容管理
- [ ] 支持自定义主题配色

## Thanks

Inspired by https://github.com/pawelmalak/flame
