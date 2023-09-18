package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status      bool   `json:"status"`
	Message     string `json:"message"`
	TotalRating string `json:"data"`
}

var totalInput int = 0 //Akumulasi pemberian rating
var akumulasi int = 0  //Akumulasi total skor rating

func main() {
	//Inisiasi Echo
	e := echo.New()

	e.GET("/rating", GetRatingController)              //Untuk melihat nilai rating terkini
	e.GET("/ratingAdd/:nilai", GetAddRatingController) //Untuk menambahkan rating
	e.GET("/reset", ResetRatingController)             //Untuk mereset attempt rating dan akumulasi rating menjadi 0

	//Memulai echo
	e.Start(":8000")
}

func GetRatingController(c echo.Context) error {

	//Melakukan update pada variabel teksKirimRating
	teksKirimRating := "Total Skor saat ini adalah: " + strconv.Itoa(akumulasi) + "dari " + strconv.Itoa(totalInput) + "input"

	return c.JSON(http.StatusOK, BaseResponse{
		Status:      true,
		Message:     "Rating terbaru sedang loading ~",
		TotalRating: teksKirimRating,
	})
}

func GetAddRatingController(c echo.Context) error {
	//Menambahkan jumlah input yang dilakukan
	totalInput++

	//Mengambil nilai rating yang diinput user
	nilai, _ := strconv.Atoi(c.Param("nilai"))

	//Memasukkan nilai baru akumulasi dari total rating ditambah nilai rating baru
	akumulasi += nilai

	//Melakukan update pada variabel teksKirimRating
	teksKirimRating := "Total Skor saat ini adalah: " + strconv.Itoa(akumulasi) + "dari " + strconv.Itoa(totalInput) + "input"

	return c.JSON(http.StatusOK, BaseResponse{
		Status:      true,
		Message:     "Rating terbaru sedang loading ~",
		TotalRating: teksKirimRating,
	})
}

func ResetRatingController(c echo.Context) error {
	//Reset nilai global totalInput dan akumulasi
	totalInput = 0
	akumulasi = 0

	//Melakukan update pada variabel teksKirimRating
	teksKirimRating := "Total Skor saat ini adalah: " + strconv.Itoa(akumulasi) + "dari " + strconv.Itoa(totalInput) + "input"

	return c.JSON(http.StatusOK, BaseResponse{
		Status:      true,
		Message:     "Rating terbaru sedang loading ~",
		TotalRating: teksKirimRating,
	})
}
