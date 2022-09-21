package models

type ContentModel struct {
	Id       int    `json:"Id"`
	TitleEn  string `json:"TitleEn"`
	TitleTh  string `json:"TitleTh"`
	Detail   string `json:"Detail"`
	Author   string `json:"Author"`
	Organize string `json:"Organize"`
}
