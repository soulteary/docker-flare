# Flare ✨

如果你觉得这个项目有帮到你，欢迎点赞✨（star）给予鼓励；如果你希望收到这个项目的更新推送，可以点击关注 👀（watch）并选择适合自己的关注模式（推荐 release）。

If you think this project works for you, please star the repo to help the developer(s). If you want to get the latest updates of the project, subscribe through the ```watch``` button and select your interested events (```Releases only``` recommended)

---

轻量、快速、美观的个人导航页面，适用于 HomeLab 或其他注重私密的场景。

A lightweight, high performance, and well-designed self-hosted navigation webpage, suit for HomeLab and other private applications.

无任何数据库依赖，应用数据完全开放透明，100% 属于用户自己。

Does not rely on any database, with application data fully transparent to the user(s)

支持在线编辑，内置 Material Design Icons 6k+ 图标，目前累计下载 1k+，期待你的反馈 :)

Supports online edits, with built-in Material Design Icons (6k+). So far over 1k downloads, and your feedbacks are appreciated :)

支持 x86 以及常见的 ARM 设备，应用资源消耗非常低：

Supports x86 architecture and common ARM devices, with extremely low resource usages:

- CPU: < 1%
- MEM: < 30M
- Docker Image: < 10M

![Flare Docker Pulls](./screenshots/docker-pulls.png)
![Flare Docker Image Size](./screenshots/docker-image-size.png)

## 快速上手 Quickstart

快速上手 Flare，需要两步：**下载**包含示例的代码、**启动**程序访问浏览器。

To Quickstart the Flare, 2 steps are required: **Download** the codes with examples included; **Run** the application and use web browser for access

### 下载包含示例的代码 Download the codes with examples included

你可以使用 `git clone` 或者选择使用 “Download ZIP” 的方式，下载包含了基础的配置示例（书签和应用）的代码：

You can use `git clone` or the Github's "Download ZIP" button, to download the code files with the sample configuration of bookmarks and applications.

```bash
git clone https://github.com/soulteary/docker-flare.git
cd docker-flare
```

`app/*yml` 目录中包含了你的书签和应用数据，你可以根据你的需求对其进行调整。如果目录中没有配置文件，应用将在首次运行的时候，进行自动创建。

The `app/*yml` path contains all the configuration samples of bookmarks and applications, and you can change them to your own settings.

### 启动程序访问浏览器 Run the program and access through web browser

启动应用非常简单，如果你习惯使用 Docker，可以执行：

To run the application, it's pretty simple do it with Docker like this:

```bash
# 可以使用最新镜像 Use the latest version
docker run --rm -it -p 5005:5005 -v `pwd`/app:/app soulteary/flare
# 也可以追求明确，使用固定版本 Use specific version for stability
docker run --rm -it -p 5005:5005 -v `pwd`/app:/app soulteary/flare:0.3.1
```

如果你习惯使用 docker-compose，只需要执行：

If you prefer docker-compose, then use the command below:

```bash
docker-compose up -d
```

如果你是 Traefik 用户，可以参考 `docker-compose.traefik.yml` 配置文件来使用。

If you want to use Traefik, check the `docker-compose.traefik.yml` configuration file as an example.

不论是哪一种方式，在命令执行完毕之后，默认情况下，我们访问浏览器的 `5005` 端口，就能看到下面的界面啦：

No matter which method being used, after running the command(s), use the web browser to open the `5005` port of either `localhost` or server IP (even internal hostname) to access the application page, as shown below:

![Flare Web UI](./screenshots/ui.png)

### 程序使用向导 User Guide

为了方便你的使用，我制作了一个简单的向导程序，在 Flare 启动之后，你可以随时访问 `/guide`，获取 Flare 基础界面功能的介绍。

A simple guide will show up at first time use, and after that you can still visit `/guide` page for instructions of Flare's various basic functions.

![Flare Guide](./screenshots/flare-guide.png)


## 程序在线编辑页面 Online Edit

为了满足随时随地编辑的需求，程序新增了“在线编辑”的页面。

