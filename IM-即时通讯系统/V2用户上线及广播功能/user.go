package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

// 创建一个用户的API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	go user.ListenMessage()
	return user
}

// 监听当前user channel的方法，一旦有消息直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
