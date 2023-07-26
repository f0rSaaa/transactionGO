package controllers

import (
	"log"
	"net/http"

	"github.com/beego/beego/orm"
	"github.com/f0rSaaaa/transactions/models"
	"github.com/gin-gonic/gin"
)

func TransferToCheckin(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	u_id := c.Param("u_id")
	log.Println(u_id)

	var t models.TransferAmount
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

	var t models.TransferAmount
	c.BindJSON(&t)
	//t.Amount the amount to be transferred to savings account
	//amount the amount that is actually present in the checkin account of the user

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
