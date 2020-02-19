package models

import (
	"fmt"
	"time"

	conn "man/ManNotes/conn"
)

// tbl_md_info表
type MDInof struct {
	MDTempid     int `gorm:"column:md_tempid;primary_key;AUTO_INCREMENT"`     //临时自增id
	MDTitle     string `gorm:"column:md_title"`     //title
	MDDes 	string `gorm:"column:md_des"`     //前30个字符的内容
	MDIMG   string `gorm:"column:md_img"`   //如果笔记内容有图片链接保存第一张图片的链接
	IsIMG   int `gorm:"column:is_img"`    //是否有图片，如果有为1  没有为0
	MDId     string    `gorm:"column:md_id"`     //内容id
	Uid string `gorm:"column:uid"`     //用户id 
	MDNotesid int `gorm:"column:md_notes_id"` //属于的笔记本分类
	MDCreatetime        int64    `gorm:"column:md_createtime"`        //首次创建时间
	MDSavetime       int64 `gorm:"column:md_savetime"`       //上次保存时间
	MDTag      string `gorm:"column:md_tag"`      //标签
	MDOpentime  int64 `gorm:"column:md_opentime"`  //上次打开时间
	MDViewTimes   int  `gorm:"column:md_viewtimes"`   //查看次数
	MDModifytimes int `gorm:"column:md_modifytimes"` //修改次数
}

func (this *MDInof) TableName() string {
	return "tbl_md_info"
}

func (this *MDInof) InsertMDNote(mdtext *MDText) error {
	orm := conn.NotesDB().Begin()
	err := orm.Create(mdtext).Error
	if err != nil {
		orm.Rollback()
		return err
	}
	err = orm.Create(this).Error
	if err != nil {
		orm.Rollback()
		return err
	}
	return orm.Commit().Error
}

//按照页数获取数据
func (this *MDInof) GetToPG(pg int, size int, uid string) ([]*MDInof,error) {
	orm := conn.NotesDB()
	dataList := make([]*MDInof, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_md_info where uid = '%s' LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList,err
}

//按照笔记本和页数获取数据
func (this *MDInof) GetNotesToPG(pg int, size int, uid string, notesid int) ([]*MDInof,error) {
	orm := conn.NotesDB()
	dataList := make([]*MDInof, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_md_info where uid = '%s' and md_notes_id = %d LIMIT %d,%d", uid, notesid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList,err
}

//通过uid mid 判断是否存在数据
func (this *MDInof) IsMD(uid,mid string) (bool,string,error) {
	orm := conn.NotesDB()
	err := orm.Where("uid = ? and md_id = ?", uid,mid).First(this).Error
	if err != nil && err.Error() == "record not found"{
		return false,"",err
	}
	return true,this.MDTitle,nil
}

//增加查看次数
func (this *MDInof) AddMDViewTimes(mid string) error {
	orm := conn.NotesDB()
	sqlStr := fmt.Sprintf("update tbl_md_info set md_viewtimes=md_viewtimes+1 where md_id='%s'; ", mid)
	return orm.Exec(sqlStr).Error
}

//增加修改次数
func (this *MDInof) AddMDModifytimes(mid string) error {
	orm := conn.NotesDB()
	sqlStr := fmt.Sprintf("update tbl_md_info set md_modifytimes=md_modifytimes+1 where md_id='%s'; ", mid)
	return orm.Exec(sqlStr).Error
}

//修改笔记名
func (this *MDInof) UpdateMDNote(mdcontent string,mid string) error {
	orm := conn.NotesDB().Begin()
	sqlMDInofStr := fmt.Sprintf("update tbl_md_info set md_title='%s',md_des='%s',md_img='%s',is_img=%d,md_opentime=%d,md_modifytimes=md_modifytimes+1 where md_id='%s'; ", 
		this.MDTitle, this.MDDes, this.MDIMG, this.IsIMG, time.Now().Unix(), mid)
	sqlMDTextStr := fmt.Sprintf("update tbl_md_text set md_content='%s' where md_id='%s'; ", mdcontent, mid)
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
func (this *MDInof) SearchTitle(word,uid string) ([]*MDInof,error) {
	orm := conn.NotesDB()
	dataList := make([]*MDInof, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_md_info where uid = '%s' and md_title like '%%%s%%' LIMIT %d,%d", uid, word, 0, 100)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList,err
}