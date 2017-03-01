package controllers

import (
	"books/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type BookController struct {
	AppController
}

func (this *BookController) Add() {
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	var book *models.Book
	err := json.Unmarshal(body, &book)
	if err != nil {
		this.Fail(101, "不能解析json")
	}
	if err = book.Add(); err != nil {
		this.Fail(400, "添加图书出错")
	}
	this.Succuess(book)
}

func (this *BookController) Update() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	var book *models.Book
	err := json.Unmarshal(body, &book)
	if err != nil {
		this.Fail(101, "不能解析json")
	}
	book.Id = id
	if err = book.Update(); err != nil {
		this.Fail(400, " 更新图书出错")
	}
	this.Succuess(book)
}

func (this *BaseController) BookList() {
	list, err := models.BookList()
	if err != nil {
		this.Fail(400, "获取图书列表失败")
	}
	this.Succuess(list)
}

func (this *BaseController) BookListByTag() {
	id := this.Ctx.Input.Param(":id")
	t, _ := strconv.Atoi(id)
	list, err := models.BookListByCategory(t)
	if err != nil {
		this.Fail(400, "获取图书列表失败")
	}
	this.Succuess(list)
}

func (this *BaseController) BookInfo() {
	id := this.Ctx.Input.Param(":id")
	bookId, _ := strconv.Atoi(id)
	book := &models.Book{Id: bookId}
	err := book.Get()
	if err != nil {
		this.Fail(400, "获取图书列表失败:"+err.Error())
	}

	type BookInfo struct {
		Book       *models.Book
		Communitys []*models.Community
	}

	cList, err := models.CommunityListByBook(bookId)
	if err != nil {
		this.Fail(400, "获取评论列表失败:"+err.Error())
	}
	res := &BookInfo{Book: book, Communitys: cList}
	this.Succuess(res)

}

func (this *BookController) BuyBook() {
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	var mybook *models.MyBook
	err := json.Unmarshal(body, &mybook)
	if err != nil {
		this.Fail(101, "不能解析json")
	}
	book, err := models.BuyBook(mybook.BookId, this.User.Id)
	if err != nil {
		this.Fail(101, "购买失败")
	}

	this.Succuess(book)
}
