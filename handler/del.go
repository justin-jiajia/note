package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justin-jiajia/note/model"
	"gorm.io/gorm"
)

// DeleteNote handles requests to delete an existing note
//	@Summary		Delete a note
//	@Description	Delete a note by its slug
//	@Tags			notes
//	@Accept			json
//	@Produce		json
//	@Param			slug	path		string	true	"Note slug"
//	@Success		200		{object}	SuccessResponse
//	@Failure		403		{object}	ErrorResponse
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/notes/{slug} [delete]
func DeleteNote(db *gorm.DB) gin.HandlerFunc {
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

		if err := db.Delete(&note).Error; err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to delete note"})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{Message: "Note deleted successfully"})
	}
}
