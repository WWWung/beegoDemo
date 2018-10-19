package models

import (
	"math"
	"test/throw"
	"time"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//Summary 简介
type Summary struct {
	ID          string    `json:"id"  db:"id"  `
	Time        time.Time `json:"time"  db:"time"  `
	HTMLContent string    `json:"htmlContent" db:"htmlContent"`
	TextContent string    `json:"textContent" db:"textContent"`
	Title       string    `json:"title" db:"title"`
	ClickNumber int       `json:"clickNumber" db:"clickNumber"`
	Type        int       `json:"type" db:"type"`
	Sort        int       `json:"sort" db:"sort"`
}

//SummaryMapper ..
type SummaryMapper struct {
	BaseMapper
}

//GetSummaryMapper ..
func GetSummaryMapper(db string) SummaryMapper {
	var mp SummaryMapper
	mp.TableName = "summary"
	if db == "" {
		db = "default"
	}
	mp.DB = gosql.Use(db)
	return mp
}

//Get ..
func (m SummaryMapper) Get(tx *sqlx.Tx, id interface{}) interface{} {
	sqlStr := "select * from " + m.TableName + " where id = ? "
	var args = []interface{}{id}

	item := &Summary{}
	r := m.getItem(tx, item, sqlStr, args...)
	return r
}

//GetList ..
func (m SummaryMapper) GetList(pageIndex int, rowsInPage int, whereStr, sort string, args ...interface{}) (r interface{}) {
	//准备参数
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	item := make([]*Summary, 0)
	sqlStr := "select * from " + m.TableName
	if sort != "" {
		sqlStr += " order by " + sort
	}
	if whereStr != "" {
		sqlStr += whereStr
	}
	err := m.getItems(nil, pageIndex, rowsInPage, &item, sqlStr, args...)
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
func (m SummaryMapper) Insert(tx *sqlx.Tx, item *Summary) (int, int) {
	sqlStr := "insert into " + m.TableName + "(id,time,htmlContent,textContent,title,clickNumber,type,sort) values (?,?,?,?,?,?,?,?)"
	var args = []interface{}{item.ID, item.Time, item.HTMLContent, item.TextContent, item.Title, item.ClickNumber, item.Type, item.Sort}
	id, count := m.insertItem(tx, sqlStr, args...)
	return id, count
}

//Update ..
func (m SummaryMapper) Update(tx *sqlx.Tx, item *Summary) int {
	sqlStr := "update " + m.TableName + " set id=?,time=?,htmlContent=?,textContent=?,title=?,clickNumber=?,type=?,sort=? where id=? "
	var args = []interface{}{item.ID, item.Time, item.HTMLContent, item.TextContent, item.Title, item.ClickNumber, item.Type, item.Sort, item.ID}
	count := m.deleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//Delete ..
func (m SummaryMapper) Delete(tx *sqlx.Tx, id interface{}) int {
	sqlStr := "delete from " + m.TableName + " where id = ? "
	var args = []interface{}{id}
	return m.deleteOrUpdateItems(tx, sqlStr, args...)
}
