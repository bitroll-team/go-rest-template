package controller

import "database/sql"

type Controller struct {
	DB *sql.DB
}

func SetupController(db *sql.DB) Controller {
	var ctrl Controller
	ctrl.DB = db
	return ctrl
}
