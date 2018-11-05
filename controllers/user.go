package controllers

import (
	"errors"
	"fmt"
	"test/models"
	"test/utils"
	"time"

	"github.com/dchest/captcha"
	"github.com/jmoiron/sqlx"
)

//UserController ..
type UserController struct {
	BaseController
}

//GetRegisterPage ..
func (c *UserController) GetRegisterPage() {
	c.isLogin()
	CaptchaID := captcha.New()
	c.Data["CaptchaId"] = CaptchaID
	c.TplName = "user/register.tpl"
}

//GetLoginPage ..
func (c *UserController) GetLoginPage() {
	c.isLogin()
	CaptchaID := captcha.New()
	c.Data["CaptchaId"] = CaptchaID
	c.TplName = "user/login.tpl"
}

//GetAdminUsersPage ..
func (c *UserController) GetAdminUsersPage() {
	c.TplName = "iframes/users.tpl"
}

//Add ..
func (c UserController) Add() string {
	verify := c.verifyCaptcha()
	if !verify {
		c.failure("验证码错误")
		return ""
	}
	item := models.User{}
	c.getItem(&item, true)
	item.NamePinYin1, item.NamePinYin2 = utils.ToPinYin1(item.Name)
	password := c.GetString("password")
	item.PwMD5 = utils.Encrypt(password)
	item.CreateTime = time.Now()
	item.LastLoginTime = time.Now()
	item.Power = 2
	mp := models.GetUserMapper("")
	mp.Tx(func(tx *sqlx.Tx) (r error) {
		defer func() {
			if err := recover(); err != nil {
				msg := utils.InterfaceToStr(err)
				r = errors.New(msg)
			}
		}()

		c.checkData(mp, tx, &item)
		_, count := mp.Insert(tx, &item)
		if count == 0 {
			panic("注册失败")
		}
		return nil
	})
	c.SetSession("user", item)
	// c.Redirect("/", 302)
	return item.ID
}

//Delete ..
func (c UserController) Delete() {
	id := c.getIDFromFormData()
	mp := models.GetUserMapper("")
	mp.Delete(nil, id)
}

//Update ..
func (c UserController) Update() {
	item := models.User{}
	c.getItem(&item, false)
	mp := models.GetUserMapper("")
	mp.Tx(func(tx *sqlx.Tx) (r error) {
		defer func() {
			if err := recover(); err != nil {
				msg := utils.InterfaceToStr(err)
				r = errors.New(msg)
			}
		}()

		count := mp.Update(tx, &item)
		if count == 0 {
			panic("Insert未成功")
		}
		return nil
	})
}

//更新最后一次登录时间
func (c UserController) updateLastLoginTime(mp models.UserMapper, item *models.User) {
	item.LastLoginTime = time.Now()
	mp.Tx(func(tx *sqlx.Tx) (r error) {
		defer func() {
			if err := recover(); err != nil {
				msg := utils.InterfaceToStr(err)
				r = errors.New(msg)
			}
		}()

		// c.checkData(mp, tx, item)
		count := mp.Update(tx, item)
		if count == 0 {
			fmt.Println(time.Now(), item.Name, "更新最后一次登录时间失败")
		}
		return nil
	})
}

//Login ..
func (c UserController) Login() (s string) {
	verify := c.verifyCaptcha()
	if !verify {
		panic("验证码错误")
	}
	if c.GetSession("user") != nil {
		panic("请勿重复登录")
	}
	mp := models.GetUserMapper("")
	whereStr := " where name=? and pwMD5=?"
	args := make([]interface{}, 0)
	name := c.GetString("name")
	args = append(args, name)
	password := c.GetString("password")
	pwMD5 := utils.Encrypt(password)
	args = append(args, pwMD5)
	r := mp.Get(nil, whereStr, args...)
	if r != nil {
		c.SetSession("user", r)
		c.updateLastLoginTime(mp, r.(*models.User))
		s = "登录成功"
	} else {
		panic("账号或密码错误")
	}
	return
}

//Logout ..
func (c UserController) Logout() string {
	c.DelSession("user")
	return "退出成功"
}

//GetItem ..
func (c UserController) GetItem() {
	name := c.getIDFromFormData()
	mp := models.GetUserMapper("")
	whereStr := " where name=?"
	args := make([]interface{}, 0)
	args = append(args, name)
	item := mp.Get(nil, whereStr, args...)
	c.success(item)
}

//GetList ..
func (c UserController) GetList() {
	pageIndex := c.getPageIndex()
	rowsInPage := c.getRowsInPage()
	order := c.getStringFromPath("sort")
	mp := models.GetUserMapper("")
	r := mp.GetList(pageIndex, rowsInPage, order)
	c.success(r)
}
func (c UserController) checkData(mp models.UserMapper, tx *sqlx.Tx, item *models.User) {
	if item.Name == "" {
		panic("请输入账号")
	}
	if item.Phone == "" {
		panic("请输入电话")
	}
	if item.Addr == "" {
		panic("请输入地址")
	}
	whereStr := " name=? "
	count := mp.GetCount(nil, whereStr, item.Name)
	if count > 0 {
		panic("账号已存在")
	}
}

//APIhandler ..
func (c UserController) APIhandler() {
	c.handler(c)
}
