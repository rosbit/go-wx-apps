# go-wx-api例程和工具

该repository是[go-wx-api](https://github.com/rosbit/go-wx-api)的例程和工具，包括
 1. samples/wx-echo-server: 该程序可以直接用于配置微信公众号**服务器配置**，并可以对公众号对话框输入做回声应答
 1. samples/wx-server: 一个可以处理菜单的公众号服务，用于微信相关操作和业务分离的场景
 1. tools/wx-menu: 创建/查询/删除微信服务号自定义菜单
 1. tools/parseAesBody: 命令行模式下测试aes加密消息的分解

## 下载、编译方法

 1. 前提：已经安装go 1.11.x及以上、git、make
 2. 进入任一文件夹，执行命令
    ```bash
    $ git clone https://github.com/rosbit/go-wx-apps
    $ cd go-wx-apps
    $ make
    ```
