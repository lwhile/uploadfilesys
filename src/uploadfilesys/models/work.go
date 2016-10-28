package models

import (
)
import "github.com/astaxie/beego/orm"

type CRUD interface {

}

type Work struct {
    ID int64
    title string
}

func init() {
    orm.RegisterModel(new(Work))
}

func Getworks() ([]*Work, error) {
    ORM := orm.NewOrm()
    var results []orm.Params
    var works []*Work
    _, err := ORM.QueryTable("works").Values(&results)
    if err != nil {
        return nil, err
    }
    for _, v := range results {
        work := &Work{}
        work.ID = v["id"].(int64)
        work.title = v["title"].(string)
        works = append(works, work)
    }
    return works, nil
}

func InsertWork(title string) error {
    ORM := orm.NewOrm()
    work := Work{title:title}
    _, err := ORM.Insert(&work)
    return err
}
