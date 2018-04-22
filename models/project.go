package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

var (
	ProjectList map[string]*Project
)

func init() {
	orm.RegisterModel(new(Project))
}

type Project struct {
	Projectid       int    `orm:"pk;auto"`
	Creatorid    string
	Lat string
	Long string
	Createdon string
	Updatedon string
	Title string
	Discription string
	Address string
	Status string
	Donepercentage string
}

func AddProject(u *Project) error {
	o := orm.NewOrm()
	_, err := o.Insert(u)
	return err
}

func GetProject(uid int) (*Project, error) {
	o := orm.NewOrm()
	project := &Project{Projectid: uid}
	err := o.Read(project)

	return project, err
}

func GetAllProject() []*Project {
	var projects []*Project
	o := orm.NewOrm()
	_, err := o.QueryTable(`project`).All(&projects)
	if err != nil {
		panic(err)
	}

	return projects
}

func UpdateProject(uid string, uu *Project) (a *Project, err error) {
	/*
	if u, ok := ProjectList[uid]; ok {
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Email != "" {
			u.Email = uu.Email
		}
		return u, nil
	}*/
	return nil, errors.New("Project Not Exist")
}
