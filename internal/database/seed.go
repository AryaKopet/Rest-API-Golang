package database

import (
	"github.com/rocksus/go-restaurant-app/internal/model"
	"github.com/rocksus/go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)

func seedDb(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Lalapan",
			OrderCode: "lalapan",
			Price:     38000,
			Type:      constant.MenuTypeFood,
		},
	}
	drinksMenu := []model.MenuItem{
		{
			Name:      "Josu",
			OrderCode: "josu",
			Price:     10000,
			Type:      constant.MenuTypeDrinks,
		},
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     5000,
			Type:      constant.MenuTypeDrinks,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "es_jeruk",
			Price:     7000,
			Type:      constant.MenuTypeDrinks,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound { //pengecekan database pertama kali
		db.Create(&foodMenu)   // menambahkan data makanan ke dalam database
		db.Create(&drinksMenu) // //menambahkan data minuman ke dalam database
	}

}
