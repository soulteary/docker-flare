![Flare banner: a colorful squirrel holding an acorn in an oil-painted landscape](.github/flare-banner.jpg)

# Flare ✨

[![CodeQL](https://github.com/soulteary/flare/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/soulteary/flare/actions/workflows/codeql-analysis.yml) [![Security Scan](https://github.com/soulteary/flare/actions/workflows/scan.yml/badge.svg)](https://github.com/soulteary/flare/actions/workflows/scan.yml) [![Release](https://github.com/soulteary/flare/actions/workflows/release.yml/badge.svg)](https://github.com/soulteary/flare/actions/workflows/release.yml) [![Go Report Card](https://raw.githubusercontent.com/soulteary/flare/main/.github/goreportcard.svg)](https://github.com/soulteary/flare/blob/main/.github/goreportcard-report.md) [![Docker Image](https://img.shields.io/docker/pulls/soulteary/flare.svg)](https://hub.docker.com/r/soulteary/flare)


如果你觉得这个项目有帮到你，欢迎点赞✨（star）给予鼓励；如果你希望收到这个项目的更新推送，可以点击关注 👀（watch）并选择适合自己的关注模式（推荐 release）。

---

轻量、快速、美观的个人导航页面，适用于 HomeLab 或其他注重私密的场景。

无任何数据库依赖，应用数据完全开放透明，100% 属于用户自己。

当前已发布稳定版为 [Flare 0.6.10](https://github.com/soulteary/flare/releases/tag/0.6.10)。本仓库提供 Docker 部署示例和使用文档，程序源码及构建发布位于 [soulteary/flare](https://github.com/soulteary/flare)。

**版本边界：** 新搜索、列表／表单／批量编辑器、Tab／折叠和列数设置已包含在
0.6.5 起的镜像中；0.6.6 进一步完善搜索模式、设置保存、天气地区选择和图标兼容。0.6.7 新增分类内部多列布局，完善中英文翻译并修复图标显示和复制。0.6.8 加入持久化管理员账号、可撤销会话、分类／条目隐私保护和导航缓存。请核对实际 `flare --version` 和 digest。
0.6.9 恢复默认免登录、启动即可使用，内部安全状态自动准备；已有私有默认保持保护。
0.6.10 修复代理覆盖 Referrer 策略时的设置保存，完善手机／平板紧凑布局，新增隐藏版本信息选项，并优化导航读取与不必要的 CSRF 处理。详见[升级说明](docs/upgrade-0.6.10.md)和[发布核验](docs/release-0.6.10-verification.md)。
使用说明见[0.6.10 文档索引](https://github.com/soulteary/flare/blob/0.6.10/docs/README.md)。
Safari、Via、Android 和 iPad 的真机兼容性仍未验证，Chromium 自动化不能替代真机结果。

支持在线编辑，内置 Material Design Icons 7.4.47 及旧版兼容图标，目前累计下载过万，期待你的反馈 :)

支持 x86 以及常见的 ARM (ARM32v6、ARM32v7、ARM64v8)设备，应用资源消耗非常低：

这些旧版截图中的资源数字不是 0.6.10 或主分支的容量保证。CPU、RSS 和镜像大小
取决于版本、平台、书签数量、天气、日志及连接并发；镜像压缩大小与容器常驻内存也不同。
请使用[性能测量方法](https://github.com/soulteary/flare/blob/main/docs/baseline-metrics.md)
记录具体数据，避免将空闲 CPU、分配量和 RSS 混为一个指标。

<img src="screenshots/docker-pulls.png" width="90%"/>
<img src="screenshots/docker-image-size.png" width="90%"/>

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

0.6.9 默认免登录，内部安全状态自动准备，无需手动初始化管理员。能访问实例的人均可浏览和修改内容；需要登录保护时使用下面的 login/private 示例。

```bash
# Docker Compose：一条命令启动
docker compose up -d flare

# 或直接运行 Docker，固定到已验证版本
docker run -d --name flare -p 5005:5005 \
  --log-driver local --log-opt max-size=10m --log-opt max-file=3 \
  -v "$PWD/app:/app" soulteary/flare:0.6.10
```

也可以将镜像替换为 `ghcr.io/soulteary/flare:0.6.10`。Docker Hub 和 GHCR 提供相同 digest 的多架构镜像，均支持 `linux/amd64`、`linux/arm64`、`linux/arm/v7` 和 `linux/arm/v6`，也都提供 `latest` 标签。

[公开首页、管理需登录](./example/login/docker-compose.yml)与[全站私有](./example/private/docker-compose.yml)示例保留登录保护，通过 `FLARE_PASS` 初始化凭据。启动前设置用户名和密码；缺少密码时 Compose 明确报错。详见[账号指南](./docs/application-account.md)。

[Traefik 示例](./example/traefik/docker-compose.yml)仍采用全站私有模式。先替换域名及 `FLARE_EXTERNAL_URL`，按账号指南在该示例目录初始化管理员后再启动，不要直接将后端暴露为免登录服务。

已有部署请阅读[0.6.10 升级说明](./docs/upgrade-0.6.10.md)。保留原有访问模式和数据挂载，不要直接用新的默认 Compose 覆盖已有私有配置。

以上运行服务的 Docker 命令和本仓库的 Compose 示例均使用 `local` 日志驱动，单个日志文件达到 `10m` 后轮转，最多保留 3 个日志文件。使用 `docker logs` 仍可查看日志。兼容 `json-file` 的配置、访问日志模式及现有部署的更新步骤见[日志配置与容量说明](./docs/logging.md)。

启动后访问 `5005` 端口即可使用；login/private/Traefik 示例按各自配置要求登录。下面是旧版界面截图：

![Flare Web UI](./screenshots/ui.png)

### 升级现有部署

0.6.9 恢复默认免登录，但不会替你删除旧 Compose 中的登录配置。请保留已有的 `FLARE_DISABLE_LOGIN=false`、`FLARE_VISIBILITY=PRIVATE` 和凭据来源；仅更新镜像即可继续私有访问。已记录的私有默认也会保留。详见[0.6.10 升级说明](./docs/upgrade-0.6.10.md)。

`.flare-security/` 必须持久化且可写，默认 `./app:/app` 整目录挂载已包含它。仅挂载三个 YAML 的部署需增加安全目录挂载。升级前停止实例并备份完整数据目录（含隐藏文件）、Compose/.env/Secret；有效旧登录配置会迁移，旧 Cookie 失效一次。不要依赖旧日志查找新生成密码。完整迁移、异常处理和回滚见[0.6.8 升级指南](./docs/upgrade-0.6.8.md)及[变更记录](./CHANGELOG.md)。


0.6.7 支持分类内部多列布局，完善设置、编辑器和导入导出的中英文翻译，修复主题下的图标对比度及图标速查表复制，并区分写入来源校验失败与登录过期。升级前备份完整数据目录及部署配置，重建容器后刷新已打开的设置和编辑页面。详见 [0.6.7 发布说明](https://github.com/soulteary/flare/releases/tag/0.6.7)。

0.6.6 新增明确的搜索模式和天气地区选择，保护设置升级与并发保存，更新 MDI 图标并保留 19 个旧图标，调整主题 favicon 和手机／平板布局。搜索 `!keyword` 快捷语法已移除；天气空地区需手动选择，旧地区仅在唯一匹配时自动迁移。升级前备份完整数据目录（含隐藏目录）、Compose 和外部密钥；回滚时恢复升级前的完整备份，不能只换旧镜像。详见 [0.6.6 发布说明](https://github.com/soulteary/flare/releases/tag/0.6.6)、[配置与迁移](docs/configuration.md)和[备份恢复](docs/backup-restore.md)。

0.6.5 新增列表／表单／批量编辑器、草稿恢复、搜索引擎和混合搜索、分类 Tab／折叠及布局控制；修正动态地址、配置解析、保存恢复、天气刷新与资源缓存。升级前先停止写入并备份完整数据目录、Compose 和外部密钥。支持的链接协议现在会被校验，未支持的协议保留在数据中但显示为 `#`。详见 [0.6.5 发布说明](https://github.com/soulteary/flare/releases/tag/0.6.5) 和[备份恢复](docs/backup-restore.md)。

0.6.4 改善移动页面布局和触摸编辑，将 HTTP 服务迁移到原生 `Fiber v3.6.1`，并将内存文件系统依赖更新到 `memfs-kit v1.1.1`。HTTP 访问日志默认只记录失败请求，可通过 `FLARE_ACCESS_LOG=errors|all|off` 或 `--access_log` 调整；启动信息和应用日志继续保留。本仓库的部署示例同步启用 Docker 日志轮转，并更新网站图标和 README 横幅。详情见 [变更记录](./CHANGELOG.md) 和 [Flare 0.6.4 发布说明](https://github.com/soulteary/flare/releases/tag/0.6.4)。

0.6.3 新增的实时时钟、秒数显示和显示时区设置继续保留。现有配置默认仍使用跟随服务器时区、包含秒数的静态显示，无需迁移。

0.6.2 的 HTTP 登录修复、浏览器书签 HTML / Flare YAML 导入导出及导入前备份继续保留。

如果反向代理终止 HTTPS、再以 HTTP 连接 Flare，请为服务设置 `FLARE_COOKIE_SECURE=true`，并始终通过 HTTPS 访问。普通 HTTP 部署无需开启该选项。

本仓库的 Compose 示例固定使用 `soulteary/flare:0.6.10`，升级版本时需显式修改 `image`；需要持续跟随稳定版可改为 `soulteary/flare:latest`。先备份现有 `./app` 数据，并确认继续使用原来的 `./app:/app` 挂载；若使用命名卷，请保留原卷名和 Compose 项目名。在原部署目录中，拉取新镜像并重建同一个服务：

```bash
docker compose pull flare
docker compose up -d --force-recreate flare
docker compose exec flare flare --version
```

需要固定版本时，将 `docker-compose.yml` 中的 `image` 改为 `soulteary/flare:0.6.10`，再执行以上命令。重建会保留原来的数据挂载。日志驱动、轮转选项和 Compose 环境变量需要重建容器才生效；发布新镜像或仅执行 `docker compose restart` 都不会自动替换正在运行的旧容器。如何核对挂载及容器实际日志配置，见[日志配置与容量说明](./docs/logging.md#更新现有部署)。

Docker 日志轮转可直接用于现有镜像。`FLARE_ACCESS_LOG` 和 `--access_log` 从 0.6.4 起可用，默认值为 `errors`；需要记录普通访问时，在服务的 `environment` 中加入 `FLARE_ACCESS_LOG=all` 后重建容器。使用旧镜像前请检查版本，或运行镜像的 `flare --help` 确认支持的参数。

使用 GHCR 时，将同一个 `image` 字段改为 `ghcr.io/soulteary/flare:0.6.10`（固定版本）或 `ghcr.io/soulteary/flare:latest`（跟随最新稳定版），其余配置和升级命令相同。

如果升级后浏览器仍显示旧的图标加载结果，可强制刷新页面或清除该站点缓存。此前为绕过问题而启用的 `FLARE_MINI_REQUEST` 可以按需保留，0.6.1 起默认模式也已恢复正常。

### 日期、时间与时区

0.6.3 起，进入 **设置 → 外观 → 模块设置**，启用日期时间后，可以选择页面加载时显示或实时更新、是否显示秒数，以及跟随服务器、跟随访问设备或指定显示时区。指定时区支持搜索建议，也可以输入 `Asia/Shanghai`、`America/New_York` 等 IANA 地理时区名称或 `UTC`；修改选项时会显示预览。

未配置 `Clock` 时，继续使用跟随服务器时区、包含秒数的静态显示。选择跟随服务器时，时区通常由容器的 `TZ` 配置；显示时区设置不会修改服务器时钟、进程全局时区或天气城市。

实时显示使用**访问设备的系统时间**，再换算到所选时区；即使选择跟随服务器，也只是跟随服务器的时区。服务器或指定时区的静态显示使用服务器时间，设备时区的静态显示只在初始化时读取一次设备时间。设备时间不准会影响实时显示；此功能不会进行 NTP 校时，也不会请求外部时间接口。

禁用 JavaScript 或浏览器不支持所需时区能力时，页面会保留可读的服务端时间。配置字段、时区优先级与降级行为见[日期、时间与时区说明](https://github.com/soulteary/flare/blob/0.6.10/docs/clock.md)。

### 程序使用向导

为了方便你的使用，我制作了一个简单的向导程序，在 flare 启动之后，你可以随时访问 `/guide`，获取 flare 基础界面功能的介绍。

![Flare Guide](./screenshots/flare-guide.png)


## 程序在线编辑页面

为了满足随时随地编辑的需求，0.6.5 起提供列表／表单／批量编辑、图标选择和草稿恢复。下面保留的截图展示早期界面，当前使用方法见[编辑器说明](https://github.com/soulteary/flare/blob/0.6.10/docs/editor.md)。

![Flare Editor](./screenshots/editor-beta.png)

工具页面地址：`/editor`

0.6.2 起，可在编辑器中选择「导入 / 导出浏览器书签与 Flare YAML」，预览后合并导入 Chrome、Edge、Firefox 等浏览器导出的书签 HTML，或 Flare YAML；也可导出 HTML / YAML 和下载最近一次导入前备份。

导入支持 UTF-8、最大 8 MiB 的文件。备份位于工作目录下的 `.flare-bookmark-backups/`；使用默认 Docker / Compose 部署时，容器中的路径是 `/app/.flare-bookmark-backups/`，对应宿主机的 `./app/.flare-bookmark-backups/`，现有的 `./app:/app` 挂载已包含该目录。自定义工作目录时，请一并持久化其中的备份目录。YAML 合并导入不等于完整配置恢复。详见 [书签迁移说明](https://github.com/soulteary/flare/blob/0.6.10/docs/bookmark-transfer.md)。

## 程序帮助页面

为了减少不必要的地址记忆负担，程序新增了一个“帮助页面”，默认展示所有的程序内的工具页面。

![Flare Help](./screenshots/flare-help.png)

工具页面地址：`/help`

## 程序性能

“快”作为 Flare 对主要优势而言，自然是需要“满分”来加持。

![Flare Lighthouse Scores](./screenshots/lighthouse.png)

0.6.4 的 Fiber 迁移、测试条件与性能取舍见[性能记录](https://github.com/soulteary/flare/blob/0.6.4/docs/performance.md)。其中的历史实验关闭了访问日志，不能直接作为所有部署的性能保证。

如果你好奇这是如何实现的，可以阅读这篇文章：[《Flare 制作记录：应用前后端性能优化》](https://soulteary.com/2022/01/19/flare-production-record-application-frontend-and-backend-performance-optimization.html)。

## 进阶文档

- [配置参考、默认值和访问模式](./docs/configuration.md)
- [备份与完整恢复](./docs/backup-restore.md)
- [反向代理和子路径边界](./docs/reverse-proxy.md)
- [故障排查与反馈信息](./docs/troubleshooting.md)
- [自定义启动参数](./docs/advanced-startup.md)
- [日志配置与 Docker 日志容量](./docs/logging.md)
- [0.6.9 升级与访问模式](./docs/upgrade-0.6.9.md)
- [管理员初始化、密码来源与会话撤销](./docs/application-account.md)
- [如何挑选和使用图标](./docs/material-design-icons.md)
- [日期、时间与时区](https://github.com/soulteary/flare/blob/0.6.10/docs/clock.md)
- [如何和 Traefik 一起使用](https://github.com/soulteary/traefik-example)

## 相比较 Flame

- 服务资源消耗极低，可以跑在任何规格的机器上，甚至是一台搭载2015年S805芯片的ARM盒子。
- 程序页面性能非常好，渲染速度更快，支持同时渲染大量（数千）书签，而不必担心风扇起飞。
- 使用声明的配置来进行导航内容的管理，无需担心数据迁移问题。
- 简化了天气数据的获取方式，不再需要申请天气网站的 `API_KEY` ，避免了不必要的成本开销。
- 简化了 Flame 中的K8S、Docker 集成等不必要的功能。
- 内置了大量风格统一、高质量的矢量图标，减少选择困难症，确保界面长期“耐看”。
- 默认免登录、启动即可使用；需要保护管理或全站访问时可选择对应登录示例。

## 关于内置图标

0.6.9 沿用 Material Design Icons 7.4.47 的 7,447 个图标，并保留 19 个旧版图标标识及原 SVG，你可以让你的每一个书签都拥有风格统一、高质量的矢量图标。

![](./screenshots/icon-cheat-sheets.png)

更多信息，可以参考 [如何挑选和使用图标](./docs/material-design-icons.md)。

## TODO

- [ ] 持续完善程序定制化功能
- [ ] 支持使用 API 进行内容管理
- [ ] 支持自定义主题配色

## Thanks

Inspired by https://github.com/pawelmalak/flame
