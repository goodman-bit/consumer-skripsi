package helpers

import (
	"bytes"
	"consumerskripsi/constans"
	"consumerskripsi/models"
	"consumerskripsi/utils"
	. "consumerskripsi/utils"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func InArray(v interface{}, in interface{}) (ok bool, i int) {
	val := reflect.Indirect(reflect.ValueOf(in))
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for ; i < val.Len(); i++ {
			if ok = v == val.Index(i).Interface(); ok {
				return
			}
		}
	}
	return
}

func BindValidateStruct(ctx echo.Context, i interface{}) error {
	if err := ctx.Bind(i); err != nil {
		return err
	}

	if err := ctx.Validate(i); err != nil {
		return err
	}
	return nil
}

func ResponseJSON(success bool, code string, msg string, result interface{}) models.Response {
	tm := time.Now()
	response := models.Response{
		Success:          success,
		StatusCode:       code,
		Result:           result,
		Message:          msg,
		ResponseDatetime: tm,
	}
	fmt.Println("Output Response:", ToString(response))

	return response
}

// ReplaceSQL ...
func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

func WorkerExtCloud(req interface{}, callbackURL string) (models.WorkerResponse, error) {
	var result models.WorkerResponse
	result.SnapshotData = req
	result.URL = callbackURL

	bodyReq, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Err workerAddMemberExtCloud - json.Marshal : ", err.Error())
		result.StatusCode = constans.EMPTY_VALUE
		result.StatusDesc = err.Error()
		return result, err
	}

	bodyReqToStr := string(bodyReq)
	fmt.Println(bodyReqToStr)

	httpReq, err := http.NewRequest("POST", callbackURL, bytes.NewBuffer(bodyReq))
	if err != nil {
		fmt.Println("Err workerAddMemberExtCloud - http.NewRequest ", callbackURL, " :", err.Error())
		result.StatusCode = constans.EMPTY_VALUE
		result.StatusDesc = err.Error()
		return result, err
	}
	defer httpReq.Body.Close()

	httpReq.Close = true
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Set("Connection", "close")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	//PRINTOUT BODY REQ
	fmt.Println("URL :>", callbackURL)
	fmt.Println("Request :>", bodyReqToStr)

	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println("Err workerTrxCallbackURL - client.Do :", err.Error())
		result.StatusCode = constans.EMPTY_VALUE
		result.StatusDesc = err.Error()
		return result, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	result.StatusCode = strconv.Itoa(resp.StatusCode)
	result.StatusDesc = resp.Status
	result.ResponseBody = bodyString

	fmt.Println("ResponseCallback:", utils.ToString(result))

	return result, nil
}

func ConditionalProgressive(startCondition string, checkInDatetime string, data models.TrxInvoiceItem) (
	string, string, bool, int64, float64) {

	schedulingStopped := false
	scheduleType := constans.EMPTY_VALUE
	delayMinutes := int64(0)
	progressivePrice := float64(0)
	progressiveType := constans.EMPTY_VALUE

	switch expr := startCondition; expr {
	case "START":
		scheduleType = "GP"
		delayMinutes = data.GracePeriod * constans.MILLISECOND
		progressiveType = constans.PROGRESSIVE_TYPE
		if data.GracePeriod == 0 {
			scheduleType = "BT"
			delayMinutes = data.BaseTime * constans.MILLISECOND
			progressivePrice = data.Price

			if data.BaseTime == 0 {
				scheduleType = "PT"
				delayMinutes = data.ProgressiveTime * constans.MILLISECOND

			}
			schedulingStopped = data.ProgressiveTime == 0
		}
	case "GP":
		scheduleType = "BT"
		delayMinutes = data.BaseTime * constans.MILLISECOND
		progressivePrice = data.Price
		progressiveType = constans.PROGRESSIVE_TYPE

		if data.BaseTime == 0 {
			scheduleType = "PT"
			delayMinutes = data.ProgressiveTime * constans.MILLISECOND
		}

		schedulingStopped = data.ProgressiveTime == 0
	case "BT":
		scheduleType = "PT"
		delayMinutes = data.ProgressiveTime * constans.MILLISECOND
	case "PT":
		scheduleType = "PT"
		delayMinutes = data.ProgressiveTime * constans.MILLISECOND
		progressivePrice = data.ProgressivePrice
		progressiveType = constans.PROGRESSIVE_TYPE
	}

	return scheduleType, progressiveType, schedulingStopped, delayMinutes, progressivePrice
}
