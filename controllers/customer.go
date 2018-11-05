package controllers

import (
	"errors"
	"test/models"
	"test/utils"
	"time"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//CustomerController ..
type CustomerController struct {
	BaseController
}

//Get ..
func (c *CustomerController) Get() {
	c.TplName = "iframes/customer.tpl"
}

//GetCustomerDetailPage ..
func (c *CustomerController) GetCustomerDetailPage() {
	c.TplName = "iframes/customerDetail.tpl"
}

//GetDetailPage ..
func (c *CustomerController) GetDetailPage() {
	item := c.GetItem()
	if item == nil {
		c.Redirect("/", 404)
		return
	}
	id := c.getIDFromFormData()
	mp := models.GetCustomerMapper("")
	mp.Tx(func(tx *sqlx.Tx) (r error) {
		defer func() {
			if err := recover(); err != nil {
				msg := utils.InterfaceToStr(err)
				r = errors.New(msg)
			}
		}()
		sqlStr := "update " + mp.TableName + " set clickNumber=(clickNumber+1) where id=?"
		_, err := gosql.WithTx(tx).Exec(sqlStr, id)
		if err != nil {
			panic(err)
		}
		return nil
	})
	obj := item.(*models.Customer)
	c.Data["title"] = obj.Title
	c.Data["clickNumber"] = obj.ClickNumber + 1
	c.Data["createTime"] = obj.CreateTime.Format("2006-01-02")
	c.Data["htmlContent"] = obj.HTMLContent

	c.isLogin()
	c.TplName = "customer/customerDetail.tpl"
}

//GetCustomerPage ..
func (c *CustomerController) GetCustomerPage() {
	c.isLogin()
	c.TplName = "customer/customer.tpl"
}

//Add ..
func (c CustomerController) Add() string {
	item := models.Customer{}
	c.getItem(&item, true)
	item.TitlePinYin1, item.TitlePinYin2 = utils.ToPinYin1(item.Title)
	item.CreateTime = time.Now()
	item.UpdateTime = time.Now()
	mp := models.GetCustomerMapper("")
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
			panic("Insert未成功")
		}
		return nil
	})

	return item.ID
}

//GetItem ..
func (c CustomerController) GetItem() interface{} {
	id := c.getIDFromFormData()
	mp := models.GetCustomerMapper("")
	item := mp.Get(nil, id)
	return item
}

//GetList ..
func (c CustomerController) GetList() interface{} {
	pageIndex := c.getPageIndex()
	rowsInPage := c.getRowsInPage()
	sort := c.getSort()
	mp := models.GetCustomerMapper("")
	r := mp.GetList(pageIndex, rowsInPage, sort)
	return r
}

//Update ..
func (c CustomerController) Update() string {
	item := models.Customer{}
	c.getItem(&item, false)
	item.UpdateTime = time.Now()
	item.TitlePinYin1, item.TitlePinYin2 = utils.ToPinYin1(item.Title)
	mp := models.GetCustomerMapper("")
	mp.Tx(func(tx *sqlx.Tx) (r error) {
		defer func() {
			if err := recover(); err != nil {
				msg := utils.InterfaceToStr(err)
				r = errors.New(msg)
			}
		}()

		// c.checkData(mp, tx, &item)
		count := mp.Update(tx, &item)
		if count == 0 {
			panic("update未成功")
		}
		return nil
	})
	return item.ID
}

//Delete ..
func (c CustomerController) Delete() {
	id := c.getIDFromFormData()
	mp := models.GetCustomerMapper("")
	mp.Delete(nil, id)
}

func (c CustomerController) checkData(mp models.CustomerMapper, tx *sqlx.Tx, item *models.Customer) {
	// count := mp.GetCount(nil, " type=? ", item.Type)
	// if count > 0 {
	// 	panic("每个类型的简介仅可有一篇文章，只能修改无法新增")
	// }
}

//APIhandler ..
func (c CustomerController) APIhandler() {
	c.handler(c)
}
