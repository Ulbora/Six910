package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddOrderComments(t *testing.T) {
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

	sdb.MockAddOrderCommentSuccess = true
	sdb.MockOrderCommentID = 9

	var c sdbi.OrderComment
	c.Comment = "test"

	res := m.AddOrderComments(&c, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddOrderCommentsfail(t *testing.T) {
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

	sdb.MockAddOrderCommentSuccess = true
	sdb.MockOrderCommentID = 9

	var c sdbi.OrderComment
	c.Comment = "test"

	res := m.AddOrderComments(&c, 5)
	if res.Success {
		t.Fail()
	}
}
func TestSix910Manager_AddOrderCommentsfail2(t *testing.T) {
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

	sdb.MockAddOrderCommentSuccess = true
	//sdb.MockOrderCommentID = 9

	var c sdbi.OrderComment
	c.Comment = "test"

	res := m.AddOrderComments(&c, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetOrderCommentList(t *testing.T) {
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

	var c sdbi.OrderComment
	c.Comment = "test"
	var cl []sdbi.OrderComment
	cl = append(cl, c)

	sdb.MockOrderCommentList = &cl

	fcl := m.GetOrderCommentList(4, 5)
	if len(*fcl) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetOrderCommentListFail(t *testing.T) {
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

	var c sdbi.OrderComment
	c.Comment = "test"
	var cl []sdbi.OrderComment
	cl = append(cl, c)

	sdb.MockOrderCommentList = &cl

	fcl := m.GetOrderCommentList(4, 5)
	if len(*fcl) != 0 {
		t.Fail()
	}
}
