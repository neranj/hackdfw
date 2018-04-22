package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*User
)

func init(){
	orm.RegisterModel(new(User))
}

type User struct {
	Id       int    `orm:"pk;auto"`
	Email    string
	Password string `json:"-"`
}

func (u *User) TableName() string {
	return "users"
}

func AddUser(u *User) error {
	o := orm.NewOrm()
	_, err := o.Insert(u)
	return err
}

func GetUser(id int) (*User, error) {
	o := orm.NewOrm()
	user := &User{Id: id}
	err := o.Read(user)

	return user, err
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
