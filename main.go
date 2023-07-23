package main

import (
	"bufio"
	"fmt"
	"mypro/controller"
	"net"
	"os"
)

func ServerHandlerErr(err error, where string) {
	if err != nil {
		fmt.Println(err, where)
		os.Exit(1)
	}
}

func ProcessMsg(conn net.Conn) {

	buff := make([]byte, 1024)

	for {

		n, readerr := conn.Read(buff)

		ServerHandlerErr(readerr, "conn.Read")

		receivedmsg := string(buff[0:n])
		fmt.Println(receivedmsg, "from ", conn.RemoteAddr())

		if receivedmsg == "exit" {
			break
		}
		response := controller.DecodeUserMsg(receivedmsg, conn.RemoteAddr().String())
		conn.Write([]byte(response))
	}
	conn.Close()
	fmt.Println("Connection Closed From ", conn.RemoteAddr())
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	ServerHandlerErr(err, "net.Listen")
	fmt.Println("Start listening at ", listen.Addr())

	for {
		conn, accepterr := listen.Accept()
		ServerHandlerErr(accepterr, "listen.Accept")
		go ProcessMsg(conn)
		linemsg, _, _ := reader.ReadLine()
		if string(linemsg) == "exit" {
			break
		}
	}

}
