# 自定义启动参数

当前应用默认启动参数，包含以下四项目：

```bash
Usage of flare:
  mr
    	是否使用请求最小化模式
  nologin
    	启用无需登陆模式 (default true)
  offline
    	启用完全离线模式
  port int
    	指定监听端口 (default 5005)
```

## 合并资源请求：`mr`

这个参数的作用是“合并页面请求资源，减少页面资源请求总数量”，在一些书签数据量极大的场景下，可以大幅提升页面渲染性能。

小伙伴可以根据自己的情况按需开启，此参数默认关闭。

启用方法，在 flare 启动参数后添加 `--mr=1`


## 免登陆模式：`nologin`

当这个参数关闭之后，用户需要在登陆之后，才能够对设置页面的内容进行调整，适合 Flare 在公网环境中使用的小伙伴。

此参数默认开启，方便在 HomeLab 或本地使用的小伙伴，减少不必要的操作。

禁用方法，在 flare 启动参数后添加 `--nologin=0`

## 离线模式：`offline`

当这个参数开启之后，flare 所有依赖公网的功能将被停用。

目前 flare 仅会调用 issue[#4](https://github.com/soulteary/docker-flare/issues/4) 中提到的 `https://wis.qq.com/weather/common` 接口，获取天气数据，开启此参数后，天气将无法获取。

此参数默认关闭，开启方法，在 flare 启动参数后添加：`--offline=1`

## 指定端口：`port`

这个参数用于自定义服务端口，使用 `docker-flare` 的小伙伴可以忽略。
