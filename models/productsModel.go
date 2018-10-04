package models

import (
	"math"
	"test/throw"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//Product 产品中心
type Product struct {
	ID           string `json:"id"  db:"id"  `
	CreateTime   string `json:"createTime"  db:"createTime"  `
	UpdateTime   string `json:"updateTime"  db:"updateTime"  `
	CreateUserID string `json:"createUserId"  db:"createUserId"  `
	UpdateUserID string `json:"updateUserId"  db:"updateUserId"  `
	HTMLContent  string `json:"htmlContent" db:"htmlContent"`
	TextContent  string `json:"textContent" db:"textContent"`
	Title        string `json:"title" db:"title"`
	ClickNumber  int    `json:"clickNumber" db:"clickNumber"`
	Brand        string `json:"brand" db:"brand"`
	Sort         int    `json:"sort" db:"sort"`
}

//ProductsMapper ..
type ProductsMapper struct {
	BaseMapper
}

//GetProductsMapper ..
func GetProductsMapper(db string) ProductsMapper {
	var mp ProductsMapper
	mp.TableName = "products"
	if db == "" {
		db = "default"
	}
	mp.DB = gosql.Use(db)
	return mp
}

//Get ..
func (m ProductsMapper) Get(tx *sqlx.Tx, id interface{}) interface{} {
	sqlStr := "select * from " + m.TableName + " where id = ? "
	var args = []interface{}{id}

	item := &Product{}
	r := m.getItem(tx, item, sqlStr, args...)
	return r
}

//GetList ..
func (m ProductsMapper) GetList(pageIndex int, rowsInPage int) (r interface{}) {
	//准备参数
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	item := make([]*Product, 0)
	sqlStr := "select * from " + m.TableName
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
func (m ProductsMapper) Insert(tx *sqlx.Tx, item *Product) (int, int) {
	sqlStr := "insert into " + m.TableName + "(id,createTime,updateTime,createUserId,updateUserId,htmlContent,textContent,title,clickNumber,brand,sort) values (?,?,?,?,?,?,?,?,?,?,?)"
	var args = []interface{}{item.ID, item.CreateTime, item.UpdateTime, item.CreateUserID, item.UpdateUserID, item.HTMLContent, item.TextContent, item.Title, item.ClickNumber, item.Brand, item.Sort}
	id, count := m.insertItem(tx, sqlStr, args...)
	return id, count
}

//Update ..
func (m ProductsMapper) Update(tx *sqlx.Tx, item *Product) int {
	sqlStr := "update " + m.TableName + " set id=?,createTime=?,updateTime=?,createUserId=?,updateUserId=?,htmlContent=?,textContent=?,title=?,clickNumber=?,brand=?,sort=? where id=? "
	var args = []interface{}{item.ID, item.CreateTime, item.UpdateTime, item.CreateUserID, item.UpdateUserID, item.HTMLContent, item.TextContent, item.Title, item.ClickNumber, item.Brand, item.Sort, item.ID}
	count := m.deleteOrUpdateItems(tx, sqlStr, args...)
	return count
}
