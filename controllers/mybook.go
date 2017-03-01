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
	this.Succuess(books)
}
