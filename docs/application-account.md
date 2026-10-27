# 管理员账号与私有访问（0.6.8）

0.6.8 使用持久化单管理员账号和可撤销服务端会话。新实例默认需要管理员且全站私有，不再默认免登录，也不会把新管理员密码写入启动日志。有效旧部署的凭据和访问模式按[升级指南](upgrade-0.6.8.md)迁移。

## 新实例：本地管理密码

根目录和 Traefik 示例使用这种方式，并显式设置启用登录和全站 PRIVATE，避免已有示例 YAML 被按旧免登录访问模式识别。服务停止时，在对应 Compose 目录执行：

```sh
docker compose run --rm --no-deps flare flare security init --username flare
docker compose up -d flare
```

终端隐藏输入密码。安全状态保存在挂载的 `/app/.flare-security/`；用户名和密码使用刚设置的值。启动日志只用于诊断。不要对已有初始化状态重新执行 `init`，也不要删除安全目录来恢复密码。

## 新实例：部署管理密码

[login 示例](../example/login/docker-compose.yml)公开首页、管理需登录；[private 示例](../example/private/docker-compose.yml)要求全站登录。进入示例目录，设置自己的用户名和强密码后启动：

```sh
export FLARE_USER=flare
read -r -s -p 'Flare password: ' FLARE_PASS; export FLARE_PASS; printf '\n'
docker compose up -d flare
```

上面的隐藏输入命令用于 Bash；也可通过受保护的 Compose `.env` 或 Secret 提供值。示例用 `${FLARE_PASS:?...}` 阻止空密码启动。`FLARE_PASS` 是应用最终有效的启动输入，需检查工作目录 `.env` 的更高优先级；不要将真实密码提交 Git。

部署模式保留 Compose/.env/Secret 对凭据的管理。有效凭据变化并重建服务后会撤销旧会话；同值重启不反复改密。希望通过网页管理密码时，在账号页验证原密码并明确切为应用管理；切换后遗留的 `FLARE_PASS` 不会覆盖新密码。切回部署管理需停机使用 `security set-credentials --mode deployment`，详见[源码账号指南](https://github.com/soulteary/flare/blob/0.6.8/docs/accounts.md)。

## 访问模式和隐私

- 新实例默认 `FLARE_DISABLE_LOGIN=false`、`FLARE_VISIBILITY=PRIVATE`，未登录不能访问导航或管理。
- 显式 `false + DEFAULT` 公开首页；分类／条目的 `private: true` 对匿名访问隐藏，搜索、跳转、编辑和导出也检查权限。
- 显式 `true + DEFAULT` 是受信内网免登录模式，允许匿名管理，不能与私有条目并存。

请用独立无 Cookie 窗口验证，不能根据管理员可见性推断匿名访问。0.6.7 及更早版本不执行逐项隐私保护；回滚见[升级指南](upgrade-0.6.8.md)。

## 会话、改密和恢复

普通重启保留未过期会话；账号页可改密、验证近期操作、退出当前或全部会话。编辑或导入提示重新验证时，保留草稿，在账号页完成验证后再重试，不要关闭来源或 CSRF 检查。

遗失密码时先停止实例、保留完整备份，再按凭据来源恢复：部署模式先修正原 Compose/.env/Secret；明确接管才使用 `--take-local-control`。本地模式可执行：

```sh
docker compose stop flare
docker compose run --rm --no-deps flare flare security reset-password
# 命令成功后再启动
docker compose up -d flare
```

维护命令支持 `--password-file` 从容器内文件读取，密码不放在明文命令参数中。恢复完整实例后，启动前执行 `flare security revoke-all`；旧备份可能包含历史会话。

TLS 在代理终止时设置 `FLARE_COOKIE_SECURE=true` 并始终使用 HTTPS。完整 origin、Host 和代理要求见[反向代理](reverse-proxy.md)。书签导出不包含完整账号、安全状态和外部部署配置。
