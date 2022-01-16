#多阶段构建方法
FROM golang:1.17 AS build
WORKDIR /httpd/
COPY main.go .

ENV  GO111MODULE=on
ENV  GOPROXY=https://goproxy.cn,direct

RUN CGO_ENABLED=0 GOOS=linux  go build -installsuffix cgo  -v -o  httpsr main.go

FROM busybox
COPY --from=build /httpd/httpsr /
EXPOSE 80
ENTRYPOINT ["/httpsr"]