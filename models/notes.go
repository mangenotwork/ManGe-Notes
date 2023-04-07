package models

// tbl_notes表
type Notes struct {
	NotesId         int    `gorm:"column:notes_id;primary_key;AUTO_INCREMENT" json:"notes_id"` //笔记本id
	UID             string `gorm:"column:uid"`                                                 //用户id
	NotesName       string `gorm:"column:notes_name" json:"notes_name"`                        //分类名笔记本名
	NotesDes        string `gorm:"column:notes_des" json:"notes_des"`                          //笔记本描述
	NotesCreatetime int64  `gorm:"column:notes_createtime"`                                    //创建时间
}

func (this *Notes) TableName() string {
	return "tbl_notes"
}
