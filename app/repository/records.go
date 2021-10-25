package repository

import (
	"database/sql"
	"sync"
)

var onceDB sync.Once

type records struct {
	DB *sql.DB
}

func NewRecords() Records {
	return &records{}
}

func (r *records) Add(event string, data interface{}) error {
	//query := "insert into log(event, data) values (?, ?)"

	return nil
}
