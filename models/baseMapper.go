package models

import (
	"fmt"
	"strconv"
	"strings"
	"test/throw"

	"database/sql"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//BaseMapper ..
type BaseMapper struct {
	TableName string
	DB        *gosql.Wrapper
}

//Tx 失败后抛出异常
func (m BaseMapper) Tx(fn func(tx *sqlx.Tx) error) {
	err := m.DB.Tx(fn)
	throw.CheckErr(err)
}

//获取一行数据，出错或没有数据都返回nil
func (m BaseMapper) getItem(tx *sqlx.Tx, dest interface{}, sqlStr string, args ...interface{}) (r interface{}) {

	defer func() {
		if err := recover(); err != nil {
			switch t := err.(type) {
			case error:
				if strings.Index(t.Error(), "no rows in result set") >= 0 {
					r = nil //没有数据都会执行到这里
					return
				}
			case *error:
				if strings.Index((*t).Error(), "no rows in result set") >= 0 {
					r = nil //没有数据都会执行到这里
					return
				}
			default:
			}
			fmt.Printf("err")
			r = nil //出错会执行到这里
		}
	}()

	var err error
	if tx == nil {
		err = m.DB.Get(dest, sqlStr, args...)
	} else {
		err = gosql.WithTx(tx).Get(dest, sqlStr, args...)
	}
	throw.CheckErr(err)

	return dest
}

func (m BaseMapper) getItems(tx *sqlx.Tx, pageIndex int, rowsInPage int, dest interface{}, sqlStr string, args ...interface{}) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err")

		}
	}()
	if rowsInPage != 0 {
		offset := m.PageIndexToLimit(pageIndex, rowsInPage)
		sqlStr = sqlStr + " limit " + strconv.Itoa(offset) + "," + strconv.Itoa(rowsInPage)
	}
	if tx == nil {
		return m.DB.Select(dest, sqlStr, args...)
	}
	return gosql.WithTx(tx).Select(dest, sqlStr, args...)
}

//PageIndexToLimit 将页码转换为limit的起始序号
func (BaseMapper) PageIndexToLimit(pageIndex int, rowsInPage int) (offset int) {
	if pageIndex < 1 {
		pageIndex = 1
	}
	offset = (pageIndex - 1) * rowsInPage
	return
}

//插入一行数据，出错抛出异常  如果id是整形，可以返回id
func (m BaseMapper) insertItem(tx *sqlx.Tx, sqlStr string, args ...interface{}) (int, int) {
	// funcName, file, line, _ := runtime.Caller(0)
	// logger.Info(sqlStr, funcName, file, line)
	// logger.InfoArgs(args...)

	var sqlResult sql.Result
	var err error

	if tx == nil {
		sqlResult, err = m.DB.Exec(sqlStr, args...)
	} else {
		sqlResult, err = gosql.WithTx(tx).Exec(sqlStr, args...)
	}
	throw.CheckErr(err)

	rows, err := sqlResult.RowsAffected()
	throw.CheckErr(err)

	lastID, err := sqlResult.LastInsertId()
	if err == nil {
		return int(lastID), int(rows)
	}
	return 0, int(rows)
}

//GetCount ..
func (m BaseMapper) GetCount(tx *sqlx.Tx, whereStr string, args ...interface{}) int {
	sqlStr := "select count(*) from " + m.TableName
	if whereStr != "" {
		sqlStr = sqlStr + " where " + whereStr
	}

	var row *sqlx.Row
	if tx == nil {
		row = m.DB.QueryRowx(sqlStr, args...)
	} else {
		row = gosql.WithTx(tx).QueryRowx(sqlStr, args...)
	}
	var count int
	err := row.Scan(&count)
	throw.CheckErr(err)
	return count
}
