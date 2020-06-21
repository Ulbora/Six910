package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddOrderTransaction(t *testing.T) {
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

	var o sdbi.Order
	o.BillingAddress = "123"
	o.StoreID = 5

	sdb.MockOrder = &o

	var ot sdbi.OrderTransaction
	ot.Amount = 10.10

	sdb.MockAddOrderTransactionSuccess = true
	sdb.MockOrderTransactionID = 9

	res := m.AddOrderTransaction(&ot, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddOrderTransactionFail(t *testing.T) {
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

	var o sdbi.Order
	o.BillingAddress = "123"
	o.StoreID = 56

	sdb.MockOrder = &o

	var ot sdbi.OrderTransaction
	ot.Amount = 10.10

	sdb.MockAddOrderTransactionSuccess = true
	sdb.MockOrderTransactionID = 9

	res := m.AddOrderTransaction(&ot, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddOrderTransactionFail2(t *testing.T) {
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

	var o sdbi.Order
	o.BillingAddress = "123"
	o.StoreID = 5

	sdb.MockOrder = &o

	var ot sdbi.OrderTransaction
	ot.Amount = 10.10

	sdb.MockAddOrderTransactionSuccess = true
	//sdb.MockOrderTransactionID = 9

	res := m.AddOrderTransaction(&ot, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetOrderTransactionList(t *testing.T) {
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

	var o sdbi.Order
	o.BillingAddress = "123"
	o.StoreID = 5

	sdb.MockOrder = &o

	var ot sdbi.OrderTransaction
	ot.Amount = 10.10

	var otl []sdbi.OrderTransaction
	otl = append(otl, ot)

	sdb.MockOrderTransactionList = &otl

	fotl := m.GetOrderTransactionList(4, 5)
	if len(*fotl) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetOrderTransactionListFail(t *testing.T) {
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

	var o sdbi.Order
	o.BillingAddress = "123"
	o.StoreID = 55

	sdb.MockOrder = &o

	var ot sdbi.OrderTransaction
	ot.Amount = 10.10

	var otl []sdbi.OrderTransaction
	otl = append(otl, ot)

	sdb.MockOrderTransactionList = &otl

	fotl := m.GetOrderTransactionList(4, 5)
	if len(*fotl) != 0 {
		t.Fail()
	}
}
