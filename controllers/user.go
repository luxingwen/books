package controllers

import (
	"books/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	BaseController
}

func (this *UserController) Register() {
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	var user *models.User
	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("register func, json unmarshal err:%v", err)
		return
	}
	fmt.Println("register body: ", string(body))
	if user.UserName == "" || user.PassWord == "" {
		return
	}
	err := user.Add()
	if err != nil {
		fmt.Println("register err:", err)
		this.Fail(2, "register err.")
	}
	this.Succuess(user)
}

func (this *UserController) Update() {
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	var user *models.User
	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("register func, json unmarshal err:%v", err)
		return
	}
	if err := user.Update(); err != nil {
		fmt.Println("update error. error: ", err)
		return
	}
	this.Succuess(user)
}

func (this *UserController) Login() {
	fmt.Println("login...")
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	var user struct {
		UserName string `json:"username"`
		PassWord string `json:"password"`
		Email    string `json:"email"`
	}
	fmt.Println("body: ", string(body))
	err := json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Login func: json unmar shal err:%v", err)
		this.Fail(1, "login failer.")
	}
	rUser, err := models.Login(user.UserName, user.Email, user.PassWord)
	if err != nil {
		fmt.Println("err: ", err)
		this.Fail(1, "login failer.")
	}
	this.Succuess(rUser)
}