Other than editing the configuration files, there is a newly added "Online Edit" page for easy use.

![Flare Editor](./screenshots/editor-beta.png)

工具页面地址：`/editor`

Tool/Page at `/editor`

## 程序帮助页面 Help Page

为了减少不必要的地址记忆负担，程序新增了一个“帮助页面”，默认展示所有的程序内的工具页面。

To relieve from navigating through the complicated menu to find certain frequently used pages, they are listed on the "Help Page", which can be accessed through the question-mark icon at the bottom.

![Flare Help](./screenshots/flare-help.png)

工具页面地址：`/help`

Page at `/help`

## 程序性能 performance

“快”作为 Flare 对主要优势而言，自然是需要“满分”来加持。

High performance is the main "selling point" of Flare, so it easily achieves full score while testing its frontend through lighthouse.

![Flare Lighthouse Scores](./screenshots/lighthouse.png)

如果你好奇这是如何实现的，可以阅读这篇文章：[《Flare 制作记录：应用前后端性能优化》](https://soulteary.com/2022/01/19/flare-production-record-application-frontend-and-backend-performance-optimization.html)。

## 进阶文档 Advance Setup Guides

- [自定义启动参数 Customize Environment Variables](./docs/advanced-startup.md)
- [关闭免登陆模式后，如何设置用户账号 Setup of user account after turning on the login mode](./docs/application-account.md)
- [如何挑选和使用图标 How to find and use MDI icons](./docs/material-design-icons.md)
- [如何和 Traefik 一起使用 How to use Flare with Traefik](https://github.com/soulteary/traefik-example)

## 相比较 Flame Compared to Flame

- 服务资源消耗极低，可以跑在任何规格的机器上，甚至是一台搭载2015年S805芯片的ARM盒子。
<br>
Low resource usage, while being able to run on "any" platform/device, even an ARM device with a S805 chip released back in 2015.
- 程序页面性能非常好，渲染速度更快，支持同时渲染大量（数千）书签，而不必担心风扇起飞。
<br>
Fast rendering, no need to worry about burning your cooling fans even with thousands of bookmarks.
- 使用声明的配置来进行导航内容的管理，无需担心数据迁移问题。
<br>
Settings are saved to a configuration file in the volume/mount, with no need to worry about upgrade or migration.
- 简化了天气数据的获取方式，不再需要申请天气网站的 `API_KEY` ，避免了不必要的成本开销。
<br>
Simplify the weather API, to prevent the cost of getting the `API_KEY`.
- 简化了 Flame 中的K8S、Docker 集成等不必要的功能。
<br>
Remove the "unnecessary"(if you care more about the performance) K8S and Docker integration
- 内置了大量风格统一、高质量的矢量图标，减少选择困难症，确保界面长期“耐看”。
<br>
With MDI built-in to provide large amount of high-quality icons with uniform style, improving the look of the pages ascetically.
- 默认使用免登陆模式，避免 HomeLab、本地使用的用户有额外的登陆操作。
<br>
Disable the login function by default, to eliminate additional operations for HomeLab and Local-only users.

## 关于内置图标 About built-in icons

程序内置了目前 [materialdesignicons.com](https://materialdesignicons.com/) 中所有的 Material Design Icons，你可以让你的每一个书签都拥有风格统一、高质量的矢量图标。

The built-in [Material Design Icons](https://materialdesignicons.com/) are built into Flare, to provide you with large amount of high-quality icons with uniform style.

![](./screenshots/icon-cheat-sheets.png)

更多信息，可以参考 [如何挑选和使用图标](./docs/material-design-icons.md)。

For more information, check [How to find and use MDI icons](./docs/material-design-icons.md)

## TODO

- [ ] 持续完善程序定制化功能 Continue to improve the customization functions
- [ ] 支持使用 API 进行内容管理 Add support for content management through API
- [ ] 支持自定义主题配色 Add support for custom colors in theme setting

## Thanks

Inspired by https://github.com/pawelmalak/flame
