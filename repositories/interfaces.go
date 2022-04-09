package repositories

import (
	"consumerskripsi/models"
)

type TrxRepository interface {
}

type TrxMongoRepository interface {
	AddTrxCheckIn(trxCheckIn models.Trx) error
	AddTrxInvoiceItems(trxInvoice models.TrxInvoiceItem) error
	IsTrxOutstandingByDocNo(docNo string) (*models.Trx, bool)
	IsTrxInvoiceItemsByIndex(docNo, productCode string) (*models.TrxInvoiceItem, bool)
	UpdateTrxInvoiceItemByIndex(trxUpdatedInvoice models.TrxForUpdateInvoice) error
}

type ProductMongoRepository interface {
	FindProductByProductCode(productCode string, ouId int64) (models.PolicyOuProductWithRules, error)
	AddPolicyOuProduct(productId int64, ouId int64, product models.PolicyOuProductWithRules) error
	AddProduct(product models.PolicyOuProductWithRules) error
}

type ProductRepository interface {
	FindProductById(idProduct string) (models.Product, error)
}

type WhitelistRouteRepository interface {
	IsWhitelistRouteByCode(ouCode string) (models.WhitelistRoute, bool)
}

type DeviceRepository interface {
	FindDeviceByDeviceId(deviceId string) (models.Device, error)
}

type PaymentMethodRepository interface {
	FindPaymentMethodIDByName(paymentMethodName string) (models.PaymentMethod, error)
	FindVendorPaymentIDByName(vendorName string) (models.PaymentVendor, error)
}
