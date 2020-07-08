package chat

import (
	"fmt"
	"io"
	"net"
)

type Chat struct {
	Users Users
}

func (chat Chat) AddUser(conn net.Conn) Users {
	chat.Users.AddUser(conn)
	return chat.Users
}

func (chat Chat) ProcessMessage(conn net.Conn, msg string) {
	chat.Users.ForEach(sendMessageToEveryone(conn, msg))
}

func sendMessageToEveryone(conn net.Conn, msg string) func(user User) {
	return func(user User) {
		var message string
		if user.conn.RemoteAddr().String() != conn.RemoteAddr().String() {
			message = fmt.Sprintf("[%s]: %s\n", conn.RemoteAddr().String(), msg)
		} else {
			message = fmt.Sprintf("[you]: %s\n", msg)
		}
		io.WriteString(user.conn, message)
	}
}
