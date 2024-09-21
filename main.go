package main

import (
	"fmt"
	"github.com/iamyusuf/gws/api"
	"github.com/iamyusuf/gws/types/model"
	"github.com/labstack/echo/v4"
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
	panicOnErr(db.AutoMigrate(&Product{}, &model.User{}), "failed to migrate")

	srv := api.NewServer(db)

	e := echo.New()
	e.GET("/", srv.Home)
	e.POST("/api/user", srv.CreateUser)
	e.GET("/api/user/:id", srv.FindUserById)

	e.Logger.Fatal(e.Start(":1323"))
}
