package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Request struct {
	Name   *string `json:"name,omitempty"`
	Family string  `json:"family,omitempty"`
	Age    int     `json:"age,omitempty"`
}

func helloHandler(c echo.Context) error {
	qp := c.QueryParam("alvani")
	fmt.Println(qp)

	auth := c.Request().Header.Get("Authorization")
	fmt.Println(auth)

	return c.JSON(http.StatusOK, "Hello")
}

func main() {
	app := echo.New()

	app.GET("/hello", helloHandler)
	app.POST("/hello", func(c echo.Context) error {
		var req Request

		if err := c.Bind(&req); err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%+v\n", req)

		return nil
	})

	if err := app.Start("0.0.0.0:1373"); err != nil {
		log.Fatal("server failed to start")
	}
}
