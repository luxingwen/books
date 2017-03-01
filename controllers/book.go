package controllers

import (
	"books/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
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

type RspBook struct {
	Id       int              `json:"id"`
	Name     string           `json:"name"`
	Autor    string           `json:"autor"`
	Desc     string           `json:"desc"`
	Pic      string           `json:"pic"`
	Money    float64          `json:"money"`
	Category *models.Category `json:"category"`
}

func (this *BaseController) BookList() {
	list, err := models.BookList()
	if err != nil {
		this.Fail(400, "获取图书列表失败")
	}
	mCategory, err := models.CategoryToMap()
	var res []*RspBook
	for _, item := range list {
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

	type Community struct {
		Id        int          `json:"id"`
		Content   string       `json:"content"`
		User      *models.User `json:"user"`
		CreatedAt time.Time    `xorm:"created"`
	}
	type BookInfo struct {
		Book       *models.Book `json:"book"`
		Communitys []*Community `json:"community"`
	}

	cList, err := models.CommunityListByBook(bookId)
	if err != nil {
		this.Fail(400, "获取评论列表失败:"+err.Error())
	}
	var userIds []int
	var communitys []*Community
	for _, item := range cList {
		userIds = append(userIds, item.UserId)
	}
	users, err := models.GetUserInId(userIds)
	if err != nil {
		this.Fail(2, err.Error())
	}
	mUser := make(map[int]*models.User)
	for _, item := range users {
		mUser[item.Id] = item
	}

	for _, item := range cList {
		if v, ok := mUser[item.UserId]; ok {
			c := &Community{Id: item.Id, Content: item.Content, CreatedAt: item.CreatedAt, User: v}
			communitys = append(communitys, c)
		}
	}

	res := &BookInfo{Book: book, Communitys: communitys}
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
