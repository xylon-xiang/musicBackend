package main

import (
	"MusicBackend/controller"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func main() {

	e := echo.New()

	e.GET("/song", SearchHandle)

	e.Logger.Fatal(e.Start(":1324"))

}

func SearchHandle(context echo.Context) error {

	keywords := context.QueryParam("keywords")
	limitStr := context.QueryParam("limit")
	page := context.QueryParam("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return context.String(http.StatusInternalServerError, "wrong")
	}

	musicInfo, err := controller.SearchNeteaseCloudMusic(keywords, limit, 1)
	if err != nil {
		return context.String(http.StatusInternalServerError, "wrong")
	}

	musicInfo2, err := controller.SearchQQCloudMusic(keywords, page)
	if err != nil {
		return context.String(http.StatusInternalServerError, "wrong")
	}

	for _, info := range *musicInfo2 {
		*musicInfo = append(*musicInfo, info)
	}

	return context.JSON(http.StatusOK, *musicInfo)
}
