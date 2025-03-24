package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestGin(t *testing.T) {
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "TestGin")
	})

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
