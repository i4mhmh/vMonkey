package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"m0nk3y/pocScanner/common"
)

func main() {
	db := common.InitDB()
	sqlDB, _ := db.DB()
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Println("code: 0128")
		}
	}(sqlDB)
	r := gin.Default()
	r = CollectionRoute(r)
	panic(r.Run())
}
