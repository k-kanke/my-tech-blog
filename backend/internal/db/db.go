package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=db user=bloguser password=blogpass dbname=blogdb port=5432 sslmode=disable"
	var db *gorm.DB
	var err error

	// 10秒以内にリトライする
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("DBに接続しました")
			return db
		}

		fmt.Printf("DB接続リトライ中... (%d/10)\n", i+1)
		time.Sleep(1 * time.Second)
	}

	panic("DBとの接続に失敗しました")
}
