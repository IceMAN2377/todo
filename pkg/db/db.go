package db

import (
	"github.com/IceMAN2377/todo.git/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(cfg *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(cfg.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{db}
}
