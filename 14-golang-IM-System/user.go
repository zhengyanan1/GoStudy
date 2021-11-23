package main

import "net"

type User struct {
	Name string
	Addr string
	C chan string
	conn net.Conn
	server *Server
}

//create a user

func NewUser(conn net.Conn, server *Server) *User  {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C: make(chan string),
		conn: conn,
		server: server,
	}

	go user.ListenMessage()

	return user
}

// Online user login
func (this *User) Online(){
	//当前连接的业务
	//	fmt.Println("连接建立成功")

	//	用户上线,将用户加入到OnlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//	广播当前用户上线消息
	this.server.BroadCast(this, "已上线")
}

// Offline user logout
func (this *User) Offline(){
	//	用户上线,将用户从OnlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//	广播当前用户上线消息
	this.server.BroadCast(this, "已下线")
}

//
func (this *User) SendMessage(msg string){
	this.conn.Write([]byte(msg))
}

// DoMessage user handle message
func (this *User) DoMessage(msg string){
	if msg == "who" {
	//	查询当前用户列表
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap{
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线\n"
			this.SendMessage(onlineMsg)
		}

		this.server.mapLock.Unlock()

	} else {
		this.server.BroadCast(this, msg)
	}
}


// ListenMessage 监听当前User的channel方法，一旦有消息，就发送给对端客户端
func (p *User) ListenMessage() {
	for {
		msg := <- p.C

		p.conn.Write([]byte(msg + "\n"))
	}
}