package models

// tbl_acc表
type ACC struct {
	UserId     string `gorm:"column:userid"`     //用户id
	Account    string `gorm:"column:account"`    //账号
	Password   string `gorm:"column:password"`   //密码
	Phone      string `gorm:"column:phone"`      //手机号
	Mail       string `gorm:"column:mail"`       //邮箱号
	Avatar     string `gorm:"column:avatar"`     //头像
	Createtime int64  `gorm:"column:createtime"` //首次创建时间
	Logintime  int64  `gorm:"column:logintime"`  //上次登录时间
	LoginIP    string `gorm:"column:loginip"`    //上次登录ip
}

func (this *ACC) TableName() string {
	return "tbl_acc"
}
