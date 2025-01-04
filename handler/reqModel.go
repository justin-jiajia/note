package handler

// CreateNoteRequest represents the request structure for creating a note
// @Description Request model for creating a new note
type CreateNoteRequest struct {
    // Title of the note
    Title string `json:"title" binding:"required"`
    // Content of the note
    Body string `json:"body"`
    // Whether the note should be encrypted
    IsEncrypted bool `json:"isEncrypted"`
    // Tag used for encryption verification
    EncryptionTag string `json:"encryptionTag"`
}

// EditNoteRequest represents the request structure for updating a note
// @Description Request model for updating an existing note
type EditNoteRequest struct {
    // Title of the note
    Title string `json:"title"`
    // Content of the note
    Body string `json:"body"`
}
