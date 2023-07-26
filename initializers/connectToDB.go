package initializers

import (
	"github.com/beego/beego/orm"
	"github.com/f0rSaaaa/transactions/models"
)

func ConnectToDB() {
	//register driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// register model
	orm.RegisterModel(new(models.SavingsAccount), new(models.CheckinAccount))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@/test?charset=utf8", 30)

	// create table
	orm.RunSyncdb("default", false, true)
}
