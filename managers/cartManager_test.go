package managers

import (
	"fmt"
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
	//cart.IPAddress = "testip"
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

	var cart *sdbi.Cart
	//cart.ID = 4
	sdb.MockCart = cart

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
	fmt.Println("added new cart: ", *res)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCart(t *testing.T) {

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

	sdb.MockCart = &cart
	sdb.MockUpdateCartSuccess = true
	res := m.UpdateCart(&cart)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCartFail1(t *testing.T) {

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

	//sdb.MockUpdateCartSuccess = true
	res := m.UpdateCart(&cart)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCartfail2(t *testing.T) {

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

	var cart2 sdbi.Cart
	cart2.ID = 4
	cart2.StoreID = 4
	sdb.MockCart = &cart2

	sdb.MockUpdateCartSuccess = true
	res := m.UpdateCart(&cart)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetCart(t *testing.T) {

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
	cart.StoreID = 4
	sdb.MockCart = &cart

	res := m.GetCart(4, 4)

	if res.ID != cart.ID {
		t.Fail()
	}
}

func TestSix910Manager_GetCartfail(t *testing.T) {

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
	cart.StoreID = 44
	sdb.MockCart = &cart

	res := m.GetCart(4, 4)

	if res.ID == cart.ID {
		t.Fail()
	}
}

func TestSix910Manager_DeleteCart(t *testing.T) {

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
	cart.CustomerID = 5
	cart.StoreID = 4
	sdb.MockCart = &cart

	sdb.MockDeleteCartSuccess = true

	res := m.DeleteCart(4, 5, 4)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteCartFail1(t *testing.T) {

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
	cart.CustomerID = 55
	cart.StoreID = 4
	sdb.MockCart = &cart

	sdb.MockDeleteCartSuccess = true

	res := m.DeleteCart(4, 5, 4)
	if res.Success {
		t.Fail()
	}
}
func TestSix910Manager_DeleteCartFail(t *testing.T) {

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
	cart.CustomerID = 5
	cart.StoreID = 4
	sdb.MockCart = &cart

	//sdb.MockDeleteCartSuccess = true

	res := m.DeleteCart(4, 5, 4)
	if res.Success {
		t.Fail()
	}
}
