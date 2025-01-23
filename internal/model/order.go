package model

type OrderStatus string

type Order struct {
	Status         OrderStatus
	ProductOreders []ProductOreder
}

type ProductOrederStatus string

type ProductOreder struct {
	OrderCode  string
	Quantity   int
	TotalPrice int64
	Status     ProductOrederStatus
}
