# Material Design Icons 使用

为了方便小伙伴使用 flare，程序内置了来自 [materialdesignicons.com](https://materialdesignicons.com/) 中所有的 Material Design Icons，你可以让你的每一个书签都拥有风格统一、高质量的矢量图标。

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

在页面中选择你喜欢的图标，鼠标点击之后，默认会将程序可以直接使用的“图标名称”保存在剪贴板，然后粘贴到配置文件中，flare 就会自动更新书签的图标啦。

（图标名称无需使用 `-` 连字符，并且大小写不敏感）
