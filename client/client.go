package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
	for {
		go receiveMessages(conn)()
		go sendMessages(conn)()
	}
}

func receiveMessages(conn *net.TCPConn) func() {
	return func() {
		for {
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print(message)
		}
	}
}

func sendMessages(conn *net.TCPConn) func() {
	return func() {
		for {
			//read in input from stdin
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')

			//send to socket
			fmt.Fprintln(conn, text)
		}
	}
}
