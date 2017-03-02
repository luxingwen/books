package main

import (
	"books/models"
	_ "books/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.SyncModels()
	beego.SetStaticPath("/upload", "upload")
	beego.Run()
}
