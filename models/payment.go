package models

//-- REQUEST WORKER BULKPREPAID APP2PAY --
type InquiryPrepaidBatchApp2Pay struct {
	BatchCode      string                            `json:"batchcode"`
	JumlahData     int                               `json:"jumlahdata"`
	TotalTransaksi float64                           `json:"totaltransaksi"`
	MerchantKey    string                            `json:"merchantkey"`
	InquiryData    []InquiryItemsPrepaidBatchApp2Pay `json:"inquirydata"`
}

type InquiryItemsPrepaidBatchApp2Pay struct {
	MerchantKey   string  `json:"merchantkey"`
	MerchantNoRef string  `json:"merchantnoref"`
	Nominal       float64 `json:"nominal"`
	ServiceFee    float64 `json:"servicefee"`
	LogTransaksi  string  `json:"logtransaksi"`
	DeviceID      string  `json:"device_id"`
}

//WORKER
type WorkerResponse struct {
	StatusCode      string      `json:"statusCode" bson:"statusCode"`
	StatusDesc      string      `json:"statusDesc" bson:"statusDesc"`
	SnapshotData    interface{} `json:"snapshotData" bson:"snapshotData"`
	URL             string      `json:"url" bson:"URL"`
	RequestDatetime string      `json:"requestDatetime"`
	ResponseBody    interface{} `json:"responseBody" bson:"responseBody"`
}

//-- WORKER APPS2PAY --
type RequestInquiryBatchApp2Pay struct {
	BatchCode string                   `json:"batchcode"`
	Detail    []RequestParkingApps2pay `json:"detail"`
}

type RequestParkingApps2pay struct {
	KodeStatus             string  `json:"kodestatus" bson:"kodestatus"`
	StatusPayment          string  `json:"statuspayment" bson:"statuspayment"`
	Description            string  `json:"description" bson:"description"`
	MerchantNoRef          string  `json:"merchantnoref" bson:"merchantnoref"`
	LogTransaksi           string  `json:"logtransaksi" bson:"logtransaksi"`
	Device_ID              string  `json:"device_id" bson:"device_id"`
	NoHeader               string  `json:"noheader" bson:"noheader"`
	BankNoRef              string  `json:"banknoref" bson:"banknoref"`
	CardType               string  `json:"cardtype" bson:"cardtype"`
	CardPan                string  `json:"cardpan" bson:"cardpan"`
	Mid                    string  `json:"mid" bson:"mid"`
	Tid                    string  `json:"tid" bson:"tid"`
	LastBalance            float64 `json:"lastbalance" bson:"lastbalance"`
	CurrentBalance         float64 `json:"currentbalance" bson:"currentbalance"`
	NamaKategoriPembayaran string  `json:"namakategoripembayaran" bson:"namakategoripembayaran"`
	TotalJual              float64 `json:"totaljual" bson:"totaljual"`
	Mdr                    float64 `json:"mdr" bson:"mdr"`
	Potongan               float64 `json:"potongan" bson:"potongan"`
	KodePromo              string  `json:"kodepromo" bson:"kodepromo"`
	PromoIssuer            string  `json:"promoissuer" bson:"promoissuer"`
	Created_at             string  `json:"created_at" bson:"created_at"`
	Settlement_datetime    string  `json:"settlement_datetime" bson:"settled_at"`
	Deduct_datetime        string  `json:"deduct_datetime" bson:"deduct_datetime"`
}

// {
// 	"batchcode":"MKP-0001"
// 	"detail":[
// 	  {
// 		"kodestatus":"200",
// 		"statuspayment":"SUCCESS",
// 		"description":"",
// 		"merchantnoref":"",
// 		"noheader":"",
// 		"banknoref":"",
// 		"cardtype":"",
// 		"cardpan":"",
// 		"mid":"",
// 		"tid":"",
// 		"lastbalance":1000,
// 		"currentbalance":2000,
// 		"namakategoripembayaran":"",
// 		"settlementdatetime":"",
// 		"deductdatetime":"",
// 		"mdr":10,
// 		"potongan":0,
// 		"kodepromo":"",
// 		"promoissuer":"",
// 		"created_at":""
// 	  },
// 	  {
// 		 "kodestatus":"200",
// 		"statuspayment":"SUCCESS",
// 		"description":"",
// 		"merchantnoref":"",
// 		"noheader":"",
// 		"banknoref":"",
// 		"cardtype":"",
// 		"cardpan":"",
// 		"mid":"",
// 		"tid":"",
// 		"lastbalance":1000,
// 		"currentbalance":2000,
// 		"namakategoripembayaran":"",
// 		"settlementdatetime":"",
// 		"deductdatetime":"",
// 		"mdr":10,
// 		"potongan":0,
// 		"kodepromo":"",
// 		"promoissuer":"",
// 		"created_at":""
// 	  },
// 	  {
// 		 "kodestatus":"200",
// 		"statuspayment":"SUCCESS",
// 		"description":"",
// 		"merchantnoref":"",
// 		"noheader":"",
// 		"banknoref":"",
// 		"cardtype":"",
// 		"cardpan":"",
// 		"mid":"",
// 		"tid":"",
// 		"lastbalance":1000,
// 		"currentbalance":2000,
// 		"namakategoripembayaran":"",
// 		"settlementdatetime":"",
// 		"deductdatetime":"",
// 		"mdr":10,
// 		"potongan":0,
// 		"kodepromo":"",
// 		"promoissuer":"",
// 		"created_at":""
// 	  }
// 	  ]
//   }
