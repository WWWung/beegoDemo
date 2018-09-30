package controllers

import (
	"test/models"
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
