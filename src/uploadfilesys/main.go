package main

import (
	_ "uploadfilesys/routers"
	"github.com/astaxie/beego"
    "uploadfilesys/controllers"
    _ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego/orm"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:wh5622@/uploadfilesys?charset=utf8")
}

func main() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/work", &controllers.WorkController{}, "get:GetWork")
    beego.Router("/work", &controllers.WorkController{}, "Post:PostWork")
    beego.Router("/file", &controllers.FileController{}, "post:PostFile")
    beego.Router("/file", &controllers.FileController{}, "get:DownloadAFile")
	beego.Run()
}

