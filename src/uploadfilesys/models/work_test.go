package models

import (
    "testing"
)

func TestInsertWork(t *testing.T) {
    titles := []string{"Java第一次实验", "Java第二次实验", "Java第三次实验","数据库实验"}
    for _, title := range titles {
        err := InsertWork(title)
        if err != nil {
            t.Error("No Pass.")
        }
    }
}
