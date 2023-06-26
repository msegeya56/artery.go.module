package migrations

import (
	"github.com/msegeya56/artery.go.module/pkg/domain/models"
	"gorm.io/gorm"
)

func RunMigrations(db gorm.DB) {

	db.AutoMigrate(&models.Art{})

	db.AutoMigrate(&models.Form{})

	db.AutoMigrate(&models.Task{})

	db.AutoMigrate(&models.Todo{})

}
