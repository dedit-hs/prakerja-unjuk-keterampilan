package models

import "time"

type Transaksi struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	KodeTransaksi string    `json:"kode_transaksi"`
	Nominal       uint      `json:"nominal"`
	RekeningID    uint      `json:"rekening_id"`
	CreatedAt     time.Time `json:"created_at"`
}
