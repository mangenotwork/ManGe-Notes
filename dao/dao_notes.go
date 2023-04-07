package dao

import (
	"fmt"

	"github.com/mangenotwork/ManGe-Notes/models"
	"github.com/mangenotwork/ManGe-Notes/object"
	"github.com/mangenotwork/ManGe-Notes/util"
)

type DaoNotes struct{}

//新增
func (this *DaoNotes) AddNotes(notes *models.Notes) (err error) {
	orm := GetConn()
	defer orm.Close()

	return orm.Create(notes).Error
}

//查询笔记本名是否存在
func (this *DaoNotes) IsNotesName(notesname, uid string) bool {
	orm := GetConn()
	orm = orm.Table(new(models.Notes).TableName())
	defer orm.Close()

	err := orm.Where("notes_name = ? and uid = ?", notesname, uid).First(this).Error
	if err != nil && err.Error() == "record not found" {
		return true
	}
	return false
}

//分页查询
func (this *DaoNotes) GetNotesPgs(uid string, pg int, size int) ([]*object.NotesInfo, error) {
	orm := GetConn()
	orm = orm.Table(new(models.Notes).TableName())
	defer orm.Close()

	dataList := make([]*object.NotesInfo, 0)
	sqlStr := fmt.Sprintf("SELECT notes_id,notes_name FROM tbl_notes where uid='%s' LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}

//分页查询笔记本信息
func (this *DaoNotes) GetNotesPgsInfo1(uid string, pg int, size int) ([]*object.NotesInfo, error) {
	orm := GetConn()
	orm = orm.Table(new(models.Notes).TableName())
	defer orm.Close()

	dataList := make([]*object.NotesInfo, 0)
	sqlStr := fmt.Sprintf("SELECT b.notes_id,b.notes_name,a.n as note_number from (SELECT md_notes_id as id,count(*)"+
		" as n FROM `tbl_md_info` where uid = '%s' GROUP BY md_notes_id) as a,`tbl_notes` as b where "+
		"b.notes_id = a.id LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}

//统计模块使用 获取笔记本数据
func (this *DaoNotes) GetAllNotes(uid string) (datas []*models.Notes, err error) {
	orm := GetConn()
	orm = orm.Table(new(models.Notes).TableName())
	defer orm.Close()
	err = orm.Where("uid=?", uid).Group("notes_id").Find(&datas).Error
	return
}

type NotesMange struct {
	models.Notes
	N int
}

//分页查询笔记本信息
func (this *DaoNotes) GetNotesPgsInfo(uid string, pg int, size int) ([]*NotesMange, error) {
	orm := GetConn()
	orm = orm.Table(new(models.Notes).TableName())
	defer orm.Close()

	dataList := make([]*NotesMange, 0)
	//SELECT B.notes_id,B.notes_name,B.notes_des,B.notes_createtime,
	//COUNT(A.md_notes_id) AS n
	//from tbl_md_info AS A RIGHT JOIN tbl_notes AS B on A.md_notes_id=B.notes_id where B.uid = 'admin' GROUP BY B.notes_id LIMIT 0,20
	sqlStr := fmt.Sprintf("SELECT B.notes_id,B.notes_name,B.notes_des,B.notes_createtime,COUNT(A.md_notes_id) AS n from"+
		" tbl_md_info AS A RIGHT JOIN tbl_notes AS B on A.md_notes_id=B.notes_id where B.uid = '%s' GROUP BY B.notes_id LIMIT %d,%d",
		uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}

type Number struct {
	N int
}

//查询默认笔记本的笔记数量
func (this *DaoNotes) NotesNumber(uid string, notesid int) (int, error) {
	orm := GetConn()
	orm = orm.Table(new(models.Notes).TableName())
	defer orm.Close()

	var number = &Number{}
	sqlStr := fmt.Sprintf("SELECT count(*) as n FROM `tbl_md_info` where uid = '%s' and md_notes_id = %d", uid, notesid)
	err := orm.Raw(sqlStr).Scan(&number).Error
	return number.N, err
}

func (this *DaoNotes) NotesAllNumber(uid string) (int, error) {
	orm := GetConn()
	orm = orm.Table(new(models.Notes).TableName())
	defer orm.Close()

	var number = &Number{}
	sqlStr := fmt.Sprintf("SELECT count(*) as n FROM `tbl_md_info` where uid = '%s'", uid)
	err := orm.Raw(sqlStr).Scan(&number).Error
	return number.N, err
}

func (this *DaoNotes) UpdateInfo(datas *object.UpdateNotes, uid string) error {
	orm := GetConn()
	orm = orm.Table(new(models.Notes).TableName())
	defer orm.Close()

	notesid, _ := new(util.Str).NumberToInt(datas.NotesID)
	sqlStr := fmt.Sprintf("update tbl_notes set notes_name='%s',notes_des='%s' where notes_id=%d and uid = '%s'; ",
		datas.NotesName, datas.NotesDes, notesid, uid)
	return orm.Exec(sqlStr).Error
}

func (this *DaoNotes) DeleteInfo(nid int, uid string) error {
	orm := GetConn()
	orm = orm.Begin()
	defer orm.Close()

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

//获取指定用户笔记本数量
func (this *DaoNotes) GetNotesCount(uid string) (int64, error) {
	orm := GetConn()
	orm = orm.Table(new(models.Notes).TableName())
	defer orm.Close()

	var count int64
	err := orm.Model(&this).Where("uid = ?", uid).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
