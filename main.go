package main

import (
	"net/http"

	"github.com/surick/go-exercises/controller"
	"github.com/surick/go-exercises/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup Controller
	controller.Startup()

	http.ListenAndServe(":8888", nil)
}