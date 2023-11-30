# gocron - 定时任务管理系统

# 项目简介

本项目基于 [ouqiang/gocron](https://github.com/ouqiang/gocron)
更新迭代后的[peng49/gocron](https://github.com/peng49/gocron)（ 新增的进程管理、LDAP用户认证、项目管理功能, 重构了前端页面）

前端使用 Element Plus + Vue3，后端使用golang的web框架macaron.v1
    
### 支持平台
> Windows、Linux、Mac OS

### 环境要求
>  MySQL



## 安装运营

###  配置文件

配置: /app/conf/app.ini

日志: /app/log/cron.log

镜像不包含gocron-node, gocron-node需要和具体业务一起构建


### 开发环境

1. 安装Go1.18+, Node.js, npm
2. 安装前端依赖 `make install-vue`
3. 启动gocron, gocron-node `make run`  浏览器可直接访问： http://localhost:5920
4. 启动node server `make run-vue`, 访问地址 http://localhost:8080， API请求会转发给gocron

`make` 编译

`make run` 编译并运行

`make package` 打包 
> 生成当前系统的压缩包 gocron-v2.x.x-darwin-amd64.tar.gz gocron-node-v2.x.x-darwin-amd64.tar.gz

`make package-all` 生成Windows、Linux、Mac的压缩包

### 命令

* gocron
    * -v 查看版本

* gocron web
    * --host 默认0.0.0.0
    * -p 端口, 指定端口, 默认5920
    * -e 指定运行环境, dev|test|prod, dev模式下可查看更多日志信息, 默认prod
    * -h 查看帮助
* gocron-node
    * -allow-root *nix平台允许以root用户运行
    * -s ip:port 监听地址  
    * -enable-tls 开启TLS    
    * -ca-file CA证书文件 
    * -cert-file 证书文件  
    * -key-file  私钥文件
    * -h 查看帮助
    * -v 查看版本

## 程序使用的组件
* Web框架 [Macaron](http://go-macaron.com/)
* 定时任务调度 [Cron](https://github.com/robfig/cron)
* ORM [Xorm](https://github.com/go-xorm/xorm)
* UI框架 [Element Plus](https://github.com/element-plus/element-plus)
* 依赖管理 [Govendor](https://github.com/kardianos/govendor)
* RPC框架 [gRPC](https://github.com/grpc/grpc)

## ChangeLog

v2.0.0
--------
+ LDAP用户认证
* 添加项目管理，项目和主机,任务关联
* 进程管理(队列消费程序)
* Vue3+ElementPlus 重构前端页面


[ouqiang/gocron ChangeLog](https://github.com/ouqiang/gocron#changelog)