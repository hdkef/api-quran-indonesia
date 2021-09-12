package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResJSON(ginctx *gin.Context, data interface{}) {
	respond, err := json.Marshal(data)
	if err != nil {
		ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}
	ginctx.Writer.Write(respond)
}
