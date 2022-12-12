package database

import "app/pkg/model"

func CreateOrder(order model.Order) {
	DB.Create(order)
}
