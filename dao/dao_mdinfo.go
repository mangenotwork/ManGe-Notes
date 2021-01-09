package dao

import (
	"fmt"
	"time"

	"github.com/mangenotwork/ManGe-Notes/models"
)

type DaoMDInof struct{}

func (this *DaoMDInof) InsertMDNote(mdtext *models.MDText, mdinfo *models.MDInof) error {
	orm := GetConn()
	orm = orm.Begin()
	defer orm.Close()

	err := orm.Create(mdinfo).Error
	if err != nil {
		orm.Rollback()
		return err
	}

	err = orm.Create(mdtext).Error
	if err != nil {
		orm.Rollback()
		return err
	}

	return orm.Commit().Error
}

//按照页数获取数据
func (this *DaoMDInof) GetToPG(pg int, size int, uid string) ([]*models.MDInof, error) {
	orm := GetConn()
	orm = orm.Table(new(models.MDInof).TableName())
	defer orm.Close()

	dataList := make([]*models.MDInof, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_md_info where uid = '%s' and md_notes_id != -1 LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}

//获取回收站笔记
func (this *DaoMDInof) GetRecycler(pg int, size int, uid string) ([]*models.MDInof, error) {
	orm := GetConn()
	orm = orm.Table(new(models.MDInof).TableName())
	defer orm.Close()

	dataList := make([]*models.MDInof, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_md_info where uid = '%s' and md_notes_id = -1 LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}

//获取草稿笔记
func (this *DaoMDInof) DraftNote(pg int, size int, uid string) ([]*models.MDInof, error) {
	orm := GetConn()
	orm = orm.Table(new(models.MDInof).TableName())
	defer orm.Close()

	dataList := make([]*models.MDInof, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_md_info where uid = '%s' and md_notes_id = -2 LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}

//按照笔记本和页数获取数据
func (this *DaoMDInof) GetNotesToPG(pg int, size int, uid string, notesid int) ([]*models.MDInof, error) {
	orm := GetConn()
	orm = orm.Table(new(models.MDInof).TableName())
	defer orm.Close()

	dataList := make([]*models.MDInof, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_md_info where uid = '%s' and md_notes_id = %d LIMIT %d,%d", uid, notesid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}

//通过uid mid 判断是否存在数据
func (this *DaoMDInof) IsMD(uid, mid string) (bool, string, error) {
	orm := GetConn()
	orm = orm.Table(new(models.MDInof).TableName())
	defer orm.Close()

	data := &models.MDInof{}
	err := orm.Where("uid = ? and md_id = ?", uid, mid).First(&data).Error
	if err != nil && err.Error() == "record not found" {
		return false, "", err
	}
	return true, data.MDTitle, nil
}

//增加查看次数
func (this *DaoMDInof) AddMDViewTimes(mid string) error {
	orm := GetConn()
	orm = orm.Table(new(models.MDInof).TableName())
	defer orm.Close()

	sqlStr := fmt.Sprintf("update tbl_md_info set md_viewtimes=md_viewtimes+1 where md_id='%s'; ", mid)
	return orm.Exec(sqlStr).Error
}

//增加修改次数
func (this *DaoMDInof) AddMDModifytimes(mid string) error {
	orm := GetConn()
	orm = orm.Table(new(models.MDInof).TableName())
	defer orm.Close()

	sqlStr := fmt.Sprintf("update tbl_md_info set md_modifytimes=md_modifytimes+1 where md_id='%s'; ", mid)
	return orm.Exec(sqlStr).Error
}

//修改笔记名
func (this *DaoMDInof) UpdateMDNote(data *models.MDInof, mdcontent string, mid string) error {
	orm := GetConn()
	orm = orm.Begin()
	defer orm.Close()

	sqlMDInofStr := fmt.Sprintf("update tbl_md_info set md_title='%s',md_des='%s',md_img='%s',is_img=%d,md_opentime=%d,md_modifytimes=md_modifytimes+1 where md_id='%s'; ",
		data.MDTitle, data.MDDes, data.MDIMG, data.IsIMG, time.Now().Unix(), mid)
	fmt.Println("[Sql] = ", sqlMDInofStr)
	sqlMDTextStr := fmt.Sprintf("update tbl_md_text set md_content='%s' where md_id='%s'; ", mdcontent, mid)
	fmt.Println("[Sql] = ", sqlMDTextStr)
	err := orm.Exec(sqlMDInofStr).Error
	if err != nil {
		orm.Rollback()
		return err
	}
	err = orm.Exec(sqlMDTextStr).Error
	if err != nil {
		orm.Rollback()
		return err
	}
	return orm.Commit().Error
}

//通过title 模糊查询笔记信息
func (this *DaoMDInof) SearchTitle(word, uid string) ([]*models.MDInof, error) {
	orm := GetConn()
	orm = orm.Table(new(models.MDInof).TableName())
	defer orm.Close()

	dataList := make([]*models.MDInof, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_md_info where uid = '%s' and md_title like '%%%s%%' LIMIT %d,%d", uid, word, 0, 100)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}

//删除笔记到回收站 md_notes_id = -1
func (this *DaoMDInof) ToDEL(mdid, uid string) error {
	orm := GetConn()
	orm = orm.Table(new(models.MDInof).TableName())
	defer orm.Close()

	sqlStr := fmt.Sprintf("update tbl_md_info set md_notes_id=-1 where md_id='%s' and uid = '%s'; ", mdid, uid)
	return orm.Exec(sqlStr).Error
}

//永久删除笔记
func (this *DaoMDInof) Schen(mdid, uid string) error {
	orm := GetConn()
	orm = orm.Table(new(models.MDInof).TableName())
	defer orm.Close()

	sqlStr := fmt.Sprintf("DELETE FROM tbl_md_info where md_id='%s' and uid = '%s' and md_notes_id in (-1,-2) ; ", mdid, uid)
	return orm.Exec(sqlStr).Error
}

//笔记转移到指定笔记本
func (this *DaoMDInof) NoteToNotes(mdid, uid string, notes int) error {
	orm := GetConn()
	orm = orm.Table(new(models.MDInof).TableName())
	defer orm.Close()

	sqlStr := fmt.Sprintf("update tbl_md_info set md_notes_id=%d where md_id='%s' and uid = '%s'; ", notes, mdid, uid)
	return orm.Exec(sqlStr).Error
}

type CountNumber struct {
	Count int64
}

//获取指定用户笔记数量
func (this *DaoMDInof) GetNoteCount(uid string) (int64, error) {
	orm := GetConn()
	orm = orm.Table(new(models.MDInof).TableName())
	defer orm.Close()

	var count int64
	err := orm.Model(&this).Where("uid = ?", uid).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

//统计模块 获取笔记本数量
func (this *DaoMDInof) GetData2Census(notesids []int) (datas []*models.MDInof, err error) {
	orm := GetConn()
	orm = orm.Table(new(models.MDInof).TableName())
	defer orm.Close()

	err = orm.Where("md_notes_id in (?)", notesids).Find(&datas).Error
	return
}
