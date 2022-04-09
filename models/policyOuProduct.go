package models

//-- FIND POLICY OU IN DYNAMO DB--

type PolicyOuProductWithRules struct {
	OuId        int64   `json:"ouId" bson:"ouId,omitempty"`
	OuCode      string  `json:"ouCode" bson:"ouCode,omitempty"`
	OuName      string  `json:"ouName" bson:"ouName,omitempty"`
	ProductId   int64   `json:"productId" bson:"productId,omitempty"`
	ProductCode string  `json:"productCode" bson:"productCode,omitempty"`
	ProductName string  `json:"productName" bson:"productName,omitempty"`
	ServiceFee  float64 `json:"serviceFee" bson:"serviceFee,omitempty"`
	Price       float64 `json:"price" bson:"price,omitempty"`
}
