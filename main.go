package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func panicOnErr(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

type Logger interface {
	Log(v any)
}

type defaultErrLogger struct{}

func NewDefaultErrLogger() Logger {
	return &defaultErrLogger{}
}

func (d *defaultErrLogger) Log(v any) {
	log.Println(v)
}

func logOnError(err error, logger Logger) {
	if err != nil {
		logger.Log(err)
	}
}

func main() {
	dsn := "host=localhost user=mdr password=secret dbname=gws port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("connected to the database")
	panicOnErr(db.AutoMigrate(&Product{}), "failed to migrate")
}
