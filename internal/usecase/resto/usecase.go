package resto

import "github.com/rocksus/go-restaurant-app/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
