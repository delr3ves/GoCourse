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
	chat.SendMessage(conn, fmt.Sprintf("[GDG Bot] New name: %s", name))

	return true
}

func ListUsers(conn net.Conn, msg string, chat Chat) bool {
	if strings.TrimSpace(msg) != "list users" {
		return false
	}
	chat.SendMessage(conn, fmt.Sprintf("[GDG Bot] There are %d users", len(chat.Users.Users)))
	chat.Users.ForEach(func(user User) {
		chat.SendMessage(conn, fmt.Sprintf("[GDG Bot]\t %s from %s", user.name, user.conn.RemoteAddr().String()))
	})
	return true
}

func WhoAmI(conn net.Conn, msg string, chat Chat) bool {
	if strings.TrimSpace(msg) != "who am i" {
		return false
	}
	you := chat.Users.Users[conn.RemoteAddr().String()]
	chat.SendMessage(conn, fmt.Sprintf("[GDG Bot] Hi: %s from %s", you.name, you.conn.RemoteAddr().String()))
	return true
}

func SendMessage(conn net.Conn, msg string, chat Chat) bool {
	chat.Users.ForEach(sendMessageToEveryone(conn, msg, chat))
	return true
}

func sendMessageToEveryone(conn net.Conn, msg string, chat Chat) func(user User) {
	return func(user User) {
		sender := chat.Users.Users[conn.RemoteAddr().String()]
		if user.Id() != conn.RemoteAddr().String() {
			chat.SendMessage(user.conn, fmt.Sprintf("[%s]: %s", sender.name, msg))
		}
	}
}
