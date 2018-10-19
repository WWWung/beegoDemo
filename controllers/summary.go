package controllers

import (
	"errors"
	"test/models"
	"test/utils"
	"time"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//SummaryController ..
type SummaryController struct {
	BaseController
}

//GetSummaryPage ..
func (c *SummaryController) GetSummaryPage() {
	c.isLogin()
	c.TplName = "summary/summary.tpl"
}

//Get ..
func (c *SummaryController) Get() {
	c.TplName = "iframes/summary.tpl"
}

//GetSummaryDetailPage ..
func (c *SummaryController) GetSummaryDetailPage() {
	c.TplName = "iframes/summaryDetail.tpl"
}

//Add ..
func (c SummaryController) Add() string {
	item := models.Summary{}
	c.getItem(&item, true)
	item.Time = time.Now()
	mp := models.GetSummaryMapper("")
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
func (c SummaryController) GetItem() interface{} {
	id := c.getIDFromFormData()
	mp := models.GetSummaryMapper("")
	item := mp.Get(nil, id)
	return item
}

//GetList ..
func (c SummaryController) GetList() interface{} {
	pageIndex := c.getPageIndex()
	rowsInPage := c.getRowsInPage()
	sort := c.getSort()
	_type := c.GetString("type")
	whereStr := ""
	mp := models.GetSummaryMapper("")
	args := make([]interface{}, 0)
	if _type != "" {
		whereStr = " where type=?"
		args = append(args, _type)
	}
	update := c.GetString("update")
	if update == "true" {
		mp.Tx(func(tx *sqlx.Tx) (r error) {
			defer func() {
				if err := recover(); err != nil {
					msg := utils.InterfaceToStr(err)
					r = errors.New(msg)
				}
			}()
			sqlStr := "update " + mp.TableName + " set clickNumber=(clickNumber+1) where type=?"
			_, err := gosql.WithTx(tx).Exec(sqlStr, args...)
			if err != nil {
				panic(err)
			}
			return nil
		})
	}

	r := mp.GetList(pageIndex, rowsInPage, whereStr, sort, args...)
	return r
}

//Update ..
func (c SummaryController) Update() string {
	item := models.Summary{}
	c.getItem(&item, false)
	item.Time = time.Now()
	mp := models.GetSummaryMapper("")
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
func (c SummaryController) Delete() {
	id := c.getIDFromFormData()
	mp := models.GetSummaryMapper("")
	mp.Delete(nil, id)
}

func (c SummaryController) checkData(mp models.SummaryMapper, tx *sqlx.Tx, item *models.Summary) {
	count := mp.GetCount(nil, " type=? ", item.Type)
	if count > 0 {
		panic("每个类型的简介仅可有一篇文章，只能修改无法新增")
	}
}

//APIhandler ..
func (c SummaryController) APIhandler() {
	c.handler(c)
}
