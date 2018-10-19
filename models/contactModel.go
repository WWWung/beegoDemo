package models

import (
	"math"
	"test/throw"
	"time"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//Word ..
type Word struct {
	ID       string    `json:"id"  db:"id"  `
	UserID   string    `json:"userId"  db:"userId"  `
	UserName string    `json:"userName"  db:"userName"  `
	Content  string    `json:"content" db:"content"`
	Time     time.Time `json:"time"  db:"time"  `
}

//WordMapper ..
type WordMapper struct {
	BaseMapper
}

//GetWordMapper ..
func GetWordMapper(db string) WordMapper {
	var mp WordMapper
	mp.TableName = "words"
	if db == "" {
		db = "default"
	}
	mp.DB = gosql.Use(db)
	return mp
}

//Get ..
func (m WordMapper) Get(tx *sqlx.Tx, whereStr string, args ...interface{}) (r interface{}) {
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	sqlStr := "select * from " + m.TableName + whereStr
	item := &Word{}
	r = m.getItem(tx, item, sqlStr, args...)
	return r
}

//GetList ..
func (m WordMapper) GetList(pageIndex int, rowsInPage int, sort string) (r interface{}) {
	//准备参数
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	item := make([]*Word, 0)
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
func (m WordMapper) Insert(tx *sqlx.Tx, item *Word) (int, int) {
	sqlStr := "insert into " + m.TableName + "(id,userId,userName,content,time) values (?,?,?,?,?)"
	var args = []interface{}{item.ID, item.UserID, item.UserName, item.Content, item.Time}
	id, count := m.insertItem(tx, sqlStr, args...)
	return id, count
}

//Update ..
func (m WordMapper) Update(tx *sqlx.Tx, item *Word) int {
	sqlStr := "update " + m.TableName + " set id=?,userId=?,userName=?,content=?,time=? where id=? "
	var args = []interface{}{item.ID, item.UserID, item.UserName, item.Content, item.Time, item.ID}
	count := m.deleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//Delete ..
func (m WordMapper) Delete(tx *sqlx.Tx, id interface{}) int {
	sqlStr := "delete from " + m.TableName + " where id = ? "
	var args = []interface{}{id}
	return m.deleteOrUpdateItems(tx, sqlStr, args...)
}
