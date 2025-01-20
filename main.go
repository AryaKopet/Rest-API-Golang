package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Pembuatan struct menu item
type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

func seedDb() {
	foodsMenu := []MenuItem{
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
		{
			Name:      "Lalapan",
			OrderCode: "lalapan",
			Price:     38000,
		},
	}
	drinksMenu := []MenuItem{
		{
			Name:      "Josu",
			OrderCode: "josu",
			Price:     10000,
		},
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     5000,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "es_jeruk",
			Price:     7000,
		},
	}

	dbAdress := "host=localhost port=5432 user=postgres password=Arya2003ok dbname=go_resto_app sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbAdress))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&MenuItem{})

}

// List menu makanan
func getFoodMenu(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		// "data": foodMenu,
	}) // respon status code 201 Created Forma tJson
}

// List menu minuman
func getDrinkMenu(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		// "data": foodMenu,
	}) // respon status code 201 Created Forma tJson
}

func main() {
	seedDb() //pemanggilan database
	e := echo.New()
	// localhost:1404/menu/food
	e.GET("/menu/foods", getFoodMenu)
	e.GET("/menu/drinks", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
