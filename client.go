package main

import (
	"fmt"
	"os"
)

// 处理错误的函数
func ClientHandlerErr(err error, where string) {
	if err != nil {
		fmt.Println(err, where)
		os.Exit(1)
	}
}

// func main() {

// 	conn, err := net.Dial("tcp", "127.0.0.1:8888")

// 	ClientHandlerErr(err, "net.Dial")

// 	//消息缓冲区
// 	//buff := make([]byte, 1024)

// 	//从命令行标准输入中获取数据
// 	reader := bufio.NewReader(os.Stdin)
// 	buff := make([]byte, 1024)
// 	//循环发送数据
// 	for {
// 		fmt.Printf(">")
// 		linemsg, _, _ := reader.ReadLine()
// 		conn.Write(linemsg)
// 		if string(linemsg) == "" {
// 			continue
// 		}
// 		if string(linemsg) == "exit" {
// 			break
// 		}
// 		n, readerr := conn.Read(buff)
// 		ClientHandlerErr(readerr, "conn.Read")
// 		receivedmsg := string(buff[0:n])
// 		fmt.Println(receivedmsg, "from ", conn.RemoteAddr())
// 	}
// 	conn.Close()
// }
