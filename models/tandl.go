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

//获取链接 列表
func (this *ToolandLink) GetLinks(pg int, size int, uid string) ([]*ToolandLink,error){
	orm := conn.NotesDB()
	dataList := make([]*ToolandLink, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_tool_link where uid='%s' and type=0 LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList,err
}

//修改收藏的链接
func (this *ToolandLink) UpdateLink() error {
	orm := conn.NotesDB()
	sqlStr := fmt.Sprintf("update tbl_tool_link set name='%s',des='%s',link='%s',ico='%s' where id=%d and uid = '%s'; ", this.Name,this.Des,this.Link,this.Ico,this.Id,this.UID)
	return orm.Exec(sqlStr).Error
}

//删除收藏的链接
func (this *ToolandLink) DELLink(uid string, linkid int) error {
	orm := conn.NotesDB()
	sqlStr := fmt.Sprintf("DELETE FROM tbl_tool_link where uid='%s' and id=%d and type=0; ",uid,linkid)
	return orm.Exec(sqlStr).Error
}

//GetAll 获取所有收藏链接
func (this *ToolandLink) GetAll(pg int, size int, uid string) ([]*ToolandLink,error){
	orm := conn.NotesDB()
	dataList := make([]*ToolandLink, 0)
	sqlStr := fmt.Sprintf("SELECT id,name,des,link,ico,tag,type FROM tbl_tool_link where uid='%s' LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList,err
}

// EDLink  修改全部字段链接信息
func (this *ToolandLink) EDLink() error {
	orm := conn.NotesDB()
	return orm.Save(&this).Error
}
