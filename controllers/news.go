package controllers

import (
	"errors"
	"test/models"
	"test/utils"
	"time"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//NewsController ..
type NewsController struct {
	BaseController
}

//Get ..
func (c *NewsController) Get() {
	c.TplName = "iframes/news.tpl"
}

//GetNewsDetailPage ..
func (c *NewsController) GetNewsDetailPage() {
	c.TplName = "iframes/newsDetail.tpl"
}

//GetDetailPage ..
func (c *NewsController) GetDetailPage() {
	item := c.GetItem()
	if item == nil {
		c.Redirect("/", 404)
		return
	}
	id := c.getIDFromFormData()
	mp := models.GetNewsMapper("")
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
	obj := item.(*models.News)
	c.Data["title"] = obj.Title
	c.Data["clickNumber"] = obj.ClickNumber + 1
	c.Data["createTime"] = obj.CreateTime.Format("2006-01-02")
	c.Data["htmlContent"] = obj.HTMLContent

	c.isLogin()
	c.TplName = "news/newsDetail.tpl"
}

//GetNewsPage ..
func (c *NewsController) GetNewsPage() {
	c.isLogin()
	c.TplName = "news/news.tpl"
}

//Add ..
func (c NewsController) Add() string {
	item := models.News{}
	c.getItem(&item, true)
	item.TitlePinYin1, item.TitlePinYin2 = utils.ToPinYin1(item.Title)
	item.CreateTime = time.Now()
	item.UpdateTime = time.Now()
	mp := models.GetNewsMapper("")
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
func (c NewsController) GetItem() interface{} {
	id := c.getIDFromFormData()
	mp := models.GetNewsMapper("")
	item := mp.Get(nil, id)
	return item
}

//GetList ..
func (c NewsController) GetList() interface{} {
	pageIndex := c.getPageIndex()
	rowsInPage := c.getRowsInPage()
	sort := c.getSort()
	mp := models.GetNewsMapper("")
	r := mp.GetList(pageIndex, rowsInPage, sort)
	return r
}

//Update ..
func (c NewsController) Update() string {
	item := models.News{}
	c.getItem(&item, false)
	item.UpdateTime = time.Now()
	item.TitlePinYin1, item.TitlePinYin2 = utils.ToPinYin1(item.Title)
	mp := models.GetNewsMapper("")
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
func (c NewsController) Delete() {
	id := c.getIDFromFormData()
	mp := models.GetNewsMapper("")
	mp.Delete(nil, id)
}

func (c NewsController) checkData(mp models.NewsMapper, tx *sqlx.Tx, item *models.News) {
	// count := mp.GetCount(nil, " type=? ", item.Type)
	// if count > 0 {
	// 	panic("每个类型的简介仅可有一篇文章，只能修改无法新增")
	// }
}

//APIhandler ..
func (c NewsController) APIhandler() {
	c.handler(c)
}
