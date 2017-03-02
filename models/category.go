package models

type Category struct {
	Id       int    `xorm:"pk autoincr" json:"id"`
	Content  string `json:"content"`
	FatherId int    `json:"fatherId"`
}

func CategoryToMap() (r map[int]*Category, err error) {
	r = make(map[int]*Category)
	err = engine.Find(&r)
	return
}

func CategoryList() (r []*Category, err error) {
	err = engine.Find(&r)
	return
}
