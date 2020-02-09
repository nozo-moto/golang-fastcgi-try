package main

import (
	"net"
	"net/http/fcgi"

	"github.com/gin-gonic/gin"
)

func main() {
    var err error
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
    //if err = r.Run(":8080"); err != nil {
    //    panic(err)
    //}
    l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	if err := fcgi.Serve(l, r); err != nil {
		panic(err)
	}
}
