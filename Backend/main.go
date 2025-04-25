package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	// json tag to serialize json body
	Name string `json:"name"`
}

func main() {
	engine := gin.New()
	engine.POST("/test", func(context *gin.Context) {
		body := "hola"
		// using BindJSON method to serialize body with struct
		if err := context.BindJSON(&body); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body)
		context.JSON(http.StatusAccepted, "hola")
	})
	engine.Run(":3000")
}
