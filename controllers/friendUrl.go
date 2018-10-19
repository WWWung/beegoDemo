package controllers

import (
	"errors"
	"test/models"
	"test/utils"

	"github.com/jmoiron/sqlx"
)

//FriendURLController ..
type FriendURLController struct {
	BaseController
}

//Get ..
func (c *FriendURLController) Get() {
	c.TplName = "iframes/friendUrl.tpl"
}

//GetFriendURLDetailPage ..
func (c *FriendURLController) GetFriendURLDetailPage() {
	c.TplName = "iframes/friendUrlDetail.tpl"
}

//Add ..
func (c FriendURLController) Add() string {
	item := models.FriendURL{}
	c.getItem(&item, true)
	mp := models.GetFriendURLMapper("")
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
func (c FriendURLController) GetItem() interface{} {
	id := c.getIDFromFormData()
	mp := models.GetFriendURLMapper("")
	item := mp.Get(nil, id)
	return item
}

//GetList ..
func (c FriendURLController) GetList() interface{} {
	pageIndex := c.getPageIndex()
	rowsInPage := c.getRowsInPage()
	sort := c.getSort()
	mp := models.GetFriendURLMapper("")
	r := mp.GetList(pageIndex, rowsInPage, sort)
	return r
}

//Update ..
func (c FriendURLController) Update() string {
	item := models.FriendURL{}
	c.getItem(&item, false)
	mp := models.GetFriendURLMapper("")
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
func (c FriendURLController) Delete() {
	id := c.getIDFromFormData()
	mp := models.GetFriendURLMapper("")
	mp.Delete(nil, id)
}

func (c FriendURLController) checkData(mp models.FriendURLMapper, tx *sqlx.Tx, item *models.FriendURL) {
	// count := mp.GetCount(nil, " type=? ", item.Type)
	// if count > 0 {
	// 	panic("每个类型的简介仅可有一篇文章，只能修改无法新增")
	// }
}

//APIhandler ..
func (c FriendURLController) APIhandler() {
	c.handler(c)
}
