package main

import (
	"quran-indonesia/handler"
	"quran-indonesia/konstanta"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	r := gin.Default()

	r.GET(konstanta.ListSuraPath(), handler.HandleListSura)
	r.GET(konstanta.AllPath(), handler.HandleAll)
	r.GET(konstanta.TakeFromPath(), handler.HandleTakeFrom)

	r.Run()
}
