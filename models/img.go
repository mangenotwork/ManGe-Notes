package models

import (
	_ "fmt"

	conn "man/ManNotes/conn"
)

// tbl_img
type IMGInfo struct {
	Id     int `gorm:"column:id;primary_key;AUTO_INCREMENT"`     //图片ID
	ImgName string `gorm:"column:imgname"` //图片名 真实名
	Imgdec     string `gorm:"column:imgdec"`     //图片描述or名称
	Uid 	string `gorm:"column:uid"` //用户id
	Time int64 `gorm:"column:time"`     //存储时间
	Date string `gorm:"column:date"`     //存储时间 日期
	Size int64 `gorm:"column:size"`     //大小
	Imgtag string `gorm:"column:imgtag"`     //图片标签
}


func (this *IMGInfo) TableName() string {
	return "tbl_img"
}

func (this *IMGInfo) CreateImg() error {
	orm := conn.NotesDB()
	return orm.Create(this).Error
}