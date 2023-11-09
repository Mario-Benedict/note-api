package conf

import (
	"github.com/Mario-Benedict/note-api/model"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.Note{})
}
