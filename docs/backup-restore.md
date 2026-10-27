# 备份与恢复

完整备份应在 Flare 和其他写入者停止时进行。保留原有挂载、卷名及 Compose 项目名。

| 类型 | 包含 | 用途与边界 |
| --- | --- | --- |
| 整个应用数据目录 | config/apps/bookmarks、应用 .env、`.flare-security/`、旧 Cookie 密钥、备份和恢复目录 | 完整恢复数据；外部 Compose、代理、环境和密钥另备份 |
| `.flare-security/` 或自定义 FLARE_STATE_DIR | 实例身份、密码哈希、会话、迁移及事务状态 | 停机完整一致恢复；启动前撤销会话，不能只还原某个文件 |
| 浏览器 HTML 导出 | 名称、链接、分类 | 跨浏览器迁移，不含账号及系统设置 |
| Flare YAML 导出 | 一个导航集合 | 导入是合并，不等于整站覆盖恢复 |
| `.flare-bookmark-backups/` | 导入前数据 | 不覆盖所有日常编辑和 config.yml |
| `.flare-editor-recovery/save-*` | 0.6.5 编辑器保存前的原始两文件 | 仅按明确恢复路径恢复，不是新草稿 |
| `.flare-search-backups/` | 0.6.6 搜索设置变更前的完整配置及 manifest | 首次升级备份永久保留，latest 仅保留上一次；不替代整站备份 |
| JSON 编辑草稿 | 0.6.5 未保存编辑数据 | 载入后检查并显式保存，不是浏览器导入文件 |

默认 `./app:/app` 挂载已包含上述目录；仅挂载三个 YAML 的部署需额外挂载安全目录。0.6.8 不再只依赖旧 Cookie 签名密钥，必须保留完整安全状态。
用包含隐藏文件的目录备份方式，别只复制三个 YAML。备份可能含密码和私有地址，应
妥善控制访问并定期在隔离部署中演练恢复。

```sh
docker compose stop flare
tar -czf "flare-data-$(date +%Y%m%d-%H%M%S).tar.gz" app
docker compose start flare
```

恢复前保存现有目录，停止所有写入者，恢复原属主和权限，启动前执行 `docker compose run --rm --no-deps flare flare security revoke-all`，成功后再启动并验证导航内容和匿名访问。普通重启不要求撤销会话，完整目录恢复则可能带回历史会话。
命名卷应按原卷名备份恢复，不要另建项目后误认为旧数据消失。
详细的编辑恢复步骤见[恢复指南](https://github.com/soulteary/flare/blob/0.6.10/docs/editor-recovery.md)。

0.6.6 的搜索备份包含 `first-upgrade-to-<target-version>.bak`、`latest.bak` 和
`manifest.json`，可能包含私密配置，请勿公开。读取页面或未改变搜索设置不会生成备份。
回滚时停止新版本及其他写入者，恢复升级前的完整数据目录和外部配置，再使用旧镜像启动；
不要依赖旧程序理解新版本已经写入的设置。详见[搜索升级保护](https://github.com/soulteary/flare/blob/0.6.10/docs/search-hybrid.md)。

## 0.6.8 回滚

停新版并隔离公网，另存当前数据，再恢复升级前的完整数据目录和外部 Compose/.env/Secret，换回记录的旧镜像；不要只降镜像并继续混用新安全状态。0.6.7 及更早版本不执行逐项 private 保护，需要同时开启登录与全站 PRIVATE 并验证，或先移除私有内容再公开。完整步骤见[升级与回滚](upgrade-0.6.8.md)。
