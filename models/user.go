package models

import (
	"errors"
	"strconv"
	"time"

)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
}

type User struct {
	Id       string
	Email    string
	Password string
}


func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Email != "" {
			u.Email = uu.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(email, password string) bool {
	for _, u := range UserList {
		if u.Email == email && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
