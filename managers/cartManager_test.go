package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddCart(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	sdb.MockCart = &cart

	var cus sdbi.Customer
	cus.ID = 2
	cus.StoreID = 3
	sdb.MockCustomer = &cus

	var ncart sdbi.Cart
	ncart.CustomerID = 2
	ncart.StoreID = 3

	res := m.AddCart(&ncart)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddCart2(t *testing.T) {

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

	var cart sdbi.Cart
	//cart.ID = 4
	sdb.MockCart = &cart

	var cus sdbi.Customer
	cus.ID = 2
	cus.StoreID = 3
	sdb.MockCustomer = &cus

	var ncart sdbi.Cart
	ncart.CustomerID = 2
	ncart.StoreID = 3
	sdb.MockAddCartSuccess = true
	sdb.MockCartID = 3

	res := m.AddCart(&ncart)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddCart2Fail1(t *testing.T) {

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

	var cart sdbi.Cart
	//cart.ID = 4
	sdb.MockCart = &cart

	var cus sdbi.Customer
	cus.ID = 2
	//cus.StoreID = 3
	sdb.MockCustomer = &cus

	var ncart sdbi.Cart
	ncart.CustomerID = 2
	ncart.StoreID = 3
	sdb.MockAddCartSuccess = true
	sdb.MockCartID = 3

	res := m.AddCart(&ncart)
	if res.Success {
		t.Fail()
	}
}
func TestSix910Manager_AddCart2fail(t *testing.T) {

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

	var cart sdbi.Cart
	//cart.ID = 4
	sdb.MockCart = &cart

	var cus sdbi.Customer
	cus.ID = 2
	cus.StoreID = 3
	sdb.MockCustomer = &cus

	var ncart sdbi.Cart
	ncart.CustomerID = 2
	ncart.StoreID = 3
	//sdb.MockAddCartSuccess = true
	//sdb.MockCartID = 3

	res := m.AddCart(&ncart)
	if res.Success {
		t.Fail()
	}
}
