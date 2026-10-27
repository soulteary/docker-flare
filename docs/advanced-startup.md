# 自定义启动参数

程序目前的启动参数，可以从环境变量或从命令行参数中解析获得。

默认推荐使用一种方案进行参数配置，如果你有特殊需求，需要混合使用。

配置优先级为：命令行参数 > 应用工作目录中的 `.env` > 环境变量。


## 环境变量

```bash
修改程序监听端口          "FLARE_PORT"
配置登陆模式下的账号      "FLARE_USER"
配置登陆模式下的密码      "FLARE_PASS"
从文件读取密码           "FLARE_PASS_FILE"
持久安全状态目录         "FLARE_STATE_DIR"
规范外部 origin          "FLARE_EXTERNAL_URL"
允许的 Host 列表         "FLARE_ALLOWED_HOSTS"
受信代理 IP/CIDR         "FLARE_TRUSTED_PROXIES"
启用或禁用程序向导        "FLARE_GUIDE"
启用程序废弃功能提示      "FLARE_DEPRECATED_NOTICE"
启用服务端请求合并功能     "FLARE_MINI_REQUEST"
禁用登陆模式             "FLARE_DISABLE_LOGIN"
启用离线模式             "FLARE_OFFLINE"
启用在线编辑器功能        "FLARE_EDITOR"
首页是否需要登陆可见      "FLARE_VISIBILITY"
强制使用仅限 HTTPS 的会话 Cookie "FLARE_COOKIE_SECURE"
HTTP 访问日志模式         "FLARE_ACCESS_LOG"
启用应用 DEBUG 日志       "FLARE_DEBUG"
```

环境变量使用示例： `FLARE_OFFLINE=0` 或 `FLARE_OFFLINE=false`

完整的环境变量列表，可以[参考 0.6.9 的配置定义](https://github.com/soulteary/flare/blob/0.6.10/config/model/cmd.go)。


## 命令行参数

使用镜像内的帮助命令查看当前版本支持的完整参数：

```bash
docker run --rm soulteary/flare:0.6.10 flare --help
```

例如，使用离线模式启动 0.6.9：

```bash
docker run --rm -it -p 5005:5005 soulteary/flare:0.6.10 flare --offline=1
```

## 功能说明

### 合并资源请求

- 环境变量：`FLARE_MINI_REQUEST`
- 命令行：`mini_request` / `m`

这个参数的作用是“合并页面请求资源，减少页面资源请求总数量”，在一些书签数据量极大的场景下，可以大幅提升页面渲染性能。

小伙伴可以根据自己的情况按需开启，此参数默认关闭。

启用方法，在 flare 启动参数后添加 `-m=1` 或者设置环境变量 `FLARE_MINI_REQUEST=1`

### 免登陆模式：`disable_login`

当这个参数关闭之后，用户需要在登陆之后，才能够对设置页面的内容进行调整，适合 Flare 在公网环境中使用的小伙伴。

0.6.9 默认开启免登录，内部安全状态自动准备，无需初始化管理员。需要登录保护时设置 `FLARE_DISABLE_LOGIN=false` 并提供凭据，见[账号指南](application-account.md)。已记录的私有默认继续保留；已有部署应保留原访问配置。

明确需要受信内网免登录时，同时设置 `FLARE_DISABLE_LOGIN=true`、`FLARE_VISIBILITY=DEFAULT`，且不得存在分类或条目的 `private: true`；该模式允许匿名管理内容。不要把它当成公网安全配置。

### HTTPS 反向代理的会话 Cookie：`cookie_secure`

当反向代理对外提供 HTTPS、再通过 HTTP 连接 Flare 时，请设置环境变量 `FLARE_COOKIE_SECURE=true`，或添加启动参数 `--cookie_secure=true`，并始终通过 HTTPS 访问。

普通 HTTP 部署可不设置此项，或设为 `false`。Flare 直接接收 TLS 请求时会自动使用 `Secure` Cookie，即使此项设为 `false` 也不会关闭该保护。该选项从 0.6.2 起可用，登录和退出使用相同的 Cookie 策略。

Docker Compose 的项目 `.env` 不会自动把变量传入容器；请在 `flare` 服务的 `environment` 中配置，例如：

```yaml
environment:
  - FLARE_COOKIE_SECURE=true
```

### HTTP 访问日志：`access_log`

从 0.6.4 起，使用环境变量 `FLARE_ACCESS_LOG` 或启动参数 `--access_log` 选择访问日志模式：

- `errors`：默认值，只记录返回错误或 HTTP 4xx、5xx 的请求。
- `all`：记录普通请求；成功的健康检查、静态资源和跳转请求仍按原有规则省略。
- `off`：关闭 HTTP 访问日志。

例如，添加 `--access_log=all` 或设置 `FLARE_ACCESS_LOG=all` 可恢复普通访问记录。启动信息和应用自身日志不受此开关影响；`FLARE_DEBUG=on` 独立开启应用 DEBUG 日志，不会切换开发资源模式，也不会自动开启全部访问日志；需要普通访问记录时使用 `FLARE_ACCESS_LOG=all`。

配置优先级、Compose 环境变量、具体过滤规则和 Docker 日志轮转见[日志配置与 Docker 日志容量](./logging.md)。修改 Compose 配置后需要重建容器。

### 日期、时间与时区

0.6.3 起，实时时钟、秒数显示和显示时区在 **设置 → 外观 → 模块设置** 中配置，或通过 `config.yml` 的 `Clock` 配置块调整。未配置时保持跟随服务器时区、包含秒数的静态显示。

容器的 `TZ` 只决定进程本地时区。实时模式使用访问设备的系统时间，再按服务器、访问设备或指定时区换算，不会修改主机时间，也不提供校时。完整字段和示例见[日期、时间与时区说明](https://github.com/soulteary/flare/blob/0.6.10/docs/clock.md)。

### 离线模式：`offline`

当这个参数开启之后，Flare 停用服务端天气和区域查询。它不阻止浏览器加载用户配置的远程图标，也不阻止打开外部书签；需要完全离线时应使用内置图标并自行约束浏览器网络。

0.6.6 的天气设置支持地区目录搜索和用户主动触发的服务器出口 IP 定位建议；离线模式停用这些服务端查询，不再自动为未配置地区的用户选择位置。

此参数默认关闭，开启方法，在 flare 启动参数后添加：`--offline=1`

### 指定端口：`port`

这个参数用于自定义服务端口，使用 `docker-flare` 的小伙伴可以忽略。


## 私有访问与完整参考

首页需要登录时，必须同时设置 `FLARE_DISABLE_LOGIN=false` 与
`FLARE_VISIBILITY=PRIVATE`。`0` 不是可见性的合法枚举，隐藏设置按钮也不是权限控制。
环境变量、CLI、默认值、生效方式以及版本适用范围见[配置参考](configuration.md)。

## 0.6.8 安全状态与凭据

安全状态默认保存在 `/app/.flare-security/`，必须持久化。`FLARE_PASS_FILE` 与非空 `FLARE_PASS` 不能并用；密码文件要在容器内可读，最多 514 字节输入，只移除一次末尾换行。账号的本地／部署管理、来源切换与维护命令见[账号指南](application-account.md)。外部 origin 和受信代理的配置见[反向代理](reverse-proxy.md)。
