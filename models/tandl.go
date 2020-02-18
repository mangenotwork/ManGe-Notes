package models

/*
		工具与链接模块
*/

import (
	"fmt"

	conn "man/ManNotes/conn"
)

// tbl_tool_link
type ToolandLink struct {
	Id     int `gorm:"column:id;primary_key;AUTO_INCREMENT"`     //笔记本id
	UID string `gorm:"column:uid"` //用户id
	Name     string `gorm:"column:name"`     //工具或链接名称
	Des 	string `gorm:"column:des"` //工具或链接描述
	Link string `gorm:"column:link"`     //工具或链接 地址
	Ico string `gorm:"column:ico"`     //工具或链接 地址 ico
	Tag string `gorm:"column:tag"`     //工具或链接 标签
	LinkType int `gorm:"column:type"` //链接 0    工具 1
}

func (this *ToolandLink) TableName() string {
	return "tbl_tool_link"
}

func (this *ToolandLink) Create() error {
	orm := conn.NotesDB()
	return orm.Create(this).Error
}

//获取tool 列表
func (this *ToolandLink) GeTools(pg int, size int, uid string) ([]*ToolandLink,error){
	orm := conn.NotesDB()
	dataList := make([]*ToolandLink, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_tool_link where uid='%s' and type=1 LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList,err
}