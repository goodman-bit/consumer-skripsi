package main

import (
	"consumerskripsi/app"
	"consumerskripsi/config"
	"consumerskripsi/constans"
	"consumerskripsi/helpers"
	"consumerskripsi/repositories"
	userRpc "consumerskripsi/repositories/testingRpc"
	"consumerskripsi/services/syncService"
	"consumerskripsi/services/testingRpcService"
	"consumerskripsi/utils"
	"errors"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"strconv"

	"context"
	"fmt"
	"log"

	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mkpproduction/mkp-sdk-go/mkp/genautonum"
	"github.com/streadway/amqp"
	"gopkg.in/go-playground/validator.v9"
)

//CustomValidator adalah
type CustomValidator struct {
	validator  *validator.Validate
	translator ut.Translator
}

var (
	echoHandler echo.Echo
	ctx         = context.Background()
)

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, row := range errs {
			return errors.New(row.Translate(cv.translator))
		}
	}

	return cv.validator.Struct(i)
}

func main() {

	echoHandler.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	echoHandler.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		result := helpers.ResponseJSON(false, strconv.Itoa(report.Code), err.Error(), nil)
		c.Logger().Error(report)
		c.JSON(report.Code, result)
	}

	if err := config.OpenConnection(); err != nil {
		panic(fmt.Sprintf("Open Connection Faild: %s", err.Error()))
	}
	//	defer config.CloseConnectionDB()
	// Mongo DB connection using database default
	mongoDB := config.ConnectMongo(ctx)

	// Connection database

	rdsConnectionLocal := config.ConnectRedisLocal()

	//Connections Postgresql
	DB := config.DBConnection()

	//RepoAutoNumber
	repoGenAutoNum := genautonum.NewRepository(DB, ctx, mongoDB)

	// Configuration Rabbit MQ
	connectionRabbitMQ, err := amqp.Dial(config.AMQPServerUrl)
	if err != nil {
		panic(err.Error())
	}
	defer connectionRabbitMQ.Close()

	channelRabbitMQ, err := connectionRabbitMQ.Channel()
	if err != nil {
		panic(err.Error())
	}
	defer channelRabbitMQ.Close()
	fmt.Println("Successful connected to RabbitMQ")

	// Configuration Repository
	repo := repositories.NewRepository(DB, mongoDB, ctx)

	// Configuration Repository and Services
	service := app.SetupApp(DB, repo, rdsConnectionLocal, channelRabbitMQ, repoGenAutoNum)

	svcSync := syncService.NewSyncService(service)

	listen, err := net.Listen("tcp", fmt.Sprintf("%s%s", ":", config.GetEnv("APP_PORT", "6000")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	testingService := testingRpcService.NewTestingRpcService(service)
	userRpc.RegisterUserServiceServer(grpcServer, &testingService)

	var queueName string
	if len(os.Args) == 1 {
		queueName = config.QUEUEName
	} else {
		queueName = os.Args[1]
	}

	rabbitMQMsgList := []string{"SCHEDULING_PROJECT"}
	condition, _ := helpers.InArray(queueName, rabbitMQMsgList)
	if condition {
		log.Println("Consumer Queue Name Running:", queueName)

		args := amqp.Table{
			"x-queue-type": "classic",
		}

		q, err := channelRabbitMQ.QueueDeclare(
			queueName, true, false, false, false, args)
		failOnError(err, "Failed to declare a queue")

		msg, err := channelRabbitMQ.Consume(
			q.Name, "", true, false, false, false, nil)
		failOnError(err, "Failed to register a consumer")

		forever := make(chan bool)

		go func() {
			for d := range msg {
				switch expr := q.Name; expr {
				case "SCHEDULING_PROJECT":
					response := svcSync.SchedulingTrx(string(d.Body))
					utils.EndLogPrint(fmt.Sprintf("%s:%s %s", "CalculateProgressive", q.Name, response))
				default:
					fmt.Printf("%s", q.Name)
				}
			}
		}()

		<-forever
	}

	if queueName == constans.PG_PARKING_CHECKIN {
		fmt.Println("Consumer Queue Name Running:", config.DEFAULT_CHANNEL_REDIS)

		fmt.Println()
		subClient := rdsConnectionLocal.Subscribe(config.DEFAULT_CHANNEL_REDIS)
		ch := subClient.Channel()

		for msg := range ch {
			log.Println("Channel:", msg.Channel, "Payload:", msg.Payload)

			switch channelName := msg.Channel; channelName {
			case config.DEFAULT_CHANNEL_REDIS:
				response := svcSync.AddTrxCheckIn(msg.Payload)
				log.Println("Response SchedulingTrx:", response)
			}
		}
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func init() {

	e := echo.New()
	echoHandler = *e
	validateCustom := validator.New()

	e.Validator = &CustomValidator{validator: validateCustom}

	e.Static("/img/*", "assets/img")
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	echoHandler.Use(middleware.Logger())
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		return
	}
}
