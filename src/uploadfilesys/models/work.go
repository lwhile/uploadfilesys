package models

import (
    "github.com/go-sql-driver/mysql"
)

type CRUD interface {

}

type Work struct {
    ID int64
    title string
}



