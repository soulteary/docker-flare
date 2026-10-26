# 配置参考与版本边界

稳定版为 **0.6.7**，包含新搜索、列表／表单／批量编辑器和布局设置。稳定部署请阅读对应
版本文档，核对运行中的版本和镜像 digest；后续源码功能合并不代表新的镜像已经发布。

配置优先级：命令行 > 应用工作目录 `.env` > 容器环境变量 > 默认值。Compose 项目
`.env` 仅用于插值，不会自动传入容器；使用服务的 `environment` 或 `env_file`。
镜像、环境变量、command 和日志驱动变化都需要重建容器，restart 不会替换镜像。

| 环境变量 | CLI | 默认值／合法值 |
| --- | --- | --- |
| `FLARE_PORT` | `--port` | `5005`；整数 1–65535 |
| `FLARE_DISABLE_LOGIN` | `--disable_login` | `true`；设为 false 才启用登录 |
| `FLARE_VISIBILITY` | `--visibility` | `DEFAULT` 或 `PRIVATE`；不可填 0/1 |
| `FLARE_USER` / `FLARE_PASS` | — | 默认用户名 flare；密码缺省时生成并输出启动日志 |
| `FLARE_EDITOR` | `--enable_editor` | true |
| `FLARE_GUIDE` | `--guide` | true |
| `FLARE_OFFLINE` | `--offline` | false；停用服务端天气和区域查询 |
| `FLARE_MINI_REQUEST` | `--mini_request` | false；内联 SVG，减少请求但增加 HTML 体积 |
| `FLARE_DEBUG` | — | 进程环境设为 on 开启应用 DEBUG 日志；不切换开发资源模式 |
| `FLARE_ACCESS_LOG` | `--access_log` | errors；可选 all/off；0.6.4 起 |
| `FLARE_COOKIE_SECURE` | `--cookie_secure` | false；HTTPS 反代终止 TLS 时设 true |
| `FLARE_COOKIE_SECRET` | `--cookie_secret` | 自动初始化安全密钥；自设至少 32 字节 |
| `FLARE_COOKIE_NAME` | `--cookie_name` | flare |
| `FLARE_DEPRECATED_NOTICE` | `--enable_notice` | true |
| `FLARE_DISABLE_CSP` | `--disable_csp` | false；不建议关闭 |

布尔值建议写字符串 `"true"`/`"false"`。`FLARE_DEBUG=on` 在正式版也可开启应用 DEBUG 日志，但不会开启开发资源模式。
启动参数修改后重启进程；外部直接修改 YAML 后也需重启，界面保存则用于后续请求。
完整字段和版本行为见[0.6.7 配置参考](https://github.com/soulteary/flare/blob/0.6.7/docs/configuration.md)。

| 目标 | DISABLE_LOGIN | VISIBILITY |
| --- | --- | --- |
| 信任内网、无需登录 | true | DEFAULT |
| 首页公开、编辑需登录 | false | DEFAULT |
| 首页和编辑均需登录 | false | PRIVATE |

隐藏设置按钮只改变外观；链接编码不是加密；分类 Tab 不是分组权限。外观 Footer 是
管理员配置的 HTML，应只由受信任的人编辑。确认访问控制时请使用未登录的新浏览器。

## 0.6.7 分类内部多列布局

在「设置 → 外观」选择分类内部多列布局，对应 `Homepage.BookmarkLayout: sections`。
默认值（空或 `groups`）仍按分类分列；`sections` 按分类纵向排列，并在每个分类内显示多列书签。
列数根据可用宽度调整，手机单列／紧凑模式继续独立生效。
完整设置见[布局说明](https://github.com/soulteary/flare/blob/0.6.7/docs/desktop-columns-and-scale.md)。

## 0.6.6 搜索与天气迁移

搜索支持混合、本地优先和选择器模式，保留旧配置的本地搜索意图，不会自动启用外部引擎。
`!keyword` 快捷语法已移除，旧 `EnableKeywords` 和引擎 `Keyword` 字段不再控制搜索；
通过目标控件选择引擎。完整行为见[搜索说明](https://github.com/soulteary/flare/blob/0.6.7/docs/search-hybrid.md)。

搜索设置实际变更时，先备份再原子替换配置，保留未知 YAML 字段。版本降级、并发修改或
备份异常会阻止保存。升级或重启后请刷新已打开的设置页面，再提交修改。

天气设置需搜索并选择具体省／市／区县，保存时重新校验候选。旧地区仅在唯一匹配时自动迁移；
无法匹配或存在歧义时保留原配置，请手动选择。空地区不再自动按 IP 定位；设置中的定位按钮
仅建议服务器出口 IP 对应的位置，受代理／VPN 影响，不代表访客位置，仍需确认并保存。
详见[天气说明](https://github.com/soulteary/flare/blob/0.6.7/docs/weather.md)。
