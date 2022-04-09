package models

type PaymentMethod struct {
	Id          int64  `json:"id"`
	PaymentName string `json:"paymentName"`
}

type PaymentVendor struct {
	Id         int64  `json:"id"`
	VendorName string `json:"vendorName"`
}
