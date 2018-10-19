package controllers

import (
	"fmt"
	"test/models"
	"test/throw"
	"test/utils"
	"time"

	"github.com/dchest/captcha"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

//AdminRegisController ..
type AdminRegisController struct {
	BaseController
}

//Get ..
func (c *AdminRegisController) Get() {
	CaptchaID := captcha.New()
	c.Data["CaptchaId"] = CaptchaID
	c.TplName = "admin/register.tpl"
}

//Regis ..
func (c *AdminRegisController) Regis() {
	verify := c.verifyCaptcha()
	if !verify {
		c.failure("验证码错误")
		return
	}
	mp := models.GetAdminUserMapper("")
	name := c.GetString("name")
	password := c.GetString("password")
	power := c.GetString("power")
	uid, err := uuid.NewV4()
	throw.CheckErr(err)
	item := models.AdminUser{}
	item.ID = uid.String()
	item.Name = name
	item.PwMD5 = utils.Encrypt(password)
	item.Power = power
	item.Time = time.Now()
	mp.Tx(func(tx *sqlx.Tx) (r error) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("err")
				c.failure(err)
			}
		}()
		c.checkData(item)
		_, count := mp.Insert(tx, &item)
		if count == 0 {
			panic("注册失败")
		}
		c.SetSession("adminUser", &item)
		data := map[string]interface{}{
			"name": name,
		}
		c.success(data)
		return nil
	})
}

func (c *AdminRegisController) checkData(item models.AdminUser) {
	mp := models.GetAdminUserMapper("")
	whereStr := "name=?"
	count := mp.GetCount(nil, whereStr, item.Name)
	if count > 0 {
		panic("账号已存在")
	}
}
