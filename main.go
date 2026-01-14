package main

import (
	"flag"
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
	storage := flag.String("storage", "memory", "storage type: memory or sqlite")
	dbPath := flag.String("db", "students.db", "path to SQLite database file")

	flag.Parse()

	var repo studentrepo.Repository

	switch *storage {
	case "sqlite":
		db, err := gorm.Open(sqlite.Open(*dbPath), new(gorm.Config))
		if err != nil {
			log.Fatalf("failed to connect database %v", err)
		}

		if err := db.AutoMigrate(new(studentsql.StudentDTO)); err != nil {
			log.Fatalf("failed to run migrations %v", err)
		}

		repo = studentsql.New(db)

		log.Printf("using sqlite storage: %s", *dbPath)
	case "memory":
		repo = studentmem.New()

		log.Println("using in-memory storage")
	default:
		log.Fatalf("unknown storage type: %s (use 'memory' or 'sqlite')", *storage)
	}

	app := echo.New()

	h := handler.NewStudent(repo)
	h.Register(app.Group("students/"))

	if err := app.Start("0.0.0.0:1373"); err != nil {
		log.Fatalf("server failed to start %v", err)
	}
}
