package main

import (
	"github.com/labstack/echo/v4"
)

func getFoodMenu(c echo.Context) error {
	return nil
}

func main() {
	e := echo.New()
	// localhost:1404/menu/food
	e.GET("/menu/food", getFoodMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
