package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "Falcon Next Using beego.me"
	c.Data["Email"] = "afalak AT gmail.com"
	c.TplName = "index.tpl"
}
