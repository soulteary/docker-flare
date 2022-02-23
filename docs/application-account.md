# 用户账号相关

flare 默认会启动免登陆模式，方便在 HomeLab 或本地使用的小伙伴。

然而，有一些小伙伴需要在公网使用，本篇文档就来展示如何设置和获取 flare 的用户和密码。

## 设置 Flare 账号和密码

我们可以通过在环境变量中设置 `FLARE_USER` 和 `FLARE_PASS` 来指定 flare 的账号和密码，下面是一个容器编排文件示例：

```yaml
version: '3.6'

services:
  flare:
    image: soulteary/flare:0.3.1
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

更多示例，可以参考[示例文件夹](../example/)。


当你使用 `docker-compose up -d` 启动应用之后，接着使用 `docker-compose ps`，就可以看到包含密码的日志输出啦：

```bash
INFO[2022-02-23T10:34:18+08:00] Flare v0.3.1-85EED66 linux/amd64 BuildDate=2022-02-23T10:22:43+0800 
INFO[2022-02-23T10:34:18+08:00]                                              
INFO[2022-02-23T10:34:18+08:00] 程序服务端口 5005                                  
INFO[2022-02-23T10:34:18+08:00] 页面请求合并 false                                 
INFO[2022-02-23T10:34:18+08:00] 启用离线模式 false                                 
INFO[2022-02-23T10:34:18+08:00] 已禁用登陆模式，用户可直接调整应用设置。                         
INFO[2022-02-23T10:34:18+08:00] 在线编辑模块启用，可以访问 /editor 来获取程序使用帮助。             
INFO[2022-02-23T10:34:18+08:00] 向导模块启用，可以访问 /guide 来获取程序使用帮助。                
INFO[2022-02-23T10:34:18+08:00] 程序已启动完毕 🚀                             
```