package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func processConn(conn net.Conn) {
	// 3. 与客户通信
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
		fmt.Print("请回复:")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}

// tcp server端

func main() {
	// 1. 本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp service on 127.0.0.1:20000 failed, err:", err)
		return
	}
	for {
		// 2. 等待别人和我连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		go processConn(conn)
	}

}
