package controllers

import (
	"github.com/astaxie/beego"
)

// AppProposalController ...
type AppProposalController struct {
	beego.Controller
}

// // Post ...
// func (o *AppProposalController) Post() {
// 	var ob models.AppProposal
// 	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
// 	appID := models.AddOne(ob)
// 	o.Data["json"] = map[string]int64{"ID": appID}
// 	o.ServeJSON()
// }

// // Get ...
// func (o *AppProposalController) Get() {
// 	appID := o.Ctx.Input.Param(":appId")
// 	if appID != "" {
// 		ob, err := models.GetOne(appID)
// 		if err != nil {
// 			o.Data["json"] = err.Error()
// 		} else {
// 			o.Data["json"] = ob
// 		}
// 	}
// 	o.ServeJSON()
// }

// // GetAll ...
// func (o *AppProposalController) GetAll() {
// 	obs := models.GetAll()
// 	o.Data["json"] = obs
// 	o.ServeJSON()
// }

// // Put ...
// func (o *AppProposalController) Put() {
// 	appID := o.Ctx.Input.Param(":appId")
// 	var ob models.Object
// 	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

// 	err := models.Update(appID, ob.Score)
// 	if err != nil {
// 		o.Data["json"] = err.Error()
// 	} else {
// 		o.Data["json"] = "update success!"
// 	}
// 	o.ServeJSON()
// }

// // Delete ...
// func (o *AppProposalController) Delete() {
// 	appID := o.Ctx.Input.Param(":appId")
// 	models.Delete(appID)
// 	o.Data["json"] = "delete success!"
// 	o.ServeJSON()
// }
