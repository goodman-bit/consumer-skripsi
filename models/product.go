package models

type Product struct {
	Id               int64   `json:"id"`
	KodeProduk       string  `json:"kodeproduk"`
	NamaProduk       string  `json:"namaproduk"`
	HargaJual        float64 `json:"hargajual"`
	IdKategoriProduk int64   `json:"idkategoriproduk"`
	IdTerminal       int64   `json:"idterminal"`
	Created_at       string  `json:"created_at"`
	ServiceFee       float64 `json:"servicefee"`
}
