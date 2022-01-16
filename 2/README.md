# 构建镜像
- 构建本地镜像
- 编写 Dockerfile 将  httpserver 容器化
- 将镜像推送至 docker 官方镜像仓库
- 通过 docker 命令本地启动 httpserver
- 通过 nsenter 进入容器查看 IP 配置

## 本地目录
```shell
.
├── Dockerfile
└── cmd
  └── linux_amd64_httpd
1 directory, 2 files
```
### 流程
Dockerfile 基于ubuntu 可执行而二进制文件linux_amd64_httpd，位于cmd目录下，
需要先基于ubuntu 然后拷贝文件目录，在执行权限授权，最后直接启动服务



