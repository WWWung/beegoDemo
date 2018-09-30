package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

//CaptchaController ..
type CaptchaController struct {
	beego.Controller
}

func (c *CaptchaController) getCaptcha() string {
	captchaID := captcha.New()
	return captchaID
}

//ReloadCaptcha ..
func (c *CaptchaController) ReloadCaptcha() {
	CaptchaID := captcha.New()
	c.Data["json"] = map[string]string{
		"captchaID": CaptchaID,
	}
	c.ServeJSON()
}

func (c *CaptchaController) verifyCaptcha() bool {
	captchaID := c.GetString("captchaId")
	captchaValue := c.GetString("captcha")
	return captcha.VerifyString(captchaID, captchaValue)
}
