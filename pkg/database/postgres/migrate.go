package postgres

import (
	"INTERN_BCC/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.City{},
		&entity.Place{},
		&entity.Ticket{},
		&entity.Guide{},
		&entity.Payment{},
	); err != nil {
		return err
	}
	return nil
}
