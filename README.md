# Flare ✨

兼顾轻量快速、界面美观的个人导航页面，适用于 HomeLab 或其他注重私密的场景。

无数据库依赖，使用简单的配置来保存数据，数据更加透明，并 100% 属于用户自己。

支持 x86 以及常见的 ARM 设备。

应用资源消耗非常低：

- CPU: < 1%
- MEM: < 30M
- Docker Image: < 10M

![](./screenshots/docker-image-size.png)


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
docker run --rm -it -p 5005:5005 -v `pwd`/app:/app soulteary/flare:0.2.4
```

如果你习惯使用 docker-compose：

```bash
docker-compose up -d
```

在命令执行完毕之后，默认访问浏览器的 `5005` 端口，就能看到下面的界面啦：

![](./screenshots/ui.png)

![](./screenshots/lighthouse.png)

## 相比较 Flame

- 服务资源消耗极低，可以跑在任何规格的机器上。
- 程序页面性能非常好，渲染速度更快，支持渲染大量书签，而不必担心风扇起飞。
- 使用声明的配置来进行导航内容的管理，无需担心数据迁移问题。
- 简化了天气数据的获取方式，不再需要申请天气网站的 `API_KEY` ，避免了不必要的成本开销。
- 简化了 Flame 中的登陆、K8S、Docker 集成等不必要的功能。

## 支持图标

支持所有的 Material Design Icons，你可以在 [materialdesignicons.com](https://materialdesignicons.com/) 或 [pictogrammers.github.io/@mdi](https://pictogrammers.github.io/@mdi/font/6.5.95/) 6700 个精心设计的图标中找到你喜欢的图标。

在配置中使用的图标名称无需使用连字符 `-`，直接输入名称即可。（大小写不敏感）

为了降低寻找和输入图标的成本，我将 `@mdi/font` 项目集成到了应用中，启动应用之后，访问 `/resources/mdi-cheat-sheets/` 可以打开图标列表页面。

在页面中点击具体图标，可以使用的图标名称就自动复制到剪贴板中啦。

![](./screenshots/icon-cheat-sheets.png)


## TODO

- [ ] 公网部署版本支持禁用设置
- [ ] 完善程序设置功能
- [ ] 支持使用 API 进行内容管理
- [ ] 支持自定义主题配色

## Thanks

Inspired by https://github.com/pawelmalak/flame
