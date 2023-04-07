package models

// tbl_md_text表
type MDText struct {
	MDId      string `gorm:"column:md_id"`      //内容id
	MDContent string `gorm:"column:md_content"` //内容
}

func (this *MDText) TableName() string {
	return "tbl_md_text"
}
