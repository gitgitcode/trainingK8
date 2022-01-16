#build http服务

- ``` 
  go build -v -o http main.go
  ```

## mac 下编译linux 可执行文件
- ``` 
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o linux_amd64_httpd main.go
  ``` 
  - -v 打印log
  - -o 输出的文件名称 

GOOS refers to the operating system (Linux, Windows, BSD, etc.),
while GOARCH refers to the architecture to build for.

GOOS 指的是操作系统(Linux、 Windows、 BSD 等) ，而 GOARCH 指的是要构建的体系结构。
Go的runtime环境变量CGO_ENABLED=1，即默认开始cgo，允许你在Go代码中调用C代码

如果标准库中是在CGO_ENABLED=1情况下编译的， 那么编译出来的最终二进制文件可能是动态链接，
所以建议设置 CGO_ENABLED=0以避免移植过程中出现的不必要问题。