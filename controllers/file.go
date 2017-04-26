package controllers

import (
	//"archive/zip"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
	"uploadfilesys/models"

	"strconv"

	"regexp"

	"github.com/astaxie/beego"
)

var stdID []string

func init() {
	stdID = make([]string, 59)
	stand := "2014241332"
	for i := 1; i <= 57; i++ {
		var suffix string
		suffix = strconv.Itoa(i)
		if i < 10 {
			suffix = "0" + suffix
		}
		stdID[i] = stand + suffix
	}
	stdID[58] = "201324133206"
}

type FileController struct {
	beego.Controller
}

type ControllerZipFile interface {
	// 压缩文件夹为zip文件
	// 存放于/statuc/upload目录下
	CompressFolder(full_folder string) error
}

func (ctrl *FileController) ComperssFolder(full_folder string, dest string) error {

	re, _ := regexp.Compile("[0-9+]")

	stdM := make(map[string]int64)

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
			numS := re.FindAll([]byte(fileinfo.Name()), -1)
			var id string
			for _, v := range numS {
				id += string(v)
				stdM[id] = fileinfo.Size()
			}
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
		return err
	}

	towrite := `已交份数: %d

没提交实验报告的同学有: %s
所有文件中:
最大文件为学号%s: %fkb,
最小文件为学号%s: %fkb

`

	var size int
	var miss string
	var max float32
	var maxHost string
	var min float32 = 99999999999999999
	var minHost string

	for _, id := range stdID {
		if _, ok := stdM[id]; !ok {
			miss += " " + id + "\n"
		} else {
			size++
			if stdM[id] > int64(max) {
				maxHost = id
				max = float32(stdM[id]) / 1024
			}
			if float32(stdM[id]) < min {
				minHost = id
				min = float32(stdM[id]) / 1024
			}
		}
	}
	result := fmt.Sprintf(towrite, size, miss, maxHost, max, minHost, min)
	// 生成分析报告
	readmeName := time.Now().Format("2006-01-02-15:04-") + "readme.txt"
	readme, _ := os.Create(readmeName)
	readme.WriteString(result)
	readme.Close()
	readme, _ = os.Open(readmeName)
	fileinfo, _ := readme.Stat()
	header, err := zip.FileInfoHeader(fileinfo)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("Header:", header)
	zipwriter, err := fpz.CreateHeader(header)
	if err != nil {
		fmt.Println("err:", err)
	}
	_, err = io.Copy(zipwriter, readme)
	if err != nil {
		fmt.Println("err:", err)
	}
	return nil
}

// return:
func (ctrl *FileController) PostFile() {
	fp, header, err := ctrl.GetFile("file")
	defer fp.Close()
	if err != nil {
		ctrl.Ctx.WriteString("获取文件描述符失败")
		return
	} else {
		title := ctrl.GetString("title")
		id := ctrl.GetString("id")
		name := ctrl.GetString("name")

		filenameEle := strings.Split(header.Filename, ".")
		ext := filenameEle[len(filenameEle)-1]
		if ext != "rar" && ext != "zip" {
			ctrl.Ctx.WriteString("不支持的文件格式")
			return
		}
		header.Filename = id + name + "." + ext
		target_folder := "static/upload/" + title + "/"
		err := os.Mkdir(target_folder, 0777)
		if err == nil || os.IsExist(err) {
			ctrl.SaveToFile("file", target_folder+header.Filename)
		}

		ctrl.Ctx.WriteString("上传成功,学委已经收到你的作业")
		return
	}
}

// 将特定目录下的文件打包成zip压缩文件
func (ctrl *FileController) DownloadAFile() {
	//token := ctrl.GetString("token")
	// Get请求的口令验证下个版本实现,现在赶着上线
	flag_download := false
	title := ctrl.GetString("title")
	println(title)
	works, err := models.GetWorks()
	if err != nil {
		ctrl.Ctx.WriteString("系统发生错误")
	}
	for _, work := range works {
		if title == work.Title {
			flag_download = true
		}
	}
	if flag_download {
		path_now, err := os.Getwd()
		if err != nil {
			ctrl.Ctx.WriteString("文件获取失败")
			return
		}
		dest_file := title + ".zip"
		ctrl.ComperssFolder(path_now+"/static/upload/"+title, path_now+"/static/upload/"+dest_file)
		//ctrl.Ctx.WriteString("ok")
		ctrl.Ctx.Output.Download(path_now+"/static/upload/"+dest_file, dest_file)
	} else {
		ctrl.Ctx.WriteString("不存在该文件")
		return
	}

}
