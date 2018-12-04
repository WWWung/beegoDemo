package controllers

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"test/models"
	"test/throw"
	"test/utils"
	"time"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

//PreviewImgController ..
type PreviewImgController struct {
	BaseController
}

//Add ..
func (c PreviewImgController) Add() string {

	//	获取上传的图片
	f, _, err := c.GetFile("file")
	defer f.Close()
	if err != nil {
		panic("err")
	}
	//	解码上传的图片
	var img image.Image
	var extName = c.getStringFromPath("extName")
	fmt.Println(extName, "============")
	if extName == "png" {
		img, err = png.Decode(f)
	} else if extName == "jpeg" {
		img, err = jpeg.Decode(f)
	} else {
		panic("文件格式错误")
	}
	if err != nil {
		panic("err")
	}
	//	打开水印图片
	wmb, _ := os.Open("file/imgs/wm.png")
	watermark, _ := png.Decode(wmb)
	defer wmb.Close()

	//把水印写到右下角，并向0坐标各偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-10, img.Bounds().Dy()-watermark.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	//生成新图片new.jpg，并设置图片质量..
	name := utils.Encrypt(strconv.Itoa(time.Now().Nanosecond())) + "." + extName
	src := "assets/imgs/" + name
	imgw, _ := os.Create(src)
	jpeg.Encode(imgw, m, &jpeg.Options{Quality: 100})

	defer imgw.Close()

	item := models.PreviewImg{}
	uid, err := uuid.NewV4()
	throw.CheckErr(err)
	id := uid.String()
	item.ID = id
	item.Src = "/" + src
	item.FileID = c.getStringFromPath("fileId")
	item.Sort = c.getIntFromPath("sort")
	mp := models.GetPreviewImgMapper("")
	mp.Tx(func(tx *sqlx.Tx) (r error) {
		defer func() {
			if err := recover(); err != nil {
				msg := utils.InterfaceToStr(err)
				r = errors.New(msg)
			}
		}()

		_, count := mp.Insert(tx, &item)
		if count == 0 {
			panic("Insert未成功")
		}
		return nil
	})
	return src
}

//GetList ..
func (c PreviewImgController) GetList() interface{} {
	pageIndex := c.getPageIndex()
	rowsInPage := c.getRowsInPage()
	sort := c.getSort()
	mp := models.GetPreviewImgMapper("")
	fileID := c.getStringFromPath("fileId")
	whereStr := ""
	args := []interface{}{}
	if fileID != "" {
		whereStr = " where fileId=? "
		args = append(args, fileID)
	}
	r := mp.GetList(pageIndex, rowsInPage, sort, whereStr, args...)
	return r
}

//Delete ..
func (c PreviewImgController) Delete() {
	fileID := c.getStringFromPath("fileId")
	mp := models.GetPreviewImgMapper("")
	r := mp.GetList(1, 999, "", " where fileId=? ", fileID)
	if r == nil {
		return
	}
	if r.(map[string]interface{})["rows"] != nil {
		for _, item := range r.(map[string]interface{})["rows"].([]*models.PreviewImg) {
			_ = os.Remove(item.Src[1:len(item.Src)])
		}
	}
	mp.Deletes(nil, " where fileId=? ", fileID)
	c.success(fileID)
}

//APIhandler ..
func (c PreviewImgController) APIhandler() {
	c.handler(c)
}
