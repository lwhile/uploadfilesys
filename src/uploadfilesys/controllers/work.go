package controllers

import (
    "github.com/astaxie/beego"
    "uploadfilesys/models"
    "strings"
)

type WorkController struct {
    beego.Controller
}

//需要用户验证
func (ctrl *WorkController) PostWork() {
    title := strings.TrimSpace(ctrl.GetString("title"))
    err := models.InsertWork(title)
    if err != nil {
        ctrl.Ctx.WriteString("发布作业失败")
        return
    }
    ctrl.Ctx.WriteString("发布作业成功")
}

func (ctrl *WorkController) GetWork() {
    var works []*models.Work
    works, err := models.GetWorks()
    if err != nil {

    }
    ctrl.Data["json"] = &works
    ctrl.ServeJSON()
}

//需要用户验证
func (ctrl *WorkController) DeleteWork() {
    title := strings.TrimSpace(ctrl.GetString("titles"))
    println("title:", title)
    num, err := models.DeleteWork(title)
    if err != nil {
        ctrl.Ctx.WriteString("删除失败,系统发生错误.")
        return
    }
    if num == 0 {
        ctrl.Ctx.WriteString("删除失败,找不到对应的作业")
        return
    }
    ctrl.Ctx.WriteString("删除成功!")
}