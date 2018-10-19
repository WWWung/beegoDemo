package models

import (
	"math"
	"test/throw"
	"time"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//User ..
type User struct {
	ID            string    `json:"id"  db:"id"  `
	CreateTime    time.Time `json:"createTime"  db:"createTime"  `
	Name          string    `json:"name"  db:"name"  `
	NamePinYin1   string    `json:"namePinYin1"  db:"namePinYin1"  `
	NamePinYin2   string    `json:"namePinYin2" db:"namePinYin2"`
	Phone         string    `json:"phone" db:"phone"`
	Addr          string    `json:"addr" db:"addr"`
	LastLoginTime time.Time `json:"lastLoginTime" db:"lastLoginTime"`
	Power         int       `json:"power" db:"power"`
	PwMD5         string    `json:"pwMD5" db:"pwMD5"`
	DwPhone       string    `json:"dwPhone" db:"dwPhone"`
}

//UserMapper ..
type UserMapper struct {
	BaseMapper
}

//GetUserMapper ..
func GetUserMapper(db string) UserMapper {
	var mp UserMapper
	mp.TableName = "user"
	if db == "" {
		db = "default"
	}
	mp.DB = gosql.Use(db)
	return mp
}

//Get ..
func (m UserMapper) Get(tx *sqlx.Tx, whereStr string, args ...interface{}) (r interface{}) {
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	sqlStr := "select * from " + m.TableName + whereStr
	item := &User{}
	r = m.getItem(tx, item, sqlStr, args...)
	return r
}

//GetList ..
func (m UserMapper) GetList(pageIndex int, rowsInPage int, sort string) (r interface{}) {
	//准备参数
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	item := make([]*User, 0)
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
func (m UserMapper) Insert(tx *sqlx.Tx, item *User) (int, int) {
	sqlStr := "insert into " + m.TableName + "(id,name,namePinYin1,namePinYin2,phone,addr,createTime,lastLoginTime,power,pwMD5,dwPhone) values (?,?,?,?,?,?,?,?,?,?,?)"
	var args = []interface{}{item.ID, item.Name, item.NamePinYin1, item.NamePinYin2, item.Phone, item.Addr, item.CreateTime, item.LastLoginTime, item.Power, item.PwMD5, item.DwPhone}
	id, count := m.insertItem(tx, sqlStr, args...)
	return id, count
}

//Update ..
func (m UserMapper) Update(tx *sqlx.Tx, item *User) int {
	sqlStr := "update " + m.TableName + " set id=?,name=?,namePinYin1=?,namePinYin2=?,phone=?,addr=?,createTime=?,lastLoginTime=?,power=?,pwMD5=?,dwPhone=? where id=? "
	var args = []interface{}{item.ID, item.Name, item.NamePinYin1, item.NamePinYin2, item.Phone, item.Addr, item.CreateTime, item.LastLoginTime, item.Power, item.PwMD5, item.DwPhone, item.ID}
	count := m.deleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//Delete ..
func (m UserMapper) Delete(tx *sqlx.Tx, id interface{}) int {
	sqlStr := "delete from " + m.TableName + " where id = ? "
	var args = []interface{}{id}
	return m.deleteOrUpdateItems(tx, sqlStr, args...)
}
