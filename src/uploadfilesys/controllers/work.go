package controllers

import (
    "github.com/astaxie/beego"
    "uploadfilesys/models"
)

type WorkController struct {
    beego.Controller
}

func (ctrl *WorkController) PostWork() {
    title := ctrl.GetString("title")
    err := models.InsertWork(title)
    if err != nil {
        ctrl.Ctx.WriteString("发布作业失败")
    }
    ctrl.Ctx.WriteString("发布作业成功")
}

func (ctrl *WorkController) GetWork() {
    var works []*models.Work
    works, err := models.GetWorks()
    if err != nil {

    }
    ctrl.Data["json"] = &works
    println(ctrl.Data["json"])
    ctrl.ServeJSON()
}
