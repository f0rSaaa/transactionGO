package main

import (
	"fmt"
	"log"

	"github.com/f0rSaaaa/transactions/controllers"
	"github.com/f0rSaaaa/transactions/initializers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}

func main() {
	fmt.Println("Transactions with Golang using Beego and Gin")

	router := gin.Default()

	router.POST("/savings", controllers.InsertSavingsData)
	router.POST("/checkin", controllers.InsertCheckinData)
	router.POST("/transfertocheckin/:u_id", controllers.TransferToCheckin)
	router.POST("/transfertosavings/:u_id", controllers.TransferToSavings)

	fmt.Println("Server running on 8080")
	//If this fails it will log the report and then exit out of the program
	log.Fatal(router.Run(), nil)

}
