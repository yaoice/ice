package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:AppsController"] = append(beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:AppsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:AppsController"] = append(beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:AppsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:AppsController"] = append(beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:AppsController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:appName`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:AppsController"] = append(beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:AppsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:appName`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:IngressesController"] = append(beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:IngressesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:IngressesController"] = append(beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:IngressesController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:ingressName`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:IngressesController"] = append(beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:IngressesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:ingressName`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:ServicesController"] = append(beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:ServicesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:ServicesController"] = append(beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:ServicesController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:serviceName`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:ServicesController"] = append(beego.GlobalControllerRouter["github.com/yaoice/ice/controllers:ServicesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:serviceName`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

}
