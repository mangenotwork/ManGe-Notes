package object

//添加网络资源 post传参
type AddLink struct {
	LinkType string `json:"type"`
	Link string `json:"link"`
	LinkName string `json:"name"`
	LinkDes string `json:"descr"`
}

//修改收藏的链接传参
type EDLinks struct {
	LinkID int `json:"link_id"`
	LinkName string `json:"link_title"`
	Link string `json:"link"`
	LinkDes string `json:"link_des"`
}

//mange 管理模块 收藏链接
type LinksInfo struct{
	Id     int   `json:"link_id"`   //笔记本id
	Name     string  `json:"link_name"`    //工具或链接名称
	Des 	string  `json:"link_des"`//工具或链接描述
	Link  string    `json:"link_url"`  //工具或链接 地址
	Ico  string   `json:"link_ico"`  //工具或链接 地址 ico
	Tag  string   `json:"link_tag"`  //工具或链接 标签
	LinkTypeStr string `json:"link_type"`
}

type EDLinksInfo struct{
	Id     string   `json:"link_id"`
	Name     string  `json:"link_name"`    //工具或链接名称
	Des 	string  `json:"link_des"`//工具或链接描述
	Link  string    `json:"link_url"`  //工具或链接 地址
	Ico  string   `json:"link_ico"`  //工具或链接 地址 ico
	Tag  string   `json:"link_tag"`  //工具或链接 标签
	LinkType string `json:"link_type"`
}