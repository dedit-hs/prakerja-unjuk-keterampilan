package routes

import (
	"prakerja/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/nasabah", controllers.GetNasabahController)
	e.GET("/nasabah/:id", controllers.GetNasabahByIdController)
	e.POST("/nasabah", controllers.AddNasabahController)

	e.GET("/rekening", controllers.GetRekeningController)
	e.GET("/rekening/:id", controllers.GetRekeningByIdController)
	e.POST("/rekening", controllers.AddRekeningController)

	e.GET("transaksi", controllers.GetTransaksiController)
	e.GET("transaksi/:nasabah_id/rekening/:rekening_id", controllers.GetTransaksiTerakhirController)
	e.POST("transaksi/kredit", controllers.KreditTransaksiController)
	e.POST("transaksi/debit", controllers.DebitTransaksiController)
	return e
}
