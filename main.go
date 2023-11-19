package main

import (
	"log"

	"github.com/1995parham-teaching/students-fall-2023/internal/domain/repository/studentrepo"
	"github.com/1995parham-teaching/students-fall-2023/internal/infra/http/handler"
	"github.com/1995parham-teaching/students-fall-2023/internal/infra/repository/studentmem"
	"github.com/1995parham-teaching/students-fall-2023/internal/infra/repository/studentsql"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("students.db"), new(gorm.Config))
	if err != nil {
		log.Fatalf("failed to connect database %v", err)
	}

	if err := db.AutoMigrate(new(studentsql.StudentDTO)); err != nil {
		log.Fatalf("failed to run migrations %v", err)
	}

	app := echo.New()

	var repo studentrepo.Repository = studentmem.New()

	h := handler.NewStudent(repo)
	h.Register(app.Group("students/"))

	if err := app.Start("0.0.0.0:1373"); err != nil {
		log.Fatalf("server failed to start %v", err)
	}
}
