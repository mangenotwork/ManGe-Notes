package dao

import (
	"fmt"

	"github.com/mangenotwork/ManGe-Notes/models"
)

type DaoIMGInfo struct {
}

func (this *DaoIMGInfo) CreateImg(imginfo *models.IMGInfo) error {
	orm := GetConn()
	orm = orm.Table(new(models.IMGInfo).TableName())
	defer orm.Close()

	return orm.Create(imginfo).Error
}

//获取用户所有图片的大小
func (this *DaoIMGInfo) GetSize(uid string) (int64, error) {
	orm := GetConn()
	orm = orm.Table(new(models.IMGInfo).TableName())
	defer orm.Close()

	type S struct {
		Size int64
	}
	var size S
	sqlStr := fmt.Sprintf("select sum(size) as size from tbl_img where uid = '%s'", uid)
	err := orm.Raw(sqlStr).Scan(&size).Error
	return size.Size, err
}

//获取我的图片
func (this *DaoIMGInfo) GetMyImg(uid string) ([]*models.IMGInfo, error) {
	orm := GetConn()
	orm = orm.Table(new(models.IMGInfo).TableName())
	defer orm.Close()

	dataList := make([]*models.IMGInfo, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_img where uid = '%s' LIMIT %d,%d", uid, 0, 100)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}

//验证图片权限
func (this *DaoIMGInfo) IsMyImg(imgid int, uid string) (bool, *models.IMGInfo) {
	orm := GetConn()
	orm = orm.Table(new(models.IMGInfo).TableName())
	defer orm.Close()

	data := &models.IMGInfo{}
	err := orm.Where("id = ? and uid = ?", imgid, uid).First(&data).Error
	if err != nil && err.Error() == "record not found" {
		return true, data
	}
	return false, data
}
