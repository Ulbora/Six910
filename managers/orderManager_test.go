package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddOrder(t *testing.T) {
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

	sdb.MockAddOrderSuccess = true
	sdb.MockOrderID = 7

	res := m.AddOrder(&o)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddOrderFail(t *testing.T) {
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

	//sdb.MockAddOrderSuccess = true
	sdb.MockOrderID = 7

	res := m.AddOrder(&o)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateOrder(t *testing.T) {
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

	sdb.MockUpdateOrderSuccess = true

	res := m.UpdateOrder(&o)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateOrderFail(t *testing.T) {
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

	sdb.MockUpdateOrderSuccess = true

	var o2 sdbi.Order
	o2.BillingAddress = "123"
	o2.StoreID = 555

	res := m.UpdateOrder(&o2)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateOrderFail2(t *testing.T) {
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

	//sdb.MockUpdateOrderSuccess = true

	res := m.UpdateOrder(&o)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetOrder(t *testing.T) {
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

	fo := m.GetOrder(3, 5)
	if fo.BillingAddress != o.BillingAddress {
		t.Fail()
	}
}

func TestSix910Manager_GetOrderfail(t *testing.T) {
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

	fo := m.GetOrder(3, 5)
	if fo.BillingAddress == o.BillingAddress {
		t.Fail()
	}
}

func TestSix910Manager_GetOrderList(t *testing.T) {
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

	var ol []sdbi.Order
	ol = append(ol, o)

	sdb.MockOrderList = &ol

	fol := m.GetOrderList(4, 5)
	if len(*fol) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetStoreOrderList(t *testing.T) {
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

	var ol []sdbi.Order
	ol = append(ol, o)

	sdb.MockOrderList = &ol

	fol := m.GetStoreOrderList(5)
	if len(*fol) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetStoreOrderListByStatus(t *testing.T) {
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

	var ol []sdbi.Order
	ol = append(ol, o)

	sdb.MockOrderList = &ol

	fol := m.GetStoreOrderListByStatus("test", 5)
	if len(*fol) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteOrder(t *testing.T) {
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
	sdb.MockDeleteOrderSuccess = true

	res := m.DeleteOrder(4, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteOrderFAil(t *testing.T) {
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
	sdb.MockDeleteOrderSuccess = true

	res := m.DeleteOrder(4, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteOrderFail2(t *testing.T) {
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
	//sdb.MockDeleteOrderSuccess = true

	res := m.DeleteOrder(4, 5)
	if res.Success {
		t.Fail()
	}
}
