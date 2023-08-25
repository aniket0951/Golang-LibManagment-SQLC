package api

import (
	"github.com/aniket0951/lib_managment/sqlc/db/api/controllers"
	"github.com/aniket0951/lib_managment/sqlc/db/api/services"
	db "github.com/aniket0951/lib_managment/sqlc/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	server.router = router

	authorService := services.NewAuthorService(store)
	bookService := services.NewBookService(store)
	userService := services.NewUserService(store)
	authorController := controllers.NewAuthorController(authorService)
	bookController := controllers.NewBookController(bookService)
	userController := controllers.NewUserController(userService)

	router.POST("/author", authorController.CreateNewAuthor)
	router.GET("/author/:id", authorController.GetAuthor)
	router.GET("/authors", authorController.GetAuthors)
	router.GET("/author/book/:id", authorController.GetAuthorWithBooks)
	router.GET("/author/book/managment/:id", authorController.GetAuthorWithBooksAndManagment)

	router.POST("/book", bookController.AddBook)
	router.GET("/book/:id", bookController.GetBook)
	router.GET("/books", bookController.GetBooks)
	router.PUT("/book/purchase/:id", bookController.PurchaseBook)

	router.POST("/user", userController.AddNewUser)
	router.GET("/user/:id", userController.GetUser)
	router.GET("/users", userController.GetUsers)
	router.PUT("/user", userController.UpdateUser)
	router.DELETE("/user/:id", userController.DeleteUser)
	router.POST("/user/purchase_book", userController.PurchaseBookForUser)
	router.POST("/user/return_book", userController.ReturnUserBook)
	router.GET("/user/books/:id", userController.GetUserWithBooks)

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
