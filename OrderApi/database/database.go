package database

import (
	"errors"
	"fmt"
	"log"
	"time"

	"assignment2.id/orderapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host   = "localhost"
	user   = "postgres"
	dbPort = "5432"
	dbName = "assignment2db"
	db     *gorm.DB
	err    error
)

func StartDB() {
	var password string
	fmt.Println("Enter db password (not hidden, be careful of shoulder surfing)")
	fmt.Scanln(&password)
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}

func CreateOrder(order *models.Order) error {
	if db == nil {
		return errors.New("DB hasn't started yet.")
	}
	var zero time.Time
	if order.OrderedAt == zero {
		order.OrderedAt = time.Now()
	}
	err := db.Create(order).Error
	if err != nil {
		return err
	}
	log.Println("New Order Data: ", order)
	return nil
}

func GetOrderById(id uint) (models.Order, error) {
	order := models.Order{}
	if db == nil {
		return order, errors.New("DB hasn't started yet.")
	}
	err := db.Model(&models.Order{}).Preload("Items").Take(&order, id).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func GetOrderByIds(ids ...uint) ([]models.Order, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	if len(ids) == 1 {
		order, err := GetOrderById(ids[0])
		return []models.Order{order}, err
	}
	if db == nil {
		return nil, errors.New("DB hasn't started yet.")
	}
	orders := make([]models.Order, 2)
	err := db.Model(&models.Order{}).Preload("Items").Find(&orders, ids).Error
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, errors.New("No ID match.")
	}
	return orders, nil
}

func UpdateOrderById(id uint, argOrder *models.Order) error {
	dbOrder, err := GetOrderById(id)
	if err != nil {
		return err
	}
	if argOrder.CustomerName != "" {
		dbOrder.CustomerName = argOrder.CustomerName
	}
	var temp time.Time
	if argOrder.OrderedAt != temp {
		dbOrder.OrderedAt = argOrder.OrderedAt
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&dbOrder).Error; err != nil {
			return err
		}
		if argOrder.Items != nil {
			if err := tx.Model(&dbOrder).Association("Items").Replace(argOrder.Items); err != nil {
				return err
			}
		}
		return nil
	})
	if err == nil {
		log.Printf("Updated order: %+v\n", dbOrder)
	}
	return err
}

func DeleteOrderById(id uint) error {
	order, err := GetOrderById(id)
	if err != nil {
		return err
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&order).Association("Items").Clear(); err != nil {
			return err
		}
		if err := tx.Delete(&order, id).Error; err != nil {
			return err
		}
		return nil
	})
	if err == nil {
		log.Println("Order with id", id, "has been successfully deleted")
	}
	return err
}
