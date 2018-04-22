package controllers

import (
	"github.com/astaxie/beego"
	"hackdfw/models"
	"encoding/json"
	"strconv"
)

type ProjectController struct {
	beego.Controller
}

// @Title CreateProject
// @Description create projects
// @Param	body		body 	models.Project	true		"body for project content"
// @Success 200 {int} models.Project
// @Failure 403 body is empty
// @router / [post]
func (u *ProjectController) Post() {
	var user models.Project
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	err := models.AddProject(&user)
	if err != nil {
		// do something
	}

	u.Data["json"] = user
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Projects
// @Success 200 {object} models.Project
// @router / [get]
func (u *ProjectController) GetAll() {
	project := models.GetAllProject()
	u.Data["json"] = project
	u.ServeJSON()
}

// @Title Get
// @Description get project by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Project
// @Failure 403 :uid is empty
// @router /:id [get]
func (u *ProjectController) Get() {
	pid := u.GetString(":id")
	if pid != "" {
		id, err := strconv.Atoi(pid)
		if err != nil {
			// do something
		}
		project, err := models.GetProject(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = project
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the project
// @Param	pid		path 	string	true		"The pid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 pid is empty
// @router /:uid [delete]
func (u *ProjectController) Delete() {
	pid := u.GetString(":pid")
//	models.DeleteProject(pid)

	if pid != "" {
		id, err := strconv.Atoi(pid)
		if err != nil {
			// do something
		}
		ifsuccess := models.DeleteProject(id)
		if ifsuccess == true {
			u.Data["json"] = map[string]bool{ "success": true }
		} else {
			u.Data["json"] = map[string]bool{ "success": false }
		}

	}

	u.Data["json"] = "delete success!"
	u.ServeJSON()
}