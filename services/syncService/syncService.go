package syncService

import (
	"consumerskripsi/constans"
	"consumerskripsi/helpers"
	"consumerskripsi/models"
	"consumerskripsi/services"
	"consumerskripsi/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strings"
)

type syncService struct {
	Service services.UsecaseService
}

func NewSyncService(service services.UsecaseService) syncService {
	return syncService{
		Service: service,
	}
}

func (svc syncService) AddTrxCheckIn(body string) error {
	fmt.Println("Start AddTrxCheckIn")
	log.Println("Request AddTrxCheckIn: ", utils.ToString(body))
	schedulingType := "START"

	bodySplit := strings.Split(body, "#")

	var trx models.Trx
	err := json.Unmarshal([]byte(bodySplit[0]), &trx)
	if err != nil {
		log.Println("Error Unmarsh : ", err)
		return err
	}

	err = svc.Service.TrxMongoRepo.AddTrxCheckIn(trx)
	if err != nil {
		log.Println("Error AddTrxCheckIn : ", err)
		return err
	}

	for _, rows := range trx.TrxInvoiceItem {

		err := svc.Service.TrxMongoRepo.AddTrxInvoiceItems(rows)
		if err != nil {
			return err
		}

		bodyRequestCalculate := fmt.Sprintf("%s%s%s%s%s", rows.DocNo, "#", schedulingType, "#", rows.ProductCode)
		go svc.SchedulingTrx(bodyRequestCalculate)

	}

	return nil
}

func (svc syncService) SchedulingTrx(body string) error {
	fmt.Println("Start SchedulingTrx")
	log.Println("Request SchedulingTrx: ", utils.ToString(body))

	pipes := strings.Split(body, "#")
	datetimeNow := utils.TimestampNow()
	resultTrx, exists := svc.Service.TrxMongoRepo.IsTrxOutstandingByDocNo(pipes[0])
	if !exists {
		log.Println("session already clear")
		return errors.New("session already clear")
	}

	resultTrxInvoice, exists := svc.Service.TrxMongoRepo.IsTrxInvoiceItemsByIndex(pipes[0], pipes[2])
	if !exists {
		log.Println("product code not exists")
		return errors.New("product code not exists")
	}

	scheduleType, progressiveType, schedulingStopped, delayMinutes, progressivePrice := helpers.ConditionalProgressive(pipes[1], resultTrx.CheckInDatetime, *resultTrxInvoice)
	if scheduleType != "OVER" {
		if resultTrxInvoice.MaxPrice > 0 && (resultTrxInvoice.TotalAmount+progressivePrice) >= resultTrxInvoice.MaxPrice {
			progressivePrice = resultTrxInvoice.MaxPrice
			progressiveType = constans.MAX_PROGRESSIVE_TYPE
			schedulingStopped = true
		}
	}

	if !schedulingStopped {
		log.Println("Scheduling has stopped:", schedulingStopped, "Scheduling Type:", scheduleType, "Delayed Messages:", delayMinutes)
		msg := amqp.Publishing{
			ContentType: "text/plan",
			Body:        []byte(fmt.Sprintf("%s%s%s%s%s", pipes[0], "#", scheduleType, "#", pipes[2])),
			Headers: amqp.Table{
				"x-delay": delayMinutes,
			},
			DeliveryMode: 2,
		}

		err := svc.Service.RabbitMQ.Publish(constans.QUEUE_PG_SCHEDULING_PROGRESSIVE, constans.QUEUE_PROGRESSIVE, false, false, msg)
		if err != nil {
			log.Println("Error: ", err.Error())
			return err
		}
	}

	if progressiveType != constans.EMPTY_VALUE {
		trxUpdate := models.TrxForUpdateInvoice{
			DocNo:            resultTrxInvoice.DocNo,
			ProductCode:      resultTrxInvoice.ProductCode,
			ProgressivePrice: progressivePrice,
			ProgressiveType:  progressiveType,
			Datetime:         datetimeNow,
		}

		err := svc.Service.TrxMongoRepo.UpdateTrxInvoiceItemByIndex(trxUpdate)
		if err != nil {
			log.Println("Error: ", err.Error())
			return err
		}
	}

	return nil

}
