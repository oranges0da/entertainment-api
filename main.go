package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oranges0da/entertainment-api/handler"
)

// RunServer start server and quit if error occurs or user quits
func RunServer(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func main() {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	handler.NewHandler(&handler.Config{
		R: router,
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	RunServer(server)

	log.Printf("Listening on port: %v\n", server.Addr)
}
