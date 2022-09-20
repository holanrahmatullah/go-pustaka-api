package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"pustaka-api/book"
	"strconv"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

//
//func (h *bookHandler) RootHandler(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"name": "Holan Rahmatullah S N",
//		"bio":  "A software engginer QA",
//	})
//}
//func (h *bookHandler) HelloHandler(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"title":    "Hello name",
//		"subtitle": "what's are yo u name?",
//	})
//
//}
//func (h *bookHandler) BookHandler(c *gin.Context) {
//	id := c.Param("id")
//	title := c.Param("title")
//
//	c.JSON(http.StatusOK, gin.H{
//		"id":    id,
//		"title": title,
//	})
//}
//func (h *bookHandler) BooksHandler(c *gin.Context) {
//	id := c.Param("id")
//
//	c.JSON(http.StatusOK, gin.H{
//		"id": id,
//	})
//}
//func (h *bookHandler) QueryHandler(c *gin.Context) {
//	title := c.Query("title")
//	price := c.Query("price")
//	c.JSON(http.StatusOK, gin.H{
//		"title": title,
//		"price": price,
//	})
//}

func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}
func (h *bookHandler) GetBooksHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBookByIDHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	bookR, err := h.bookService.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	bookResponse := convertToBookResponse(bookR)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

func (h *bookHandler) CreateBooksHandler(c *gin.Context) {
	var bookRequest book.CreateBookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, conditon: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}
	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
		//"title": bookInput.Title,
		//"price": bookInput.Price,
		//"sub_title": bookInput.SubTitle,
	})

}

func (h *bookHandler) UpdateBooksHandler(c *gin.Context) {
	var bookRequest book.UpdateBookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, conditon: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.Update(id, bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	})

}

func (h *bookHandler) DeleteBooksHandler(c *gin.Context) {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.Delete(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	})

}
