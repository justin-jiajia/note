package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justin-jiajia/note/model"
	"gorm.io/gorm"
)

// ViewNote handles requests to view a single note by its slug
//
//	@Summary		View a note
//	@Description	Get a note by its slug
//	@Tags			notes
//	@Accept			json
//	@Produce		json
//	@Param			slug	path		string	true	"Note slug"
//	@Success		200		{object}	NoteResponse
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/notes/{slug} [get]
func ViewNote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")

		var note model.Note
		result := db.Where("slug = ?", slug).First(&note)

		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: "Note not found"})
			return
		}
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to retrieve note"})
			return
		}

		c.JSON(http.StatusOK, NoteResponse{
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
