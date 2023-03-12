package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sing3demons/gin/contrib/logger"
	"github.com/sing3demons/gin/contrib/middleware"
	"github.com/sing3demons/gin/contrib/response"
)

func main() {
	logger, loggerClose, err := logger.Logger("logs")
	if err != nil {
		fmt.Printf("%v", err.Error())
		os.Exit(1)
	}
	defer loggerClose()

	r := gin.Default()
	// middleware
	r.Use(middleware.ZapLogger(logger))
	r.Use(middleware.RecoveryWithZap(logger, true))

	r.GET("/", func(ctx *gin.Context) {
		response.ResponseJsonWithLogger(ctx, http.StatusOK, gin.H{
			"status": "OK",
		})
	})
	r.Run(":8080")
}
