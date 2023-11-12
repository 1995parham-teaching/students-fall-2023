package main

import (
	"log"

	"github.com/1995parham-teaching/students-fall-2023/internal/domain/repository/studentrepo"
	"github.com/1995parham-teaching/students-fall-2023/internal/infra/http/handler"
	"github.com/1995parham-teaching/students-fall-2023/internal/infra/repository/studentmem"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	var repo studentrepo.Repository = studentmem.New()

	h := handler.NewStudent(repo)
	h.Register(app.Group("students/"))

	if err := app.Start("0.0.0.0:1373"); err != nil {
		log.Fatal("server failed to start")
	}
}
