# STARTUP

## 如何使用

### 编写配置文件

```json
[
    {
        "enabled": true,
        "command": "echo",
        "args": "hello startup",
        "type": "forking"
    },
    {
        "enabled": true,
        "command": "ping",
        "args": "127.0.0.1 -n 4"
    }
]
```

配置文件模板如上，[ ]数组内部，每一个对象代表一个运行项目。

`enabled` 布尔：`true`、 `false`

`command` 字符串：可执行文件名称或路径

`args` (可选，默认为`空`)字符串：可执行文件参数

`type` (可选，默认为`simple`)字符串：`simple`执行指令，并等待指令完成、 `forking`执行指令，不等待指令完成

### 安装地址

#### windows

1. 将本程序可执行文件与配置文件放置在`windows` 自启动路径`C:\ProgramData\Microsoft\Windows\Start Menu\Programs\StartUp`下，即可开机启动