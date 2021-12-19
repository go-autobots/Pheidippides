# Pheidippides
IOT Queue Broker, 以马拉松缔造者命名。

## use
```shell
git clone project
kratos run
```
> 需安装 kratos 官网看下

或者：
到 main 文件自己跑, 服务就起来了

## demo
[mock client demo](internal/mockclient/client.go)

## ps
简单起见先不适配第三方 logger 控制台(标准)输出

## todo
- [X] config proto
- [X] api proto 
    - [X] kratos 脚手架 生成 pb.go 及 service
- [X] config.yaml
- [X] queue interface
- [X] redis implement
  - [X] redis docker compose file 
- [X] wire(看是否需要，可以不用其实)
    - [X] fix some problem
- [X] 启动
- [ ] health check
