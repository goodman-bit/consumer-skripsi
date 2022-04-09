package models

import "time"

// GLOBAL ERROR MSG
type ErrorMsg struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type WhitelistRoute struct {
	OuCode    string `json:"ouCode" bson:"ouCode"`
	SuffixUrl string `json:"suffixUrl" bson:"suffixUrl"`
}

type ResponseConfirmPayment struct {
	Message          string    `json:"message"`
	ResponseDatetime time.Time `json:"responseDatetime"`
	Result           struct {
		AcquiringID     int     `json:"acquiringId"`
		Amount          float64 `json:"amount"`
		BankNoRef       string  `json:"bankNoRef"`
		CardPan         string  `json:"cardPan"`
		CardType        string  `json:"cardType"`
		CorporateName   string  `json:"corporateName"`
		CreatedAt       string  `json:"createdAt"`
		CurrentBalance  float64 `json:"currentBalance"`
		DeviceID        string  `json:"deviceId"`
		Discount        float64 `json:"discount"`
		LastBalance     float64 `json:"lastBalance"`
		Mdr             float64 `json:"mdr"`
		NoHeader        string  `json:"noHeader"`
		PaymentCategory string  `json:"paymentCategory"`
		PaymentFee      float64 `json:"paymentFee"`
		PromoCode       string  `json:"promoCode"`
		PromoIssuer     string  `json:"promoIssuer"`
		ServiceFee      float64 `json:"serviceFee"`
		StatusPayment   string  `json:"statusPayment"`
		StatusCode      string  `json:"statusCode"`
	} `json:"result"`
	Status  string `json:"status"`
	Success bool   `json:"success"`
}
