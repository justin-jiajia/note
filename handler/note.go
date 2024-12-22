package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justin-jiajia/note/model"
	"gorm.io/gorm"
)

type NoteResponse struct {
	ID            uint   `json:"id"`
	Slug          string `json:"slug"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	IsEncrypted   bool   `json:"isEncrypted"`
	EncryptionTag string `json:"encryptionTag,omitempty"`
	CreatedAt     int64 `json:"createdAt"`
	UpdatedAt     int64 `json:"updatedAt"`
}

func ViewNote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")

		var note model.Note
		result := db.Where("slug = ?", slug).First(&note)
		
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
			return
		}
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve note"})
			return
		}

		c.JSON(http.StatusOK, NoteResponse{
			ID:          note.ID,
			Slug:        note.Slug,
			Title:       note.Title,
			Body:        note.Body,
			IsEncrypted: note.IsEncrypted,
			CreatedAt:   note.CreatedAt.Unix(),
			UpdatedAt:   note.UpdatedAt.Unix(),
		})
	}
}

func EditNote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var note model.Note
		if err := db.Where("slug = ?", c.Param("slug")).First(&note).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
			return
		}

		// Verify encryption tag for encrypted notes
		if note.IsEncrypted {
			providedTag := c.GetHeader("X-Encryption-Tag")
			if providedTag != note.EncryptionTag {
				c.JSON(http.StatusForbidden, gin.H{"error": "Invalid encryption tag"})
				return
			}
		}

		var input struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		note.Title = input.Title
		note.Body = input.Body

		if err := db.Save(&note).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
			return
		}

		c.JSON(http.StatusOK, NoteResponse{
			ID:          note.ID,
			Slug:        note.Slug,
			Title:       note.Title,
			Body:        note.Body,
			IsEncrypted: note.IsEncrypted,
			CreatedAt:   note.CreatedAt.Unix(),
			UpdatedAt:   note.UpdatedAt.Unix(),
		})
	}
}

func DeleteNote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var note model.Note
		if err := db.Where("slug = ?", c.Param("slug")).First(&note).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
			return
		}

		// Verify encryption tag for encrypted notes
		if note.IsEncrypted {
			providedTag := c.GetHeader("X-Encryption-Tag")
			if providedTag != note.EncryptionTag {
				c.JSON(http.StatusForbidden, gin.H{"error": "Invalid encryption tag"})
				return
			}
		}

		if err := db.Delete(&note).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
	}
}
