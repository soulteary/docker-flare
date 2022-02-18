# 用户账号相关

flare 默认会启动免登陆模式，方便在 HomeLab 或本地使用的小伙伴。

然而，有一些小伙伴需要在公网使用，本篇文档就来展示如何设置和获取 flare 的用户和密码。

## 设置 Flare 账号和密码

我们可以通过在环境变量中设置 `FLARE_USER` 和 `FLARE_PASS` 来指定 flare 的账号和密码，下面是一个容器编排文件示例：

```yaml
version: '3.6'

services:
  flare:
    image: soulteary/flare:0.2.8
    restart: always
    # 默认无需添加任何参数，如有特殊需求
    # 可阅读文档 https://github.com/soulteary/docker-flare/blob/main/docs/advanced-startup.md
    # 启用账号登陆模式
    command: flare --nologin=0
    environment:
      # 如需开启用户登陆模式，需要先设置 `nologin` 启动参数为 `0`
      # 如开启 `nologin`，未设置 FLARE_USER，则默认用户为 `flare`
      - FLARE_USER=flare
      # 指定你自己的账号密码，如未设置 `FLARE_USER`，则会默认生成密码并展示在应用启动日志中
      - FLARE_PASS=your_password
    ports:
      - 5005:5005
    volumes:
      - ./app:/app
```

当你使用 `docker-compose up -d` 启动应用之后，接着使用 `docker-compose ps`，就可以看到包含密码的日志输出啦：

```bash
2022/02/18 12:24:50 用户未指定 `FLARE_USER`，使用默认用户名 flare
2022/02/18 12:24:50 用户未指定 `FLARE_PASS`，自动生成应用密码 91a02de0b1e95e5f
2022/02/18 12:24:50 
2022/02/18 12:24:50 Flare v0.2.8-59DA7C1 linux/amd64 BuildDate=2022-02-18T11:17:47+0800
2022/02/18 12:24:50 
2022/02/18 12:24:50 服务端口 5005
2022/02/18 12:24:50 合并页面资源请求 false
2022/02/18 12:24:50 启用应用离线模式 false
2022/02/18 12:24:50 启用免登陆模式 true
2022/02/18 12:24:50 启用教程模块，可以访问 /guide 来获取“使用向导”。
2022/02/18 12:24:50 应用已启动 🚀
```