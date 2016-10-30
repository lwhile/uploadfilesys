package models

import "github.com/astaxie/beego"

func CheckIsAdmin(adminname, password string) {
    adminname_conf := beego.AppConfig.String("adminname")
    adminpassword_conf := beego.AppConfig.String("adminpassword")
    if adminname_conf == adminname && adminpassword_conf == password {
        return true
    }
    return false
}
