# 用户账号相关

flare 默认会启动免登陆模式，方便在 HomeLab 或本地使用的小伙伴。

然而，有一些小伙伴需要在公网使用，本篇文档就来展示如何设置和获取 flare 的用户和密码。

## 设置 Flare 账号和密码

我们可以通过在环境变量中设置 `FLARE_USER` 和 `FLARE_PASS` 来指定 flare 的账号和密码，下面是一个容器编排文件示例：

```yaml
version: '3.6'

services:
  flare:
    image: soulteary/flare
    restart: always
    # 默认无需添加任何参数，如有特殊需求
    # 可阅读文档 https://github.com/soulteary/docker-flare/blob/main/docs/advanced-startup.md
    # 启用账号登陆模式
    command: flare --disable_login=0
    environment:
      # 启用用户登陆模式后，如未设置 FLARE_USER，则默认用户为 flare
      - FLARE_USER=flare
      # 指定自己的密码；如未设置 FLARE_PASS，则生成密码并展示在启动日志中
      - FLARE_PASS=your_password
    ports:
      - 5005:5005
    volumes:
      - ./app:/app
```

更多示例，可以参考[示例文件夹](../example/)。

如果反向代理终止 HTTPS、再通过 HTTP 连接 Flare，请在服务的 `environment` 中添加 `FLARE_COOKIE_SECURE=true`，并通过 HTTPS 访问。普通 HTTP 部署无需开启该选项。完整配置优先级和参数见[自定义启动参数](./advanced-startup.md#https-反向代理的会话-cookiecookie_secure)。

在执行命令前，不妨执行 `docker pull soulteary/flare` 确认使用的应用镜像是新版本。

使用 `docker compose up -d flare` 启动后，可以通过以下命令查看服务日志；未指定 `FLARE_PASS` 时，可从启动日志中获取生成的密码。自行指定密码时，使用设置的密码即可。

```bash
docker compose logs flare
```
