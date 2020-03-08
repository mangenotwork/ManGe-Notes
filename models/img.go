package models

import (
	"fmt"

	conn "man/ManNotes/conn"
)

// tbl_img
type IMGInfo struct {
	Id     int `gorm:"column:id;primary_key;AUTO_INCREMENT"`     //图片ID
	ImgName string `gorm:"column:imgname"` //图片名 图片地址
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

//获取我的图片
func (this *IMGInfo) GetMyImg(uid string) ([]*IMGInfo,error) {
	orm := conn.NotesDB()
	dataList := make([]*IMGInfo, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_img where uid = '%s' LIMIT %d,%d", uid, 0, 100)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList,err
}

//验证图片权限
func (this *IMGInfo) IsMyImg(imgid int,uid string) (bool,*IMGInfo){
	orm := conn.NotesDB()
	err := orm.Where("id = ? and uid = ?", imgid, uid).First(this).Error
	if err != nil && err.Error() == "record not found"{
		return true,this
	}
	return false,this
}
