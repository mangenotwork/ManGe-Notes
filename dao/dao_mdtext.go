package dao

import (
	"fmt"

	"github.com/mangenotwork/ManGe-Notes/models"
)

type DaoMDText struct{}

//通过mid查询MD内容
func (this *DaoMDText) GetMDTxt(mid string) (string, error) {
	orm := GetConn()
	orm = orm.Table(new(models.MDText).TableName())
	defer orm.Close()

	data := &models.MDText{}
	err := orm.Where("md_id = ?", mid).First(data).Error
	if err != nil {
		fmt.Println("通过mid查询MD内容错误", err)
		return "", err
	}
	return data.MDContent, nil
}
