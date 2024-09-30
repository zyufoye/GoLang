package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个Server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	//...当前连接的业务
	fmt.Println("链接建立成功！")
}

// 给当前的Server类绑定一个启动服务器的方法
func (this *Server) Start() {
	//socket listen
	//以tcp形式监听当前服务器的Ip和Port
	Listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer Listener.Close()

	for {
		//accept
		conn, err := Listener.Accept()
		if err != nil {
			fmt.Println("Listener.Accept err:", err)
			continue
		}
		//do hander 做当前客户端的业务
		go this.Handler(conn)
	}

}
