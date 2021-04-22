package book

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// book struct that will implement IBook interface
type bookController struct {
	bookservice BookService
}

// NewBookController return bookController struct with IBook service
// this method is the entry to this controller
func NewBookController(bookservice BookService) *bookController {
	return &bookController{bookservice: bookservice}
}

// NewBook returns the new created book
func (b *bookController) UploadNewBook(c echo.Context) error {

	book := NewBookRequest{}
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}
	src, err := file.Open()
	if err != nil {
		panic(err)
	}

	book.File = src

	if err := c.Bind(&book); err != nil {
		panic(err)
	}

	b.bookservice.NewBookProject(book)

	return c.String(http.StatusOK, "book has been uploaded")

}