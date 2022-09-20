package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
)

func main() {
	dsn := "root:warouw1945@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Errors")
	}
	//fmt.Println("Database Connection Succeed ")
	db.AutoMigrate(&book.Book{})

	//book := book.Book{}
	//book.Title = "Majalah Mudah"
	//book.Price = 10
	//book.Rating = 9
	//book.Description = "Merupakan buku panduan yang bagus dan mudah"
	//
	//err = db.Create(&book).Error
	//if err != nil {
	//	fmt.Println("=======================")
	//	fmt.Println("Error creating book record")
	//	fmt.Println("============================")
	//}

	//dibawah ini processing data sebelum di lakukan action atau melakukan query
	//var books []book.Book
	//err = db.Debug().Where("rating = ?", 9).Find(&books).Error
	//if err != nil {
	//
	//	fmt.Println("=======================")
	//	fmt.Println("Error finding book record")
	//	fmt.Println("============================")
	//}
	//
	//for _, b := range books {
	//	//melihat data atau mengambil data
	//	fmt.Println("Title:", b.Title)
	//	fmt.Printf(`Book objerct %v`, b)
	//
	//	// mengupdate data
	//
	//	b.Title = "Amplay Version"
	//	err = db.Debug().Save(&books).Error
	//	if err != nil {
	//		fmt.Println("=======================")
	//		fmt.Println("Error updeating book record")
	//		fmt.Println("============================")
	//	}
	//}

	//dibawah ini processing data sebelum di lakukan action atau melakukan query
	//var book book.Book
	//err = db.Debug().Where("id = ?", 2).Find(&book).Error
	//if err != nil {
	//
	//	fmt.Println("=======================")
	//	fmt.Println("Error finding book record")
	//	fmt.Println("============================")
	//}
	//
	//// mengupdate data
	//
	//book.Title = "Tangkuban Version"
	//book.Price = 1000
	//err = db.Debug().Save(&book).Error
	//if err != nil {
	//	fmt.Println("=======================")
	//	fmt.Println("Error updeating book record")
	//	fmt.Println("============================")
	//}

	// delete book
	//var book book.Book
	//err = db.Debug().Where("id = ?", 1).Find(&book).Error
	//if err != nil {
	//	fmt.Println("=======================")
	//	fmt.Println("Error finding book record")
	//	fmt.Println("============================")
	//
	//}

	//err = db.Debug().Delete(&book).Error
	//if err != nil {
	//	fmt.Println("=======================")
	//	fmt.Println("Error Deleteing book record")
	//	fmt.Println("============================")
	//}

	// memanggil repository
	bookRepository := book.NewRepository(db)

	//books, err := bookRepository.FindAll()
	//for _, book := range books {
	//	fmt.Println("Title: ", book.Title)
	//}

	//book, err := bookRepository.FindByID(2)

	//fmt.Println("Title: ", book.Title)

	//book := book.Book{
	//	Title:       "Amreican Union",
	//	Description: "ini booookk mahal",
	//	Price:       100000,
	//	Rating:      100,
	//	Discount:    50,
	//}
	//bookRepository.Create(book)

	bookService := book.NewService(bookRepository)
	//bookHandler := handler.NewBookHandler(bookService)
	bookHandler := handler.NewBookHandler(bookService)
	//bookRequest := book.BookRequest{
	//	Title: "Spain Union",
	//	Price: "100000",
	//}
	//bookService.Create(bookRequest)

	router := gin.Default()

	V1 := router.Group("/v1")
	//V1.GET("/", bookHandler.RootHandler)
	//V1.GET("/hello", bookHandler.HelloHandler)
	//V1.GET("/books/:id", bookHandler.BooksHandler)
	//V1.GET("/book/:id/:title", bookHandler.BookHandler)
	//V1.GET("/query", bookHandler.QueryHandler)
	V1.GET("/books", bookHandler.GetBooksHandler)
	V1.GET("/books/:id", bookHandler.GetBookByIDHandler)
	V1.POST("/books", bookHandler.CreateBooksHandler)
	V1.PUT("books/:id", bookHandler.UpdateBooksHandler)
	V1.DELETE("books/:id", bookHandler.DeleteBooksHandler)

	router.Run()
}
