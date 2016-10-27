package main

import (
	_ "uploadfilesys/routers"
	"github.com/astaxie/beego"
)

func main() {
    beego.Router("/file")
	beego.Run()
}

