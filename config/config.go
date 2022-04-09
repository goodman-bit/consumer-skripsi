package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DBDriver = GetEnv("DB_DRIVER", "postgres")
	DBName   = GetEnv("DB_NAME", "local")
	DBHost   = GetEnv("DB_HOST", "localhost")
	DBPort   = GetEnv("DB_PORT", "5432")
	DBUser   = GetEnv("DB_USER", "root")
	DBPass   = GetEnv("DB_PASS", "")
	SSLMode  = GetEnv("SSL_MODE", "disable")

	SALT_KEY = GetEnv("SALT_KEY")

	REDISHost = GetEnv("REDIS_HOST")
	REDISPass = GetEnv("REDIS_PASS")
	REDISPort = GetEnv("REDIS_PORT")

	REDISHostLocal = GetEnv("REDIS_HOST_LOCAL")
	REDISPassLocal = GetEnv("REDIS_PASS_LOCAL")
	REDISPortLocal = GetEnv("REDIS_PORT_LOCAL")

	MONGOHost = GetEnv("MONGO_HOST")
	MONGOPort = GetEnv("MONGO_PORT")
	MONGODB   = GetEnv("MONGO_DB")

	QUEUEName     = GetEnv("QUEUE_NAME")
	AMQPServerUrl = GetEnv("AMQP_SERVER_URL")

	MERCHANT_KEY          = GetEnv("MERCHANT_KEY")
	DEFAULT_CHANNEL_REDIS = GetEnv("DEFAULT_CHANNEL_REDIS")
	URL_APPS2PAY_DEV      = GetEnv("URL_APPS2PAY_DEV")
	CHANNELS_TRX_REDIS    = GetEnv("CHANNELS_TRX_REDIS")
)

func GetEnv(key string, value ...string) string {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error Load file .env not found")
	}

	if os.Getenv(key) != "" {
		return os.Getenv(key)
	} else {
		if len(value) > 0 {
			return value[0]
		}
		return ""
	}
}
