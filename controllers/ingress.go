package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/yaoice/ice/models"
)

// Operations about ingress
type IngressesController struct {
	beego.Controller
}

// @Title Create
// @Description create new ingress
// @Param   project     path    string  true        "The project you want to get"
// @Param   body        body    models.Ingress  true    "The ingress content"
// @Success 200 {app} models.Ingress
// @Failure 400 :body is empty
// @router / [post]
func (this *IngressesController) Post() {
	var ingress models.Ingress
	project := this.GetString(":project")
	json.Unmarshal(this.Ctx.Input.RequestBody, &ingress)
	result, err := models.CreateIngress(project, &ingress)
	if err != nil {
		logs.Error("create ingress %s response err: %v", ingress.AppName, err)
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = err
	} else {
		this.Data["json"] = result
	}
	this.ServeJSON()
}

// @Title Get
// @Description get ingress by name
// @Param   project     path    string  true        "The project you want to get"
// @Param   ingressName     path    string  true        "The ingress you want to get"
// @Success 200 {ingress} models.Ingress
// @Failure 400 :ingressName is empty
// @router /:ingressName [get]
func (this *IngressesController) Get() {
	project := this.GetString(":project")
	ingressName := this.GetString(":ingressName")
	ingress, err := models.GetIngress(project, ingressName)
	if err != nil {
		logs.Error("get ingress %s response err: %v", ingressName, err)
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = err
	} else {
		this.Data["json"] = ingress
	}
	this.ServeJSON()
}

// @Title Delete
// @Description delete ingress by name
// @Param   project     path    string  true        "The project you want to get"
// @Param   ingressName     path    string  true        "The ingress you want to delete"
// @Success 200 {ingress} models.Ingress
// @Failure 400 :ingressName is empty
// @router /:ingressName [delete]
func (this *IngressesController) Delete() {
	project := this.GetString(":project")
	ingressName := this.GetString(":ingressName")
	err := models.DeleteIngress(project, ingressName)
	if err != nil {
		logs.Error("delete ingress %s response err: %v", ingressName, err)
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = err
	} else {
		s := fmt.Sprintf("delete ingress %s success!", ingressName)
		this.Data["json"] = s
	}
	this.ServeJSON()
}
