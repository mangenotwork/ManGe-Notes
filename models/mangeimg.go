package models

import (
	"fmt"

	conn "github.com/mangenotwork/ManGe-Notes/conn"
)

// tbl_sc_img
type SCIMGInfo struct {
	Id         int    `gorm:"column:id;primary_key;AUTO_INCREMENT"` //图片ID
	ImgName    string `gorm:"column:imgurl;unique"`                 //图片地址
	Imgdec     string `gorm:"column:imgdec"`                        //图片描述or名称
	Uid        string `gorm:"column:uid"`                           //用户id
	Time       int64  `gorm:"column:time"`                          //存储时间 时间戳
	Date       string `gorm:"column:date"`                          //存储时间 日期
	Size       int64  `gorm:"column:size"`                          //大小
	Imgtag     string `gorm:"column:imgtag"`                        //图片标签
	Like       int64  `gorm:"column:like"`                          //点赞数
	Collection int64  `gorm:"column:collection"`                    //收藏数
	Usecount   int64  `gorm:"column:usecount"`                      //被使用数
	View       int64  `gorm:"column:view"`                          //查看数
}

func (this *SCIMGInfo) TableName() string {
	return "tbl_sc_img"
}

func (this *SCIMGInfo) AddSCImg() error {
	orm := conn.NotesDB()
	return orm.Create(this).Error
}

//获取漫鸽图库
func (this *SCIMGInfo) MangeImgList() ([]*SCIMGInfo, error) {
	orm := conn.NotesDB()
	dataList := make([]*SCIMGInfo, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_sc_img LIMIT %d,%d", 0, 100)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}
