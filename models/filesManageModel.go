package models

import (
	"math"
	"test/throw"
	"time"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//FilesManage ..
type FilesManage struct {
	ID             string    `json:"id"  db:"id"  `
	Type           int       `json:"type"  db:"type"  `
	URL            string    `json:"url"  db:"url"  `
	Name           string    `json:"name"  db:"name"  `
	NamePinYin1    string    `json:"namePinYin1"  db:"namePinYin1"  `
	NamePinYin2    string    `json:"namePinYin2" db:"namePinYin2"`
	Description    string    `json:"description" db:"description"`
	ExtName        string    `json:"extName" db:"extName"`
	UploadTime     time.Time `json:"uploadTime" db:"uploadTime"`
	Sort           int       `json:"sort" db:"sort"`
	DownloadNumber int       `json:"downloadNumber" db:"downloadNumber"`
	Size           int       `json:"size" db:"size"`
	Rank           int       `json:"rank" db:"rank"`
	Preview        int       `json:"preview" db:"preview"`
}

//FilesManageMapper ..
type FilesManageMapper struct {
	BaseMapper
}

//GetFilesManageMapper ..
func GetFilesManageMapper(db string) FilesManageMapper {
	var mp FilesManageMapper
	mp.TableName = "filesManage"
	if db == "" {
		db = "default"
	}
	mp.DB = gosql.Use(db)
	return mp
}

//Get ..
func (m FilesManageMapper) Get(tx *sqlx.Tx, id interface{}) interface{} {
	sqlStr := "select * from " + m.TableName + " where id = ? "
	var args = []interface{}{id}

	item := &FilesManage{}
	r := m.getItem(tx, item, sqlStr, args...)
	return r
}

//GetList ..
func (m FilesManageMapper) GetList(pageIndex int, rowsInPage int, whereStr string, order string, args ...interface{}) (r interface{}) {
	//准备参数
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	item := make([]*FilesManage, 0)
	sqlStr := "select * from " + m.TableName
	if whereStr != "" {
		sqlStr += whereStr
	}
	if order != "" {
		sqlStr += " order by " + order
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
func (m FilesManageMapper) Insert(tx *sqlx.Tx, item *FilesManage) (int, int) {
	sqlStr := "insert into " + m.TableName + "(id,type,url,name,namePinYin1,namePinYin2,description,extName,uploadTime,sort,size,downloadNumber,rank,preview) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	var args = []interface{}{item.ID, item.Type, item.URL, item.Name, item.NamePinYin1, item.NamePinYin2, item.Description, item.ExtName, item.UploadTime, item.Sort, item.Size, item.DownloadNumber, item.Rank, item.Preview}
	id, count := m.insertItem(tx, sqlStr, args...)
	return id, count
}

//Update ..
func (m FilesManageMapper) Update(tx *sqlx.Tx, item *FilesManage) int {
	sqlStr := "update " + m.TableName + " set id=?,type=?,url=?,name=?,namePinYin1=?,namePinYin2=?,description=?,extName=?,uploadTime=?,sort=?,downloadNumber=?,size=?,rank=?,preview=? where id=? "
	var args = []interface{}{item.ID, item.Type, item.URL, item.Name, item.NamePinYin1, item.NamePinYin2, item.Description, item.ExtName, item.UploadTime, item.Sort, item.DownloadNumber, item.Size, item.Rank, item.Preview, item.ID}
	count := m.deleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//Delete ..
func (m FilesManageMapper) Delete(tx *sqlx.Tx, id interface{}) int {
	sqlStr := "delete from " + m.TableName + " where id = ? "
	var args = []interface{}{id}
	return m.deleteOrUpdateItems(tx, sqlStr, args...)
}
