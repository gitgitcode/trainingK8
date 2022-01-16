# 构建基于ubuntu系统执行
# 执行http服务 文件
FROM ubuntu
ENV DEBIAN_FRONTEND=noninteractive
RUN sed -i "s/^exit 101$/exit 0/" /usr/sbin/policy-rc.d
#DEBIAN_FRONTEND这个环境变量，告知操作系统应该从哪儿获得用户输入。如果设置为”noninteractive”，
#你就可以直接运行命令，而无需向用户请求输入（所有操作都是非交互式的）
RUN \
apt-get update && \
apt-get install -y ca-certificates && \
apt-get install -y curl && \
apt-get install -y vim && \
rm -rf /var/lib/apt/lists/*

ADD ./cmd/linux_amd64_httpd linux_amd64_httpd
ENTRYPOINT ["/linux_amd64_httpd"]