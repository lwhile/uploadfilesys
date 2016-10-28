package controllers

import (
    //"archive/zip"
    "github.com/astaxie/beego"
    "os"
    "path/filepath"
    "archive/zip"
    "io"
)

type FileController struct {
    beego.Controller
}

type ControllerZipFile interface {
    // 压缩文件夹为zip文件
    // 存放于/statuc/upload目录下
    CompressFolder(full_folder string) error
}

func (ctrl *FileController) ComperssFolder(full_folder string, dest string) error {
    fp, err := os.Create(dest)
    defer fp.Close()
    if err != nil {
        return err
    }
    fpz := zip.NewWriter(fp)
    defer fpz.Close()

    // 获取待压缩文件列表
    err = filepath.Walk(full_folder, func(path string, fileinfo os.FileInfo, err error) error {
        if !fileinfo.IsDir() {
            file, err1 := os.Open(path)
            defer file.Close()
            if err1 != nil {
                return err1
            }

            header, _ := zip.FileInfoHeader(fileinfo)
            zipwriter, err1 := fpz.CreateHeader(header)
            if err1 != nil {
                return err1
            }

            _, err1 = io.Copy(zipwriter, file)
            if err1 != nil {
                return err1
            }

        }
        return nil
    })
    if err != nil {
        println(err)
        return err
    }
    return nil
}

func (ctrl *FileController) PostFile() {
    fp, header, err := ctrl.GetFile("file")
    println(fp)
    defer fp.Close()
    if err != nil {
        // todo: 处理错误
    } else {
        title := ctrl.GetString("title")
        target_folder := "static/upload/" + title + "/"
        err := os.Mkdir(target_folder, 0777)
        if err == nil || os.IsExist(err) {
            ctrl.SaveToFile("file", target_folder + header.Filename)
        }
        ctrl.Ctx.WriteString("1")
    }
}

// 将特定目录下的文件打包成zip压缩文件
func (ctrl *FileController) DownloadAFile() {
    title := ctrl.GetString("title")
    path_now, err := os.Getwd()
    if err != nil {
        // todo: 处理错误
    }
    dest_file := title + ".zip"
    ctrl.ComperssFolder(path_now + "/static/upload/第三次", path_now + "/static/upload/" + dest_file)
}


