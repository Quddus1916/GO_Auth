package config

import (
	"github.com/joho/godotenv"
	"jwt-auth.com/types"
	"os"
)

func Initconfig() types.Config {
	godotenv.Load()
	var config types.Config
	config.Port = os.Getenv("PORT")
	config.MongoUrl = os.Getenv("MONGOURL")

	return config

}

func Getconfig() types.Config {
	return Initconfig()
}
