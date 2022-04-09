package models

type Sync struct {
	TrxList []TrxItems `json:"trxList" bson:"trxList"`
}

type TrxItems struct {
	DocNo            string                   `json:"docNo" bson:"docNo"`
	DocDate          string                   `json:"docDate" bson:"docDate"`
	DeviceId         string                   `json:"deviceId" bson:"deviceId"`
	Gate             string                   `json:"gate" bson:"gate"`
	CardNumberUUID   string                   `json:"cardNumberUuid" bson:"cardNumberUuid"`
	CardNumber       string                   `json:"cardNumber" bson:"cardNumber"`
	TypeCard         string                   `json:"typeCard" bson:"typeCard"`
	ExtLocalDatetime string                   `json:"extLocalDatetime" bson:"extLocalDatetime"`
	GrandTotal       float64                  `json:"grandTotal" bson:"grandTotal"`
	ServiceFee       float64                  `json:"serviceFee" bson:"serviceFee"`
	ProductId        int64                    `json:"productId" bson:"productId"`
	ProductCode      string                   `json:"productCode" bson:"productCode"`
	ProductName      string                   `json:"productName" bson:"productName"`
	Price            float64                  `json:"price" bson:"price"`
	ProductData      PolicyOuProductWithRules `json:"productData" bson:"productData"`
	OuId             int64                    `json:"ouId" bson:"ouId"`
	OuName           string                   `json:"ouName" bson:"ouName"`
	OuCode           string                   `json:"ouCode" bson:"ouCode"`
	MDR              float64                  `json:"mdr" bson:"mdr"`
	MID              string                   `json:"mid" bson:"mid"`
	TID              string                   `json:"tid" bson:"tid"`
	CardType         string                   `json:"cardType" bson:"cardType"`
	CardPan          string                   `json:"cardPan" bson:"cardPan"`
	LastBalance      float64                  `json:"lastBalance" bson:"lastBalance"`
	ResponseWorker   string                   `json:"responseWorker" bson:"responseWorker"`
	CurrentBalance   float64                  `json:"currentBalance" bson:"currentBalance"`
	LogTrans         string                   `json:"logTrans" bson:"logTrans"`
	MerchantKey      string                   `json:"merchantKey" bson:"merchantKey"`
	StatusTransaksi  string                   `json:"statusTransaksi" bson:"statusTransaksi"`
	QrText           *string                  `json:"qrText" bson:"qrText"`
	FlagSyncData     bool                     `json:"flagSyncData" bson:"flagSyncData"`
}
