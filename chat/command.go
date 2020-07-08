package chat

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

func UpdateUser(conn net.Conn, msg string, chat Chat) bool {
	var re = regexp.MustCompile(`(?m)set name (.+)(\n)`)
	match := re.FindStringSubmatch(msg)
	if len(match) < 2 {
		return false
	}
	name := strings.TrimSpace(match[1])
	chat.Users.UpdateUser(User{conn: conn, name: name})
	return true
}

func ListUsers(conn net.Conn, msg string, chat Chat) bool {
	if strings.TrimSpace(msg) != "list users"{
		return false
	}
	chat.SendMessage(conn, fmt.Sprintf("There are %d users\n", len(chat.Users.Users)))
	chat.Users.ForEach(func(user User) {
		chat.SendMessage(conn, fmt.Sprintf("\t %s from %s\n", user.name, user.conn.RemoteAddr().String()))
	})
	return true
}

func WhoAmI(conn net.Conn, msg string, chat Chat) bool {
	if strings.TrimSpace(msg) != "who am i" {
		return false
	}
	you := chat.Users.Users[conn.RemoteAddr().String()]
	chat.SendMessage(conn, fmt.Sprintf("Hi: %s from %s\n", you.name, you.conn.RemoteAddr().String()))
	return true
}

func SendMessage(conn net.Conn, msg string, chat Chat) bool {
	chat.Users.ForEach(sendMessageToEveryone(conn, msg, chat))
	return true
}

func sendMessageToEveryone(conn net.Conn, msg string, chat Chat) func(user User) {
	return func(user User) {
		sender := chat.Users.Users[conn.RemoteAddr().String()]
		var message string
		if user.Id() != conn.RemoteAddr().String() {
			message = fmt.Sprintf("[%s]: %s\n", sender.name, msg)
		} else {
			message = fmt.Sprintf("[you]: %s\n", msg)
		}
		chat.SendMessage(user.conn, message)
	}
}
