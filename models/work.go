package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type CRUD interface {
}

type Work struct {
	Id    int64 `orm:"pk;auto"`
	Title string
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:1234@/uploadfilesys?charset=utf8")
	orm.RegisterModel(new(Work))

}

func GetWorks() ([]*Work, error) {
	ORM := orm.NewOrm()
	var results []orm.Params
	var works []*Work
	_, err := ORM.QueryTable("work").Values(&results)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		work := &Work{}
		work.Id = v["Id"].(int64)
		work.Title = v["Title"].(string)
		works = append(works, work)
	}
	return works, nil
}

func InsertWork(title string) error {
	ORM := orm.NewOrm()
	work := &Work{Title: title}
	_, err := ORM.Insert(work)
	return err
}

func DeleteWork(Title string) (int64, error) {
	ORM := orm.NewOrm()
	nums, err := ORM.Delete(&Work{Title: Title})
	return nums, err
}
