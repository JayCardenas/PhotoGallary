package photos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Create Request"})

}

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get request"})

}
