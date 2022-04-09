package generalRepository

import (
	"consumerskripsi/constans"
	"consumerskripsi/models"
	"consumerskripsi/repositories"

	"gopkg.in/mgo.v2/bson"
)

type whitelistRouteRepository struct {
	RepoDB repositories.Repository
}

func NewWhitelistRouteRepository(repoDB repositories.Repository) whitelistRouteRepository {
	return whitelistRouteRepository{
		RepoDB: repoDB,
	}
}

func (ctx whitelistRouteRepository) IsWhitelistRouteByCode(ouCode string) (models.WhitelistRoute, bool) {
	var result models.WhitelistRoute

	filter := bson.M{
		"ouCode": ouCode,
	}

	collect := ctx.RepoDB.MongoDB.Collection(constans.WHITELIST_ROUTE)
	data := collect.FindOne(ctx.RepoDB.Context, filter)
	if data.Err() != nil {
		return result, false
	}

	err := data.Decode(&result)
	if err != nil {
		return result, false
	}

	return result, true
}
