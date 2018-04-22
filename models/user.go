package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"fmt"
)

var (
	UserList map[string]*User
)

func init(){
	orm.RegisterModel(new(User))
}

type User struct {
	Id       int    `orm:"pk"`
	Email    string
	Password string
}


func AddUser(u User) int {
	o := orm.NewOrm()
	fmt.Println(o.Insert(&u))

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

func DeleteProject(uid int) bool {
	o := orm.NewOrm()
	_,err := o.Delete(&Project{ Projectid: uid})
	if(err == nil){
		return false
	}else{
		return true
	}
}


