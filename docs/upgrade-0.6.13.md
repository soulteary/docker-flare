# 升级到 Flare 0.6.13

应用发布见 [Flare 0.6.13](https://github.com/soulteary/flare/releases/tag/0.6.13)。本仓库更新部署示例，不创建新的部署仓库版本。

## 从 0.6.11 或 0.6.12 升级

1. 停止实例，备份整个 `app`（包括 `.flare-security/`、YAML 和所有隐藏文件）、Compose、外部环境变量及 Secret，并记录旧镜像 digest。
2. 保留原项目名、数据挂载和访问配置，只将镜像更新为 `soulteary/flare:0.6.13` 或 `ghcr.io/soulteary/flare:0.6.13`。
3. 执行 `docker compose pull flare` 和 `docker compose up -d --force-recreate flare`。仅 `restart` 不会替换镜像。
4. 执行 `docker compose exec flare flare --version`，应显示 `0.6.13`；用匿名窗口核对原访问模式，检查登录、设置保存、编辑及重启后状态。

随仓 `app/bookmarks.yml` 的占位链接原先引用不存在的分类 `b`，现改为已声明的 `string-name`。现有用户数据请自行修正实际无效引用，不要用仓库示例覆盖个人书签。

这两版更新 fasthttp、minify 和 kit 依赖，无新增配置或数据迁移。根目录示例继续默认公开且免登录，可访问实例的人能浏览和编辑；登录、私有和 Traefik 示例继续保持原有保护。已有私有部署不得改用根目录配置覆盖 `FLARE_DISABLE_LOGIN=false`、`FLARE_VISIBILITY=PRIVATE`、凭据来源及 Cookie/反代设置。不要重新初始化或删除已有安全状态。

## 更早版本与回退

从 0.6.10 或更早版本升级，先阅读 [0.6.11 来源校验](upgrade-0.6.11.md)、[0.6.9 默认访问行为](upgrade-0.6.9.md)和 [0.6.8 安全状态迁移](upgrade-0.6.8.md)。旧版不一定执行当前逐项隐私控制。

回退时停止新版，使用保留的旧镜像与升级前完整数据、外部配置备份。不要只替换镜像、丢弃安全目录或同时用两个版本写同一目录。四份 Compose 保留完整 `./app:/app` 挂载，避免将安全状态或自动书签备份留在容器可写层。
