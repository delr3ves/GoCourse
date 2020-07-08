package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"github.com/delr3ves/GoCourse/chat"
)

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err)
	}
	defer listener.Close()

	log.Printf("Chat server started on :8888")

	chatRoom := chat.Chat{
		Users: chat.Users{
			Users: make(map[string]chat.User),
		},
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err)
			continue
		}

		io.WriteString(conn, "Bienvenido al Chat de GDG Marbella!\n")

		chatRoom.AddUser(conn)

		go func() {
			for {
				msg, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					if err == io.EOF {
						chatRoom.Users.RemoveUser(conn)
					} else {
						log.Println(err)
					}
					continue
				}
				chatRoom.ProcessMessage(conn, msg)
			}
		}()
	}
}
