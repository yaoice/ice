package controllers

import (
	"github.com/yaoice/ice/models"
	"encoding/json"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "fmt"
)

// Operations about app
type AppsController struct {
	beego.Controller
}

// @Title Create
// @Description create new app
// @Param   project     path    string  true        "The project you want to get"
// @Param   body        body    models.App  true    "The app content"
// @Success 200 {app} models.App.AppName
// @Failure 400 :body is empty
// @router / [post]
func (this *AppsController) Post() {
    var app models.App
    project := this.GetString(":project")
    json.Unmarshal(this.Ctx.Input.RequestBody, &app)
    result, err := models.CreateApp(project, &app)
    if err != nil {
        logs.Error("create app %s response err: %v", app.AppName, err)
        this.Ctx.Output.SetStatus(400)
        this.Data["json"] = err
    } else {
        this.Data["json"] = result
    }
    this.ServeJSON()
}

// @Title Get
// @Description get app by appname
// @Param   project     path    string  true        "The project you want to get"
// @Param   appName     path    string  true        "The appname you want to get"
// @Success 200 {app} models.App
// @Failure 400 :appName is empty
// @router /:appName [get]
func (this *AppsController) Get() {
    project := this.GetString(":project")
    appName := this.GetString(":appName")
    app, err := models.GetApp(project, appName)
    if err != nil {
        logs.Error("get app %s response err: %v", appName, err)
        this.Ctx.Output.SetStatus(400)
        this.Data["json"] = err
    } else {
        this.Data["json"] = app
    }
    this.ServeJSON()
}

// @Title GetAll
// @Description get all apps
// @Param   project     path    string  true        "The project you want to get"
// @Success 200 {app} models.App
// @Failure 400 :appName is empty
// @router / [get]
func (this *AppsController) GetAll() {
    project := this.GetString(":project")
	apps, err := models.GetAllApp(project)
    if err != nil {
        logs.Error("get all app response err: %v", err)
        this.Data["json"] = err
    } else {
	    this.Data["json"] = apps
    }
	this.ServeJSON()
}

/*
// @Title Update
// @Description update app by appname
// @Param   project     path    string  true        "The project you want to get"
// @Param   appName     path    string  true        "The app you want to update"
// @Param   body        body    models.App   true        "The body"
// @Success 200 {app} models.App
// @Failure 400 :appName is empty
// @router /:appName [put]
func (this *AppsController) Put() {

}
*/

// @Title Delete
// @Description delete app by appname
// @Param   project     path    string  true        "The project you want to get"
// @Param   appName     path    string  true        "The app you want to delete"
// @Success 200 {app} models.App
// @Failure 400 :appName is empty
// @router /:appName [delete]
func (this *AppsController) Delete() {
    project := this.GetString(":project")
    appName := this.GetString(":appName")
    err := models.DeleteApp(project, appName)
    if err != nil {
        logs.Error("delete app %s response err: %v", appName, err)
        this.Ctx.Output.SetStatus(400)
        this.Data["json"] = err
    } else {
        s := fmt.Sprintf("delete app %s success!", appName)
        this.Data["json"] = s
    }
    this.ServeJSON()
}
