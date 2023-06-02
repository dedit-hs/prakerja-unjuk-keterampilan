package controllers

import (
	"net/http"
	"prakerja/config"
	"prakerja/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetNasabahController(c echo.Context) error {
	var nasabah []models.Nasabah
	result := config.DB.Preload("Rekening").Find(&nasabah)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.FailResponse{
			Message: "Terjadi kesalahan server, Gagal mendapatkan data nasabah.",
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    nasabah,
	})
}

func GetNasabahByIdController(c echo.Context) error {
	var nasabah models.Nasabah
	id := c.Param("id")
	result := config.DB.Preload("Rekening").First(&nasabah, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.FailResponse{
			Message: "Terjadi kesalahan server, gagal mendapatkan data nasabah.",
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    nasabah,
	})
}

func AddNasabahController(c echo.Context) error {
	var newNasabah models.Nasabah
	c.Bind(&newNasabah)

	var nasabah models.Nasabah
	cekNasabah := config.DB.Where("nik = ?", &newNasabah.NIK).First(&nasabah)
	if cekNasabah.Error == gorm.ErrRecordNotFound {
		result := config.DB.Create(&newNasabah)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, models.FailResponse{
				Message: "Terjadi kesalahan server, data nasabah gagal disimpan.",
			})
		}
		return c.JSON(http.StatusCreated, models.BaseResponse{
			Message: "Success",
			Data:    newNasabah,
		})
	}
	return c.JSON(http.StatusBadRequest, models.FailResponse{
		Message: "Failed, NIK sudah terdaftar.",
	})
}
