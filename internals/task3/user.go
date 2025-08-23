package task3

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type UserManager struct {
	users map[int]*User
}

func NewUserManager() *UserManager {
	return &UserManager{
		users: make(map[int]*User),
	}
}

func (u *UserManager) AddUser(user *User) error {
	if _, ok := u.users[user.ID]; ok {
		return fmt.Errorf("User dengan ID %d sudah terdaftar", user.ID)
	}
	u.users[user.ID] = user
	fmt.Println("Berhasil menambahkan user")
	return nil
}

func (u *UserManager) GetUser(id int) (*User, error) {
	if user, ok := u.users[id]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("User dengan ID %d tidak ditemukan", id)
}
