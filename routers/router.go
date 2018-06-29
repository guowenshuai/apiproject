// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/guowenshuai/apiproject/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	beego.Info("running router init...")


	// JobController
	beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:JobController"] = append(beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:JobController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:JobController"] = append(beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:JobController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:jobId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:JobController"] = append(beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:JobController"],
		beego.ControllerComments{
			Method:           "ByBoss",
			Router:           `/ByBoss`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:JobController"] = append(beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:JobController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:JobController"] = append(beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:JobController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:jobId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:JobController"] = append(beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:JobController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:jobId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	// WebsocketController
	beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:WebSocketController"] = append(beego.GlobalControllerRouter["github.com/guowenshuai/apiproject/controllers:WebSocketController"],
		beego.ControllerComments{
			Method:           "Upgrade",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.Info("running router init over...")

}

func init() {

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/job",
			beego.NSInclude(
				&controllers.JobController{},
			)),
		beego.NSNamespace("/ws",
			beego.NSInclude(
				&controllers.WebSocketController{},
			)),
	)
	beego.AddNamespace(ns)
}