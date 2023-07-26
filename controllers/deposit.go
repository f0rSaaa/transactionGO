package controllers

import (
	"log"
	"net/http"

	"github.com/beego/beego/orm"
	"github.com/f0rSaaaa/transactions/models"
	"github.com/gin-gonic/gin"
)

func InsertSavingsData(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	var deposit models.DepositSavings
	c.BindJSON(&deposit)

	_, err := o.Raw("insert into savings_account(user_id, savings) values(?,?);", deposit.UserId, deposit.Savings).Exec()
	if err != nil {
		log.Println("Cannot insert in the table")
		log.Println(err)
		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Data cannot be inserted in database",
		})
		return
	}

	// fmt.Println(reflect.TypeOf(deposit.Savings))
	// fmt.Println(deposit.Savings)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Data successfully received",
		"your_id": deposit.UserId,
	})
}

func InsertCheckinData(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	var deposit models.DepositCheckin
	c.BindJSON(&deposit)

	_, err := o.Raw("insert into checkin_account(user_id, checkin) values(?,?);", deposit.UserId, deposit.Checkin).Exec()
	if err != nil {
		log.Println("Cannot insert in the table")
		log.Println(err)
		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Data cannot be inserted in database",
		})
		return
	}

	// fmt.Println(reflect.TypeOf(deposit.Savings))
	// fmt.Println(deposit.Savings)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Data successfully received",
		"your_id": deposit.UserId,
	})
}
