package models

import (
	"fmt"

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

//获取用户所有图片的大小
func (this *IMGInfo) GetSize(uid string) (int64,error) {
	
	orm := conn.NotesDB()
	type S struct{
		Size int64
	}
	var size S
	sqlStr := fmt.Sprintf("select sum(size) as size from tbl_img where uid = '%s'", uid)
	err := orm.Raw(sqlStr).Scan(&size).Error
	return size.Size,err
}