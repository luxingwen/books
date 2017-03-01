package controllers

import (
	"books/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type CommunityController struct {
	AppController
}

func (this *CommunityController) Add() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	var com *models.Community
	err := json.Unmarshal(body, &com)
	if err != nil {
		this.Fail(101, "不能解析json")
	}
	com.BookId = id
	com.UserId = this.User.Id
	if err = com.Add(); err != nil {
		this.Fail(400, "评论出错")
	}
	this.Succuess(com)
}
