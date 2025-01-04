package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justin-jiajia/note/model"
	"gorm.io/gorm"
)

// EditNote handles requests to update an existing note
//	@Summary		Update a note
//	@Description	Update a note's title and body by its slug
//	@Tags			notes
//	@Accept			json
//	@Produce		json
//	@Param			slug	path		string			true	"Note slug"
//	@Param			note	body		EditNoteRequest	true	"Note update data"
//	@Success		200		{object}	NoteResponse
//	@Failure		400		{object}	ErrorResponse
//	@Failure		403		{object}	ErrorResponse
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/notes/{slug} [put]
func EditNote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var note model.Note
		if err := db.Where("slug = ?", c.Param("slug")).First(&note).Error; err != nil {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: "Note not found"})
			return
		}

		// Verify encryption tag for encrypted notes
		if note.IsEncrypted {
			providedTag := c.GetHeader("X-Encryption-Tag")
			if providedTag != note.EncryptionTag {
				c.JSON(http.StatusForbidden, ErrorResponse{Error: "Invalid encryption tag"})
				return
			}
		}

		var input struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
			return
		}

		note.Title = input.Title
		note.Body = input.Body

		if err := db.Save(&note).Error; err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to update note"})
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
