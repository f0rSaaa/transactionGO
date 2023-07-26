package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/beego/beego/orm"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type SavingsAccount struct {
	Id      int    `orm:"auto"`
	UserId  string `orm:"unique"`
	Savings int
}

type CheckinAccount struct {
	Id      int    `orm:"auto"`
	UserId  string `orm:"unique"`
	Checkin int
}

type DepositSavings struct {
	UserId  string
	Savings int
}

type DepositCheckin struct {
	UserId  string
	Checkin int
}

type TransferAmount struct {
	Amount int
}

func init() {
	//register driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// register model
	orm.RegisterModel(new(SavingsAccount), new(CheckinAccount))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@/test?charset=utf8", 30)

	// create table
	orm.RunSyncdb("default", false, true)
}

func InsertSavingsData(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	var deposit DepositSavings
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

	var deposit DepositCheckin
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

func TransferToCheckin(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	u_id := c.Param("u_id")
	log.Println(u_id)

	var t TransferAmount
	c.BindJSON(&t)
	//t.Amount the amount to be transferred to checkin account
	//amount the amount that is actually present in the savings account of the user

	var amountS int
	var amountC int

	o.Begin()

	err := o.Raw("select savings from savings_account where user_id = ?", u_id).QueryRow(&amountS)
	// log.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Cannot Transfer to checkin",
		})
		o.Rollback()
		return
	}

	err = o.Raw("select checkin from checkin_account where user_id = ?", u_id).QueryRow(&amountC)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Cannot Transfer to checkin",
		})
		o.Rollback()
		return
	}

	//new amount in the savings table
	new_amount_savings := amountS - t.Amount
	new_amount_checkin := amountC + t.Amount

	_, err = o.Raw("update savings_account set savings = ? where user_id = ?", new_amount_savings, u_id).Exec()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Cannot Transfer to checkin",
		})
		o.Rollback()
		return
	}

	_, err = o.Raw("update checkin_account set checkin = ? where user_id = ?", new_amount_checkin, u_id).Exec()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Cannot Transfer to checkin",
		})
		o.Rollback()
		return
	}

	o.Commit()
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Transfer Succesfull",
	})
}

func TransferToSavings(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	u_id := c.Param("u_id")
	log.Println(u_id)

	var t TransferAmount
	c.BindJSON(&t)
	//t.Amount the amount to be transferred to checkin account
	//amount the amount that is actually present in the savings account of the user

	var amountS int
	var amountC int

	o.Begin()

	err := o.Raw("select checkin from checkin_account where user_id = ?", u_id).QueryRow(&amountC)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Cannot Transfer to savings",
		})
		o.Rollback()
		return
	}

	err = o.Raw("select savings from savings_account where user_id = ?", u_id).QueryRow(&amountS)
	// log.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Cannot Transfer to savings",
		})
		o.Rollback()
		return
	}

	//new amount in the savings table
	new_amount_checkin := amountC - t.Amount
	new_amount_savings := amountS + t.Amount

	_, err = o.Raw("update checkin_account set checkin = ? where user_id = ?", new_amount_checkin, u_id).Exec()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Cannot Transfer to savings",
		})
		o.Rollback()
		return
	}

	_, err = o.Raw("update savings_account set savings = ? where user_id = ?", new_amount_savings, u_id).Exec()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Cannot Transfer to savings",
		})
		o.Rollback()
		return
	}

	o.Commit()
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Transfer Succesfull",
	})
}

func main() {
	fmt.Println("Transactions with Golang")

	router := gin.Default()

	router.POST("/savings", InsertSavingsData)
	router.POST("/checkin", InsertCheckinData)
	router.POST("/transfertocheckin/:u_id", TransferToCheckin)
	router.POST("/transfertosavings/:u_id", TransferToSavings)

	fmt.Println("Server running on 8080")
	//If this fails it will log the report and then exit out of the program
	log.Fatal(router.Run("localhost:8080"), nil)

}
