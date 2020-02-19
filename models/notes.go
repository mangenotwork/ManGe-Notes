package models

import (
	"fmt"

	conn "man/ManNotes/conn"
)

// tbl_notes表
type Notes struct {
	NotesId     int `gorm:"column:notes_id;primary_key;AUTO_INCREMENT" json:"notes_id"`     //笔记本id
	UID string `gorm:"column:uid"` //用户id
	NotesName     string `gorm:"column:notes_name" json:"notes_name"`     //分类名笔记本名
	NotesDes 	string `gorm:"column:notes_des" json:"notes_des"` //笔记本描述
	NotesCreatetime int64 `gorm:"column:notes_createtime"`     //创建时间
}

func (this *Notes) TableName() string {
	return "tbl_notes"
}

//新增
func (this *Notes) AddNotes() (err error) {
	orm := conn.NotesDB()
	return orm.Create(this).Error
}

//查询笔记本名是否存在
func (this *Notes) IsNotesName(notesname,uid string) bool {
	orm := conn.NotesDB()
	err := orm.Where("notes_name = ? and uid = ?", notesname, uid).First(this).Error
	if err != nil && err.Error() == "record not found"{
		return true
	}
	return false
}

//分页查询
func (this *Notes) GetNotesPgs(uid string, pg int, size int) ([]*Notes,error){
	orm := conn.NotesDB()
	dataList := make([]*Notes, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_notes where uid='%s' LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList,err
}

//这是一个测试
func (this *Notes) TestDB() (err error) {
	orm := conn.NotesDB()
	return orm.Create(this).Error
}
//这是一个测试
func TestNotes(){
	u := &Notes{NotesName:"aaa",
				NotesCreatetime:12345789123,
			}
	err := u.TestDB()
	fmt.Println(err)
}