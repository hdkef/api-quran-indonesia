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

func HandleTakeFrom(ginctx *gin.Context) {
	suraString := ginctx.Params.ByName(konstanta.SURA)
	takeString := ginctx.Params.ByName(konstanta.TAKEVALUE)
	fromString := ginctx.Params.ByName(konstanta.FROMVALUE)

	suraInt, err := strconv.Atoi(suraString)
	if err != nil {
		utils.ResErr(ginctx, http.StatusBadRequest, err)
		return
	}

	takeInt, err := strconv.Atoi(takeString)
	if err != nil {
		utils.ResErr(ginctx, http.StatusBadRequest, err)
		return
	}

	fromInt, err := strconv.Atoi(fromString)
	if err != nil {
		utils.ResErr(ginctx, http.StatusBadRequest, err)
		return
	}

	result, err := data.UnmarshallQuranCSV(func(q interface{}) (interface{}, bool, bool) {

		sura, _ := strconv.Atoi(q.(models.Quran).Sura)
		aya, _ := strconv.Atoi(q.(models.Quran).Aya)

		if sura == suraInt && aya >= fromInt {
			if aya >= fromInt+takeInt {
				return q, false, true
			}
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
