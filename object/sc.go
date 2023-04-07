package object

type UpLoadImgData struct{
	ImgName string `json:"imgname"`
	ImgTag string `json:"imgnametag"`
}

//添加网络图片链接
type LinkImgData struct {
	ImgName string `json:"imgname"`
	ImgTag string `json:"imgnametag"`
	ImgLink string `json:"imglink"`
}

//IMG的显示数据
type ImgInfoShow struct {
	ImgId int `json:"img_id"`
	ImgUrl string `json:"img_url"`
	ImgName string `json:"img_name"`
	ImgTag string `json:"img_tag"`
	ImgCreate string `json:"img_create"`
}

//漫鸽图库 图片数据列表显示
type ManGeImgList struct {
	ImgId int `json:"img_id"`
	ImgUrl string `json:"img_url"`
	ImgName string `json:"img_name"`
}

//漫鸽图库图片信息显示