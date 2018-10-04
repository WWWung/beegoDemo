package controllers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"test/throw"
	"test/utils"

	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
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
	e := utils.InterfaceToStr(err)
	c.Data["json"] = map[string]interface{}{
		"code": 1,
		"data": e,
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

func (c *BaseController) getIDFromFormData() (id interface{}) {
	ids := c.Input()["data[id]"]
	if len(ids) > 0 {
		id = ids[0]
	}
	return
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

func (c *BaseController) getMethod() (m string) {
	method := c.getDataFromPath("m")
	if method != nil {
		m = method.(string)
	}
	return
}

func (c *BaseController) getItem(item interface{}, setID bool) {
	jsonStr := c.getDataFromPath("data")
	if jsonStr == nil {
		panic("数据格式错误")
	}
	j := jsonStr.(string)
	err := json.Unmarshal([]byte(j), item)
	if setID {
		uid, err := uuid.NewV4()
		throw.CheckErr(err)
		id := uid.String()
		v := reflect.ValueOf(item).Elem()
		v.FieldByName("ID").Set(reflect.ValueOf(id))
	}
	throw.CheckErr(err)
}

func (c *BaseController) apply(i interface{}, methodName string) []reflect.Value {
	obj := reflect.ValueOf(i)
	params := make([]reflect.Value, 0)
	m := obj.MethodByName(methodName)
	fmt.Println("apply 方法" + methodName)
	if !m.IsValid() {
		panic("方法" + methodName + "不存在")
	}
	r := m.Call(params)
	return r
}

func (c BaseController) handler(i interface{}) {
	defer func() {
		if err := recover(); err != nil {
			c.failure(err)
		}
	}()
	m := c.getMethod()
	r := c.apply(i, m)
	if len(r) == 0 {
		c.success(nil)
	} else {
		c.success(r[0].Interface())
	}
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
