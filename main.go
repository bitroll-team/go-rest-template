package main

import (
	"com/example/controller"
	"com/example/database"
	"com/example/router"
	"fmt"
	"log"
)

func main() {
	log.Println("Starting...")
	config := ReadEnvVars()

	// setup database

	db, err := database.DBConnect(config.DBConnStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.RunMigrations(db, MigrationsFS, MigrationsPath); err != nil {
		log.Fatal(err)
	}
	log.Println("Migrations succeded")

	// controller

	conn := controller.SetupController(db)

	// router

	router := router.SetupRouter(&conn)
	err = router.Run(fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		log.Fatal(err)
	}
}
