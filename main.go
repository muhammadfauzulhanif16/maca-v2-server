package main

import (
	"maca/auth"
	"maca/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(auth.CORS())

	routes.Book(r)

	if err := r.Run(); err != nil {
		return
	}
}
