package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip string
	Port int

//	在线用户列表
	OnlineMap map[string]*User
	mapLock sync.RWMutex

//	消息广播的channel
	Message chan string
}

// NewServer 创建一个server的接口
func NewServer(ip string, port int) *Server{
	server := &Server{
		Ip: ip,
		Port: port,
		OnlineMap: make(map[string]*User),
		Message: make(chan string),
	}

	return server
}

//监听Message广播消息channel的goroutine,一旦有消息就发送给全部在线User
func (this *Server) ListenMessage() {
	for {
		msg := <- this.Message

	//	将message发送给全部user
		this.mapLock.Lock()
		for _, user := range this.OnlineMap{
			user.C <- msg
		}
		this.mapLock.Unlock()
	}
}

//广播消息的方法
func (this *Server)BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn){

	user := NewUser(conn, this)

	user.Online()

//	接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF{
				fmt.Println("Conn Read err:", err)
				return
			}

			//去掉用户消息最后的'\n'
			msg := string(buf[: n - 1])
			user.DoMessage(msg)
		}
	}()

//	当前handler阻塞
	select {}
}

// Start 启动服务器的接口
func(this *Server) Start(){
	//	socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil{
		fmt.Println("net.Listen err:", err)
		return
	}
	//	close listen socket
	defer listener.Close()

	//启动监听Message的goroutine
	go this.ListenMessage()

	for {
		//	accept
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("listener accept err:", err)
			continue
		}

		//do handler
		go this.Handler(conn)
	}
}