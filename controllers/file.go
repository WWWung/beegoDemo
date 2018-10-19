package controllers

import (
	"strconv"
	"strings"
	"test/utils"
	"time"
)

//FileController ..
type FileController struct {
	BaseController
}

//UploadFile ..
func (c *FileController) UploadFile() {
	f, h, err := c.GetFile("file")
	defer f.Close()
	if err != nil {
		c.failure("err")
		return
	}
	s := strings.Split(h.Filename, ".")
	//	获取文件扩展名
	etName := s[len(s)-1]
	name := utils.Encrypt(strconv.Itoa(time.Now().Nanosecond())) + "." + etName
	folderF := c.getDataFromPath("type")
	var folder string
	//	根据文件类型例如img、doc、pdf等等存放到不同文件夹
	if folderF != nil {
		folder = folderF.(string)
	} else {
		panic("文件信息错误")
	}
	src := "assets/" + folder + "/" + name
	c.SaveToFile("file", src)
	data := map[string]string{
		"src": "/" + src,
	}
	c.success(data)
}
