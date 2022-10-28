package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type ItemBody struct {
	ItemCode    string `example:"SOMECODE"`
	Description string `example:"Some description."`
	Quantity    uint   `example:"1"`
}
type OrderBody struct {
	CustomerName string `example:"Test"`
	Items        []ItemBody
	OrderedAt    time.Time `example:"2019-11-09T21:21:46+00:00"`
}
type Item struct {
	ID          uint   `gorm:"primaryKey" example:"1"`
	ItemCode    string `gorm:"not null;type:varchar(8192)" example:"Contoh"`
	Description string `gorm:"type:varchar(8192)" example:"Some description."`
	Quantity    uint   `gorm:"not null" example:"1"`
	OrderID     uint   `example:"1"`
}
type Order struct {
	ID           uint   `gorm:"primaryKey" example:"1"`
	CustomerName string `gorm:"type:varchar(8192)" example:"Contoh"`
	Items        []Item
	OrderedAt    time.Time `gorm:"not null" example:"2019-11-09T21:21:46+00:00"`
}

var ErrItemCodeEmpty error = errors.New("ItemCode kosong.")
var ErrCustomerNameEmpty error = errors.New("CustomerName kosong.")

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ItemCode == "" {
		err = ErrItemCodeEmpty
	}
	return
}
func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.CustomerName == "" {
		err = ErrCustomerNameEmpty
	}
	return
}
