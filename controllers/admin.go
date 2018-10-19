package controllers

import (
	"test/models"

	"github.com/astaxie/beego"
)

//AdminController ..
type AdminController struct {
	beego.Controller
}

//Get ..
func (c *AdminController) Get() {
	v := c.GetSession("adminUser")
	if v == nil {
		c.Redirect("/admin/login", 302)
		return
	}
	realValue := v.(*models.AdminUser)
	c.Data["Name"] = realValue.Name
	c.TplName = "admin/admin.tpl"
}
