package main

import (
	"fmt"
	"github.com/iamyusuf/gws/api"
	"github.com/iamyusuf/gws/services"
	"github.com/iamyusuf/gws/storage"
	"github.com/iamyusuf/gws/types/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func panicOnErr(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	dsn := "host=localhost user=mdr password=secret dbname=gws port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	db, err := storage.NewDB(dsn)

	panicOnErr(err, "could not connect to db!")
	panicOnErr(db.AutoMigrate(&model.User{}), "failed to migrate")

	srv := api.NewServer(db)
	e := echo.New()

	logWriter, err := services.NewLogWriter()

	if err == nil {
		e.Logger.SetOutput(logWriter)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", srv.Home)
	e.POST("/api/user", srv.CreateUser)
	e.PUT("/api/user/:id", srv.UpdateUser)
	e.GET("/api/user/:id", srv.FindUserById)
	e.DELETE("/api/user/:id", srv.DeleteUserById)

	e.Logger.Fatal(e.Start(":1323"))
}
