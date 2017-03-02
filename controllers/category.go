package controllers

import (
	"books/models"
	"fmt"
)

type CategoryController struct {
	BaseController
}

func (this *CategoryController) List() {
	list, err := models.CategoryList()
	if err != nil {
		fmt.Println("category List...", err)
		this.Fail(400, err.Error())
	}
	this.Succuess(list)
}
