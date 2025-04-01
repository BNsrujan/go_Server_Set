package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	Port string
}

func (e *envConfig) LoadConfig() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("ENV file not found")
	}

	// _, ok := os.LookupEnv("PORT")

	// if !ok {
	// 	log.Panic("PORT is not loaded")
	// }

	e.Port = loadString("PORT", "8000")

}

var Config envConfig

func init() {
	Config.LoadConfig()
}

func loadString(value string, fallback string) string {

	val, ok := os.LookupEnv(value)

	if !ok {
		log.Panic("App_PORT is not loaded")
		return fallback
	}
	return val
}
