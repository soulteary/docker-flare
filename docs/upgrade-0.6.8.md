# 升级到 0.6.8 与回滚

应用版本为 `0.6.8`，本部署仓库发行标签为 `v0.6.8`。本次新增持久化管理员账号、可撤销服务端会话、分类／条目隐私保护及导航缓存。三个 YAML 继续保存内容与设置，没有内容数据库迁移。

## 升级前

1. 记录旧镜像标签和 digest。停止 Flare 及其他写入者，完整备份 `app/`，包括隐藏目录、`.flare-cookie-secret` 和已有安全状态；另备份 Compose、工作目录 `.env`、Secret 和代理配置。
2. 保留原 `./app:/app`、命名卷名称和 Compose 项目名；该整目录挂载已覆盖 `.flare-security/`。只挂载三个 YAML 的部署，另加 `./security:/app/.flare-security`，或通过 `FLARE_STATE_DIR` 指向持久目录。
3. 将原部署 `image` 改为 `soulteary/flare:0.6.8` 或 `ghcr.io/soulteary/flare:0.6.8`，保留原有效用户名、密码和访问模式。不要把本仓库的新安装示例直接覆盖到旧部署。
4. 在旧实例停止后，可执行只读检查：

```sh
docker compose run --rm --no-deps flare flare security migration-check --json
```

## 更新和核验

```sh
docker compose pull flare
docker compose up -d --force-recreate flare
docker compose exec flare flare --version
```

有效旧登录部署自动迁移；旧 Cookie 失效一次，重新登录即可。部署管理凭据仍由原 Compose/.env/Secret 管理；网页明确转为本地管理后，遗留 FLARE_PASS 不覆盖新密码。

有效旧免登录且无 private 标记的配置可保留匿名管理。无法区分的新旧默认配置不会自动公开；经确认的旧默认部署可在停机备份后按[源码升级指南](https://github.com/soulteary/flare/blob/0.6.8/docs/login-private-upgrade.md)采用 `security adopt-legacy-defaults`，不要盲目运行。

免登录与 private 内容并存时业务暂停；需开启登录并提供有效凭据，或停机逐项确认哪些 private 标记应移除。程序不自动清除标记，PRIVATE 与禁用登录等矛盾配置也不会降为公开。

用独立无 Cookie 窗口检查导航、搜索、旧跳转、编辑器和导出；确认 private 分类／条目不向匿名用户泄露。再重建容器，核验账号、实例身份和会话持久化，确认后再开放公网。

含 YAML 锚点、别名和合并键的文件继续按有效值读取，但网页编辑／导入覆盖为只读；原始 YAML 导出保持字节，不自动展开或重写。完整子路径、多副本和网络盘并发写不在本阶段支持范围。

## 回滚与恢复

停新版并隔离公网，另存升级后的完整数据；恢复升级前完整工作目录、外部配置、Secret 和挂载，再使用记录的旧镜像。不能仅降镜像并混用新版安全状态。

旧版可解析 YAML 不代表执行 private 保护。0.6.7 及更早版本若需保护私有内容，必须同时启用登录与全站 PRIVATE 并实际验证，或先移除私有内容再公开。恢复旧 Cookie 密钥可能让旧 Cookie 重新有效。

恢复一致的 0.6.8 完整备份后，启动前执行 `docker compose run --rm --no-deps flare flare security revoke-all`，成功后再启动并检查；普通目录复制可能带回历史会话。仅恢复 YAML 不等于完整实例恢复。

详细操作见[应用发布说明](https://github.com/soulteary/flare/releases/tag/0.6.8)、[账号指南](application-account.md)及[备份恢复](backup-restore.md)。
