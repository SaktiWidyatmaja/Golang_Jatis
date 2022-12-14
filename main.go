package main

import (
	"eCommerce/internal/delivery"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	delivery.Routes(r)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
