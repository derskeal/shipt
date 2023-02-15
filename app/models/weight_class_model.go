package models

import "gorm.io/gorm"

type WeightClass struct {
	gorm.Model
	Name  string
	Price int64
}
