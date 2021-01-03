package models

import (
	"fmt"

	conn "github.com/mangenotwork/ManGe-Notes/conn"
)

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

func (this *ACC) CreateUser() error {
	orm := conn.NotesDB()
	return orm.Create(this).Error
}

//通过用户账号查询用户信息
func (this *ACC) GetACCinfo(acc string) (err error) {
	orm := conn.NotesDB()
	err = orm.Where("account = ?", acc).First(this).Error
	return
}

//更新最后登陆时间和最后登陆ip
func (this *ACC) UpdateLastLogin(uid string, lasttime int64, lsatip string) (err error) {
	orm := conn.NotesDB()
	fmt.Println(conn.GetDBConns("notes"))
	updateData := map[string]interface{}{"logintime": lasttime, "loginip": lsatip}
	return orm.Model(&this).Where("userid = ?", uid).Updates(updateData).Error
}

//判断用户账号是否存在
func (this *ACC) ACCIsAccount(acc string) bool {
	orm := conn.NotesDB()
	err := orm.Where("account = ?", acc).First(this).Error
	if err != nil && err.Error() == "record not found" {
		return true
	}
	return false
}

//判断用户绑定手机号是否存在
func (this *ACC) ACCIsPhone(phone string) {

}

//判断用户绑定邮箱是否存在
func (this *ACC) ACCIsMail(mail string) {

}
