package handler

import (
	"quran-indonesia/data"
	"quran-indonesia/models"
	"quran-indonesia/utils"

	"github.com/gin-gonic/gin"
)

var LISTSURAH []models.List

func init() {
	LISTSURAH = data.UnmarshallListCSV()
}

func HandleListSura(ginctx *gin.Context) {
	utils.ResJSON(ginctx, LISTSURAH)
}
