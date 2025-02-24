package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/d3vus/projects/go/noteai/config"
	"gitlab.com/d3vus/projects/go/noteai/handler"
	"gitlab.com/d3vus/projects/go/noteai/model"
)

func main() {

	// Initiate the OpenAI configuration
	cfg := config.LoadConfig()
	if cfg.Openaikey == "" {

		fmt.Println("Provide a Key to Further any chat with GPT")
		return
	}

	// Intiate the Store
	store := model.NewNoteStore()
	noteHandler := handler.NewNoteHandler(store, cfg.Openaikey)

	// Initiate Gin router
	router := gin.Default()

	// Serve Static files
	router.Static("/static","./static")

	// Load HTML template file
	router.LoadHTMLGlob("template/*")

	// Intiate Routing
	router.GET("/", func(c *gin.Context) {
		
		c.HTML(http.StatusOK, "notes.html", nil)
	} )

	api := router.Group("/api")
	{
		api.POST("/notes", noteHandler.CreateNewNote)
	}
	
	// Now Start the Engine
	router.Run(":8080")
	
}
