# 日志配置与 Docker 日志容量

从 Flare 0.6.4 起，HTTP 访问日志默认只记录错误请求。应用仍将日志输出到标准输出或标准错误，Docker 部署应同时为容器配置日志轮转，避免长期运行时日志持续占用磁盘。

本文中的 `FLARE_ACCESS_LOG` 和 `--access_log` 从 Flare 0.6.4 起可用。可选用 `soulteary/flare:0.6.7`、`ghcr.io/soulteary/flare:0.6.7` 或其他不低于 0.6.4 的已发布版本，并检查 `flare --version`；旧镜像不支持该开关。Docker 日志轮转配置可用于现有镜像。

## 访问日志模式

使用 `FLARE_ACCESS_LOG` 或 `--access_log` 选择模式：

- `errors`：默认值。记录返回错误的请求以及 HTTP 4xx、5xx 响应，正常请求不写访问日志。
- `all`：记录普通请求，同时保留已有的降噪规则：`/ping`、`/favicon.ico`，以及以 `/assets/`、`/redir/` 为前缀的路径，在没有返回错误且状态码小于 400 时跳过；这些路径的错误和 4xx、5xx 响应仍会记录。
- `off`：关闭 HTTP 访问日志。

例如，临时启用普通请求日志：

```sh
FLARE_ACCESS_LOG=all flare
```

也可以使用启动参数：

```sh
flare --access_log=all
```

配置优先级从高到低为：显式指定的 `--access_log`、应用工作目录内 `.env` 的 `FLARE_ACCESS_LOG`、进程环境变量 `FLARE_ACCESS_LOG`、默认值 `errors`。

模式值会忽略大小写和两端空白。显式配置的空值或其他无效值会产生警告，并使用 `errors`。

应用 `.env` 示例：

```dotenv
FLARE_ACCESS_LOG=errors
```

默认 Docker 部署的应用工作目录是 `/app`，对应持久化目录 `./app`，因此应用读取的是 `./app/.env`。Compose 项目目录中的 `.env` 用于 Compose 变量替换，不会自动成为容器的环境变量；在 Compose 中配置环境变量时，应将其写入 `services.flare.environment`。

```yaml
services:
  flare:
    environment:
      FLARE_ACCESS_LOG: errors
```

访问日志模式独立于应用日志等级。启动信息和应用自身的 INFO、WARN、ERROR 日志继续记录；`off` 只关闭 HTTP 访问日志。`FLARE_DEBUG=on` 仍用于启用应用 DEBUG 日志，不会自动将访问日志模式改为 `all`。

## Docker 日志轮转

在现有 `services.flare` 中加入以下配置：

```yaml
services:
  flare:
    logging:
      driver: local
      options:
        max-size: "10m"
        max-file: "3"
```

`local` 驱动在单个日志文件达到 `10m` 后轮转，最多保留 3 个日志文件，并默认压缩轮转日志。实际磁盘占用会受到压缩和存储开销影响。这项配置管理的是容器的全部标准输出和标准错误，因此即使关闭访问日志，也能约束其他应用日志的保留量。

如果现有日志采集工具依赖 `json-file`，可使用相同限额：

```yaml
services:
  flare:
    logging:
      driver: json-file
      options:
        max-size: "10m"
        max-file: "3"
```

两个驱动均可通过 `docker logs` 查看日志。不要手动删除、截断或修改 Docker 管理的日志文件；由 Docker 的日志驱动执行轮转和清理。

## 更新现有部署

先确认正在运行的容器如何持久化 `/app`，并备份对应数据。默认部署使用 `./app:/app`，其中包含书签、设置和其他持久化文件；重建时应保留这个挂载。若使用命名卷，继续使用原来的卷名和 Compose 项目名，不要新建一个空卷替代。

在原部署目录中检查当前挂载：

```sh
docker inspect "$(docker compose ps -q flare)" --format '{{range .Mounts}}{{println .Type .Source "->" .Destination}}{{end}}'
```

确认数据挂载后，向原来的 Compose 文件加入 `logging` 配置。需要改变访问日志模式时，再加入 `FLARE_ACCESS_LOG` 或调整启动参数。使用原来的项目名及其他 Compose 参数重建同一个服务：

```sh
docker compose up -d --force-recreate flare
```

日志驱动、轮转选项和 Compose 环境变量在创建容器时生效。仅运行 `docker compose restart flare` 不会应用这些 Compose 配置变更。

Docker 日志轮转可直接用于现有镜像；访问日志开关需要 Flare 0.6.4 或更新版本。源码合并不代表镜像已经发布；使用固定版本时，先将 `image` 改为 `soulteary/flare:0.6.7` 或 `ghcr.io/soulteary/flare:0.6.7`，再拉取镜像。继续使用 `latest` 时，也应检查实际版本或 `flare --help` 是否包含 `--access_log`。

选择已发布镜像后，更新镜像、重建并核对版本：

```sh
docker compose pull flare
docker compose up -d --force-recreate flare
docker compose exec flare flare --version
```

重建后，检查容器实际使用的日志配置，并查看最近的日志：

```sh
docker inspect "$(docker compose ps -q flare)" --format '{{json .HostConfig.LogConfig}}'
docker logs --tail 50 "$(docker compose ps -q flare)"
```

采用上面的 `local` 配置时，第一条命令的结果应包含 `Type` 为 `local`，以及 `max-size` 为 `10m`、`max-file` 为 `3`。默认 `errors` 模式下，普通成功访问不产生访问日志；应用启动信息仍可见，错误请求仍会记录。

Docker 行为详见官方文档：[local 驱动](https://docs.docker.com/engine/logging/drivers/local/)、[json-file 驱动](https://docs.docker.com/engine/logging/drivers/json-file/)、[Compose 重建容器](https://docs.docker.com/reference/cli/docker/compose/up/)和 [Compose restart](https://docs.docker.com/reference/cli/docker/compose/restart/)。
