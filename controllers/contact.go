package controllers

import (
	"errors"
	"test/models"
	"test/utils"
	"time"

	"github.com/jmoiron/sqlx"
)

//ContactController ..
type ContactController struct {
	BaseController
}

//Get ..
func (c *ContactController) Get() {
	c.isLogin()
	c.TplName = "contact/contact.tpl"
}

//GetAdminContactPage ..
func (c *ContactController) GetAdminContactPage() {
	c.TplName = "iframes/contact.tpl"
}

//GetList ..
func (c ContactController) GetList() interface{} {
	pageIndex := c.getPageIndex()
	rowsInPage := c.getRowsInPage()
	// searchkey := c.getStringFromPath("searchKey")
	// var sqlArgs = []interface{}{}
	// var whereStr string
	// if searchkey != "" {
	// 	whereStr = " where type = ? "
	// 	sqlArgs = append(sqlArgs, searchkey)
	// }
	mp := models.GetWordMapper("")
	var sort = " time desc"
	r := mp.GetList(pageIndex, rowsInPage, sort)
	return r
}

//Add ..
func (c ContactController) Add() string {
	var p = c.CheckPower()
	if p == 0 {
		panic("未登录")
	}
	user := c.GetSession("user").(*models.User)
	item := models.Word{}
	c.getItem(&item, true)
	item.UserID = user.ID
	item.UserName = user.Name
	item.Time = time.Now()
	mp := models.GetWordMapper("")
	mp.Tx(func(tx *sqlx.Tx) (r error) {
		defer func() {
			if err := recover(); err != nil {
				msg := utils.InterfaceToStr(err)
				r = errors.New(msg)
			}
		}()

		// c.checkData(mp, tx, &item)
		_, count := mp.Insert(tx, &item)
		if count == 0 {
			panic("Insert未成功")
		}
		return nil
	})

	return item.ID
}

//Delete ..
func (c ContactController) Delete() {
	id := c.getIDFromFormData()
	mp := models.GetWordMapper("")
	mp.Delete(nil, id)
}

//APIhandler ..
func (c ContactController) APIhandler() {
	c.handler(c)
}
