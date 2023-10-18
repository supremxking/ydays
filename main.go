package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"ydays/db"
)

func init() {
	env.Load()
	fmt.Println("key = ", env.Get("key"))
	db.Create_db("ines.fabiao@ynov.com", "Fabiao", "Ines", "Ines345")
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("src/html/*")
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", gin.H{
			//"phrase": "c'est moi aurelie",
			//"title":  "test",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"phrase": "c'est moi aurelie",
			"title":  "test",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
