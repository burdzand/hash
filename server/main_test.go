package main

import (
	"testing"
	"time"

	"github.com/burdzand/hash/server/connections/database"
	"github.com/burdzand/hash/server/models"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	DB = database.GetConn()
	DB.LogMode(true)
	migration()
}

func TestCreateConnection(t *testing.T) {
	DB = database.GetConn()
}

func migration() {
	DB.DropTableIfExists(
		&models.User{},
		&models.Product{},
	)
	DB.AutoMigrate(
		&models.User{},
		&models.Product{},
	)
}

func TestUserCreate(t *testing.T) {
	user := models.User{First_name: "Hash", Lirst_name: "Test", Date_of_birth: time.Now()}
	res := DB.Save(&user)
	if res.Error != nil {
		t.Errorf("User not created.... %t ", res.Error)
	}
}

func TestProductCreate(t *testing.T) {
	product := models.Product{Price_in_cents: 2000}
	res := DB.Save(&product)
	if res.Error != nil {
		t.Errorf("Product not created.... %t ", res.Error)
	}
}

func TestUserBirth(t *testing.T) {
	user := models.User{}
	if DB.First(&user, 1).RecordNotFound() {
		t.Errorf("User not foud")
	}
	if !user.CheckIsBirthDay() {
		t.Errorf("User not a Birthday")
	}
}

func TestCheckDiscountBirthday(t *testing.T) {
	user := models.User{}
	if DB.First(&user, 1).RecordNotFound() {
		t.Errorf("User not Found")
	}
	product := models.Product{}
	if DB.First(&product, 1).RecordNotFound() {
		t.Errorf("Product not Found")
	}
	discount := getDiscount(&user, &product)
	if discount.ValueInCents != 1900 {
		t.Errorf("Discount Birthday do not Match")
	}
}

func TestCheckDiscountNoBirthday(t *testing.T) {
	user := models.User{First_name: "Hash 2", Lirst_name: "Test 2", Date_of_birth: time.Now().AddDate(0, 0, -1)}
	DB.Save(&user)
	product := models.Product{}
	DB.First(&product, 1)

	discount := getDiscount(&user, &product)
	if discount.ValueInCents != product.Price_in_cents {
		t.Errorf("Discount Birthday do not Match")
	}
}

func TestCheckBlackFriday(t *testing.T) {
	blackday, _ := time.Parse("2006-01-02", "2019-11-25")
	if !checkBlackFriday(blackday) {
		t.Errorf("The Date is not a blackfriday Day")
	}
}

func TestCheckNotBlackFriday(t *testing.T) {
	blackday, _ := time.Parse("2006-01-02", "2019-11-24")
	if checkBlackFriday(blackday) {
		t.Errorf("The Date is not a blackfriday Day")
	}
}
