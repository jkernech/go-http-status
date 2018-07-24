package main

import (
  "time"

  "github.com/gin-gonic/gin"
)


var startTime time.Time

// Index handler process incoming HTTP requests of the homepage.
func Index(c *gin.Context) {
  c.JSON(200, gin.H{
    "name": "status",
    "startedAt": startTime,
    "uptime": time.Since(startTime).Seconds(),
  })
}

func main() {
  startTime = time.Now()

  r := gin.Default()

  r.GET("/", Index)
  r.Run() // listen and serve by default on 0.0.0.0:8080
}
