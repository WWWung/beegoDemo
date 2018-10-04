package controllers

import (
	"errors"
	"test/models"
	"test/utils"

	"github.com/jmoiron/sqlx"
)

//ProductsController ..
type ProductsController struct {
	BaseController
}

//Get ..
func (c *ProductsController) Get() {
	c.TplName = "iframes/products.tpl"
}

//ProductsList ..
func (c *ProductsController) ProductsList() {
	pageIndex := c.getPageIndex()
	rowsInPage := c.getRowsInPage()
	mp := models.GetProductsMapper("")
	r := mp.GetList(pageIndex, rowsInPage)
	c.success(r)
}

//Add ..
func (c ProductsController) Add() string {
	item := models.Product{}
	c.getItem(&item, true)
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
	// return "add"
}

//Update ..
func (c ProductsController) Update() {
	item := models.Product{}
	c.getItem(&item, false)
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
