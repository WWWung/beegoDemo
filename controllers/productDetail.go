package controllers

//ProductsDetailController ..
type ProductsDetailController struct {
	BaseController
}

//Get ..
func (c *ProductsDetailController) Get() {
	c.TplName = "iframes/productDetail.tpl"
}
