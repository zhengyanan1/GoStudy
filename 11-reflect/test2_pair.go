package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//pair<type:*File, value:"/dev/tty"文件描述符>
	var tty, err = os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	//pair<type, value:>
	var r io.Reader
	//pair<type:*File, value:"/dev/tty"文件描述符>
	r = tty

	var w io.Writer
	//pair<type:*File, value:"/dev/tty"文件描述符>
	w = r.(io.Writer)
	w.Write([]byte("HELLO THIS IS A TEST!!!\n"))
}