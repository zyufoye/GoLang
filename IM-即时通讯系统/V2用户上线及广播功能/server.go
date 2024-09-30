package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int
	//此处新增OnlineMap和Message属性
	OnlineMap map[string]*User
	maplock   sync.RWMutex

	//广播消息队列
	Message chan string
}

// 创建一个Server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部在线User
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		//将msg发送给全部在线User
		this.maplock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.maplock.Unlock()
	}
}

// 广播消息的方法
func (this *Server) Broadcast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg
}

// 完善处理流程，包含添加用户进onlinemap，还有广播用户上线消息
func (this *Server) Handler(conn net.Conn) {
	//...当前连接的业务
	fmt.Println("链接建立成功！")
	//用户上线，将用户加入到OnlineMap中
	user := NewUser(conn)
	this.maplock.Lock()
	this.OnlineMap[user.Name] = user
	this.maplock.Unlock()

	//广播当前用户上线消息
	this.Broadcast(user, "已上线")

	//当前handler阻塞，不要死亡
	select {}
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

	//在这里启动监听message的goroutine
	go this.ListenMessager()
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
