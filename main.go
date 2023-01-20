package main

import (
 "github.com/gin-gonic/gin"
 "net/http"
)

func main() {
  g := gin.Default()
  g.LoadHTMLGlob("templates/*.html")
  g.GET("/", func(c *gin.Context){ 
    c.HTML(http.StatusOK,"paha.html",nil)
  })
  g.RunTLS(":8080","./cert/server.pem", "./cert/server-key.pem")
}