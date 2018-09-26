package main

import (
	_ "todo/routers"
	"github.com/astaxie/beego"
	_ "todo/controllers"
)

func main() {
	beego.Run()
}

