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
//	@Success		200		{object}	NoteWithHistoriesResponse
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

		var histories []model.NoteHistory
		if err := db.Where("note_id = ?", note.ID).Find(&histories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to retrieve note histories"})
			return
		}

		// Convert histories to []SingleNote
		var singleNoteHistories []SingleNote
		for _, history := range histories {
			singleNoteHistories = append(singleNoteHistories, SingleNote{
				Title:     history.Title,
				Body:      history.Body,
				CreatedAt: history.CreatedAt.Unix(),
			})
		}

		c.JSON(http.StatusOK, NoteWithHistoriesResponse{
			NoteResponse: NoteResponse{
				SingleNote: SingleNote{
					Title:     note.Title,
					Body:      note.Body,
					CreatedAt: note.CreatedAt.Unix(),
				},
				ID:             note.ID,
				Slug:           note.Slug,
				IsEncrypted:    note.IsEncrypted,
				UpdatedAt:      note.UpdatedAt.Unix(),
				EncryptionSalt: note.EncryptionSalt,
				EncryptionTag:  note.EncryptionTag,
			},
			Histories: singleNoteHistories,
		})
	}
}
