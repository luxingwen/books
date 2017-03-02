package controllers

import (
	"books/models"
)

type MyBookController struct {
	AppController
}

func (this *MyBookController) MyBookList() {
	ids, err := models.MyBookIds(this.User.Id)
	if err != nil {
		this.Fail(2, err.Error())
	}
	books, err := models.BookInId(ids)
	if err != nil {
		this.Fail(2, err.Error())
	}

	mCategory, err := models.CategoryToMap()
	var res []*RspBook
	for _, item := range books {
		book := &RspBook{Id: item.Id, Name: item.Name, Autor: item.Autor, Desc: item.Desc, Pic: item.Pic, Money: item.Money}
		if v, ok := mCategory[item.Category]; ok {
			book.Category = v
		} else {
			book.Category = &models.Category{Id: 0, Content: "其它分类", FatherId: 0}
		}
		res = append(res, book)
	}
	this.Succuess(res)
}
