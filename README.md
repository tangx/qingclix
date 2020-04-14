# qingclix
青云常用操作命令行

## Prepare

1. 准备青云授权文件, 路径为 `~/.qingcloud/config.yaml` 。
  + 使用该路径，是为了保证与青云本身的命令行工具对齐，而非额外配置。
  + 配置路径在 `global.client.go` 中硬编码。
2. 准备 `qingclix` 使用的**服务器预设参数**文件 `~/.qingclix/config.json` 。
  + 复制 `docs/config.json` 到 `~/.qingclix/config.json`。
  + 配置路径在 `global.client.go` 中硬编码。

## Usage 

```bash
$ qingclix help  

青云控制台的操作常用操作复杂，
  例如，新购机器、更换操作系统 等
  实现目标根据预设参数或配置，快速实现日常操作

Usage:
  qingclix [flags]
  qingclix [command]

Available Commands:
  buy         根据预设信息购买机器
  help        Help about any command

Flags:
  -c, --count int       设置购买数量 (default 1)
  -h, --help            help for qingclix
      --skip_contract   强制跳过合约购买过程。 true: 跳过
  -v, --verbose int     logrus 日志等级。 0: Panic, 4: Info, 6: Trace.  (default 4)

Use "qingclix [command] --help" for more information about a command.
```

## Todo

**预设值购买**
+ [x] 预设值服务器购买
+ [x] 预设值硬盘购买与绑定
+ [x] 预设值服务器、硬盘合约购买与绑定
+ [x] 支持强制跳过合约购买
+ [x] 支持批量购买

**自定义购买**
+ [ ] 获取与保存网络、用户密钥等信息
+ [ ] 处理服务器与硬盘的关联关系
+ [ ] 保存选择配置到 预设值 

**克隆已存在服务器**
+ [ ] 保存选择配置到 预设值 


## 删除
+ [ ] 删除时必须确认(ex 输入要删除实例的名字或 ID)


## 使用到的库

+ `struct` 转 `url.Values`: `github.com/tangx/go-querystring/query`
