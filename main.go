package main

import (
	"final-project/configs"
	"final-project/database"
	"final-project/routers"
)

func main() {
	config := configs.LoadEnv()
	database.StartDB(&config)
	routers.StartServer().Run(config.Port)
}
