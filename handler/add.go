package handler

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justin-jiajia/note/model"
	"gorm.io/gorm"
)

// generateSlug creates a random 4-character slug using a-z
//
//	@Summary		Generate unique slug
//	@Description	Creates a random 4-character string for use as a note slug
//	@Return			string Random 4-character slug
func generateSlug() string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 4)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// CreateNote handles requests to create a new note
//
//	@Summary		Create a note
//	@Description	Create a new note with title, body, and optional encryption
//	@Tags			notes
//	@Accept			json
//	@Produce		json
//	@Param			note	body		CreateNoteRequest	true	"Note creation data"
//	@Success		201		{object}	NoteResponse
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/notes [post]
func CreateNote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CreateNoteRequest
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
			return
		}

		// Generate unique slug
		var exists bool
		slug := generateSlug()
		for attempts := 0; attempts < 10; attempts++ {
			err := db.Model(&model.Note{}).Select("count(*) > 0").Where("slug = ?", slug).Find(&exists).Error
			if err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to check slug"})
				return
			}
			if !exists {
				break
			}
			slug = generateSlug()
		}
		if exists {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to generate unique slug"})
			return
		}

		note := model.Note{
			Title:                     input.Title,
			Body:                      input.Body,
			Slug:                      slug,
			IsEncrypted:               input.IsEncrypted,
			EncryptionTag:             input.EncryptionTag,
			EncryptionVerificationTag: input.EncryptionVerificationTag,
			EncryptionSalt:            input.EncryptionSalt,
		}

		if err := db.Create(&note).Error; err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create note"})
			return
		}

		c.JSON(http.StatusCreated, NoteResponse{
			ID:             note.ID,
			Slug:           note.Slug,
			Title:          note.Title,
			Body:           note.Body,
			IsEncrypted:    note.IsEncrypted,
			CreatedAt:      note.CreatedAt.Unix(),
			UpdatedAt:      note.UpdatedAt.Unix(),
			EncryptionSalt: note.EncryptionSalt,
			EncryptionTag:  note.EncryptionTag,
		})
	}
}
