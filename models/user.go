package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int   `json:"id" gorm:"primatry_key"`
	NAME     string `json:"name"`
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
}