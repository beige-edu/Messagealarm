package auth

import (
	beegojwt "PrometheusAlert/pkg/util"
	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

func (c *Controller) Prepare() {
	tokenString := c.Ctx.Request.Header.Get("Authorization")
	et := beegojwt.EasyToken{}
	valido, _, _ := et.ValidateToken(tokenString)
	if !valido {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "Permission Deny"
		c.ServeJSON()
	}
	return
}
