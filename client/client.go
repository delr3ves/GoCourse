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
		[]string{},
		core.NewMessageSender(conn),
	)

	for {
		go receiveMessages(conn, &chatWindow)()
	}
}

func receiveMessages(conn *net.TCPConn, chatWindow *core.ChatWindow) func() {
	return func() {
		for {
			message, _ := bufio.NewReader(conn).ReadString('\n')
			chatWindow.PrintMessage(strings.TrimRight(message, "\n"))
		}
	}
}
