package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Pembuatan struct menu item
type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

// Pembuatan list menu item
func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	}) // respon status code 201 Created Forma tJson
}

func main() {
	e := echo.New()
	// localhost:1404/menu/food
	e.GET("/menu/food", getFoodMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
