package ping

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	fmt.Print("it is comming")
	ctx.JSON(http.StatusOK, gin.H{
		"text": "hello!",
	})
}
