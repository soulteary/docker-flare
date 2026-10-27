# 管理员账号与私有访问（0.6.9）

0.6.9 恢复默认免登录与公开导航，直接 `docker compose up -d flare` 即可使用，内部安全状态自动准备，不需要管理员。能访问实例的人均可修改内容。需要登录时才配置账户；持久化单管理员、可撤销会话及私有资源保护继续保留，也不会将密码写入日志。已有部署见[升级指南](upgrade-0.6.9.md)。

## 新实例：本地管理密码

这是一种可选的账户设置方式。Traefik 示例显式启用登录与全站 PRIVATE，首次启动前在其 Compose 目录执行下面命令。根目录默认示例不需要这些步骤；若要改为登录模式，先在 Compose environment 中设置 `FLARE_DISABLE_LOGIN: "false"`，全站私有时再设置 `FLARE_VISIBILITY: "PRIVATE"`。

```sh
docker compose run --rm --no-deps flare flare security init --username flare
docker compose up -d flare
```

终端隐藏输入两次密码；新密码要求 15–128 个 Unicode 字符且不能是弱密码。安全状态保存在挂载的 `/app/.flare-security/`；用户名和密码使用刚设置的值。启动日志只用于诊断。不要对已有初始化状态重新执行 `init`，也不要删除安全目录来恢复密码。

如果实例已在默认免登录模式启动过，内部状态已存在，应将上述 `security init --username flare` 替换为 `security set-credentials --mode local --username flare`，再以启用登录的配置启动。不要删除状态目录重新初始化。

## 新实例：通过启动配置设置密码

[login 示例](../example/login/docker-compose.yml)公开首页、管理需登录；[private 示例](../example/private/docker-compose.yml)要求全站登录。进入示例目录，设置自己的用户名和强密码后启动：

```sh
export FLARE_USER=flare
read -r -s -p 'Flare password: ' FLARE_PASS; export FLARE_PASS; printf '\n'
docker compose up -d flare
```

上面的隐藏输入命令用于 Bash；也可通过受保护的 Compose `.env` 或 Secret 提供值。示例用 `${FLARE_PASS:?...}` 阻止空密码启动。`FLARE_PASS` 是应用最终有效的启动输入，需检查工作目录 `.env` 的更高优先级；不要将真实密码提交 Git。

凭据模式取决于已有状态和安装来源：新安装默认本地管理，迁移的旧凭据默认部署管理。需要持续由部署管理时，在首次初始化前显式设置 `FLARE_CREDENTIAL_MODE=deployment`；已初始化实例使用停机命令 `security set-credentials --mode deployment` 切换，不能只改环境变量。部署模式保留 Compose/.env/Secret 对凭据的管理。有效凭据变化并重建服务后会撤销旧会话；同值重启不反复改密。希望通过网页管理密码时，在账号页验证原密码并明确切为应用管理；切换后遗留的 `FLARE_PASS` 不会覆盖新密码。切回部署管理需停机使用 `security set-credentials --mode deployment`，详见[源码账号指南](https://github.com/soulteary/flare/blob/0.6.9/docs/accounts.md)。

## 访问模式和隐私

- 默认 `FLARE_DISABLE_LOGIN=true`、`FLARE_VISIBILITY=DEFAULT`，无需账户即可访问和修改，且自动准备内部状态。已记录的私有默认仍保留。
- 显式 `false + DEFAULT` 公开首页；分类／条目的 `private: true` 对匿名访问隐藏，搜索、跳转、编辑和导出也检查权限。
- 显式 `true + DEFAULT` 是受信内网免登录模式，允许匿名管理，不能与私有条目并存。

请用独立无 Cookie 窗口验证，不能根据管理员可见性推断匿名访问。0.6.7 及更早版本不执行逐项隐私保护；回滚见[升级指南](upgrade-0.6.9.md)。

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
