package controllers

import (
	"strconv"
	"test/throw"

	"github.com/astaxie/beego"
)

//BaseController ..
type BaseController struct {
	CaptchaController
}

// 返回数据格式：
// 	成功：
// 	{
// 		"code": 0,
// 		"data": data map[string]interface{} 返回的数据
// 	}
// 	失败:
// 	{
// 		"code": 1,
// 		"data": interface{} 失败描述语句
// 	}

//Success ..
func (c *BaseController) success(data interface{}) {
	c.Data["json"] = map[string]interface{}{
		"code": 0,
		"data": data,
	}
	c.ServeJSON()
}

//Failure ..
func (c *BaseController) failure(err interface{}) {
	c.Data["json"] = map[string]interface{}{
		"code": 1,
		"data": err,
	}
	c.ServeJSON()
}

func (c *BaseController) getDataFromPath(key string) interface{} {
	m := c.Input()
	if m[key] == nil {
		return nil
	}
	return m[key][0]
}

func (c *BaseController) getPageIndex() (i int) {
	defer func() {
		if err := recover(); err != nil {
			i = 0
		}
	}()
	var err error
	p := c.getDataFromPath("pageIndex")
	i, err = strconv.Atoi(p.(string))
	throw.CheckErr(err)
	return
}

func (c *BaseController) getRowsInPage() (i int) {
	defer func() {
		if err := recover(); err != nil {
			i = 0
		}
	}()
	var err error
	p := c.getDataFromPath("rowsInPage")
	i, err = strconv.Atoi(p.(string))
	throw.CheckErr(err)
	return
}

//===============================================

//MainController ..
type MainController struct {
	beego.Controller
}

//Get ..
func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

//ReturnName ..
func (c *MainController) ReturnName() {
	c.Ctx.Output.Body([]byte("wwwung"))
}
