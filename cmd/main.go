package main

import (
	"fmt"
	"github.com/brightoneqq/go-pegasus/pkg/router"
	"github.com/brightoneqq/go-pegasus/pkg/setting"
	"net/http"
)

func main() {
	//router := gin.Default()
	//router.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})
	x := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        x,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
