package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddAddress(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	m := sm.GetNew()

	var cus sdbi.Customer
	cus.StoreID = 5

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.CustomerID = 6
	sdb.MockAddAddressSuccess = true
	sdb.MockAddressID = 3
	res := m.AddAddress(&add, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddAddressFail1(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	m := sm.GetNew()

	var cus sdbi.Customer
	cus.StoreID = 55

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.CustomerID = 6
	sdb.MockAddAddressSuccess = true
	sdb.MockAddressID = 3
	res := m.AddAddress(&add, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddAddressFail2(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	m := sm.GetNew()

	var cus sdbi.Customer
	cus.StoreID = 5

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.CustomerID = 6
	//sdb.MockAddAddressSuccess = true
	sdb.MockAddressID = 3
	res := m.AddAddress(&add, 5)
	if res.Success {
		t.Fail()
	}
}
