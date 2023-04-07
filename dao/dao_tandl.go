package dao

import (
	"fmt"

	"github.com/mangenotwork/ManGe-Notes/models"
)

type DaoToolandLink struct{}

func (this *DaoToolandLink) Create(toolandLink *models.ToolandLink) error {
	orm := GetConn()
	orm = orm.Table(new(models.ToolandLink).TableName())
	defer orm.Close()

	return orm.Create(toolandLink).Error
}

//获取tool 列表
func (this *DaoToolandLink) GeTools(pg int, size int, uid string) ([]*models.ToolandLink, error) {
	orm := GetConn()
	orm = orm.Table(new(models.ToolandLink).TableName())
	defer orm.Close()

	dataList := make([]*models.ToolandLink, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_tool_link where uid='%s' and type=1 LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}

//获取链接 列表
func (this *DaoToolandLink) GetLinks(pg int, size int, uid string) ([]*models.ToolandLink, error) {
	orm := GetConn()
	orm = orm.Table(new(models.ToolandLink).TableName())
	defer orm.Close()

	dataList := make([]*models.ToolandLink, 0)
	sqlStr := fmt.Sprintf("SELECT * FROM tbl_tool_link where uid='%s' and type=0 LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}

//修改收藏的链接
func (this *DaoToolandLink) UpdateLink(data *models.ToolandLink) error {
	orm := GetConn()
	orm = orm.Table(new(models.ToolandLink).TableName())
	defer orm.Close()

	sqlStr := fmt.Sprintf("update tbl_tool_link set name='%s',des='%s',link='%s',ico='%s' where id=%d and uid = '%s'; ",
		data.Name, data.Des, data.Link, data.Ico, data.Id, data.UID)
	return orm.Exec(sqlStr).Error
}

//删除收藏的链接
func (this *DaoToolandLink) DELLink(uid string, linkid int) error {
	orm := GetConn()
	orm = orm.Table(new(models.ToolandLink).TableName())
	defer orm.Close()

	sqlStr := fmt.Sprintf("DELETE FROM tbl_tool_link where uid='%s' and id=%d and type=0; ", uid, linkid)
	return orm.Exec(sqlStr).Error
}

//GetAll 获取所有收藏链接
func (this *DaoToolandLink) GetAll(pg int, size int, uid string) ([]*models.ToolandLink, error) {
	orm := GetConn()
	orm = orm.Table(new(models.ToolandLink).TableName())
	defer orm.Close()

	dataList := make([]*models.ToolandLink, 0)
	sqlStr := fmt.Sprintf("SELECT id,name,des,link,ico,tag,type FROM tbl_tool_link where uid='%s' LIMIT %d,%d", uid, pg, size)
	err := orm.Raw(sqlStr).Scan(&dataList).Error
	return dataList, err
}

// EDLink  修改全部字段链接信息
func (this *DaoToolandLink) EDLink(toolandLink *models.ToolandLink) error {
	orm := GetConn()
	orm = orm.Table(new(models.ToolandLink).TableName())
	defer orm.Close()

	return orm.Save(&toolandLink).Error
}
