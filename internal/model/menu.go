package model

type MenuType string

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}
