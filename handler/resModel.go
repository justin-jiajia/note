package handler

// NoteResponse represents the API response structure for note operations
// @Description Response model for note data
type NoteResponse struct {
	// Unique identifier for the note
	ID uint `json:"id"`
	// URL-friendly identifier for the note
	Slug string `json:"slug"`
	// Title of the note
	Title string `json:"title"`
	// Content of the note
	Body string `json:"body"`
	// Whether the note is encrypted
	IsEncrypted bool `json:"is_encrypted"`
	// Tag used for encryption verification (only included if note is encrypted)
	EncryptionTag string `json:"encryption_tag,omitempty"`
	// Unix timestamp of note creation
	CreatedAt int64 `json:"created_at"`
	// Unix timestamp of last note update
	UpdatedAt int64 `json:"updated_at"`
	// Salt used for client-side encryption (only included if note is encrypted)
	EncryptionSalt string `json:"encryption_salt,omitempty"`
}

// ErrorResponse represents the error response structure
// @Description Error response model
type ErrorResponse struct {
	// Error message
	Error string `json:"error"`
}

// SuccessResponse represents the success response structure
// @Description Success response model
type SuccessResponse struct {
	// Success message
	Message string `json:"message"`
}
