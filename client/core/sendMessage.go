package core

import (
	"fmt"
	"net"
)

type MessageSender struct {
	conn net.Conn
}

func NewMessageSender(conn net.Conn) MessageSender {
	return MessageSender{conn: conn}
}

func (sender MessageSender) sendMessage(message string) {
	fmt.Fprintln(sender.conn, message)
}
