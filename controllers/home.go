package controllers

//HomeController ..
type HomeController struct {
	BaseController
}

//Get ..
func (c *HomeController) Get() {
	// u := c.GetSession("user")
	// if u != nil {
	// 	user := u.(*models.User)
	// 	c.Data["tip"] = "您好，<strong>" + user.Name + "</strong>先生 <a href=\"javascript:;\" id=\"logout\">退出</a>"
	// } else {
	// 	c.Data["tip"] = "<a href=\"/register\">注册</a><a href=\"/login\">登录</a>"
	// }
	c.isLogin()
	c.TplName = "index.tpl"
}
