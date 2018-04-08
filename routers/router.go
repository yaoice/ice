// @APIVersion 1.0.0
// @Title Ice API
// @Description Ice is an k8s api proxy
// @Contact yao3690093@gmail.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/yaoice/ice/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/apis/v1",
       beego.NSNamespace("/projects",
            beego.NSInclude(
                &controllers.ProjectsController{},
            ),
        ),
        beego.NSNamespace("/projects/:project/apps",
			beego.NSInclude(
				&controllers.AppsController{},
			),
		),
        beego.NSNamespace("/projects/:project/ingresses",
            beego.NSInclude(
                &controllers.IngressesController{},
            ),
        ),
        beego.NSNamespace("/projects/:project/services",
            beego.NSInclude(
                &controllers.ServicesController{},
            ),
        ),
        beego.NSNamespace("/flavors",
            beego.NSInclude(
                &controllers.FlavorsController{},
            ),
        ),
        beego.NSNamespace("/images",
            beego.NSInclude(
                &controllers.ImagesController{},
            ),
        ),
	)
	beego.AddNamespace(ns)
}
