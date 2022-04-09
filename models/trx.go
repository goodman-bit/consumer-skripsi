package models

type Trx struct {
	DocNo                          string                `json:"docNo" bson:"docNo"`
	DocDate                        string                `json:"docDate" bson:"docDate"`
	ExtDocNo                       string                `json:"extDocNo" bson:"extDocNo"`
	CheckInDatetime                string                `json:"checkInDatetime" bson:"checkInDatetime"`
	CheckOutDatetime               string                `json:"checkOutDatetime" bson:"checkOutDatetime"`
	DeviceIdIn                     string                `json:"deviceIdIn" bson:"deviceIdIn"`
	DeviceId                       string                `json:"device_id" bson:"deviceId"`
	GateIn                         string                `json:"gateIn" bson:"gateIn"`
	GateOut                        string                `json:"gateOut" bson:"gateOut"`
	MemberCode                     string                `json:"memberCode" bson:"memberCode"`
	MemberName                     string                `json:"memberName" bson:"memberName"`
	MemberType                     string                `json:"memberType" bson:"memberType"`
	CardNumberUUIDIn               string                `json:"cardNumberUuidIn" bson:"cardNumberUuidIn"`
	CardNumberIn                   string                `json:"cardNumberIn" bson:"cardNumberIn"`
	CardNumberUUID                 string                `json:"cardNumberUuid" bson:"cardNumberUuid"`
	CardNumber                     string                `json:"cardNumber" bson:"cardNumber"`
	TypeCard                       string                `json:"typeCard" bson:"typeCard"`
	BeginningBalance               float64               `json:"beginningBalance" bson:"beginningBalance"`
	ExtLocalDatetime               string                `json:"extLocalDatetime" bson:"extLocalDatetime"`
	GrandTotal                     float64               `json:"grandTotal" bson:"grandTotal"`
	ProductId                      int64                 `json:"productId" bson:"productId"`
	ProductCode                    string                `json:"productCode" bson:"productCode"`
	ProductName                    string                `json:"productName" bson:"productName"`
	ProductData                    string                `json:"productData" bson:"productData"`
	RequestData                    string                `json:"requestData" bson:"requestData"`
	RequestOutData                 string                `json:"requestOutData" bson:"requestOutData"`
	OuId                           int64                 `json:"ouId" bson:"ouId"`
	OuName                         string                `json:"ouName" bson:"ouName"`
	OuCode                         string                `json:"ouCode" bson:"ouCode"`
	OuSubBranchId                  int64                 `json:"ouSubBranchId" bson:"ouSubBranchId"`
	OuSubBranchName                string                `json:"ouSubBranchName" bson:"ouSubBranchName"`
	OuSubBranchCode                string                `json:"ouSubBranchCode" bson:"ouSubBranchCode"`
	MainOuId                       int64                 `json:"mainOuId" bson:"mainOuId"`
	MainOuCode                     string                `json:"mainOuCode" bson:"mainOuCode"`
	MainOuName                     string                `json:"mainOuName" bson:"mainOuName"`
	CheckInTime                    int64                 `json:"checkInTime" bson:"checkInTime"`
	CheckOutTime                   int64                 `json:"checkOutTime" bson:"checkOutTime"`
	DurationTime                   int64                 `json:"durationTime" bson:"durationTime"`
	Price                          float64               `json:"price" bson:"price"`
	BaseTime                       int64                 `json:"baseTime" bson:"baseTime"`
	ProgressiveTime                int64                 `json:"progressiveTime" bson:"progressiveTime"`
	ProgressivePrice               float64               `json:"progressivePrice" bson:"progressivePrice"`
	IsPct                          string                `json:"isPct" bson:"isPct"`
	ProgressivePct                 int64                 `json:"progressivePct" bson:"progressivePct"`
	MaxPrice                       float64               `json:"maxPrice" bson:"maxPrice"`
	Is24H                          string                `json:"is24H" bson:"is24H"`
	OvernightTime                  string                `json:"overnightTime" bson:"overnightTime"`
	OvernightPrice                 float64               `json:"overnightPrice" bson:"overnightPrice"`
	GracePeriod                    int64                 `json:"gracePeriod" bson:"gracePeriod"`
	ServiceFee                     float64               `json:"serviceFee" bson:"serviceFee"`
	VehicleNumberIn                string                `json:"vehicleNumberIn" bson:"vehicleNumberIn"`
	VehicleNumberOut               string                `json:"vehicleNumberOut" bson:"vehicleNumberOut"`
	LogTrans                       string                `json:"logTrans" bson:"logTrans"`
	QrText                         string                `json:"qrText" bson:"qrText"`
	MerchantKey                    string                `json:"merchantKey" bson:"merchantKey"`
	TrxInvoiceItem                 []TrxInvoiceItem      `json:"trxInvoiceItem" bson:"trxInvoiceItem"`
	RequestAddTrxInvoiceDetailItem *TrxInvoiceDetailItem `json:"requestAddTrxInvoiceDetailItem" bson:"requestAddTrxInvoiceDetailItem"`
	ResponseSyncTrx                *string               `json:"responseSyncTrx" bson:"responseSyncTrx"`
	FlagErrorCallAPI               bool                  `json:"flagErrorCallAPI" bson:"flagErrorCallAPI"`
	ErrorMsg                       string                `json:"errorMsg" bson:"errorMsg"`
	CallbackPayment                *CallbackPayment      `json:"callbackPayment" bson:"callbackPayment"`
	LastUpdatedAt                  string                `json:"lastUpdatedAt" bson:"lastUpdatedAt"`
	MemberData                     *TrxMember            `json:"memberData" bson:"memberData"`
	RefDocNo                       string                `json:"refDocNo" bson:"refDocNo"`
	ShiftCode                      string                `json:"shiftCode" bson:"shiftCode"`
	Username                       string                `json:"username" bson:"username"`
	TrxAddInfo                     interface{}           `json:"trxAddInfo" bson:"trxAddInfo"`
	FlagTrxFromCloud               bool                  `json:"flagTrxFromCloud" bson:"flagTrxFromCloud"`
}

