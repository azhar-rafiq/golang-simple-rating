package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status      bool        `json:"status"`
	Message     string      `json:"message"`
	TotalRating interface{} `json:"data"`
}

var totalInput int = 0
var akumulasi int = 0

// var teksKirimRating string = "Total Skor saat ini adalah: %akumulasi dari %totalInput input"

func main() {
	e := echo.New()

	e.GET("/rating", GetRatingController)
	e.GET("/rating/:nilai", GetAddRatingController)
	e.Start(":8000")

	fmt.Println("Selamat datang di sistem Rating Input! ^-^)/ ")
}

func GetRatingController(c echo.Context) error {

	teksKirimRating := "Total Skor saat ini adalah: " + strconv.Itoa(akumulasi) + "dari " + strconv.Itoa(totalInput) + "input"

	return c.JSON(http.StatusOK, BaseResponse{
		Status:      true,
		Message:     "Rating terbaru sedang loading ~",
		TotalRating: teksKirimRating,
	})
}

func GetAddRatingController(c echo.Context) error {
	totalInput++
	nilai, _ := strconv.Atoi(c.Param("nilai"))

	// akumulasiBaru := akumulasi + nilai

	akumulasi += nilai

	teksKirimRating := "Total Skor saat ini adalah: " + strconv.Itoa(akumulasi) + "dari " + strconv.Itoa(totalInput) + "input"

	return c.JSON(http.StatusOK, BaseResponse{
		Status:      true,
		Message:     "Rating terbaru sedang loading ~",
		TotalRating: teksKirimRating,
	})
}
