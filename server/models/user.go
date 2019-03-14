package models

import (
	"time"
)

type User struct {
	ID            int
	Date_of_birth time.Time
	First_name    string
	Lirst_name    string
}

func (User) TableName() string {
	return "app_user"
}

func (u User) CheckIsBirthDay() bool {
	return u.Date_of_birth.Day() == time.Now().Day() && u.Date_of_birth.Month() == time.Now().Month()
}
