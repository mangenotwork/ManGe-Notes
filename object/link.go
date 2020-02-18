package object

//添加网络资源 post传参
type AddLink struct {
	LinkType string `json:"type"`
	Link string `json:"link"`
	LinkName string `json:"name"`
	LinkDes string `json:"descr"`
}