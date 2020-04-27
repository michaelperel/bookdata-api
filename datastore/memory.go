package datastore

import (
	"fmt"
	"github.com/matt-FFFFFF/bookdata-api/loader"
)

// Books is the memory-backed datastore used by the API
// It contains a single field 'Store', which is (a pointer to) a slice of loader.BookData struct pointers
type Books struct {
	Store *[]*loader.BookData `json:"store"`
}

// Initialize is the method used to populate the in-memory datastore.
// At the beginning, this simply returns a pointer to the struct literal.
// You need to change this to load data from the CSV file
func (b *Books) Initialize() {
	b.Store = &loader.BooksLiteral
}

// GetAllBooks returns the entire dataset, subjet to the rudimentary limit & skip parameters
func (b *Books) GetAllBooks(limit, skip int) *[]*loader.BookData {
	if limit == 0 || limit > len(*b.Store) {
		limit = len(*b.Store)
	}
	ret := (*b.Store)[skip:limit]
	return &ret
}

func (b *Books) AddBook(book loader.BookData) {
	updatedStore := append(*b.Store, &book)
	b.Store = &updatedStore
}

func (b *Books) DeleteBook(isbn string) error {
	i := -1
	for j, book := range *b.Store {
		if book.ISBN == isbn {
			i = j
			break
		}
	}
	if i == -1 {
		return fmt.Errorf("book with '%s' ISBN not found.", isbn)
	}
	// delete without memory leak
	// https://github.com/golang/go/wiki/SliceTricks
	copy((*b.Store)[i:], (*b.Store)[i+1:])
	(*b.Store)[len(*b.Store)-1] = nil // or the zero value of T
	*b.Store = (*b.Store)[:len(*b.Store)-1]
	return nil
}