type TrxInvoiceItem struct {
	DocNo            string  `json:"docNo" bson:"docNo"`
	ProductId        int64   `json:"productId" bson:"productId"`
	ProductCode      string  `json:"productCode" bson:"productCode"`
	ProductName      string  `json:"productName" bson:"productName"`
	IsPctServiceFee  string  `json:"isPctServiceFee" bson:"isPctServiceFee"`
	ServiceFee       float64 `json:"serviceFee" bson:"serviceFee"`
	ServiceFeeMember float64 `json:"serviceFeeMember" bson:"serviceFeeMember"`
	Price            float64 `json:"price" bson:"price"`
	BaseTime         int64   `json:"baseTime" bson:"baseTime"`
	ProgressiveTime  int64   `json:"progressiveTime" bson:"progressiveTime"`
	ProgressivePrice float64 `json:"progressivePrice" bson:"progressivePrice"`
	IsPct            string  `json:"isPct" bson:"isPct"`
	ProgressivePct   int64   `json:"progressivePct" bson:"progressivePct"`
	MaxPrice         float64 `json:"maxPrice" bson:"maxPrice"`
	Is24H            string  `json:"is24H" bson:"is24H"`
	OvernightTime    string  `json:"overnightTime" bson:"overnightTime"`
	OvernightPrice   float64 `json:"overnightPrice" bson:"overnightPrice"`
	GracePeriod      int64   `json:"gracePeriod" bson:"gracePeriod"`
	TotalAmount      float64 `json:"totalAmount" bson:"totalAmount"`
}

type TrxInvoiceDetailItem struct {
	DocNo         string  `json:"docNo" bson:"docNo"`
	ProductCode   string  `json:"productCode" bson:"productCode"`
	InvoiceAmount float64 `json:"invoiceAmount" bson:"invoiceAmount"`
	CreatedAt     string  `json:"createdAt" bson:"createdAt"`
	CreatedDate   string  `json:"createdDate" bson:"createdDate"`
}

type CallbackPayment struct {
	Kodestatus             string  `json:"kodestatus"`
	Statuspayment          string  `json:"statuspayment"`
	Description            string  `json:"description"`
	Merchantnoref          string  `json:"merchantnoref"`
	Noheader               string  `json:"noheader"`
	Banknoref              string  `json:"banknoref"`
	Cardtype               string  `json:"cardtype"`
	Cardpan                string  `json:"cardpan"`
	Mid                    string  `json:"mid"`
	Tid                    string  `json:"tid"`
	Lastbalance            float64 `json:"lastbalance"`
	Currentbalance         float64 `json:"currentbalance"`
	Namakategoripembayaran string  `json:"namakategoripembayaran"`
	SettlementDatetime     string  `json:"settlement_datetime"`
	DeductDatetime         string  `json:"deduct_datetime"`
	Mdr                    float64 `json:"mdr"`
	Potongan               float64 `json:"potongan"`
	Kodepromo              string  `json:"kodepromo"`
	Promoissuer            string  `json:"promoissuer"`
	CreatedAt              string  `json:"created_at"`
}

type TrxMember struct {
	DocNo              string  `json:"docNo" bson:"docNo"`
	PartnerCode        string  `json:"partnerCode" bson:"partnerCode"`
	FirstName          string  `json:"firstName" bson:"firstName"`
	LastName           string  `json:"lastName" bson:"lastName"`
	RoleType           string  `json:"roleType" bson:"roleType"`
	PhoneNumber        string  `json:"phoneNumber" bson:"phoneNumber"`
	Email              string  `json:"email" bson:"email"`
	Active             string  `json:"active" bson:"active"`
	ActiveAt           string  `json:"activeAt" bson:"activeAt"`
	NonActiveAt        *string `json:"nonActiveAt" bson:"nonActiveAt"`
	OuId               int64   `json:"ouId" bson:"ouId"`
	TypePartner        string  `json:"typePartner" bson:"typePartner"`
	CardNumber         string  `json:"cardNumber" bson:"cardNumber"`
	VehicleNumber      string  `json:"vehicleNumber" bson:"vehicleNumber"`
	RegisteredDatetime string  `json:"registeredDatetime" bson:"registeredDatetime"`
	DateFrom           string  `json:"dateFrom" bson:"dateFrom"`
	DateTo             string  `json:"dateTo" bson:"dateTo"`
	ProductId          int64   `json:"productId" bson:"productId"`
	ProductCode        string  `json:"productCode" bson:"productCode"`
}

type TrxForUpdateInvoice struct {
	DocNo            string  `json:"docNo"`
	ProductCode      string  `json:"productCode"`
	ProgressivePrice float64 `json:"progressivePrice"`
	ProgressiveType  string  `json:"progressiveType"`
	Datetime         string  `json:"datetime"`
}
