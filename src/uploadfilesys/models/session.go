package models

import "github.com/astaxie/beego"

func CheckIsAdmin(token string) {
    admintoken_conf := beego.AppConfig.String("admintoken")
    if admintoken_conf == token {
        return true
    }
    return false
}
