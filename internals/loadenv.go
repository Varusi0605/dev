package internals

import (
	"shopping-site/pkg/loggers"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		loggers.FatalLog.Fatal(err)
	}
}
