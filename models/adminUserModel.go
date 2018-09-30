package models

import (
	"time"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//AdminUser ..
type AdminUser struct {
	ID    string    `json:"id" db:"id"`
	Name  string    `json:"name" db:"name"`
	Power string    `json:"power" db:"power"`
	PwMD5 string    `json:"pwMD5" db:"pwMD5"`
	Time  time.Time `json:"time" db:"time"`
}

//AdminUserMapper ..
type AdminUserMapper struct {
	BaseMapper
}

//GetAdminUserMapper ..
func GetAdminUserMapper(db string) AdminUserMapper {
	var mp AdminUserMapper
	mp.TableName = "adminUser"
	if db == "" {
		db = "default"
	}
	mp.DB = gosql.Use(db)
	return mp
}

//Get ..
func (m AdminUserMapper) Get(tx *sqlx.Tx, name interface{}, pwMD5 interface{}) interface{} {
	sqlStr := "select * from " + m.TableName + " where name = ? and pwMD5 = ? "
	var args = []interface{}{name, pwMD5}

	item := &AdminUser{}
	r := m.getItem(tx, item, sqlStr, args...)
	return r
}

//Insert ..
func (m AdminUserMapper) Insert(tx *sqlx.Tx, item *AdminUser) (int, int) {
	sqlStr := "insert into " + m.TableName + "(id,name,power,pwMD5,time) values (?,?,?,?,?)"
	var args = []interface{}{item.ID, item.Name, item.Power, item.PwMD5, item.Time}
	id, count := m.insertItem(tx, sqlStr, args...)
	return id, count
}
