package routers

import (
	"beego_one_app/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/welcome", &controllers.WelcomeController{}, "get:WelcomeIndex")
}
