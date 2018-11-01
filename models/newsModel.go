package models

import (
	"math"
	"test/throw"
	"time"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//News 简介
type News struct {
	ID           string    `json:"id"  db:"id"  `
	CreateTime   time.Time `json:"createTime"  db:"createTime"  `
	UpdateTime   time.Time `json:"updateTime"  db:"updateTime"  `
	HTMLContent  string    `json:"htmlContent" db:"htmlContent"`
	TextContent  string    `json:"textContent" db:"textContent"`
	Title        string    `json:"title" db:"title"`
	ClickNumber  int       `json:"clickNumber" db:"clickNumber"`
	Sort         int       `json:"sort" db:"sort"`
	TitlePinYin1 string    `json:"titlePinYin1" db:"titlePinYin1"`
	TitlePinYin2 string    `json:"titlePinYin2" db:"titlePinYin2"`
	Original     int       `json:"original" db:"original"`
}

//NewsMapper ..
type NewsMapper struct {
	BaseMapper
}

//GetNewsMapper ..
func GetNewsMapper(db string) NewsMapper {
	var mp NewsMapper
	mp.TableName = "news"
	if db == "" {
		db = "default"
	}
	mp.DB = gosql.Use(db)
	return mp
}

//Get ..
func (m NewsMapper) Get(tx *sqlx.Tx, id interface{}) interface{} {
	sqlStr := "select * from " + m.TableName + " where id = ? "
	var args = []interface{}{id}

	item := &News{}
	r := m.getItem(tx, item, sqlStr, args...)
	return r
}

//GetList ..
func (m NewsMapper) GetList(pageIndex int, rowsInPage int, sort string) (r interface{}) {
	//准备参数
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	item := make([]*News, 0)
	sqlStr := "select * from " + m.TableName
	if sort != "" {
		sqlStr += " order by " + sort
	}
	err := m.getItems(nil, pageIndex, rowsInPage, &item, sqlStr)
	throw.CheckErr(err)
	total := m.GetCount(nil, "")
	var pageCount int
	if rowsInPage != 0 {
		pageCountF := float64(total) / float64(rowsInPage)
		pageCount = int(math.Ceil(pageCountF))
	} else {
		pageCount = 0
	}
	r = PagingData(pageIndex, rowsInPage, pageCount, total, item)
	return
}

//Insert ..
func (m NewsMapper) Insert(tx *sqlx.Tx, item *News) (int, int) {
	sqlStr := "insert into " + m.TableName + "(id,createTime,updateTime,htmlContent,textContent,title,clickNumber,sort,titlePinYin1,titlePinYin2,original) values (?,?,?,?,?,?,?,?,?,?,?)"
	var args = []interface{}{item.ID, item.CreateTime, item.UpdateTime, item.HTMLContent, item.TextContent, item.Title, item.ClickNumber, item.Sort, item.TitlePinYin1, item.TitlePinYin2, item.Original}
	id, count := m.insertItem(tx, sqlStr, args...)
	return id, count
}

//Update ..
func (m NewsMapper) Update(tx *sqlx.Tx, item *News) int {
	sqlStr := "update " + m.TableName + " set id=?,createTime=?,updateTime=?,htmlContent=?,textContent=?,title=?,clickNumber=?,sort=?,titlePinYin1=?,titlePinYin2=?,original=? where id=? "
	var args = []interface{}{item.ID, item.CreateTime, item.UpdateTime, item.HTMLContent, item.TextContent, item.Title, item.ClickNumber, item.Sort, item.TitlePinYin1, item.TitlePinYin2, item.Original, item.ID}
	count := m.deleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//Delete ..
func (m NewsMapper) Delete(tx *sqlx.Tx, id interface{}) int {
	sqlStr := "delete from " + m.TableName + " where id = ? "
	var args = []interface{}{id}
	return m.deleteOrUpdateItems(tx, sqlStr, args...)
}
