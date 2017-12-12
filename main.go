package main

import (
	_ "beego_one_app/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "./static")
	beego.SetStaticPath("/docs", "./static/docs")
	beego.Run()
}
