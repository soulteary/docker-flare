# 备份与恢复

完整备份应在 Flare 和其他写入者停止时进行。保留原有挂载、卷名及 Compose 项目名。

| 类型 | 包含 | 用途与边界 |
| --- | --- | --- |
| 整个应用数据目录 | config/apps/bookmarks、应用 .env、Cookie 密钥、备份和恢复目录 | 完整恢复数据；外部 Compose、代理、环境和密钥另备份 |
| 浏览器 HTML 导出 | 名称、链接、分类 | 跨浏览器迁移，不含账号及系统设置 |
| Flare YAML 导出 | 一个导航集合 | 导入是合并，不等于整站覆盖恢复 |
| `.flare-bookmark-backups/` | 导入前数据 | 不覆盖所有日常编辑和 config.yml |
| `.flare-editor-recovery/save-*` | 开发版编辑器保存前的原始两文件 | 仅按明确恢复路径恢复，不是新草稿 |
| JSON 编辑草稿 | 开发版未保存工作数据 | 载入后检查并显式保存，不是浏览器导入文件 |

默认 `./app:/app` 挂载已包含上述目录。Cookie 密钥也要保留，否则已有会话失效。
用包含隐藏文件的目录备份方式，别只复制三个 YAML。备份可能含密码和私有地址，应
妥善控制访问并定期在隔离部署中演练恢复。

```sh
docker compose stop flare
tar -czf "flare-data-$(date +%Y%m%d-%H%M%S).tar.gz" app
docker compose start flare
```

恢复前保存现有目录，停止所有写入者，恢复原属主和权限，再启动并验证导航内容。
命名卷应按原卷名备份恢复，不要另建项目后误认为旧数据消失。
详细的开发版编辑恢复步骤见[恢复指南](https://github.com/soulteary/flare/blob/main/docs/editor-recovery.md)。
