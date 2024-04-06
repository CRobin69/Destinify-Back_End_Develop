package postgres

import (
	"github.com/CRobin69/Destinify-Back_End_Develop/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db = db.Debug()
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.City{},
		&entity.Place{},
		&entity.Culinary{},
		&entity.Ticket{},
		&entity.Guide{},
		&entity.Order{},
		&entity.Transaction{},
		&entity.Comment{},
		); err != nil {
		return err
	}
	return nil
}
