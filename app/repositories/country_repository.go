package repositories

import (
	"shipt/app/config"
	"shipt/app/models"
)

func FindCountryByCode(code string) (models.Country, error) {
	var country models.Country
	result := config.DB.Where("code = ?", code).First(&country)
	return country, result.Error
}
