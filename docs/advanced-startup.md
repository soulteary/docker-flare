# 自定义启动参数

程序目前的启动参数，可以从环境变量或从命令行参数中解析获得。

默认推荐使用一种方案进行参数配置，如果你有特殊需求，需要混合使用。

请注意，命令行参数拥有更高的优先级，会覆盖环境变量中的参数。


## 环境变量

```bash
修改程序监听端口          "FLARE_PORT"
配置登陆模式下的账号      "FLARE_USER"
配置登陆模式下的密码      "FLARE_PASS"
启用或禁用程序向导        "FLARE_GUIDE"
启用程序废弃功能提示      "FLARE_DEPRECATED_NOTICE"
启用服务端请求合并功能     "FLARE_MINI_REQUEST"
禁用登陆模式             "FLARE_DISABLE_LOGIN"
启用离线模式             "FLARE_OFFLINE"
启用在线编辑器功能        "FLARE_EDITOR"
首页是否需要登陆可见      "FLARE_VISIBILITY"
```

环境变量使用示例： `FLARE_OFFLINE=0` 或 `FLARE_OFFLINE=false`

完全的环境变量列表，可以[参考这里](https://github.com/soulteary/flare/blob/main/model/cmd.go)。


## 命令行参数

```
docker run --rm -it -p 5005:5005 soulteary/flare:0.4.0 flare -h
找不到配置文件config，创建默认配置。

Flare v0.4.0-9D209E38A3BC7637E14E0762EFC7E4BC8F39D03E linux/amd64 BuildDate=2022-09-17T12:42:33Z

支持命令:
  -p, --port int            指定监听端口 (default 5005)
  -g, --guide               启用应用向导 (default true)
  -s, --visibility string   调整网站整体可见性 (default "DEFAULT")
  -m, --mini_request        使用请求最小化模式
  -o, --offline             启用离线模式
  -l, --disable_login       禁用账号登陆 (default true)
  -n, --enable_notice       启用废弃日志警告 (default true)
  -e, --enable_editor       启用编辑器 (default true)
  -c, --disable_csp         禁用CSP
  -v, --version             显示应用版本号
  -h, --help                显示帮助
```

## 功能说明

### 合并资源请求

- 环境变量：`FLARE_MINI_REQUEST`
- 命令行：`mini_request` / `m`

这个参数的作用是“合并页面请求资源，减少页面资源请求总数量”，在一些书签数据量极大的场景下，可以大幅提升页面渲染性能。

小伙伴可以根据自己的情况按需开启，此参数默认关闭。

启用方法，在 flare 启动参数后添加 `-m=1` 或者设置环境变量 `FLARE_MINI_REQUEST=1`
### 免登陆模式：`nologin`

当这个参数关闭之后，用户需要在登陆之后，才能够对设置页面的内容进行调整，适合 Flare 在公网环境中使用的小伙伴。

此参数默认开启，方便在 HomeLab 或本地使用的小伙伴，减少不必要的操作。

禁用方法，在 flare 启动参数后添加 `--nologin=0`

### 离线模式：`offline`

当这个参数开启之后，flare 所有依赖公网的功能将被停用。

目前 flare 仅会调用 issue[#4](https://github.com/soulteary/docker-flare/issues/4) 中提到的 `https://wis.qq.com/weather/common` 接口获取天气数据，和 IPIP 的地理位置接口获取地理信息，开启此参数后，天气将无法获取。

此参数默认关闭，开启方法，在 flare 启动参数后添加：`--offline=1`

### 指定端口：`port`

这个参数用于自定义服务端口，使用 `docker-flare` 的小伙伴可以忽略。
