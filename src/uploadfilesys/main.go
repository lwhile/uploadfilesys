package main

import (
	_ "uploadfilesys/routers"
	"github.com/astaxie/beego"
    "uploadfilesys/controllers"
)

func main() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/work", &controllers.WorkController{}, "get:GetWork")
    beego.Router("/work", &controllers.WorkController{}, "Post:PostWork")
    beego.Router("/file", &controllers.FileController{}, "post:PostFile")
    beego.Router("/file", &controllers.FileController{}, "get:DownloadAFile")
	beego.Run()
}

