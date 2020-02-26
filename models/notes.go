package models

import (
	"fmt"

	conn "man/ManNotes/conn"
	object "man/ManNotes/object"
	util "man/ManNotes/util"
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
func (this *Notes) GetNotesPgs(uid string, pg int, size int) ([]*object.NotesInfo,error){
	orm := conn.NotesDB()
	dataList := make([]*object.NotesInfo, 0)
	sqlStr := fmt.Sprintf("SELECT notes_id,notes_name FROM tbl_notes where uid='%s' LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList,err
}

//分页查询笔记本信息
func (this *Notes) GetNotesPgsInfo1(uid string, pg int, size int) ([]*object.NotesInfo,error){
	orm := conn.NotesDB()
	dataList := make([]*object.NotesInfo,0)
	sqlStr := fmt.Sprintf("SELECT b.notes_id,b.notes_name,a.n as note_number from (SELECT md_notes_id as id,count(*)"+
		" as n FROM `tbl_md_info` where uid = '%s' GROUP BY md_notes_id) as a,`tbl_notes` as b where "+
		"b.notes_id = a.id LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList,err
}

type NotesMange struct{
	Notes
	N int
}

//分页查询笔记本信息
func (this *Notes) GetNotesPgsInfo(uid string, pg int, size int) ([]*NotesMange,error){
	orm := conn.NotesDB()
	dataList := make([]*NotesMange,0)
	//sqlStr := fmt.Sprintf("SELECT b.notes_id,b.notes_name,b.notes_des,b.notes_createtime,a.n as n from (SELECT md_notes_id as id,count(*)"+
	//	" as n FROM `tbl_md_info` where uid = '%s' GROUP BY md_notes_id) as a,`tbl_notes` as b where "+
	//	"b.notes_id = a.id LIMIT %d,%d", uid, pg, size)
	sqlStr := fmt.Sprintf("SELECT B.notes_id,B.notes_name,B.notes_des,B.notes_createtime,COUNT(A.md_notes_id) AS n from"+
		" tbl_md_info AS A RIGHT JOIN tbl_notes AS B on A.md_notes_id=B.notes_id where B.uid = '%s' GROUP BY B.notes_id LIMIT %d,%d",
		uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList,err
}


type Number struct{
		N int
	}

//查询默认笔记本的笔记数量
func (this *Notes) NotesNumber(uid string, notesid int) (int,error) {
	orm := conn.NotesDB()

	var number = &Number{}
	sqlStr := fmt.Sprintf("SELECT count(*) as n FROM `tbl_md_info` where uid = '%s' and md_notes_id = %d", uid, notesid)
	err := orm.Raw(sqlStr).Scan(&number).Error
	return number.N,err
}

func (this *Notes) NotesAllNumber(uid string) (int,error) {
	orm := conn.NotesDB()

	var number = &Number{}
	sqlStr := fmt.Sprintf("SELECT count(*) as n FROM `tbl_md_info` where uid = '%s'", uid)
	err := orm.Raw(sqlStr).Scan(&number).Error
	return number.N,err
}

func (this *Notes) UpdateInfo(datas *object.UpdateNotes,uid string) error {
	orm := conn.NotesDB()
	notesid,_ := new(util.Str).NumberToInt(datas.NotesID)
	sqlStr := fmt.Sprintf("update tbl_notes set notes_name='%s',notes_des='%s' where notes_id=%d and uid = '%s'; ", 
		datas.NotesName, datas.NotesDes, notesid, uid)
	return orm.Exec(sqlStr).Error
}

func (this *Notes) DeleteInfo(nid int,uid string) error {
	orm := conn.NotesDB().Begin()
	sqlStrNotes := fmt.Sprintf("DELETE FROM tbl_notes where notes_id=%d and uid = '%s'; ", nid, uid)
	//将某笔记本的笔记转移到默认笔记本
	sqlMDInofStr := fmt.Sprintf("update tbl_md_info set md_notes_id=0 where md_notes_id=%d and uid = '%s'; ", nid, uid)
	err := orm.Exec(sqlStrNotes).Error
	if err != nil {
		orm.Rollback()
		return err
	}
	err = orm.Exec(sqlMDInofStr).Error
	if err != nil {
		orm.Rollback()
		return err
	}
	return orm.Commit().Error
}