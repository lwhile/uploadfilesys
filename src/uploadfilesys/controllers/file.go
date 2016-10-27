package controllers

import "github.com/astaxie/beego"

type FileController struct {
    beego.Controller
}

func (ctrl *FileController) PostFile() {
    fp, header, err := ctrl.GetFile("file")
    defer fp.Close()
    if err != nil {
        // todo: 处理错误
    } else {
        ctrl.SaveToFile("file", "static/upload/" + header.Filename)
        // todo: 返回状态
    }
}
