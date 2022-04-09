package app

import (
	"consumerskripsi/repositories"
	"consumerskripsi/repositories/generalRepository"
	"consumerskripsi/repositories/productRepository"
	"consumerskripsi/repositories/trxRepository"
	"consumerskripsi/services"
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/mkpproduction/mkp-sdk-go/mkp/genautonum"
	"github.com/streadway/amqp"
)

func SetupApp(DB *sql.DB, repo repositories.Repository, redisLocal *redis.Client, rabbitMQ *amqp.Channel, repoGenAutoNum genautonum.Repository) services.UsecaseService {

	trxRepo := trxRepository.NewTrxRepository(repo)
	trxMongoRepo := trxRepository.NewTrxMongoRepository(repo)
	whitelistRepo := generalRepository.NewWhitelistRouteRepository(repo)
	deviceRepo := generalRepository.NewDeviceRepository(repo)
	paymentMethodRepo := generalRepository.NewPaymentMethodRepository(repo)
	productMongoRepo := productRepository.NewProductMongoRepository(repo)
	genAutoNumRepo := genautonum.NewGenerateAutonumberRepository(repoGenAutoNum)

	// Servicesc
	usecaseSvc := services.NewUsecaseService(DB, trxRepo, trxMongoRepo, whitelistRepo, deviceRepo, paymentMethodRepo,
		productMongoRepo, redisLocal, rabbitMQ, genAutoNumRepo)

	return usecaseSvc
}
