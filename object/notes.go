package object


//创建笔记本
type CNotes struct {
	NotesName string `json:"notes_name"`
	Description string `json:"description"`
}

//返回给页面的笔记本列表信息
type NotesInfo struct {
	NotesID string `json:"notes_id"`
	NotesName string `json:"notes_name"`
	NoteNumber int `json:"note_number"`
}

type ReturnNotes struct {
	All int `json:"notes_all"` //全部笔记
	Def int `json:"notes_def"` //默认笔记
	Dar int `json:"notes_dar"` //草稿
	Rec int `json:"notes_rec"` //回收站
	Notes []*NotesInfo `json:"notes"`
}

//返回给mange 的笔记本信息
type MangeNotesList struct{
	NotesID int `json:"notes_id"`
	NotesName string `json:"notes_name"`
	NotesDes string `json:"notes_des"`
	CreateTime string `json:"create"`
	NoteNumber int `json:"note_number"`
} 

//修改笔记本信息传参
type UpdateNotes struct{
	NotesID string `json:"notes_id"`
    NotesName string `json:"notes_name"`
    NotesDes string `json:"notes_des"`
}
