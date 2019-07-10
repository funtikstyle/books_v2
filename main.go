package main

import (
	"books_v2/handler"
	"books_v2/service"
	"books_v2/store"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	books := store.NewBooks()
	srv := service.NewService(books)
	handlers := handler.NewHandlers(srv)
	router := httprouter.New()

	router.GET("/user/:id", handlers.GetContact)
	router.GET("/users/list/", handlers.GetContactList)
	router.POST("/user/", handlers.AddContact)
	router.PUT("/user/:id", handlers.UpdateContact)
	router.DELETE("/user/:id", handlers.DeleteContact)

	fmt.Println("Start server")
	log.Fatal(http.ListenAndServe(":8088", router))
}
