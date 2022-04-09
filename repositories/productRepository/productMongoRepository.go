package productRepository

import (
	"consumerskripsi/constans"
	"consumerskripsi/models"
	"consumerskripsi/repositories"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type productMongoRepository struct {
	RepoDB repositories.Repository
}

func NewProductMongoRepository(repo repositories.Repository) productMongoRepository {
	return productMongoRepository{
		RepoDB: repo,
	}
}

func (ctx productMongoRepository) FindProductByProductCode(productCode string, ouId int64) (models.PolicyOuProductWithRules, error) {
	var result models.PolicyOuProductWithRules

	filter := bson.M{
		"ouId":        ouId,
		"productCode": productCode,
	}

	data := ctx.RepoDB.MongoDB.Collection(constans.TABLE_POLICY_OU_PRODUCT).FindOne(ctx.RepoDB.Context, filter)
	if data.Err() != nil {
		fmt.Println(data.Err())
		return result, errors.New("Policy Ou Product Not Found")
	}

	err := data.Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (ctx productMongoRepository) AddPolicyOuProduct(productId int64, ouId int64, product models.PolicyOuProductWithRules) error {
	filter := bson.M{
		"ouId":      ouId,
		"productId": productId,
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndReplaceOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	collect := ctx.RepoDB.MongoDB.Collection(constans.TABLE_POLICY_OU_PRODUCT)
	result := collect.FindOneAndReplace(ctx.RepoDB.Context, filter, product, &opt)
	if result.Err() != nil {
		return result.Err()
	}

	return nil

}

func (ctx productMongoRepository) AddProduct(product models.PolicyOuProductWithRules) error {
	_, err := ctx.RepoDB.MongoDB.Collection(constans.TABLE_POLICY_OU_PRODUCT).
		InsertOne(ctx.RepoDB.Context, product)

	if err != nil {
		return err
	}
	return nil
}
