package model

import (
	"time"

	"github.com/Mario-Benedict/note-api/db"
)

type Note struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title" gorm:"size:256;not null"`
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}

type NoteModel struct{}

func (NoteModel) All() ([]Note, error) {
	var notes []Note

	err := db.GetDB().Find(&notes).Error

	if err != nil {
		return notes, err
	}

	return notes, err
}

func (NoteModel) GetByID(id int) (Note, error) {
	var note Note

	err := db.GetDB().Where("id = ?", id).First(&note).Error

	return note, err
}

func (NoteModel) Create(title string, content string) error {
	note := Note{
		Title:   title,
		Content: content,
	}

	err := db.GetDB().Create(&note).Error

	return err
}

func (NoteModel) Update(id int, title string, content string) error {
	var note Note

	err := db.GetDB().Where("id = ?", id).First(&note).Error

	if err != nil {
		return err
	}

	note = Note{
		Title:   title,
		Content: content,
	}

	err = db.GetDB().Where("id = ?", id).Updates(&note).Error
	return err
}

func (NoteModel) Delete(id int) error {
	var note Note

	err := db.GetDB().Where("id = ?", id).First(&note).Error

	if err != nil {
		return err
	}

	err = db.GetDB().Where("id = ?", id).Delete(&Note{}).Error

	return err
}
