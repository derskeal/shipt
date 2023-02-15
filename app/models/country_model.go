package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Code       string
	Name       string
	IsEUMember bool
}
