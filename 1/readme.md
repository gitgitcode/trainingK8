# 简单http服务器

- 环境变量 可以在 执行前设置
    
    % VERSION=293 go run main.go
    
- 输出WriteHeader 在先
  
  w.WriteHeader(http.StatusOK)
  
  w.Write([]byte(msg))
  
  

- [ ] [golangbyexample 获取ip](https://golangbyexample.com/golang-ip-address-http-request/)

- [ ] [参考获取ip](https://github.com/polaris1119/goutils/blob/master/ip.go)

