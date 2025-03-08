package main

import (
	"github.com/gin-gonic/gin"

)

func testHandler(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"status": "return test"})
}
