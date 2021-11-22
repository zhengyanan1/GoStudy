package main

import "net"

type User struct {
	Name string
	Addr string
	C chan string
	conn net.Conn
}

//create a user

func NewUser(conn net.Conn) *User  {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C: make(chan string),
		conn: conn,
	}

	go user.ListenMessage()

	return user
}

// ListenMessage 监听当前User的channel方法，一旦有消息，就发送给对端客户端
func (p *User) ListenMessage() {
	for {
		msg := <- p.C

		p.conn.Write([]byte(msg + "\n"))
	}
}