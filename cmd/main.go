package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rocksus/go-restaurant-app/internal/database"
	"github.com/rocksus/go-restaurant-app/internal/delivery/rest"
	mRepo "github.com/rocksus/go-restaurant-app/internal/repository/menu"
	rUsecase "github.com/rocksus/go-restaurant-app/internal/usecase/resto"
)

const (
	dbAdress = "host=localhost port=5432 user=postgres password=Arya2003ok dbname=go_resto_app sslmode=disable"
)

func main() {
	e := echo.New()

	db := database.GetDB(dbAdress)

	menuRepo := mRepo.GetRepository(db)

	restoUseCase := rUsecase.GetUseCase(menuRepo)

	h := rest.NewHandler(restoUseCase)

	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":14045"))
}
