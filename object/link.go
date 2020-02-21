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