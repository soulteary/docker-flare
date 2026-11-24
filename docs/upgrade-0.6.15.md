# 升级到 Flare 0.6.15

本仓库的默认、登录、私有和 Traefik 部署示例使用已发布的 [Flare 0.6.15](https://github.com/soulteary/flare/releases/tag/0.6.15)。应用镜像标签不带 `v`。

1. 停止写入并备份完整 `app/` 目录（含 `.flare-security/`）、Compose、`.env` 和外部密钥。
2. 保留原有项目名、卷、挂载路径和访问配置，将服务的 `image` 更新为 `soulteary/flare:0.6.15`。
3. 在原部署目录运行 `docker compose pull flare` 和 `docker compose up -d --force-recreate flare`。
4. 运行 `docker compose exec flare flare --version`，确认版本为 `0.6.15`。检查导航读取、设置保存和重启后的状态，并用独立无 Cookie 窗口核对匿名访问权限。

0.6.15 包含导航文件热加载和无效文件恢复、TLS 反向代理下的同源写入修复以及 30 天登录会话。完整变更和行为说明以[应用发布记录](https://github.com/soulteary/flare/releases/tag/0.6.15)和[对应版本文档](https://github.com/soulteary/flare/blob/0.6.15/docs/README.md)为准。

已于 2026-11-24 核对 Docker Hub 的公开多架构索引：

```text
soulteary/flare:0.6.15@sha256:ffe1f052726b39df7540ff80583f4c75755becb07e8356b212b5348ee28761da
```

索引包含 Linux amd64、arm64、arm/v6 和 arm/v7。示例继续使用明确的版本标签；需要不可变部署时可以使用上述完整引用。此核对覆盖公开镜像元数据与部署配置，不代表已完成所有浏览器的真机测试。

需要回滚时，停止实例并恢复升级前的完整数据与配置备份，再切换镜像；不要只换旧镜像而继续使用新版本已修改的数据。
