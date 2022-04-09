package utils

import (
	"database/sql"
	"time"

	"github.com/labstack/echo"
)

func DateFormatToTime(v string) time.Time {
	time, _ := time.Parse("2006-01-02", v)
	return time
}

func TimeStampFormatToTime(v string) time.Time {
	time, _ := time.Parse("2006-01-02 15:04:05", v)
	return time
}

func GolangFormatToTime(v string) time.Time {
	time, _ := time.Parse("2006-01-02T15:04:05Z", v)
	return time
}

func TimestampNow() string {
	return time.Now().Format("2006-01-02 15:04:05")

}

func DateNow() string {
	return time.Now().Format("2006-01-02")
}

func GetUsernameToken(ctx echo.Context) string {
	// user := ctx.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)

	// return claims["username"].(string)
	return "DASHBOARD MKP"
}

func DBTransaction(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Rollback Panic
		} else if err != nil {
			tx.Rollback() // err is not nill
		} else {
			err = tx.Commit() // err is nil
		}
	}()
	err = txFunc(tx)
	return err
}
