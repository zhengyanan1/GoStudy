package main

import (
	"net"
	"strings"
)

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
	} else if len(msg) > 7 && msg[:7] == "rename|"{
	//	消息格式：rename|zhangsan
		newName := strings.Split(msg, "|")[1]
	//	判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok{
			this.SendMessage("用户名" + newName + "被使用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMessage("用户名已更新成功为" + newName + "\n")
		}

	} else if len(msg) > 4 && msg[:3] == "to|"{
	//	消息格式：to|zhangsan|消息内容

		//	1.获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMessage("消息格式不正确，请使用: \"to|zhangsan|nihao\" 格式\n")
			return
		}

		//	2 根据用户名获取User对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok{
			this.SendMessage("用户" + remoteName + "不在线\n")
			return
		}

		//	3。获取消息内容，发送给该用户
		content := strings.Split(msg, "|")[2]
		if content == ""{
			this.SendMessage("无消息内容，请重发\n")
			return
		}

		remoteUser.SendMessage(this.Name + "说:" + content + "\n")

	} else if len(msg) > 0 {
		this.server.BroadCast(this, msg)
	} else {

	}
}


// ListenMessage 监听当前User的channel方法，一旦有消息，就发送给对端客户端
func (p *User) ListenMessage() {
	for {
		msg := <- p.C

		p.conn.Write([]byte(msg + "\n"))
	}
}