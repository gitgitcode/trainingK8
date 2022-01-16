# 构建镜像
- 构建本地镜像
- 编写 Dockerfile 将  httpserver 容器化
- 将镜像推送至 docker 官方镜像仓库
- 通过 docker 命令本地启动 httpserver
- 通过 nsenter 进入容器查看 IP 配置

## 本地目录
```shell
.
├── Dockerfile1
└── cmd
  └── linux_amd64_httpd
1 directory, 2 files
```
### 流程
Dockerfile 基于ubuntu 可执行而二进制文件linux_amd64_httpd，位于cmd目录下，
需要先基于ubuntu 然后拷贝文件目录，在执行权限授权，最后直接启动服务

## 改进后
```shell
.
├── Dockerfile
├── Dockerfile1
├── README.md
├── cmd
│   └── linux_amd64_httpd
└── main.go


```
- > docker build -t gohttp:v1   -f  Dockerfile  .
 ```
 [+] Building 9.7s (12/12) FINISHED                                                                         
  => [internal] load build definition from Dockerfile                                                  0.5s
  => => transferring dockerfile: 346B                                                                  0.0s
  => [internal] load .dockerignore                                                                     0.3s
  => => transferring context: 2B                                                                       0.0s
  => [internal] load metadata for docker.io/library/busybox:latest                                     0.0s
  => [internal] load metadata for docker.io/library/golang:1.17                                        2.0s
  => CACHED [stage-1 1/2] FROM docker.io/library/busybox                                               0.0s
  => [build 1/4] FROM docker.io/library/golang:1.17@sha256:0fa6504d3f1613f554c42131b8bf2dd1b2346fb69c  0.0s
  => [internal] load build context                                                                     0.3s
  => => transferring context: 29B                                                                      0.0s
  => CACHED [build 2/4] WORKDIR /httpd/                                                                0.0s
  => CACHED [build 3/4] COPY main.go .                                                                 0.0s
  => [build 4/4] RUN CGO_ENABLED=0 GOOS=linux  go build -installsuffix cgo  -v -o  httpsr main.go      5.4s
  => [stage-1 2/2] COPY --from=build /httpd/httpsr /                                                   0.6s
  => exporting to image                                                                                0.5s
  => => exporting layers
   => => writing image sha256:eb1cc7565032843aa5075cabe5524f5d0b93979c01319334b38a8d0813b7afcf          0.0s
#完成
> docker images 

```
## 多阶段构建方法
```dockerfile

#多阶段构建方法
FROM golang:1.17 AS build
#基于  golang:1.17 
WORKDIR /httpd/ 
#本地目录
COPY main.go .
# 拷贝 源文件到容器目录
# copy 不支持 ../ 相对目录
#不会创创建新目录 必须写全 abc/1 cb/1 1 这个目录必须写
ENV  GO111MODULE=on
ENV  GOPROXY=https://goproxy.cn,direct
# 环境变量

RUN CGO_ENABLED=0 GOOS=linux  go build -installsuffix cgo  -v -o  httpsr main.go

#镜像 busybox 
FROM busybox
COPY --from=build /httpd/httpsr /
# COPY 命令把前一阶段构建的产物拷贝到另一个镜像中
#COPY --from=0 
EXPOSE 80
#暴露 80 端口
# 要将 EXPOSE 和在运行时使用 -p <宿主端口>:<容器端口> 区分开来。-p，是映射宿主端口和容器端口，换句话说，就是将容器的对应端口服务公开给外界访问，而 EXPOSE 仅仅是声明容器打算使用什么端口而已，并不会自动在宿主进行端口映射。
ENTRYPOINT ["/httpsr"]
#执行 
```
