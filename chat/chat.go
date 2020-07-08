package chat

import (
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
	chainOfResponsibility := []func(conn net.Conn, msg string, chat Chat) bool {
		UpdateUser,
		ListUsers,
		WhoAmI,
		SendMessage,
	}
	for _, action := range chainOfResponsibility{
		processed := action(conn, msg, chat)
		if (processed) {
			return
		}
	}
}

func (chat Chat) SendMessage(conn net.Conn, msg string) {
	io.WriteString(conn, msg)
}

