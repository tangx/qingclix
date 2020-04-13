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

Usage:
  qingclix [flags]
  qingclix [command]

Available Commands:
  buy         根据预设信息购买机器
  help        Help about any command

Flags:
  -h, --help   help for qingclix

Use "qingclix [command] --help" for more information about a command.
```

## Todo

**预设值购买**
+ [x] 预设值服务器购买
+ [x] 预设值硬盘购买与绑定
+ [x] 预设值服务器、硬盘合约购买与绑定

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
