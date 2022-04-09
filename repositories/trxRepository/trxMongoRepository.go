package trxRepository

import (
	"consumerskripsi/constans"
	"consumerskripsi/models"
	"consumerskripsi/repositories"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type trxMongoRepository struct {
	RepoDB repositories.Repository
}

func NewTrxMongoRepository(repoDB repositories.Repository) trxMongoRepository {
	return trxMongoRepository{
		RepoDB: repoDB,
	}
}

func (ctx trxMongoRepository) AddTrxCheckIn(trxCheckIn models.Trx) error {
	_, err := ctx.RepoDB.MongoDB.Collection(constans.TRX_OUTSTANDING_COLLECTION).InsertOne(ctx.RepoDB.Context, trxCheckIn)
	if err != nil {
		return err
	}
	return nil
}

func (ctx trxMongoRepository) AddTrxInvoiceItems(trxInvoice models.TrxInvoiceItem) error {
	_, err := ctx.RepoDB.MongoDB.Collection(constans.TRX_INVOICE_ITEM_COLLECTION).
		InsertOne(ctx.RepoDB.Context, trxInvoice)
	if err != nil {
		return err
	}

	return nil
}

func (ctx trxMongoRepository) IsTrxOutstandingByDocNo(docNo string) (*models.Trx, bool) {
	var result models.Trx
	filter := bson.M{
		"docNo": docNo,
	}
	err := ctx.RepoDB.MongoDB.Collection(constans.TRX_OUTSTANDING_COLLECTION).
		FindOne(ctx.RepoDB.Context, filter).Decode(&result)
	if err != nil {
		return nil, false
	}

	return &result, true
}

func (ctx trxMongoRepository) IsTrxInvoiceItemsByIndex(docNo, productCode string) (*models.TrxInvoiceItem, bool) {
	var result models.TrxInvoiceItem
	filter := bson.M{
		"docNo":       docNo,
		"productCode": productCode,
	}

	err := ctx.RepoDB.MongoDB.Collection(constans.TRX_INVOICE_ITEM_COLLECTION).
		FindOne(ctx.RepoDB.Context, filter).Decode(&result)
	if err != nil {
		return nil, false
	}

	return &result, true
}

func (ctx trxMongoRepository) UpdateTrxInvoiceItemByIndex(trxUpdatedInvoice models.TrxForUpdateInvoice) error {
	var update interface{}
	var updateTrxInvoiceOutstanding interface{}

	filter := bson.M{
		"docNo":       trxUpdatedInvoice.DocNo,
		"productCode": trxUpdatedInvoice.ProductCode,
	}

	if trxUpdatedInvoice.ProgressiveType == constans.PROGRESSIVE_TYPE {
		update = bson.M{
			"$inc": bson.M{"totalAmount": trxUpdatedInvoice.ProgressivePrice},
		}

		updateTrxInvoiceOutstanding = bson.M{
			"$inc": bson.M{
				"trxInvoiceItem.$[elemX].totalAmount": trxUpdatedInvoice.ProgressivePrice,
			},
		}

	} else if trxUpdatedInvoice.ProgressiveType == constans.MAX_PROGRESSIVE_TYPE {
		update = bson.M{
			"$set": bson.M{"totalAmount": trxUpdatedInvoice.ProgressivePrice},
		}

		updateTrxInvoiceOutstanding = bson.M{
			"$set": bson.M{
				"trxInvoiceItem.$[elemX].totalAmount": trxUpdatedInvoice.ProgressivePrice,
			},
		}

	}

	upsert := true
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := ctx.RepoDB.MongoDB.Collection(constans.TRX_OUTSTANDING_COLLECTION).UpdateOne(ctx.RepoDB.Context, bson.M{
		"docNo": trxUpdatedInvoice.DocNo,
	}, bson.M{
		"$set": bson.M{"lastUpdatedAt": trxUpdatedInvoice.Datetime},
	}, &opt)
	if err != nil {
		return err
	}

	_, err = ctx.RepoDB.MongoDB.Collection(constans.TRX_INVOICE_ITEM_COLLECTION).UpdateOne(ctx.RepoDB.Context, filter, update, &opt)
	if err != nil {
		return err
	}

	filterForOutstanding := bson.M{
		"elemX.docNo":       trxUpdatedInvoice.DocNo,
		"elemX.productCode": trxUpdatedInvoice.ProductCode,
	}

	_, err = ctx.RepoDB.MongoDB.Collection(constans.TRX_OUTSTANDING_COLLECTION).
		UpdateOne(ctx.RepoDB.Context, bson.M{"docNo": trxUpdatedInvoice.DocNo}, updateTrxInvoiceOutstanding, options.Update().SetArrayFilters(
			options.ArrayFilters{Filters: []interface{}{
				filterForOutstanding,
			},
			},
		))

	if err != nil {
		return err
	}

	return nil
}
