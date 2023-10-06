package models

// Request represents a user request entity (e.g., issue or return).
type Request struct {
	ID          int    `json:"id"`
	BookISBN    string `json:"book_isbn"`
	ReaderEmail string `json:"reader_email"`
	RequestType string `json:"request_type"` // Issue or Return
}
