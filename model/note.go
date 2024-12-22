package model

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID            uint           `gorm:"primarykey"`
	Slug          string         `gorm:"size:255;uniqueIndex"`
	Title         string         `gorm:"size:255;not null"`
	Body          string         `gorm:"type:text"`
	IsEncrypted   bool          `gorm:"default:false"`
	EncryptionTag string         `gorm:"size:255"` // For client-side encryption verification
	CreatedAt     time.Time      `gorm:"not null"`
	UpdatedAt     time.Time      `gorm:"not null"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type NoteHistory struct {
	ID        uint      `gorm:"primarykey"`
	NoteID    uint      `gorm:"not null"`
	Title     string    `gorm:"size:255;not null"`
	Body      string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"not null"` // When this version was created
}

// BeforeUpdate hook to create history record
func (n *Note) BeforeUpdate(tx *gorm.DB) error {
	// Create history record
	history := NoteHistory{
		NoteID:    n.ID,
		Title:     n.Title,
		Body:      n.Body,
		CreatedAt: time.Now(),
	}
	return tx.Create(&history).Error
}
