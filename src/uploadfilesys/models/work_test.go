package models

import (
    "testing"
)

func TestInsertAndGetWorkAndDeleteWork(t *testing.T) {
    titles := []string{"Java第一次实验", "Java第二次实验", "Java第三次实验", "数据库实验"}
    for _, title := range titles {
        err := InsertWork(title)
        if err != nil {
            t.Error(err)
        }
    }

    works, err := GetWorks()
    if err != nil {
        t.Error(err)
    }
    if len(works) != 4 {
        t.Error(len(works), "插入失败")
    }

    for _, title := range titles {
        nums, err := DeleteWork(title)
        if nums != 1 && err != nil {
            t.Error(err)
        }
    }

}

