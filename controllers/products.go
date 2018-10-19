package controllers

import (
	"errors"
	"fmt"
	"test/models"
	"test/utils"
	"time"

	"github.com/jmoiron/sqlx"
)

//ProductsController ..
type ProductsController struct {
	BaseController
}

//Get ..
func (c *ProductsController) Get() {
	r := c.getList()
	rs := r.(map[string]interface{})
	c.Data["pageIndex"] = rs["pageIndex"]
	c.Data["total"] = rs["total"]
	c.Data["pageCount"] = rs["pageCount"]
	// fmt.Println(rs)
	c.isLogin()
	c.TplName = "products/products.tpl"
}

//GetProductDetailPage ..
func (c *ProductsController) GetProductDetailPage() {
	id := c.getDataFromPath("id")
	mp := models.GetProductsMapper("")
	item := mp.Get(nil, id)
	if item == nil {
		c.Redirect("/", 404)
		return
	}
	obj := item.(*models.Product)
	obj.ClickNumber++
	fmt.Println(obj.ClickNumber)
	mp.Tx(func(tx *sqlx.Tx) (r error) {
		defer func() {
			if err := recover(); err != nil {
				msg := utils.InterfaceToStr(err)
				r = errors.New(msg)
			}
		}()
		count := mp.Update(tx, obj)
		if count == 0 {
			panic("出现错误")
		}
		return nil
	})
	c.Data["title"] = obj.Title
	updateTime := obj.UpdateTime.Format("2006-01-02")
	c.Data["updateTime"] = updateTime
	c.Data["htmlContent"] = obj.HTMLContent
	c.Data["clickNumber"] = obj.ClickNumber
	c.Data["brand"] = obj.Brand
	c.isLogin()
	c.TplName = "products/productDetail.tpl"
}

//GetAdminPage ..
func (c *ProductsController) GetAdminPage() {
	c.TplName = "iframes/products.tpl"
}

//Add ..
func (c ProductsController) Add() string {
	item := models.Product{}
	c.getItem(&item, true)
	item.TitlePinYin1, item.TitlePinYin2 = utils.ToPinYin1(item.Title)
	item.CreateTime = time.Now()
	item.UpdateTime = time.Now()
	mp := models.GetProductsMapper("")
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

//Update ..
func (c ProductsController) Update() {
	item := models.Product{}
	c.getItem(&item, false)
	item.TitlePinYin1, item.TitlePinYin2 = utils.ToPinYin1(item.Title)
	item.UpdateTime = time.Now()
	mp := models.GetProductsMapper("")
	mp.Tx(func(tx *sqlx.Tx) (r error) {
		defer func() {
			if err := recover(); err != nil {
				msg := utils.InterfaceToStr(err)
				r = errors.New(msg)
			}
		}()

		c.checkData(mp, tx, &item)
		count := mp.Update(tx, &item)
		if count == 0 {
			panic("Insert未成功")
		}
		return nil
	})
}

//GetItem ..
func (c ProductsController) GetItem() {
	id := c.getIDFromFormData()
	mp := models.GetProductsMapper("")
	item := mp.Get(nil, id)
	c.success(item)
}

//	默认rowsInPage=10
func (c ProductsController) getList() (r interface{}) {
	pageIndex := c.getPageIndex()
	rowsInPage := c.getRowsInPage()
	if rowsInPage == 0 {
		rowsInPage = 10
	}
	mp := models.GetProductsMapper("")
	r = mp.GetList(pageIndex, rowsInPage, "")
	return
}

//GetList ..
func (c ProductsController) GetList() {
	pageIndex := c.getPageIndex()
	rowsInPage := c.getRowsInPage()
	sort := c.getSort()
	mp := models.GetProductsMapper("")
	r := mp.GetList(pageIndex, rowsInPage, sort)
	c.success(r)
}

//Delete ..
func (c ProductsController) Delete() {
	id := c.getIDFromFormData()
	mp := models.GetProductsMapper("")
	mp.Delete(nil, id)
}

func (c ProductsController) checkData(mp models.ProductsMapper, tx *sqlx.Tx, item *models.Product) {
	if item.ID == "" {
		panic("id不能为空")
	}
	if item.Title == "" {
		panic("标题不能为空")
	}
}

//APIhandler ..
func (c ProductsController) APIhandler() {
	c.handler(c)
}
