package dto

type CreateTransactionDTO struct {
	ProdukID int    `json:"produk_id" binding:"required"`
	Tipe     string `json:"tipe" binding:"required,oneof=IN OUT"`
	Jumlah   int    `json:"jumlah" binding:"required,gt=0"`
}