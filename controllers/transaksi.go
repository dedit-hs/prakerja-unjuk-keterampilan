package controllers

import (
	"net/http"
	"prakerja/config"
	"prakerja/models"

	"github.com/labstack/echo/v4"
)

func GetTransaksiController(c echo.Context) error {
	var transaksi []models.Transaksi

	result := config.DB.Find(&transaksi)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.FailResponse{
			Message: "Terjadi kesalahan server, gagal mendapatkan data transaksi.",
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    transaksi,
	})
}

func GetTransaksiTerakhirController(c echo.Context) error {
	var transaksi []models.Transaksi
	nasabah_id := c.Param("nasabah_id")
	rekening_id := c.Param("rekening_id")

	result := config.DB.Joins("JOIN rekenings ON rekenings.id = transaksis.rekening_id").Joins("Join nasabahs ON nasabahs.id = rekenings.nasabah_id").Where("nasabahs.id = ? AND rekenings.id = ?", nasabah_id, rekening_id).Order("created_at").Limit(5).Find(&transaksi)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.FailResponse{
			Message: "Terjadi kesalahan server, gagal mendapatkan transaksi terakhir.",
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    transaksi,
	})
}

func KreditTransaksiController(c echo.Context) error {
	var newTransaksi models.Transaksi
	var rekening models.Rekening
	c.Bind(&newTransaksi)

	cekRekening := config.DB.First(&rekening, newTransaksi.RekeningID)

	if cekRekening.Error == nil {
		result := config.DB.Create(&newTransaksi)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, models.FailResponse{
				Message: "Terjadi kesalahan server, gagal menyimpan transaksi.",
			})
		}

		newSaldo := rekening.Saldo + newTransaksi.Nominal
		config.DB.Model(&rekening).Where("id = ?", newTransaksi.RekeningID).Update("saldo", newSaldo)
		return c.JSON(http.StatusCreated, models.BaseResponse{
			Message: "Success",
			Data:    newTransaksi,
		})
	}
	return c.JSON(http.StatusBadRequest, models.FailResponse{
		Message: "Terjadi kesalahan, rekening tidak ditemukan.",
	})

}

func DebitTransaksiController(c echo.Context) error {
	var newTransaksi models.Transaksi
	var rekening models.Rekening
	c.Bind(&newTransaksi)

	cekRekening := config.DB.First(&rekening, newTransaksi.RekeningID)

	if cekRekening.Error == nil {
		result := config.DB.Create(&newTransaksi)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, models.FailResponse{
				Message: "Terjadi kesalahan server, gagal menyimpan transaksi.",
			})
		}

		newSaldo := rekening.Saldo - newTransaksi.Nominal
		if newSaldo <= 10000 {
			return c.JSON(http.StatusBadRequest, models.FailResponse{
				Message: "Saldo anda tidak mencukupi.",
			})
		}
		config.DB.Model(&rekening).Where("id = ?", newTransaksi.RekeningID).Update("saldo", newSaldo)
		return c.JSON(http.StatusCreated, models.BaseResponse{
			Message: "Success",
			Data:    newTransaksi,
		})
	}
	return c.JSON(http.StatusBadRequest, models.FailResponse{
		Message: "Terjadi kesalahan, rekening tidak ditemukan.",
	})

}
