package handler

import "github.com/gin-gonic/gin"

func HandleListSura(ginctx *gin.Context) {
	ginctx.Writer.Write([]byte("HandleListSura"))
}
