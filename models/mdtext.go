package models

import (
	"fmt"

    conn "man/ManNotes/conn"
)

// tbl_md_text表
type MDText struct {
	MDId     string `gorm:"column:md_id"`     //内容id
	MDContent     string `gorm:"column:md_content"`     //内容
}


func (this *MDText) TableName() string {
	return "tbl_md_text"
}

//通过mid查询MD内容
func (this *MDText) GetMDTxt(mid string) (string,error) {
	orm := conn.NotesDB()
	err := orm.Where("md_id = ?", mid).First(this).Error
	if err != nil{
		fmt.Println("通过mid查询MD内容错误",err)
		return "",err
	}
	return this.MDContent,nil
}