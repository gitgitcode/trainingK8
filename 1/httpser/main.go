package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
  rt time.Duration = 10
  wt time.Duration = 10
)

func main(){
	server(rt,wt)
}

func server(rt time.Duration , wt time.Duration  ) *http.Server {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/healthz", healthz)
	s := &http.Server{
	  	//Addr : ":80",
	  	ReadTimeout: rt * time.Second,
	  	WriteTimeout: wt * time.Second,
	  	MaxHeaderBytes: 1<<20,
	  }
	  log.Fatalln(s.ListenAndServe())
	  return s
}



func healthz(w http.ResponseWriter, r *http.Request){
	msg := "200"
	w.WriteHeader(http.StatusOK)
	settHeader(w,r)
	_, err := w.Write([]byte(msg))
	if err != nil {
		return
	}
}


//hello 方法 打印
//接收客户端 request，并将 request 中带的 header 写入 response header
//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func hello(w http.ResponseWriter, r *http.Request){
	//定义返回信息
	var body []string
	//定义返回头状态码
	var code = http.StatusBadRequest

	body = append(body, fmt.Sprintf("StatusBadRequest: %d \n", code))
	body = append(body, fmt.Sprintf("客户端 Url: %s \n", r.URL ))
	log.Println(r.Host)
	ip, err := GetIp(r)
	if err != nil {
		log.Println(err)
		body = append(body, fmt.Sprintf("客户端 IP: %s \n", err ))
	}else{
		body = append(body, fmt.Sprintf("客户端 IP: %s \n", ip ))
	}

	settHeader(w,r)
	w.WriteHeader(code)
	//先写头
	for _,v := range body{
		//最后写返回
		_, err :=  w.Write([]byte(v))
		if err != nil{
			log.Printf( "w.Write :%s",err)
		}
	}

}

//设置头信息
func settHeader(w http.ResponseWriter, r *http.Request){
	//header := r.Header.Clone()
	s :=os.Getenv("VERSION")
	log.Println("version",s)
	w.Header().Add("VERSION",s)
	for i,v := range r.Header {
		//log.Printf("header ->%s:%s",i,v)
		for _, s := range v {
			w.Header().Add(i,s)
		}
	}
}


func GetIp(r *http.Request) (string ,error)  {
	// GetIp
	xRealIp := r.Header.Get("X-Real-IP")
	if ip:= net.ParseIP(xRealIp) ; ip != nil {
		return ip.String() ,nil
	}

	//var ipstr []string
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIp := strings.Split(ips,",")
	for _,ip := range splitIp{
		netIp := net.ParseIP(ip)
		if netIp != nil {
			return  netIp.String(),nil
			 //ipstr = append(ipstr, netIp.String() )
		}
	}

	ip, _ ,err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "",err
	}
	netIp := net.ParseIP(ip)
	if netIp != nil {
		return netIp.String(),nil
	}
	return "",fmt.Errorf("no Ip ")
}