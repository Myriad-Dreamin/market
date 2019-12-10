# Minimum Market
market backend powered by minimum

## interface definition
https://github.com/Myriad-Dreamin/market/blob/master/control

## Documentation
API文档：

https://myriaddreamin.docs.apiary.io/

https://github.com/Myriad-Dreamin/market/blob/master/apiary.apib

## Installation

如果你需要在物理机上编译，那么你需要先下载[golang](https://golang.org/dl/)和[gcc](https://gcc.gnu.org/)，并设置下列环境变量：

```bash
$ export GO111MODULE=on GOPROXY=https://goproxy.io
```

物理机的手动编译方法如下：

```bash
$ go install github.com/Myriad-Dreamin/go-magic-package/package-attach-to-path
$ python makefile.py generate
go generate ./types/custom-fields.go
go generate ./types/error.go
go generate ./types/goods_type.go
go generate ./test/goods_test.go
go generate ./package-probe/init.go
go generate ./package-probe/probe_test.go
...
$ go build -o srv
```

你也可以使用Docker已经打包好的编译环境

```bash
$ python makefile.py image
docker build --tag minimum-market/backend:alpine .
Sending build context to Docker daemon  30.33MB
Step 1/17 : FROM golang:alpine AS build-env
---> 69cf534c966a
Step 2/17 : RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
---> Running in ce132172a09d
Removing intermediate container ce132172a09d
---> 376ee4fd7653
Step 3/17 : ADD . /market
---> e670225b4744
Step 4/17 : WORKDIR /market
---> Running in 8cf0497e7f93
Removing intermediate container 8cf0497e7f93
---> f09a519cf0e2
Step 5/17 : ENV GO111MODULE=on GOPROXY=https://goproxy.io
---> Running in a28b2d6316e5
Removing intermediate container a28b2d6316e5
---> 7b770ce73343
Step 6/17 : RUN apk update && apk add gcc && apk add musl-dev
---> Running in 2576257587fd
fetch http://mirrors.ustc.edu.cn/alpine/v3.10/main/x86_64/APKINDEX.tar.gz
fetch http://mirrors.ustc.edu.cn/alpine/v3.10/community/x86_64/APKINDEX.tar.gz
v3.10.3-75-g38169cc48e [http://mirrors.ustc.edu.cn/alpine/v3.10/main]
v3.10.3-68-ge1e42c5d6c [http://mirrors.ustc.edu.cn/alpine/v3.10/community]
OK: 10341 distinct packages available
(1/10) Installing binutils (2.32-r0)
(2/10) Installing gmp (6.1.2-r1)
(3/10) Installing isl (0.18-r0)
(4/10) Installing libgomp (8.3.0-r0)
(5/10) Installing libatomic (8.3.0-r0)
(6/10) Installing libgcc (8.3.0-r0)
(7/10) Installing mpfr3 (3.1.5-r1)
(8/10) Installing mpc1 (1.1.0-r0)
(9/10) Installing libstdc++ (8.3.0-r0)
(10/10) Installing gcc (8.3.0-r0)
Executing busybox-1.30.1-r2.trigger
OK: 93 MiB in 25 packages
(1/1) Installing musl-dev (1.1.22-r3)
OK: 103 MiB in 26 packages
Removing intermediate container 2576257587fd
---> 940342f6b3df
Step 7/17 : RUN GOOS=linux GOARCH=amd64 go build -v -o ../server .
---> Running in 86503fb594eb
go: downloading github.com/Myriad-Dreamin/functional-go v0.0.0-20191102095642-532e6dc9bfd5
go: downloading go.uber.org/zap v1.12.0
...
github.com/Myriad-Dreamin/market
Removing intermediate container 86503fb594eb
---> a531e1989773
Step 8/17 : FROM alpine
latest: Pulling from library/alpine
Digest: sha256:c19173c5ada610a5989151111163d28a67368362762534d8a8121ce95cf2bd5a
Status: Downloaded newer image for alpine:latest
---> 965ea09ff2eb
Step 9/17 : RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
---> Using cache
---> 59428e652e83
Step 10/17 : WORKDIR /
---> Using cache
---> 222b3b9b4b0d
Step 11/17 : RUN apk add -U tzdata
---> Using cache
---> 06df42051739
Step 12/17 : RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
---> Using cache
---> fb0711f09d04
Step 13/17 : COPY --from=build-env /server /usr/local/bin/server
---> a9607105879c
Step 14/17 : COPY docker/run.sh /usr/local/bin/run.sh
---> e9337ea67d62
Step 15/17 : RUN chmod +x /usr/local/bin/run.sh
---> Running in f7f2b82936d7
Removing intermediate container f7f2b82936d7
---> a9b29f3ce0bb
Step 16/17 : EXPOSE 23336
---> Running in ea870b13b1f7
Removing intermediate container ea870b13b1f7
---> 5edca078b777
Step 17/17 : CMD run.sh
---> Running in 877c54f3ae80
Removing intermediate container 877c54f3ae80
---> 26107e66e56e
Successfully built 26107e66e56e
Successfully tagged minimum-market/backend:alpine
```

#### 测试Docker镜像

新起一个临时的mysql实例

```bash
$ python makefile.py run_testdb
docker run -id -p 23306:3306 -e MYSQL_ROOT_PASSWORD=12345678 -e MYSQL_DATABASE=market -e MYSQL_USER=madmin -e MYSQL_PASSWORD=12345678 --name backend-testdb mysql:5.7
dd1c326a9b52d482d1fc9ebd85a3ff640ad16fbc697816d7e8257df84aadb320
```

创建一个host的docker容器，并使用docker/config.toml中的示例配置

```bash
$ python makefile.py run docker/config.toml
docker run -id -v /home/kamiyoru/work/gosrc/src/github.com/Myriad-Dreamin/market/docker/config.toml:/config.toml --network=host --name backend minimum-market/backend:alpine
39feeefbc7b8abbbdf2c7fc766b02f8c2e2cd77385f075119f59072d4f9a6949
```

查看已经创建的容器

```bash
$ docker ps -a
CONTAINER ID        IMAGE                                 COMMAND                  CREATED             STATUS                   PORTS                                NAMES
39feeefbc7b8        minimum-market/backend:alpine         "/bin/sh -c run.sh"      7 seconds ago       Up 7 seconds                                                  backend
dd1c326a9b52        mysql:5.7                             "docker-entrypoint.s…"   54 seconds ago      Up 53 seconds            33060/tcp, 0.0.0.0:23306->3306/tcp   backend-testdb

```

查看启动日志

```bash
$ docker logs 39f
2019-12-11T00:25:57.532+0800	Debug	 server/fs.go:38 		{"level": "info", "_msg": "host path", "name": "GoodsPicturePath", "creating-path": "//"}
2019-12-11T00:25:57.536+0800	Debug	 server/fs.go:38 		{"level": "info", "_msg": "host path", "name": "NeedsPicturePath", "creating-path": "//"}
2019-12-11T00:25:57.549+0800	Debug	 server/db.go:41 		{"level": "info", "_msg": "connected to database", "connection-type": "mysql", "user": "madmin", "database": "market", "charset": "utf8mb4", "location": "Local"}
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /v1/statistic/fee         --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (5 handlers)
[GIN-debug] GET    /v1/statistic/count       --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (5 handlers)
[GIN-debug] PUT    /v1/needs/:nid            --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (6 handlers)
[GIN-debug] DELETE /v1/needs/:nid            --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (6 handlers)
[GIN-debug] DELETE /v1/needs/:nid/force      --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (6 handlers)
[GIN-debug] PUT    /v1/needs/:nid/picture    --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (6 handlers)
[GIN-debug] GET    /v1/needs/:nid/picture    --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (5 handlers)
[GIN-debug] GET    /v1/needs/:nid/inspect    --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (6 handlers)
[GIN-debug] POST   /v1/needs                 --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (5 handlers)
[GIN-debug] DELETE /v1/goods/:goid           --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (6 handlers)
[GIN-debug] DELETE /v1/goods/:goid/force     --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (6 handlers)
[GIN-debug] PUT    /v1/goods/:goid/picture   --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (6 handlers)
[GIN-debug] GET    /v1/goods/:goid/picture   --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (5 handlers)
[GIN-debug] GET    /v1/goods/:goid/inspect   --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (6 handlers)
[GIN-debug] PUT    /v1/goods/:goid           --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (6 handlers)
[GIN-debug] POST   /v1/goods                 --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (5 handlers)
[GIN-debug] POST   /v1/user/:id/goods/:goid/buy --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (5 handlers)
[GIN-debug] POST   /v1/user/:id/goods/:goid/confirm-buy --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (5 handlers)
[GIN-debug] POST   /v1/user/:id/needs/:nid/sell --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (5 handlers)
[GIN-debug] POST   /v1/user/:id/needs/:nid/confirm-sell --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (5 handlers)
[GIN-debug] PUT    /v1/user/:id/password     --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (6 handlers)
[GIN-debug] PUT    /v1/user/:id              --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (6 handlers)
[GIN-debug] DELETE /v1/user/:id              --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (6 handlers)
[GIN-debug] GET    /v1/user/:id              --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (4 handlers)
[GIN-debug] POST   /v1/user                  --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (4 handlers)
[GIN-debug] POST   /v1/login                 --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (4 handlers)
[GIN-debug] GET    /v1/user-list             --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (4 handlers)
[GIN-debug] GET    /v1/needs/:nid            --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (4 handlers)
[GIN-debug] GET    /v1/needs-list            --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (4 handlers)
[GIN-debug] GET    /v1/goods/:goid           --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (4 handlers)
[GIN-debug] GET    /v1/goods-list            --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (4 handlers)
[GIN-debug] GET    /ping                     --> github.com/Myriad-Dreamin/market/lib/controller/gin.ToGinHandler.func1 (4 handlers)
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "provider/model/objectDB", "name": "*splayer.ObjectDB"}
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "provider/model", "name": "*splayer.Provider"}
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "global/httpEngine", "name": "*gin.Engine"}
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "provider/model/enforcer", "name": "*casbin.SyncedEnforcer"}
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "provider/model/goodsDB", "name": "*splayer.GoodsDB"}
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "middleware/route-auth", "name": "*controller.Middleware"}
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "provider/model/needsDB", "name": "*splayer.NeedsDB"}
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "provider/model/userDB", "name": "*splayer.UserDB"}
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "provider/service", "name": "*service.Provider"}
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "provider/router", "name": "*router.Provider"}
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "global/configuration", "name": "*config.ServerConfig"}
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "middleware/jwt", "name": "*jwt.Middleware"}
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "global/logger", "name": "*logger.kitLogger"}
2019-12-11T00:25:57.653+0800	Debug	 module/module.go:98 		{"level": "debug", "_msg": "module installed", "path": "middleware/cors", "name": "gin.HandlerFunc"}
[GIN-debug] Listening and serving HTTP on :23336
# 使用服务探针检测容器生命状态
$ curl http://localhost:23336/ping
{"message":"pong"}
```

停止运行，并删除容器

```bash
$ python makefile.py stop_run
docker stop backend
backend
$ python makefile.py remove_run
docker rm backend
backend
$ python makefile.py stop_testdb
docker stop backend-testdb
backend-testdb
$ python makefile.py remove_testdb
docker rm backend-testdb
backend-testdb
```

#### 从docker-compose.yml创建服务容器组

docker-compose在windows下是docker安装时内置的命令，但在linux下，你需要从repo中下载，链接为[docker-compose](https://github.com/docker/compose/releases)。

```bash
$ curl -L https://github.com/docker/compose/releases/download/1.25.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   617    0   617    0     0    596      0 --:--:--  0:00:01 --:--:--   596
100 16.2M  100 16.2M    0     0   440k      0  0:00:37  0:00:37 --:--:-- 2143k
$ chmod +x /usr/local/bin/docker-compose
```

安装完了以后，可以从`docker-compose.yml`运行服务容器组。

makefile会从context中自动读取以下参数填充模板`docker-compose.template.yml`从而生成`docker-compose.yml`，所以如果需要修改服务容器配置，请修改template。

context参数如下：

+ `redis-root-password`: redis root的密码，默认值为`12345678`

+ `mysql-root-password`: mysql root的密码，默认值为`12345678`

+ `conf-path`: mysql conf文件夹在本机的相对路径，默认值为`testdb/conf`

+ `logs-path`: mysql logs文件夹在本机的相对路径，默认值为`testdb/logs`

+ `data-path`: mysql data文件夹在本机的相对路径，默认值为`testdb/data`

+ `node-name`: 后端服务的镜像名，默认值为`minimum-market/backend:alpine`

+ `instance-name`: 实例名，在此处无用

+ `target-port`: 服务暴露端口，默认值为`23335`

+ `config-file`: 配置文件在本机的相对路径，默认值为`config.toml`

+ `mysql-norm-password`: mysql 普通admin的密码，默认值为`12345678`

如何设置context，可以使用`python cli.py apply_context`命令也可以手动修改`.market-env.json`

部署方式如下：

```bash
# 部署服务容器
$ python makefile.py up
# 启动服务容器
$ python makefile.py start
$ docker ps -a
CONTAINER ID        IMAGE                                 COMMAND                  CREATED             STATUS                   PORTS                                NAMES
6ab58bddf7f1        minimum-market/backend:alpine         "/bin/sh -c run.sh"      6 minutes ago       Up 58 seconds            0.0.0.0:23335->23336/tcp             market_server_1
466ca6db23ed        mysql:5.7                             "docker-entrypoint.s…"   6 minutes ago       Up About a minute        33060/tcp, 0.0.0.0:23306->3306/tcp   market_mysql_1
b77c4e24e3c6        redis                                 "docker-entrypoint.s…"   6 minutes ago       Up 59 seconds            0.0.0.0:23379->6379/tcp              market_redis_1
# 使用服务探针检测容器生命状态
$ curl http://localhost:23335/ping
{"message":"pong"}
# 停止服务容器
$ python makefile.py stop
# 销毁服务容器
$ python makefile.py down
# 清除临时配置的服务器持久化路径
$ python makefile.py clean
```

## Minimum-Cli/Makefile

python环境如下：

```bash
$ python --version
Python 2.7.15
$ python3 --version
Python 3.7.2
```

使用`cli.py`需要安装`fire`

```bash
pip install fire
```

`python-fire`的文档参见[python-fire](https://github.com/google/python-fire)

如果你仅仅安装，则不需要安装`fire`，因为`makefile.py`仅仅依赖于`pymake.py`，而`pymake`只需要原生的`python2/3`模块。

`cli`有如下命令：

```plain
NAME                                      
    cli.py                                
                                          
SYNOPSIS                                  
    cli.py COMMAND | VALUE                
                                          
COMMANDS                                  
    COMMAND is one of the following:      
                                          
     apply_context                        
                                          
     create_pure_service                  
                                          
     create_router                        
                                          
     create_service                       
                                          
     create_template                      
                                          
     fast_generate                        
                                          
     generate                             
                                          
     hello                                
                                          
     install                              
                                          
     make                                 
                                          
     replace                              
                                          
     template_to                          
                                          
     templates_to                         
                                          
     test
```

#### apply_context

```bash
NAME
    cli.py apply_context

SYNOPSIS
    cli.py apply_context <flags> [KEY_VALUES]...

POSITIONAL ARGUMENTS
    KEY_VALUES

FLAGS
    --f=F (required)
    --c=C (required)
```

此命令用于帮助设置`makefile`的context，context保存在`.market-env.json`中，你可以手动修改。

```python
$ python3 cli.py apply_context mysql-root-password=aabb1234
# .market-env.json
{
    "mysql-root-password": "aabb1234"
}
```
