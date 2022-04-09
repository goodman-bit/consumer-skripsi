package generalRepository

import (
	"consumerskripsi/models"
	"consumerskripsi/repositories"
)

type paymentMethodRepository struct {
	RepoDB repositories.Repository
}

func NewPaymentMethodRepository(repoDB repositories.Repository) paymentMethodRepository {
	return paymentMethodRepository{
		RepoDB: repoDB,
	}
}

func (ctx paymentMethodRepository) FindPaymentMethodIDByName(paymentMethodName string) (models.PaymentMethod, error) {
	var result models.PaymentMethod

	query := `SELECT id, uraian as namametodepembayaran 
	FROM metodepembayaran
	WHERE uraian ILIKE $1 `

	err := ctx.RepoDB.DB.QueryRow(query, "%"+paymentMethodName+"%").
		Scan(&result.Id, &result.PaymentName)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (ctx paymentMethodRepository) FindVendorPaymentIDByName(vendorName string) (models.PaymentVendor, error) {
	var result models.PaymentVendor
	query := `SELECT id, uraian as namavendorpembayaran
	FROM vendorpembayaran
	WHERE uraian ILIKE $1`

	err := ctx.RepoDB.DB.QueryRow(query, "%"+vendorName+"%").
		Scan(&result.Id, &result.VendorName)
	if err != nil {
		return result, err
	}

	return result, nil
}
