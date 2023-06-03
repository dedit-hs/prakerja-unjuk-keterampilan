package controllers

import (
	"net/http"
	"prakerja/config"
	"prakerja/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetRekeningController(c echo.Context) error {
	var rekening []models.Rekening
	result := config.DB.Preload("Transaksi").Find(&rekening)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.FailResponse{
			Message: "Terjadi kesalahan server, Gagal mendapatkan data rekening.",
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    rekening,
	})
}

func GetRekeningByIdController(c echo.Context) error {
	var rekening models.Rekening
	id := c.Param("id")

	result := config.DB.Preload("Transaksi").First(&rekening, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.FailResponse{
			Message: "Terjadi kesalahan server, gagal mendapatkan data rekening.",
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    rekening,
	})
}

func AddRekeningController(c echo.Context) error {
	var newRekening models.Rekening
	c.Bind(&newRekening)

	var nasabah models.Nasabah
	cekNasabah := config.DB.Where("id = ?", &newRekening.NasabahID).First(&nasabah)
	if cekNasabah.Error == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusBadRequest, models.FailResponse{
			Message: "Gagal membuat rekening, Nasabah tidak ditemukan.",
		})
	}

	var rekening models.Rekening
	cekRekening := config.DB.Where("no_rekening = ?", &newRekening.NoRekening).First(&rekening)
	if cekRekening.Error == gorm.ErrRecordNotFound {
		result := config.DB.Create(&newRekening)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, models.FailResponse{
				Message: "Terjadi kesalahan server, rekening gagal disimpan.",
			})
		}
		return c.JSON(http.StatusCreated, models.BaseResponse{
			Message: "Success",
			Data:    newRekening,
		})
	}
	return c.JSON(http.StatusBadRequest, models.FailResponse{
		Message: "Failed, No Rekening sudah terdaftar.",
	})
}
