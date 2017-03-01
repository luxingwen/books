package routers

import (
	"books/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/api/login", &controllers.UserController{}, "post:Login")

	apiNs := beego.NewNamespace("/api",
		beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
		beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
		beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
		beego.NSRouter("/books", &controllers.BaseController{}, "get:BookList"),
		beego.NSRouter("/books/:id([0-9]+)", &controllers.BaseController{}, "get:BookListByTag"),
		beego.NSRouter("/book", &controllers.BookController{}, "post:Add"),
		beego.NSRouter("/book/:id([0-9]+)", &controllers.BookController{}, "put:Update"),
		beego.NSRouter("/book/info/:id([0-9]+)", &controllers.BookController{}, "get:BookInfo"),
		beego.NSRouter("/book/buy", &controllers.BookController{}, "post:BuyBook"),
		beego.NSRouter("/shopcar/:id([0-9]+)", &controllers.ShopCarController{}, "post:Add;put:Remove"),
		beego.NSRouter("/book/info/:id([0-9]+)/community", &controllers.CommunityController{}, "post:Add"),
	)
	beego.AddNamespace(apiNs)
}
