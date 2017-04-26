package models

import "github.com/astaxie/beego"

func CheckIsAdmin(token string) bool {
    admintoken_conf := beego.AppConfig.String("admintoken")
    println(admintoken_conf, " ", token)
    if admintoken_conf == token {
        return true
    }
    return false
}
