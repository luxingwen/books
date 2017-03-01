package controllers

import (
	"books/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type ShopCarController struct {
	AppController
}

func (this *ShopCarController) Add() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	var com *models.ShopCar
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	err := json.Unmarshal(body, &com)
	if err != nil {
		this.Fail(400, err.Error())
	}
	if id != com.BookId {
		this.Fail(400, "非法操作")
	}
	com.UserId = this.User.Id
	if err := com.Add(); err != nil {
		this.Fail(400, "添加购物车出错")
	}
	this.Succuess(com)
}

func (this *ShopCarController) Remove() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	var com *models.ShopCar
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	err := json.Unmarshal(body, &com)
	if err != nil {
		this.Fail(400, err.Error())
	}
	if id != com.BookId {
		this.Fail(400, "非法操作")
	}
	com.UserId = this.User.Id
	if err := com.Del(); err != nil {
		this.Fail(400, "remove购物车出错")
	}
	this.Succuess(com)
}

func (this *ShopCarController) List() {
	ids, err := models.MyShopcarBookIds(this.User.Id)
	if err != nil {
		this.Fail(2, err.Error())
	}
	books, err := models.BookInId(ids)
	if err != nil {
		this.Fail(2, err.Error())
	}
	this.Succuess(books)
}
