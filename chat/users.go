package chat

import (
	"net"
	"sync"
)

type User struct {
	name string
	conn net.Conn
}

func (user User) Id() string {
	return user.conn.RemoteAddr().String()
}

type Users struct {
	Users map[string]User
	mutex sync.Mutex
}

func (users Users) ForEach(action func(user User)) {
	for _, user := range users.Users {
		action(user)
	}
}

func (users Users) AddUser(conn net.Conn) User {
	user := User{
		name: conn.RemoteAddr().String(),
		conn: conn,
	}
	users.UpdateUser(user)
	return user
}

func (users Users) UpdateUser(user User) User {
	users.ensureSync(func() {
		users.Users[user.Id()] = user
	})
	return user
}

func (users Users) RemoveUser(conn net.Conn) {
	users.ensureSync(func() {
		delete(users.Users, conn.RemoteAddr().String())
	})
}

func (users Users) ensureSync(action func()) {
	users.mutex.Lock()
	action()
	users.mutex.Unlock()

}
