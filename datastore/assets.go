package datastore

import (
	"encoding/csv"
	"fmt"
	"github.com/matt-FFFFFF/bookdata-api/loader"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

type Assets struct {
	Store *[]*loader.BookData `json:"store"`
}

func (a *Assets) Initialize() {
	f, _ := os.Open(a.csvFile())
	defer f.Close()

	lines, _ := csv.NewReader(f).ReadAll()

	bookData := make([]*loader.BookData, len(lines))

	for i, l := range lines {
		d := &loader.BookData{
			BookID:       l[0],
			Title:        l[1],
			Authors:      l[2],
			ISBN:         l[4],
			ISBN13:       l[5],
			LanguageCode: l[6],
		}

		f, _ := strconv.ParseFloat(l[3], 64)
		d.AverageRating = f

		n, _ := strconv.Atoi(l[7])
		d.NumPages = n

		n, _ = strconv.Atoi(l[8])
		d.Ratings = n

		n, _ = strconv.Atoi(l[9])
		d.Reviews = n

		bookData[i] = d
	}
	a.Store = &bookData
}

func (a *Assets) GetAllBooks(limit, skip int) *[]*loader.BookData {
	if limit == 0 || limit > len(*a.Store) {
		limit = len(*a.Store)
	}
	ret := (*a.Store)[skip:limit]
	return &ret
}

func (a *Assets) csvFile() string {
	_, filename, _, _ := runtime.Caller(0)
	rootDir := filepath.Dir(filepath.Dir(filename))
	return filepath.Join(rootDir, "assets", "books.csv")
}

func (a *Assets) AddBook(book loader.BookData) {
	updatedStore := append(*a.Store, &book)
	a.Store = &updatedStore
}

func (a *Assets) DeleteBook(isbn string) error {
	i := -1
	for j, book := range *a.Store {
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
	copy((*a.Store)[i:], (*a.Store)[i+1:])
	(*a.Store)[len(*a.Store)-1] = nil // or the zero value of T
	*a.Store = (*a.Store)[:len(*a.Store)-1]
	return nil
}
