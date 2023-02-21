package mysql

import (
	"fmt"
	"go-playground/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dao struct {
	DB *gorm.DB
}

func GetDB() *Dao {
	db, err := gorm.Open(mysql.Open(GetDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &Dao{DB: db}
}

func GetDSN() string {
	con, err := config.NewConfig("")
	if err != nil {
		panic(err)
	}

	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf(dsn,
		con.MySQL.User,
		con.MySQL.Password,
		con.MySQL.Host,
		con.MySQL.Port,
		con.MySQL.Db,
	)
}
