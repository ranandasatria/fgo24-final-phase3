package dto

type CreateProductRequest struct {
	Nama       string `json:"nama" binding:"required"`
	GambarURL  string `json:"gambar_url" binding:"required,url"`
	KategoriID int    `json:"kategori_id" binding:"required"`
	Stok       int    `json:"stok" binding:"required,gte=0"`
	HargaBeli  int    `json:"harga_beli" binding:"required,gte=0"`
	HargaJual  int    `json:"harga_jual" binding:"required,gte=0"`
}
