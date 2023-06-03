package models

import "time"

type Rekening struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	NasabahID  uint        `json:"nasabah_id"`
	NoRekening uint        `json:"no_rekening"`
	Saldo      uint        `json:"saldo"`
	Transaksi  []Transaksi `json:"transaksi" gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
