package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/letanthang/nc_student/config"
	"github.com/letanthang/nc_student/route"
	mw "github.com/letanthang/nc_student/middleware"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(mw.SimpleLogger())
	route.All(e)

	log.Println(e.Start(":9090"))
}
