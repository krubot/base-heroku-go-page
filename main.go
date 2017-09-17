
package main

import (
        "log"
        "net/http"
        "os"
        "github.com/gin-gonic/gin"
        _ "github.com/heroku/x/hmetrics/onload"
)

func main() {

	// heroku sets the PORT environment variable      
        port := os.Getenv("PORT")
        if port == "" {
        	log.Fatal("$PORT must be set")
	}

	// configure gin router with webpage html
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("index.html")
	router.GET("/", func( content *gin.Context) {
		content.HTML(http.StatusOK,"index.html",nil)
	})
	router.Run(":"+port)

}
