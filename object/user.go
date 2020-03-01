package object

//用户注册上传的信息
type UserRegInfo struct {
	Acc string `json:"acc"`
	Password string `json:"password"`
	InviteCode string `json:"invite_code"`
	Phone string `json:"phone"`
}

//用户账号密码登录
type Logininfo struct {
	LoginAcc string `json:"loginacc"`
	LoginPassword string `json:"loginpassword"`
}

type UserBasisInfo struct {
	UName string `json:"user_name"`
	UAvatar string `json:"user_avatar"`
	LastLogin string `json:"last_login"`
	LastLoginIp string `json:"last_loginip"`
}

//用户是使用空间  笔记数量 素材存储空间 笔记本数量
type UsedInfo struct{
	NoteMax int64 
	NoteNow int64
	NotePercent int64
	SpaceMax string
	SpaceNow string
	SpacePercent int64
	NotesMax int64
	NotesNow int64
	NotesPercent int64
}