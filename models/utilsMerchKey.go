package models

//-- INPUT DECRYPT UTILS --
type RequestDecryptMerchKey struct {
	CryptoText string `json:"cryptotext" validate:"required"`
}

//-- RESPONSE ENCRYPT UTILS --
type ResultEncryptMerchKey struct {
	Result string `json:"result"`
}

//-- RESPONSE DECRYPT & INPUT ENCRYPT UTILS --
type MerchantKey struct {
	OuId   int64  `json:"ouId" validate:"required"`
	OuName string `json:"ouName" validate:"required"`
	OuCode string `json:"ouCode" validate:"required"`
}
