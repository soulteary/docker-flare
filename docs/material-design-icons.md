# Material Design Icons 使用

为了方便小伙伴使用 flare，0.6.6 内置 Material Design Icons 7.4.47 的 7,447 个图标，并保留 19 个已被上游移除的旧图标标识及原 SVG（共 7,466 个名称），你可以让你的每一个书签都拥有风格统一、高质量的矢量图标。

在示例的书签配置文件中，我们可以看到如何为应用配置图标。

```yaml
links:
  - name: "Regexp 101"
    link: "https://regex101.com/"
    icon: "ladybug"
    desc: "在线正则表达式"
  - name: "JSON2Go"
    link: "https://mholt.github.io/json-to-go/"
    icon: "google"
    desc: "快速生成结构体"
```

启动 flare 之后，使用浏览器访问 `/icons/` 可以打开图标列表页面。

在页面中选择你喜欢的图标，鼠标点击之后，默认会将程序可以直接使用的“图标名称”保存在剪贴板，然后粘贴到配置文件中。直接修改 YAML 后需要重启 Flare 才会重新读取；通过在线编辑器保存的修改用于后续页面请求。

（图标名称无需使用 `-` 连字符，并且大小写不敏感）

旧书签中的 `discord`、`desktopmac`、`tabletandroid` 等名称无需重写。请以运行中 `/icons/` 的列表为准，上游网站中新加入的图标不一定已随当前镜像提供。
