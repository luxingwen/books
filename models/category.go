package models

type Category struct {
	Id       int    `xorm:"pk autoincr" json:"id"`
	Content  string `json:"content"`
	FatherId int
}
