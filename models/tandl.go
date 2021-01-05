package models

/*
	工具与链接模块
*/

// tbl_tool_link
type ToolandLink struct {
	Id       int    `gorm:"column:id;primary_key;AUTO_INCREMENT"` //笔记本id
	UID      string `gorm:"column:uid"`                           //用户id
	Name     string `gorm:"column:name"`                          //工具或链接名称
	Des      string `gorm:"column:des"`                           //工具或链接描述
	Link     string `gorm:"column:link"`                          //工具或链接 地址
	Ico      string `gorm:"column:ico"`                           //工具或链接 地址 ico
	Tag      string `gorm:"column:tag"`                           //工具或链接 标签
	LinkType int    `gorm:"column:type"`                          //链接 0    工具 1
}

func (this *ToolandLink) TableName() string {
	return "tbl_tool_link"
}
