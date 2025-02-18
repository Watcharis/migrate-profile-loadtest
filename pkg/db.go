package pkg

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	// dsn := fmt.Sprintf(
	// 	"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
	// 	"lotto",
	// 	"P@ssw0rd",
	// 	"(34.126.127.110:3306)", // public sit
	// 	"ktb_glo",
	// )

	// dsn := fmt.Sprintf(
	// 	"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
	// 	"user",
	// 	"longpass",
	// 	"(localhost:3306)",
	// 	"lotto",
	// )

	// dsn := fmt.Sprintf(
	// 	"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
	// 	"appusr",
	// 	"L0@dT3$t",
	// 	"(34.142.175.151:3306)", // public loadtest
	// 	"ktb_glo",
	// )

	// dsn := fmt.Sprintf(
	// 	"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
	// 	"appusr",
	// 	"L0@dT3$t",
	// 	"(10.98.33.43:3306)", // private loadtest
	// 	"ktb_glo",
	// )

	// dsn := fmt.Sprintf(
	// 	"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
	// 	"lotto",
	// 	"P@ssw0rd",
	// 	"(ktb-lotto-mysql-dev.lotto.nonprod.gcp.ktbcloud:3306)",
	// 	"ktb_glo",
	// )

	dsn := fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
		"appusr",
		"L0@dT3$t",
		"(35.240.199.154:3306)", // public loadtest-v8 replica
		"ktb_glo",
	)

	// dsn := fmt.Sprintf(
	// 	"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
	// 	"appusr",
	// 	"L0@dT3$t",
	// 	"(34.143.214.226:3306)", // public loadtest-v8
	// 	"ktb_glo",
	// )

	log.Println("Initialing database with dsn")

	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	// config connection pools
	sqlDB, err := db.DB()
	if err != nil {
		log.Panic(err)
	}
	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetConnMaxLifetime(time.Duration(5 * time.Minute))

	// returns database statistics
	log.Printf("Database is running!")

	return db
}
