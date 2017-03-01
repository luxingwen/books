package controllers

import (
	"books/models"
	"fmt"
)

type AppController struct {
	BaseController
	User *models.User
}

func (this *AppController) Prepare() {
	// if models.Dev {
	// 	user, err := models.GetUser(1)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	this.SetSession("user", user)
	// }

	token := this.GetString("token", "")
	fmt.Println("token...:", token)
	if token != "" {
		user, err := models.GetUserByToken(token)
		if err != nil {
			fmt.Println("getuser bytoken err. ", err)
		} else {
			this.SetSession("user", user)
		}
	}
	user := this.GetSession("user")

	if user == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "msg": "无效用户，请登陆"}
		this.ServeJSON()
		this.StopRun()
	}
	if v, ok := user.(*models.User); ok {
		this.User = v
	}

	fmt.Println("user----> ", user)
}
