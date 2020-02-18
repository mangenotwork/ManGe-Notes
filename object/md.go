package object

//创建MD笔记
type CMDData struct {
	Title string `json:"md_title"`
	Detail string `json:"md_detail"`
	NotesId string `json:"notesid"`
	Tags string `json:"tags"`
}

//返回接口的笔记信息
type ReturnNoteInfo struct {
	Title     string `json:"title"`     //title
	Des string `json:"desc"`  //描述
	IsImg int `json:"isimg"` //是否有图片
	ImgLink string `json:"imglink"`//图片链接
	Id     string    `json:"id"`     //id
	Savetime       string `json:"time"`       //上次保存时间
	Tags      string `json:"tags"`      //标签
	ViewTimes   int  `json:"view_times"`   //查看次数
	Modifytimes int `json:"modify_times"` //修改次数
}
