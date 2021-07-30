package handler

import (
	"net/http"
	"quran-indonesia/data"
	"quran-indonesia/konstanta"
	"quran-indonesia/models"
	"quran-indonesia/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleAll(ginctx *gin.Context) {

	suraString := ginctx.Params.ByName(konstanta.SURA)

	suraInt, err := strconv.Atoi(suraString)
	if err != nil {
		utils.ResErr(ginctx, http.StatusBadRequest, err)
		return
	}

	result, err := data.UnmarshallQuranCSV(func(q models.Quran) (models.Quran, bool, bool) {

		sura, _ := strconv.Atoi(q.Sura)

		if sura == suraInt {
			return q, true, false
		}
		if sura > suraInt {
			return models.Quran{}, false, true
		}
		return models.Quran{}, false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResJSON(ginctx, result)

}
