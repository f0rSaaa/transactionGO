package models

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
