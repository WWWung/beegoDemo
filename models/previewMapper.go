package models

import (
	"math"
	"test/throw"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//PreviewImg 针对pdf文件上传图片作为预览
type PreviewImg struct {
	ID     string `json:"id"  db:"id"  `
	Src    string `json:"src"  db:"src"  `
	Src2   string `json:"src2"  db:"src2"  `
	Sort   int    `json:"sort"  db:"sort"  `
	Mark   int    `json:"mark"  db:"mark"  `
	FileID string `json:"fileId" db:"fileId"`
}

//PreviewImgMapper ..
type PreviewImgMapper struct {
	BaseMapper
}

//GetPreviewImgMapper ..
func GetPreviewImgMapper(db string) PreviewImgMapper {
	var mp PreviewImgMapper
	mp.TableName = "preview_img"
	if db == "" {
		db = "default"
	}
	mp.DB = gosql.Use(db)
	return mp
}

//Get ..
func (m PreviewImgMapper) Get(tx *sqlx.Tx, id interface{}) interface{} {
	sqlStr := "select * from " + m.TableName + " where id = ? "
	var args = []interface{}{id}

	item := &PreviewImg{}
	r := m.getItem(tx, item, sqlStr, args...)
	return r
}

//GetList ..
func (m PreviewImgMapper) GetList(pageIndex int, rowsInPage int, sort string, whereStr string, sqlArgs ...interface{}) (r interface{}) {
	//准备参数
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	item := make([]*PreviewImg, 0)
	sqlStr := "select * from " + m.TableName
	sqlStr += whereStr
	if sort != "" {
		sqlStr += " order by " + sort
	} else {
		sqlStr += " order by sort "
	}
	err := m.getItems(nil, pageIndex, rowsInPage, &item, sqlStr, sqlArgs...)
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
func (m PreviewImgMapper) Insert(tx *sqlx.Tx, item *PreviewImg) (int, int) {
	sqlStr := "insert into " + m.TableName + "(id,src,src2,sort,mark,fileId) values (?,?,?,?,?,?)"
	var args = []interface{}{item.ID, item.Src, item.Src2, item.Sort, item.Mark, item.FileID}
	id, count := m.insertItem(tx, sqlStr, args...)
	return id, count
}

//Update ..
func (m PreviewImgMapper) Update(tx *sqlx.Tx, item *PreviewImg) int {
	sqlStr := "update " + m.TableName + " set id=?,src=?,src2=?,sort=?,mark=?,fileId=? where id=? "
	var args = []interface{}{item.ID, item.Src, item.Src2, item.Sort, item.Mark, item.FileID, item.ID}
	count := m.deleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//Delete ..
func (m PreviewImgMapper) Delete(tx *sqlx.Tx, id interface{}) int {
	sqlStr := "delete from " + m.TableName + " where id = ? "
	var args = []interface{}{id}
	return m.deleteOrUpdateItems(tx, sqlStr, args...)
}

//Deletes ..
func (m PreviewImgMapper) Deletes(tx *sqlx.Tx, whereStr string, args ...interface{}) int {
	sqlStr := "delete from " + m.TableName + whereStr
	return m.deleteOrUpdateItems(tx, sqlStr, args...)
}
