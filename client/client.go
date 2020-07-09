package main

import (
	"bufio"
	"fmt"
	"github.com/delr3ves/GoCourse/client/core"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "host", "port")
		os.Exit(1)
	}
	host := os.Args[1]
	port := os.Args[2]
	servAddr := fmt.Sprintf("%s:%s", host, port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		fmt.Print("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Print("Dial failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	chatWindow := core.NewChatWindow(
		core.NewMessageSender(conn),
	)
	messageProcessor := &core.MessageReceivedCallback{}

	go receiveMessages(conn, messageProcessor)

	chatWindow.Init(messageProcessor)
}

func receiveMessages(conn *net.TCPConn, messageProcessor *core.MessageReceivedCallback) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		messageProcessor.OnMessageReceived(strings.TrimRight(message, "\n"))
	}
}
