package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/1995parham-teaching/students-fall-2023/internal/domain/model"
	"github.com/1995parham-teaching/students-fall-2023/internal/domain/repository/studentrepo"
	"github.com/1995parham-teaching/students-fall-2023/internal/infra/repository/studentmem"
	"github.com/labstack/echo/v4"
)

type Request struct {
	Name   string `json:"name,omitempty"`
	Family string `json:"family,omitempty"`
}

func main() {
	app := echo.New()

	var repo studentrepo.Repository = studentmem.New()

	app.GET("/student", func(c echo.Context) error {
		students := repo.Get(c.Request().Context(), studentrepo.GetCommand{
			ID:           nil,
			FirstName:    nil,
			LastName:     nil,
			EntranceYear: nil,
		})

		return c.JSON(http.StatusOK, students)
	})

	app.PUT("/student", func(c echo.Context) error {
		var req Request

		if err := c.Bind(&req); err != nil {
			fmt.Println(err)
		}
		// we have the filled request

		id := rand.Uint64() % 1_000_000
		repo.Add(c.Request().Context(), model.Student{
			ID:           id,
			FirstName:    req.Name,
			LastName:     req.Family,
			EntranceYear: 0,
			Courses:      []model.Course{},
		})

		return c.JSON(http.StatusCreated, id)
	})

	if err := app.Start("0.0.0.0:1373"); err != nil {
		log.Fatal("server failed to start")
	}
}
