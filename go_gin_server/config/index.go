package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	AppPort string
}

// type
func (e *envConfig) LoadConfig() {
	// we can not call the function directly we have to make the variable for it
	err := godotenv.Load(".env")

	if err != nil {
		// it err is not null there is an err
		// Panic is the highest type of error
		log.Panic("ENV file not loaded")
	}

	//before runing this you will load the file first and the lookupenv
	// val, ok := os.LookupEnv("App_PORT")

	// if !ok {
	// 	log.Panic("App_PORT is not loaded")
	// }

	e.AppPort = loadString("App_PORT", ":8080")
	// e.DB_PATH = loadString("DB_PATH","https://192.185/u")
}

// make shore that it is capitalized because if it's not we can not access this in different place in the folder
var Config envConfig

// the function name should be the same ,this function  will fun before anything run
func init() {
	Config.LoadConfig()
}

func loadString(val string, fallback string) string {
	// value and the boolean function
	val, ok := os.LookupEnv(val)

	if !ok {
		log.Panic("App_PORT is not loaded")
		return fallback
	}

	return val
}
