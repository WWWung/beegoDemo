package controllers

//DownloadCenterController ..
type DownloadCenterController struct {
	BaseController
}

//Get ..
func (c *DownloadCenterController) Get() {
	c.isLogin()
	c.TplName = "download/download.tpl"
}

//GetPriviewPage ..
func (c *DownloadCenterController) GetPriviewPage() {
	c.TplName = "download/preview2.tpl"
}

//GetPriviewSettingPage ..
func (c *DownloadCenterController) GetPriviewSettingPage() {
	c.TplName = "iframes/preSetting.tpl"
}
