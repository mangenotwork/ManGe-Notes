package dao

import (
	"fmt"

	"github.com/mangenotwork/ManGe-Notes/models"
)

type DaoSCIMGInfo struct{}

func (this *DaoSCIMGInfo) AddSCImg(scimgInfo *models.SCIMGInfo) error {
	orm := GetConn()
	orm = orm.Table(new(models.SCIMGInfo).TableName())
	defer orm.Close()

	return orm.Create(scimgInfo).Error
}

//获取漫鸽图库
func (this *DaoSCIMGInfo) MangeImgList() ([]*models.SCIMGInfo, error) {
	orm := GetConn()
	orm = orm.Table(new(models.SCIMGInfo).TableName())
	defer orm.Close()

	dataList := make([]*models.SCIMGInfo, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_sc_img LIMIT %d,%d", 0, 100)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}
