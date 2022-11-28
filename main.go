package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"

	"github.com/dime-makarov/DmPillsReminder/dataaccess"
)

func main() {

	cfg := mysql.Config{
		User:   DatabaseUser,
		Passwd: DatabasePassword,
		Net:    "tcp",
		Addr:   DatabaseAddress,
		DBName: DatabaseName,
	}

	err := dataaccess.InitializeDB(cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/prescriptions", GetPrescriptionsApi)

	router.Run("localhost:8080")
}

func GetPrescriptionsApi(c *gin.Context) {
	prescriptions, err := dataaccess.GetPrescriptions()

	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, prescriptions)
}
