package models

//-- FIND DEVICE IN DYNAMO DB --
type DeviceDB struct {
	DeviceID    string `json:"deviceID" bson:"deviceId"`
	MerchantKey string `json:"merchantKey" bson:"merchantKey"`
}

// Device ...
type Device struct {
	ID           int     `json:"id"`
	Device_ID    string  `json:"device_id"`
	MID          *string `json:"mid"`
	TID          *string `json:"tid"`
	IDCorporate  int     `json:"idcorporate"`
	AcquiringMID *string `json:"acquiringmid"`
	AcquiringTID *string `json:"acquiringtid"`
	MerchantKey  string  `json:"merchantkey"`
	Created_at   string  `json:"created_at"`
	Updated_at   *string `json:"updated_at"`
}
