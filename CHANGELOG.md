# 变更记录

项目应用中的所有变更记录，将在这个文件中进行存档。

## 目录

* [0.6.11](#0611---2026-10-28)
* [0.6.10](#0610---2026-10-27)
* [0.6.9](#069---2026-10-27)
* [0.6.8](#068---2026-10-27)
* [0.6.7](#067---2026-10-27)
* [0.6.6](#066---2026-10-26)
* [0.6.5](#065---2026-10-26)
* [0.6.4](#064---2026-10-25)
* [0.6.3](#063---2026-10-25)
* [0.6.2](#062---2026-10-25)
* [0.6.1](#061---2026-10-25)
* [0.5.1](#051---2024-01-06)
* [0.5.0](#050---2024-01-05)
* [0.4.3](#043---2023-12-07)
* [0.4.2](#042---2023-12-06)
* [0.4.1](#041---2023-05-07)
* [0.4.0](#040---2022-09-17)
* [0.3.3](#033---2022-03-10)
* [0.3.2](#032---2022-02-23)
* [0.3.1](#031---2022-02-23)
* [0.2.10](#0210---2022-02-21)
* [0.2.9](#029---2022-02-20)
* [0.2.8](#028---2022-02-18)
* [0.2.7](#027---2022-02-17)
* [0.2.6](#026---2022-02-17)
* [0.2.5](#025---2022-02-16)
* [0.2.4](#024---2022-02-11)
* [0.2.3](#023---2022-02-07)
* [0.2.1](#021---2022-02-07)
* [0.2.0](#020---2022-01-30)
* [0.1.0](#010---2022-01-27)

# 0.6.11 - 2026-10-28

同步 [Flare 0.6.11](https://github.com/soulteary/flare/releases/tag/0.6.11)。部署仓库标签为 `v0.6.11`，应用镜像标签为 `0.6.11`（无 `v`），程序与镜像仍由源码仓库构建发布。

- [部署] 四份 Compose 固定为 `soulteary/flare:0.6.11`，保留访问模式、数据挂载、反向代理和日志轮转配置。
- [诊断] `SOURCE_REJECTED` 响应新增稳定原因码、可读说明、调整建议及请求 ID，可区分协议、Host、端口、不透明 Origin、重复来源头与跨站提交。
- [编辑器] 保存失败时展示详细诊断并保留草稿；旧格式响应继续显示本地化说明。
- [安全] 未放宽来源保护，未声明可信的代理转发头仍不采信，跨站和恶意写入继续拒绝。
- [文档] 同步当前版本引用，补充[升级说明](docs/upgrade-0.6.11.md)和[发布核验](docs/release-0.6.11-verification.md)。

升级时保留完整数据目录、访问配置和安全状态，拉取镜像后重建容器。遇到来源错误时附上 `error.reason` 与 `request_id`，但不要发布 Cookie、密码或完整私有 Referer。

# 0.6.10 - 2026-10-27

同步 [Flare 0.6.10](https://github.com/soulteary/flare/releases/tag/0.6.10)。部署仓库标签为 `v0.6.10`，应用镜像标签为 `0.6.10`（无 `v`），程序与镜像仍由源码仓库构建发布。

- [部署] 四份 Compose 固定为 `soulteary/flare:0.6.10`，保留根目录默认免登录、login/private/Traefik 的原有访问模式、数据挂载与日志轮转。
- [保存] 修复代理响应头覆盖 `Referrer-Policy: no-referrer` 时，正常设置表单可能被 `SOURCE_REJECTED` 拒绝；不放宽来源和 CSRF 校验。
- [布局] 可选紧凑模式覆盖至 1024 CSS 像素，统一图文对齐，长描述单行省略，工具按钮不再遮挡书签；保留默认布局、列数及字体大小选项。
- [外观与性能] 新增隐藏版本信息选项，优化导航搜索和分类读取，减少无需 CSRF 的浏览请求中的额外处理。
- [文档] 同步当前版本引用，补充[升级说明](docs/upgrade-0.6.10.md)和[发布核验](docs/release-0.6.10-verification.md)。历史核验记录保留原始版本语义。

升级时保留原有安全状态、访问配置与数据挂载，拉取镜像后重建容器。Host／协议／端口不一致仍需修正代理配置；实体 Safari／Via／Android／iPad 兼容性不作未经测试的承诺。

# 0.6.9 - 2026-10-27

同步 [Flare 0.6.9](https://github.com/soulteary/flare/releases/tag/0.6.9)。部署仓库标签为 `v0.6.9`，应用镜像标签为 `0.6.9`。

- 四份 Compose 固定到 `soulteary/flare:0.6.9`。根目录恢复默认免登录，启动自动准备内部状态，不要求手动初始化。
- login/private/Traefik 保留登录保护及原有数据挂载、日志轮转配置。
- 更新快速开始、账号、配置与故障排查，补充 [0.6.9 升级说明](docs/upgrade-0.6.9.md)，明确已有私有部署必须保留访问配置。
- 继续保留私有内容冲突检查、CSRF、状态故障保护、密码和会话机制。部署核验见[验收记录](docs/release-0.6.9-verification.md)。

# 0.6.8 - 2026-10-27

同步 [Flare 0.6.8](https://github.com/soulteary/flare/releases/tag/0.6.8)。本仓库沿用 `v` 前缀，对应 `v0.6.8`；应用镜像标签为 `0.6.8`（无 `v`），由源码仓库构建发布。

- [部署] 四份 Compose 示例固定为 `soulteary/flare:0.6.8`；根目录和 Traefik 示例使用管理员初始化，login/private 示例要求显式部署密码，分别提供公开首页或全站私有访问。
- [账号] 新实例默认需要管理员且全站私有；支持持久化单管理员、密码来源管理、恢复、可撤销会话，普通重启保留有效会话。
- [隐私] 分类和条目的 private 标记执行权限检查，完善编辑器草稿、导入导出及重新认证提示。
- [性能] 合并持久会话读取，缓存导航 HTML 片段和请求设置，恢复公开静态资源缓存。
- [升级] 补充 `.flare-security/` 持久化、旧 Cookie 一次失效、旧配置迁移、YAML 别名网页只读、备份恢复和回滚边界；移除默认免登录和日志生成密码的过时说明。

升级前停止实例并备份完整目录及外部配置；仅挂载三个 YAML 的部署需要新增安全目录挂载。详情见[0.6.8 升级与回滚](docs/upgrade-0.6.8.md)。源码仓库的最终七平台归档、四平台容器及浏览器预检和正式发布均已通过；本仓库部署示例的验收记录见[部署核验](docs/release-0.6.8-verification.md)。实体 Safari/Via/Android/iPad 和树莓派兼容性不作未经测试的承诺。

# 0.6.7 - 2026-10-27

同步 [Flare 0.6.7](https://github.com/soulteary/flare/releases/tag/0.6.7)。本仓库沿用 `v` 前缀，对应 `v0.6.7`；应用镜像和二进制由源码仓库发布，镜像标签为 `0.6.7`（无 `v`）。

- [部署] 四份 Compose 示例固定为 `soulteary/flare:0.6.7`，README、GHCR 示例、帮助命令和当前版本文档链接同步更新。
- [布局] 支持书签分类内部多列布局，改善桌面、手机和平板上的分类展示。
- [翻译] 接入 i18n-kit，完善设置、编辑器、导入导出及错误提示的中英文翻译。
- [图标] 修复不同主题下编辑器图标对比度，同步 MDI 速查表与图标库，复制时保留 Flare 图标名称。
- [修正] 区分写入来源校验失败与登录过期，完善导航错误信息的 HTML 转义；改进 Go 静态检查及基准分析输出目录限制。

升级前停止写入并备份完整数据目录、Compose 和外部密钥；重建容器后刷新已打开的设置和编辑页面。应用七个归档及双镜像仓库四平台产物已校验，详见[应用发布说明与验收](https://github.com/soulteary/flare/releases/tag/0.6.7)。Safari／Via／Android／iPad 真机兼容性仍未验证。

# 0.6.6 - 2026-10-26

同步 [Flare 0.6.6](https://github.com/soulteary/flare/releases/tag/0.6.6)。本仓库沿用 `v` 前缀，对应 `v0.6.6`；应用镜像与二进制继续由源码仓库发布，镜像标签为 `0.6.6`（无 `v`）。

- [部署] 四份 Compose 示例固定为 `soulteary/flare:0.6.6`，保留数据挂载、访问模式和日志轮转；README、GHCR 示例、帮助命令及当前配置文档同步到 0.6.6。
- [搜索] 明确混合／本地优先／选择器模式，移除 `!keyword` 快捷语法；补充配置版本、并发保护、原子保存和私有升级备份说明。
- [天气] 更新地区选择、唯一匹配迁移及用户主动确认服务器出口 IP 建议的说明。
- [图标与布局] MDI 更新至 7.4.47，保留 19 个旧图标的原 SVG；主题 favicon 和手机／平板布局随应用更新。
- [配置] 统一 YAML v3 解析与校验；旧 CSV 编辑提交继续返回 409，防止绕过版本冲突保护。

升级前停止写入，备份完整数据目录（含隐藏文件）、Compose 和外部密钥；回滚需恢复完整升级前备份，不能仅替换镜像。Safari／Via／Android／iPad 真机兼容性仍未验证。公开产物及升级回滚验证见[应用验收记录](https://github.com/soulteary/flare/blob/fa477897ca2faa7c249368d7172dd7c9648d3115/docs/release-0.6.6-verification.md)。

# 0.6.5 - 2026-10-26

同步 [Flare 0.6.5](https://github.com/soulteary/flare/releases/tag/0.6.5)。应用镜像和二进制由源码仓库发布，本仓库同步部署文档；此记录不表示部署仓库的版本标签已经创建。

- [新增] 列表／表单／批量编辑器、分类管理、图标选择、中英文界面、未保存提示、版本冲突保护与 JSON 草稿导出恢复。
- [新增] 可配置搜索引擎和混合搜索、分类 Tab／折叠、移动和桌面布局控制。
- [修正] 导航文字和属性转义、链接和图片协议校验、请求级动态地址、编码跳转及 IPv6／Base64 处理。
- [修正] 隔离无效环境字段，保留认证设置，修复旧布尔参数别名，校验修改请求来源，保留失败保存的表单并保护恢复副本。
- [优化] 天气异步刷新、请求合并、过期数据与失败退避；改进资源缓存升级、解析快照、分类渲染、编辑器和资源元数据路径。
- [文档] 补充配置优先级、备份恢复、反代边界、问题定位和可复现性能报告。性能报告以审计前主分支而非 0.6.4 为基准，RSS 并非全面下降，详见[实测限制](https://github.com/soulteary/flare/blob/0.6.5/metrics-reports/audit-resource-comparison-2026-10-26.md)。

升级前备份完整数据目录和外部配置，保留原挂载并重建容器。未支持的链接协议仍保存在数据中，但渲染为 `#`；单条书签 private 标记不提供单条访问控制。Safari／Via／Android／iPad 真机验收仍未验证；完整子路径部署、多用户和上传图标仍属独立功能。

# 0.6.4 - 2026-10-25

同步 [Flare 0.6.4](https://github.com/soulteary/flare/releases/tag/0.6.4)，本仓库沿用 `v` 前缀，对应版本标签为 `v0.6.4`。应用镜像和二进制由源码仓库发布，本仓库同步部署文档和发行记录。

- [修正] 改善移动页面布局和触摸编辑，覆盖分类选择、书签字段保留与保存失败后的重试。详见 [flare #126](https://github.com/soulteary/flare/pull/126)。
- [维护] HTTP 服务由 Echo 迁移为原生 `Fiber v3.6.1`，保留原有会话、资源缓存和 Range 请求行为，完善错误响应、请求状态和表单处理；归档性能比较、跳转路径优化和复现方法。详见 [flare #128](https://github.com/soulteary/flare/pull/128) 和[性能记录](https://github.com/soulteary/flare/blob/0.6.4/docs/performance.md)。
- [维护] 内存文件系统依赖由 `memfs-kit v1.1.0` 更新至 `v1.1.1`。
- [优化] HTTP 访问日志默认只记录返回错误或状态为 4xx、5xx 的请求；新增 `FLARE_ACCESS_LOG=errors|all|off` 与 `--access_log`，保留应用启动信息和异常日志。详见 [flare #125](https://github.com/soulteary/flare/issues/125) 和 [flare #130](https://github.com/soulteary/flare/pull/130)。
- [部署] 四份 Compose 示例和 README 的 Docker 启动命令默认使用 `local` 日志驱动，单个日志文件上限 `10m`，最多保留 3 个文件；补充兼容 `json-file` 的配置和现有部署的重建、检查步骤。详见 [docker-flare #189](https://github.com/soulteary/docker-flare/pull/189) 和[日志配置与容量说明](./docs/logging.md)。
- [外观] 更新圆形松鼠网站图标和 README 横幅。详见 [flare #127](https://github.com/soulteary/flare/pull/127)、[flare #129](https://github.com/soulteary/flare/pull/129) 和 [docker-flare #188](https://github.com/soulteary/docker-flare/pull/188)。
- [文档] 固定版本镜像、帮助命令和源码文档链接同步至 0.6.4，明确访问日志开关的最低版本和升级方式。

升级注意：无需迁移现有书签和设置。先备份并保留原有数据挂载，拉取镜像后重建容器；仅重启不会应用新镜像、日志驱动或 Compose 环境变量。默认成功访问不再产生访问日志，需要保留普通访问记录时设置 `FLARE_ACCESS_LOG=all`。Docker 轮转配置可用于旧镜像，访问日志开关从 0.6.4 起可用。0.6.2 的 HTTPS 反代 Cookie、书签备份持久化要求，以及 0.6.3 的时钟默认行为继续适用。

# 0.6.3 - 2026-10-25

同步 [Flare 0.6.3](https://github.com/soulteary/flare/releases/tag/0.6.3)，本仓库沿用 `v` 前缀，对应版本标签为 `v0.6.3`。应用镜像和二进制由源码仓库发布，本仓库同步部署文档和发行记录。

- [新增] 外观的模块设置支持可选的实时时钟、秒数显示，以及跟随服务器、跟随访问设备或指定 IANA 显示时区；提供时区搜索建议和预览。日期、星期、时间和分时问候语按同一时区换算。详见 [flare #123](https://github.com/soulteary/flare/pull/123) 和[日期、时间与时区说明](https://github.com/soulteary/flare/blob/0.6.3/docs/clock.md)。
- [优化] 实时时钟在页面隐藏时暂停，恢复显示时重新读取当前设备时间；浏览器能力不足或禁用 JavaScript 时保留服务端输出。现有配置继续使用跟随服务器时区、包含秒数的静态显示。
- [维护] 内存文件系统依赖由 `memfs-kit v1.0.0` 更新至 `v1.1.0`。详见 [flare #124](https://github.com/soulteary/flare/pull/124)。
- [维护] Echo 更新至 `v5.4.1`，纳入上游安全补丁。详见 [Echo v5.4.1 发布说明](https://github.com/labstack/echo/releases/tag/v5.4.1)。
- [文档] 固定版本部署示例同步至 0.6.3，补充时钟入口、时间来源及默认行为，完善 Compose 命令、Traefik 示例链接和书签备份持久化路径。

升级注意：实时显示使用访问设备的系统时间，再换算到所选显示时区，不提供 NTP 或网络校时。无需迁移现有配置；0.6.2 的 HTTP Cookie 配置和书签导入备份要求继续适用。拉取镜像后须重建容器，仅重启不会更新镜像。

# 0.6.2 - 2026-10-25

同步 [Flare 0.6.2](https://github.com/soulteary/flare/releases/tag/0.6.2)，本仓库沿用 `v` 前缀，对应版本标签为 `v0.6.2`。应用镜像和二进制由源码仓库发布，本仓库同步部署文档和发行记录。

- [修正] 修复 0.6.0 / 0.6.1 在普通 HTTP 下正确输入凭据后仍返回登录页的问题，明确会话 Cookie 属性和登录退出行为。详见 [flare #122](https://github.com/soulteary/flare/pull/122)。
- [新增] 编辑器支持浏览器书签 HTML 和 Flare YAML 的导入、预览、去重策略、合并和导出；覆盖应用与分类书签，导入前自动备份，支持下载最近一次导入前备份。详见 [flare #119](https://github.com/soulteary/flare/pull/119)。
- [维护] 内存文件系统迁移到 memfs-kit v1.0.0。
- [发布] Docker Hub 发布成功后自动同步四平台镜像到 GHCR，并验证版本、源码提交及 manifest digest 一致。
- [文档] 固定版本镜像示例更新至 0.6.2，补充 HTTPS 反代 Cookie 和书签备份说明。

升级注意：HTTPS 反代以 HTTP 连接 Flare 时，请设置 `FLARE_COOKIE_SECURE=true` 并通过 HTTPS 访问。书签导入备份目录 `.flare-bookmark-backups/` 应持久化；YAML 合并导入不等于完整配置恢复。拉取镜像后须重建容器，仅重启不会更新镜像。实时钟及可选显示时区不属于 0.6.2。

# 0.6.1 - 2026-10-25

同步 [Flare 0.6.1](https://github.com/soulteary/flare/releases/tag/0.6.1)，本仓库的版本标签为 `v0.6.1`。

- [修正] 修复 0.6.0 默认模式下内置 MDI 图标请求返回 500 的问题。Echo v5 的静态文件服务要求文件支持定位读取，原内存文件实现不满足该接口；现在通过独立的可定位读取器提供生成的 SVG。https://github.com/soulteary/docker-flare/issues/184 https://github.com/soulteary/flare/pull/117
- [修正] 图标使用与内容对应的 ETag，缺失资源返回 404；避免错误响应被长期缓存，以及共享 ETag 将不存在的资源错误地变成 304。
- [优化] 完善图标缓存的并发访问，以及 GET、HEAD、Range、条件请求和主题切换的回归验证；请求最小化模式继续支持内联 SVG。
- [发布] 新增 GHCR 镜像源 `ghcr.io/soulteary/flare:0.6.1` 和 `:latest`，与 Docker Hub 镜像保持相同 digest；四个平台已通过公开访问验证。后续 Flare 稳定版本会在 Docker Hub 发布完成后自动同步到 GHCR。
- [文档] 更新固定版本示例和升级步骤。默认 Compose 配置继续使用 `soulteary/flare:latest`；更新现有服务需要执行 `docker compose pull flare` 和 `docker compose up -d flare`。

更早版本的应用变更请参考 [0.5.2](https://github.com/soulteary/flare/releases/tag/0.5.2) 和 [0.6.0](https://github.com/soulteary/flare/releases/tag/0.6.0) 的发布说明。

# 0.5.1 - 2024-01-06

- [修正] 修正 .env 文件读取时的逻辑。
- [优化] 完善程序解析参数时的测试覆盖率。

# 0.5.0 - 2024-01-05

- [优化] 使用 Go 新版日志工具替换三方日志库，性能提升。
- [修正] 修正一个 HTML Tag 笔误。（ 感谢 @LightAPIs 同学 #37 ）
- [修正] 修正华为手机访问时的样式问题。（ 感谢 @LightAPIs 同学 #38 ）
- [修正] 修正一个潜在的程序隐患，让程序运行更加稳定。
- [优化] 完善程序解析参数时的测试覆盖率。

# 0.4.3 - 2023-12-07

- [修正] 修正因为 Golang 版本更新，导致的模版解析出现的展示异常。

# 0.4.2 - 2023-12-06

这个版本，特别感谢 @LightAPIs 的提交。

- [新增] 根据环境，进行动态链接的展示和切换。（感谢 @hanuxxx 、@ztzheng 两位同学的反馈）
- [修正] 修正了一些使用场景下，加密链接转换的问题 https://github.com/soulteary/docker-flare/issues/119 （感谢 @sungamma 的反馈）
- [修正] 编辑器现在可以自由排序内容了 https://github.com/soulteary/docker-flare/issues/124 （感谢 @elvescn 的反馈）
- [修正] 编辑器的标签写反的问题。https://github.com/soulteary/docker-flare/issues/48 （感谢 @shp2814560）反馈
- [修正] 编辑器的展示问题和某些情况下读取 `.env` 文件存在的问题，多实例部署 `session` 干扰的问题，错误提示的一些问题。
- [优化] Golang 版本升级至最新的 1.21，相关依赖进行升级调整，GitHub 构建工具修正。

# 0.4.1 - 2023-05-07

- [修正] 感谢 @LightAPIs，修正了 Windows 环境下的 ICON 生成展示问题，及三处界面细节问题。https://github.com/soulteary/flare/releases/tag/0.4.1 https://github.com/soulteary/docker-flare/issues/70 （感谢 @liushuaiiu 的反馈）

# 0.4.0 - 2022-09-17

- [新增] 支持关闭程序的 CSP 策略，允许用户自由选择添加“统计脚本”（如GA）。 https://github.com/soulteary/docker-flare/issues/72 （感谢 @onmpen 反馈问题）
- [新增] 支持了新的架构 ARM32v6，现在程序可以跑在更多的设备上了。
- [新增] 新增了 Linux 和 macOS 操作系统可以直接运行的程序。https://github.com/soulteary/flare/releases/tag/0.4.0 https://github.com/soulteary/docker-flare/issues/35 （感谢 @ToSeeAll 的提问）
- [修正] 重新调整了登陆相关逻辑，确保各种浏览器都能够正确处理跳转逻辑。https://github.com/soulteary/docker-flare/issues/68 （感谢 @hzmabin 反馈问题）
- [修正] 修正全站需要登陆的时候，不展示自定义页脚的问题。https://github.com/soulteary/docker-flare/issues/77 （感谢 @404gods 反馈问题）
- [修正] 修正关闭组件时，不展示日期的问题。https://github.com/soulteary/docker-flare/issues/44 （感谢 @shuax 反馈问题）
- [优化] 拆分和重构程序，进行部分功能调整和问题修正，提升项目质量，为接下来支持新功能做准备。 https://github.com/soulteary/flare/compare/v0.0.17...0.4.0
- [优化] 改进构建方式，将构建迁至 GitHub Action，让构建更透明，大家用的更放心。🎉

最后，感谢 @@JKnuts 在我不在的日子里，为其他同学解答问题♥️

# 0.3.3 - 2022-03-10

- [新增] 支持从 envfile 中设置应用启动选项。
- [修正] 现在支持完全从 env 中设置应用启动配置，实现诸如禁用离线模式，设置用户账号密码等。
- [修正] DockerHub 中镜像版本不是最新内容。https://github.com/soulteary/docker-flare/issues/43  （感谢 @flashliao @shuax @ZhXWei 反馈 ）
- [修正] 禁用界面应用选项后，帮助界面不展示内容。 https://github.com/soulteary/docker-flare/issues/39 （感谢 @harrison-guo-chn 反馈）
- [优化] 禁用 Chrome 默认弹出烦人的翻译对话框。 https://github.com/soulteary/docker-flare/issues/41 （感谢 @shuax 反馈）
- [优化] 新增界面选项，不再强制界面中的英文展示为大写。 https://github.com/soulteary/docker-flare/issues/31 （感谢 @llussy @kscu 反馈）
- [优化] 拆分和重构程序构建脚本，为自定义主题做准备。

# 0.3.2 - 2022-02-23

- [修正] 解决不能正确设置主题的问题，参考 https://github.com/soulteary/docker-flare/issues/26 感谢 @kx 反馈。

# 0.3.1 - 2022-02-23

当前版本新增了四个功能，并进行了大规模的重构，整理和开放了一部分代码。

- [新增] 图标支持使用外链、支持针对未填写图标名称或外链图标地址的项目，根据连接主机名称进行 favicon 获取和展示。 https://github.com/soulteary/docker-flare/issues/14 （感谢 @lmm214 @eallion 反馈）
- [新增] 将首页内容设置为登陆后可见。 https://github.com/soulteary/docker-flare/issues/23 （ 感谢 @OTritium 反馈）
- [新增] 将首页左下角的工具按钮（设置、帮助）在界面中进行隐藏。 https://github.com/soulteary/docker-flare/issues/20 （感谢 @fgprodigal 反馈）
- [新增] 支持将页面展示的链接进行简单的编码处理，避免直接被搜索引擎或机器人记录。
- [优化] 重构和持续迭代项目，整理了 10% 左右的代码到 https://github.com/soulteary/flare

# 0.2.10 - 2022-02-21

当前版本主要修正了编辑器的使用问题，以及更新了帮助页面展示的功能链接。

- [修正] 修正编辑器不能使用的问题。 https://github.com/soulteary/docker-flame/issues/2 （感谢 @Mantyke 反馈）
- [优化] 完善帮助页面的功能链接。

# 0.2.9 - 2022-02-20

当前版本新增了在线编辑功能，以及一个用于聚合内部工具地址的帮助页面，完善了程序启动时和环境变量以及命令行参数的交互，并进行了大量重构工作。

- [新增] 新增在线编辑器功能，解决一些小伙伴提到的“随时随地”编辑的需求。 https://github.com/soulteary/docker-flare/issues/11
- [新增] 考虑到接下来随着程序功能完善，内部工具链接会越来越多，新增一个“帮助页面”，将程序内置链接集中存放和展示。 https://github.com/soulteary/docker-flare/issues/19 https://github.com/soulteary/docker-flare/issues/14#issuecomment-1044430872 (感谢 @lmm214 提醒 )
- [优化] 调整部分程序内部链接，并新增废弃提醒（可通过命令行或环境变量进行关闭），完善程序日志输出实现，重构部分功能，为之后完全开源做准备。
- [优化] 重构程序初始化相关功能，程序目前支持使用命令行参数和环境变量两种方式来进行初始化。 https://github.com/soulteary/docker-flare/issues/5 （感谢 @flashliao 建议）

# 0.2.8 - 2022-02-18

当前版本优化了三个用户反馈的小细节。

- [优化] 添加了搜索返回功能。 https://github.com/soulteary/docker-flare/issues/10 （感谢 @eallion 建议）
- [优化] 简单调整了移动端展示，下一步计划先开放模版和自定义主题。 https://github.com/soulteary/docker-flare/issues/6 （感谢 @flashliao 建议）
- [新增] 为容器镜像添加了 `latest` 标签。 https://github.com/soulteary/docker-flare/issues/8 https://github.com/soulteary/docker-flare/issues/7 （感谢 @viticis @eallion 建议）

## 0.2.7 - 2022-02-17

当前版本优化了一个小细节，以及修正了一个用户反馈的问题。

- [优化] 在非离线模式下，当用户首次使用程序，或程序配置中的 Location 字段为空的时候，将尝试使用 IPIP 的免费接口来自动进行定位，减少不必要的用户输入操作。（感谢 IPIP.net 高春辉大叔的帮忙！）
- [修正] 程序出现设置选项回显不正确的问题，https://github.com/soulteary/docker-flare/issues/7 （感谢 @viticis 小伙伴的反馈）

## 0.2.6 - 2022-02-17

当前版本新增了一个简单的功能，“使用向导”，相比较阅读文章，十几秒钟的交互引导会更直观一些。

- [新增] 用户使用向导，访问 `/guide` 即可使用，如果不需要这个功能，可以在环境变量中配置 `FLARE_GUIDE=0` 来关闭功能。

## 0.2.5 - 2022-02-16

当前版本调整内容比较多，程序配置文件新增了一些配置项，请留意文档或阅读博客文章。

如果你不希望手动调整程序配置文件，可以在备份当前数据目录 `app` 后，将其中的程序配置文件 `app/config.yml` 删除，重启应用，让程序自动生成新版配置文件。

- [新增] 用户登陆，让公网部署的小伙伴不必担心被陌生人随便修改配置的问题。
- [新增] 免登陆模式，让 HomeLab、本地运行 Flare 的小伙伴，保持使用上的简单。
- [新增] 离线模式，让 Flare 不访问任何公网资源。目前 Flare 仅会调用 issue[#4](https://github.com/soulteary/docker-flare/issues/4) 中提到的 `https://wis.qq.com/weather/common` 接口。
- [新增] 允许用户在页脚自定义展示内容，方便设置诸如个人信息、备案号、用户自己的网站统计脚本等。
- [新增] 允许用户对首页的问候语进行自定义设置，支持设置固定问候语，或根据早晨、中午、下午、晚上几个不同时间段展示不同的内容。
- [优化] 完善应用启动日志输出，方便用户反馈问题。
- [优化] 整理和重构部分代码实现，优化应用渲染性能。

## 0.2.4 - 2022-02-11

解决用户 @ember-zhang 提出的特殊链接渲染和跳转问题，支持了服务端跳转[#3](https://github.com/soulteary/docker-flare/issues/3)。

## 0.2.3 - 2022-02-07

针对应用链接的样式进行调整，让“应用”之间的距离稍微大一些。

修正了用户 @SandZhSand 反馈的在 Firefox 下“应用”书签的一个样式兼容性问题。

## 0.2.1 - 2022-02-07

完善构建脚本，支持在常见的 ARM 设备上直接运行。

## 0.2.0 - 2022-01-30

项目数据存储相关功能重构，使用 yaml 配置进行数据持久化，让应用更加轻量。

## 0.1.0 - 2022-01-27

项目初始化。编写简单的文档，完成基础展示功能，使用 SQLite 进行数据持久化，添加链接图标辅助工具。
