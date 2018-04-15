package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/yaoice/ice/models"
)

// Operations about service
type ServicesController struct {
	beego.Controller
}

// @Title Create
// @Description create new service
// @Param   project     path    string  true        "The project you want to get"
// @Param   body        body    models.Service  true    "The service content"
// @Success 200 {service} models.Service
// @Failure 400 :body is empty
// @router / [post]
func (this *ServicesController) Post() {
	var service models.Service
	project := this.GetString(":project")
	json.Unmarshal(this.Ctx.Input.RequestBody, &service)
	result, err := models.CreateService(project, &service)
	if err != nil {
		logs.Error("create service %s response err: %v", service.AppName, err)
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = err
	} else {
		this.Data["json"] = result
	}
	this.ServeJSON()
}

// @Title Get
// @Description get service by name
// @Param   project     path    string  true        "The project you want to get"
// @Param   serviceName     path    string  true        "The service you want to get"
// @Success 200 {service} models.Service
// @Failure 400 :serviceName is empty
// @router /:serviceName [get]
func (this *ServicesController) Get() {
	project := this.GetString(":project")
	serviceName := this.GetString(":serviceName")
	service, err := models.GetService(project, serviceName)
	if err != nil {
		logs.Error("get service %s response err: %v", serviceName, err)
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = err
	} else {
		this.Data["json"] = service
	}
	this.ServeJSON()
}

// @Title Delete
// @Description delete service by name
// @Param   project     path    string  true        "The project you want to get"
// @Param   serviceName     path    string  true        "The service you want to delete"
// @Success 200 {service} models.Service
// @Failure 400 :serviceName is empty
// @router /:serviceName [delete]
func (this *ServicesController) Delete() {
	project := this.GetString(":project")
	serviceName := this.GetString(":serviceName")
	err := models.DeleteService(project, serviceName)
	if err != nil {
		logs.Error("delete service %s response err: %v", serviceName, err)
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = err
	} else {
		s := fmt.Sprintf("delete service %s success!", serviceName)
		this.Data["json"] = s
	}
	this.ServeJSON()
}
