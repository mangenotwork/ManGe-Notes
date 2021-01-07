package dao

import (
	"github.com/mangenotwork/ManGe-Notes/models"
)

type DaoACC struct{}

func (this *DaoACC) CreateUser(newuser *models.ACC) error {
	orm := GetConn()
	orm = orm.Table(new(models.ACC).TableName())
	defer orm.Close()

	return orm.Create(newuser).Error
}

//通过用户账号查询用户信息
func (this *DaoACC) GetACCinfo(acc string) (data *models.ACC, err error) {
	orm := GetConn()
	orm = orm.Table(new(models.ACC).TableName())
	defer orm.Close()

	data = &models.ACC{}
	err = orm.Where("account = ?", acc).First(&data).Error
	return
}

//更新最后登陆时间和最后登陆ip
func (this *DaoACC) UpdateLastLogin(uid string, lasttime int64, lsatip string) (err error) {
	orm := GetConn()
	orm = orm.Table(new(models.ACC).TableName())
	defer orm.Close()

	updateData := map[string]interface{}{"logintime": lasttime, "loginip": lsatip}
	return orm.Model(&this).Where("userid = ?", uid).Updates(updateData).Error
}

//判断用户账号是否存在
func (this *DaoACC) ACCIsAccount(acc string) bool {
	orm := GetConn()
	orm = orm.Table(new(models.ACC).TableName())
	defer orm.Close()

	err := orm.Where("account = ?", acc).First(this).Error
	if err != nil && err.Error() == "record not found" {
		return true
	}
	return false
}

//判断用户绑定手机号是否存在
func (this *DaoACC) ACCIsPhone(phone string) {

}

//判断用户绑定邮箱是否存在
func (this *DaoACC) ACCIsMail(mail string) {

}
