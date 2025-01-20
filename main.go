package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbAdress = "host=localhost port=5432 user=postgres password=Arya2003ok dbname=go_resto_app sslmode=disable"
)

func main() {
	seedDb() //pemanggilan database
	e := echo.New()
	// localhost:1404/menu/food
	e.GET("/menu", getMenu)
	e.Logger.Fatal(e.Start(":14045"))
}

type MenuType string

const (
	MenuTypeFood   = "Food"
	MenuTypeDrinks = "Drink"
)

// Pembuatan struct menu item
type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}

func seedDb() {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Lalapan",
			OrderCode: "lalapan",
			Price:     38000,
			Type:      MenuTypeFood,
		},
	}
	drinksMenu := []MenuItem{
		{
			Name:      "Josu",
			OrderCode: "josu",
			Price:     10000,
			Type:      MenuTypeDrinks,
		},
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     5000,
			Type:      MenuTypeDrinks,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "es_jeruk",
			Price:     7000,
			Type:      MenuTypeDrinks,
		},
	}

	fmt.Println(foodMenu, drinksMenu)
	db, err := gorm.Open(postgres.Open(dbAdress))
	if err != nil {
		panic(err)
	}

	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound { //pengecekan database pertama kali
		db.AutoMigrate(&MenuItem{}) // membuat migrasi otomatis kedalam database
		db.Create(foodMenu)         // menambahkan data makanan ke dalam database
		db.Create(drinksMenu)       // //menambahkan data minuman ke dalam database
	}

}

func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	db, err := gorm.Open(postgres.Open(dbAdress))
	if err != nil {
		panic(err)
	}

	var menuData []MenuItem
	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})

}
