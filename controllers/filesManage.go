package controllers

import (
	"errors"
	"io"
	"os"
	"test/models"
	"test/throw"
	"test/utils"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//FilesManageController ..
type FilesManageController struct {
	BaseController
}

//Get ..
func (c *FilesManageController) Get() {
	c.TplName = "iframes/files.tpl"
}

//GetFileDetailPage ..
func (c *FilesManageController) GetFileDetailPage() {
	c.TplName = "iframes/fileDetail.tpl"
}

//GetList ..
func (c FilesManageController) GetList() {
	pageIndex := c.getPageIndex()
	rowsInPage := c.getRowsInPage()
	order := c.getStringFromPath("sort")
	searchkey := c.getStringFromPath("searchKey")
	var sqlArgs = []interface{}{}
	var whereStr string
	if searchkey != "" {
		whereStr = " where type = ? "
		sqlArgs = append(sqlArgs, searchkey)
	}
	mp := models.GetFilesManageMapper("")
	r := mp.GetList(pageIndex, rowsInPage, whereStr, order, sqlArgs...)
	c.success(r)
}

//Delete ..
func (c FilesManageController) Delete() {
	id := c.getIDFromFormData()
	mp := models.GetFilesManageMapper("")
	mp.Delete(nil, id)
}

func (c FilesManageController) checkData(mp models.FilesManageMapper, tx *sqlx.Tx, item *models.FilesManage) {
	if item.ID == "" {
		panic("id不能为空")
	}
	if item.Name == "" {
		panic("名字不能为空")
	}
}

//Update ..
func (c FilesManageController) Update() {
	item := models.FilesManage{}
	c.getItem(&item, false)
	item.NamePinYin1, item.NamePinYin2 = utils.ToPinYin1(item.Name)
	mp := models.GetFilesManageMapper("")
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
func (c FilesManageController) GetItem() {
	id := c.getIDFromFormData()
	mp := models.GetFilesManageMapper("")
	item := mp.Get(nil, id)
	c.success(item)
}

//Add ..
func (c FilesManageController) Add() string {
	item := models.FilesManage{}
	c.getItem(&item, true)
	item.NamePinYin1, item.NamePinYin2 = utils.ToPinYin1(item.Name)
	mp := models.GetFilesManageMapper("")
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

//Download ..
func (c FilesManageController) Download() {
	_url := c.getStringFromPath("url")
	mp := models.GetFilesManageMapper("")
	mp.Tx(func(tx *sqlx.Tx) (r error) {
		defer func() {
			if err := recover(); err != nil {
				msg := utils.InterfaceToStr(err)
				r = errors.New(msg)
			}
		}()
		sqlStr := "update filesManage set downloadNumber=(downloadNumber+1) where url=?"
		_, err := gosql.WithTx(tx).Exec(sqlStr, _url)
		if err != nil {
			panic(err)
		}
		return nil
	})
	_type := c.getStringFromPath("type")
	name := c.getStringFromPath("name")
	dName := c.getStringFromPath("dName")
	suffix := c.getStringFromPath("suffix")
	if _type == "" || name == "" {
		panic("格式错误")
	}
	addr := "assets/" + _type + "/" + name
	file, err := os.Open(addr)
	throw.CheckErr(err)
	defer file.Close()
	w := c.Ctx.ResponseWriter
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("content-disposition", "attachment; filename="+dName+"."+suffix)
	_, err = io.Copy(w, file)
	throw.CheckErr(err)
}

//Preview ..
func (c FilesManageController) Preview() {
	// addr := "assets/" + "pdfs" + "/" + "e1676318c81cd46d56233730f7ab5b66.pdf"
	// file, err := os.Open(addr)
	// throw.CheckErr(err)
	// defer file.Close()
	// w := c.Ctx.ResponseWriter
	// _, err = io.Copy(w, file)
	// throw.CheckErr(err)
}

//APIhandler ..
func (c FilesManageController) APIhandler() {
	c.handler(c)
}
