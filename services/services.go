package services

import (
	"consumerskripsi/repositories"
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/mkpproduction/mkp-sdk-go/mkp/genautonum"
	"github.com/streadway/amqp"
)

type UsecaseService struct {
	RepoDB            *sql.DB
	TrxRepo           repositories.TrxRepository
	TrxMongoRepo      repositories.TrxMongoRepository
	WhitelistRepo     repositories.WhitelistRouteRepository
	DeviceRepo        repositories.DeviceRepository
	PaymentMethodRepo repositories.PaymentMethodRepository
	ProductMongoRepo  repositories.ProductMongoRepository
	RedisClientLocal  *redis.Client
	RabbitMQ          *amqp.Channel
	GenAutoNumRepo    genautonum.GenerateAutonumberRepository
}

func NewUsecaseService(
	RepoDB *sql.DB,
	TrxRepo repositories.TrxRepository,
	TrxMongoRepo repositories.TrxMongoRepository,
	WhitelistRepo repositories.WhitelistRouteRepository,
	DeviceRepo repositories.DeviceRepository,
	PaymentMethodRepo repositories.PaymentMethodRepository,
	ProductMongoRepo repositories.ProductMongoRepository,
	RedisClientLocal *redis.Client,
	rabbitMQ *amqp.Channel,
	GenAutoNumRepo genautonum.GenerateAutonumberRepository,

) UsecaseService {
	return UsecaseService{
		RepoDB:            RepoDB,
		TrxRepo:           TrxRepo,
		TrxMongoRepo:      TrxMongoRepo,
		WhitelistRepo:     WhitelistRepo,
		DeviceRepo:        DeviceRepo,
		PaymentMethodRepo: PaymentMethodRepo,
		ProductMongoRepo:  ProductMongoRepo,
		RedisClientLocal:  RedisClientLocal,
		RabbitMQ:          rabbitMQ,
		GenAutoNumRepo:    GenAutoNumRepo,
	}
}
