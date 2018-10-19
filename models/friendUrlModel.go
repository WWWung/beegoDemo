package models

import (
	"math"
	"test/throw"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//FriendURL ..
type FriendURL struct {
	ID   string `json:"id"  db:"id"  `
	Name string `json:"name" db:"name"`
	URL  string `json:"url" db:"url"`
	Sort int    `json:"sort" db:"sort"`
}

//FriendURLMapper ..
type FriendURLMapper struct {
	BaseMapper
}

//GetFriendURLMapper ..
func GetFriendURLMapper(db string) FriendURLMapper {
	var mp FriendURLMapper
	mp.TableName = "friendUrl"
	if db == "" {
		db = "default"
	}
	mp.DB = gosql.Use(db)
	return mp
}

//Get ..
func (m FriendURLMapper) Get(tx *sqlx.Tx, id interface{}) interface{} {
	sqlStr := "select * from " + m.TableName + " where id = ? "
	var args = []interface{}{id}

	item := &FriendURL{}
	r := m.getItem(tx, item, sqlStr, args...)
	return r
}

//GetList ..
func (m FriendURLMapper) GetList(pageIndex int, rowsInPage int, sort string) (r interface{}) {
	//准备参数
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	item := make([]*FriendURL, 0)
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
func (m FriendURLMapper) Insert(tx *sqlx.Tx, item *FriendURL) (int, int) {
	sqlStr := "insert into " + m.TableName + "(id,name,url,sort) values (?,?,?,?)"
	var args = []interface{}{item.ID, item.Name, item.URL, item.Sort}
	id, count := m.insertItem(tx, sqlStr, args...)
	return id, count
}

//Update ..
func (m FriendURLMapper) Update(tx *sqlx.Tx, item *FriendURL) int {
	sqlStr := "update " + m.TableName + " set id=?,name=?,url=?,sort=? where id=? "
	var args = []interface{}{item.ID, item.Name, item.URL, item.Sort, item.ID}
	count := m.deleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//Delete ..
func (m FriendURLMapper) Delete(tx *sqlx.Tx, id interface{}) int {
	sqlStr := "delete from " + m.TableName + " where id = ? "
	var args = []interface{}{id}
	return m.deleteOrUpdateItems(tx, sqlStr, args...)
}
