package models

// tbl_md_info表
type MDInof struct {
	MDTempid      int    `gorm:"column:md_tempid;primary_key;AUTO_INCREMENT"` //临时自增id
	MDTitle       string `gorm:"column:md_title"`                             //title
	MDDes         string `gorm:"column:md_des"`                               //前30个字符的内容
	MDIMG         string `gorm:"column:md_img"`                               //如果笔记内容有图片链接保存第一张图片的链接
	IsIMG         int    `gorm:"column:is_img"`                               //是否有图片，如果有为1  没有为0
	MDId          string `gorm:"column:md_id"`                                //内容id
	Uid           string `gorm:"column:uid"`                                  //用户id
	MDNotesid     int    `gorm:"column:md_notes_id"`                          //属于的笔记本分类
	MDCreatetime  int64  `gorm:"column:md_createtime"`                        //首次创建时间
	MDSavetime    int64  `gorm:"column:md_savetime"`                          //上次保存时间
	MDTag         string `gorm:"column:md_tag"`                               //标签
	MDOpentime    int64  `gorm:"column:md_opentime"`                          //上次打开时间
	MDViewTimes   int    `gorm:"column:md_viewtimes"`                         //查看次数
	MDModifytimes int    `gorm:"column:md_modifytimes"`                       //修改次数
}

func (this *MDInof) TableName() string {
	return "tbl_md_info"
}
