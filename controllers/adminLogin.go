package controllers

import (
	"test/models"
	"test/utils"

	"github.com/dchest/captcha"
)

//AdminLoginController ..
type AdminLoginController struct {
	BaseController
}

//Get ..
func (c *AdminLoginController) Get() {
	CaptchaID := captcha.New()
	c.Data["CaptchaId"] = CaptchaID
	c.TplName = "admin/login.tpl"
}

//Loginin ..
func (c *AdminLoginController) Loginin() {
	if !c.verifyCaptcha() {
		c.failure("验证码错误")
		return
	}
	mp := models.GetAdminUserMapper("")
	name := c.GetString("name")
	password := c.GetString("password")
	pwMD5 := utils.Encrypt(password)
	item := mp.Get(nil, name, pwMD5)
	if item == nil {
		c.failure("账号或者密码错误")
		return
	}
	c.SetSession("adminUser", item)
	data := map[string]interface{}{
		"name": item.(*models.AdminUser).Name,
	}
	c.success(data)
}
