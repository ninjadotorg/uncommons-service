package routers

import (
	"github.com/ninjadotorg/uncommons-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/app-proposal",
			beego.NSInclude(
				&controllers.AppProposalController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
