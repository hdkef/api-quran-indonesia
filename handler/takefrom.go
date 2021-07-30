package handler

import "github.com/gin-gonic/gin"

func HandleTakeFrom(ginctx *gin.Context) {
	ginctx.Writer.Write([]byte("HandleTakeFrom"))
}
