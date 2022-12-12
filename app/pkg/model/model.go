package model

type User struct {
	Id    string `gorm:"primaryKey;column:id;type:varchar(255)"`
	Name  string `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Phone string `json:"phone" gorm:"column:phone;type:varchar(255);not null"`
}

type Book struct {
	Name   string `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Amount int    `json:"amount" gorm:"column:amount;type:int;not null"`
}
