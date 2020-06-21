package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddOrderItem(t *testing.T) {
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

	var i sdbi.OrderItem
	i.BackOrdered = true

	sdb.MockAddOrderItemSuccess = true
	sdb.MockOrderItemID = 6

	res := m.AddOrderItem(&i, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddOrderItemFAil2(t *testing.T) {
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

	var i sdbi.OrderItem
	i.BackOrdered = true

	sdb.MockAddOrderItemSuccess = true
	//sdb.MockOrderItemID = 6

	res := m.AddOrderItem(&i, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddOrderItemFail(t *testing.T) {
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

	var i sdbi.OrderItem
	i.BackOrdered = true

	sdb.MockAddOrderItemSuccess = true
	sdb.MockOrderItemID = 6

	res := m.AddOrderItem(&i, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateOrderItem(t *testing.T) {
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

	var i sdbi.OrderItem
	i.BackOrdered = true

	sdb.MockUpdateOrderItemSuccess = true

	res := m.UpdateOrderItem(&i, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateOrderItemFail(t *testing.T) {
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

	var i sdbi.OrderItem
	i.BackOrdered = true

	sdb.MockUpdateOrderItemSuccess = true

	res := m.UpdateOrderItem(&i, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateOrderItemFail2(t *testing.T) {
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

	var i sdbi.OrderItem
	i.BackOrdered = true

	//sdb.MockUpdateOrderItemSuccess = true

	res := m.UpdateOrderItem(&i, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetOrderItem(t *testing.T) {
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

	var i sdbi.OrderItem
	i.BackOrdered = true

	sdb.MockOrderItem = &i

	fi := m.GetOrderItem(4, 5)
	if fi.BackOrdered != i.BackOrdered {
		t.Fail()
	}
}

func TestSix910Manager_GetOrderItemFail(t *testing.T) {
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

	var i sdbi.OrderItem
	i.BackOrdered = true

	sdb.MockOrderItem = &i

	fi := m.GetOrderItem(4, 5)
	if fi.BackOrdered == i.BackOrdered {
		t.Fail()
	}
}

func TestSix910Manager_GetOrderItemList(t *testing.T) {
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

	var i sdbi.OrderItem
	i.BackOrdered = true

	sdb.MockOrderItem = &i

	var il []sdbi.OrderItem
	il = append(il, i)

	sdb.MockOrderItemList = &il

	fil := m.GetOrderItemList(4, 5)
	if len(*fil) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetOrderItemListFail(t *testing.T) {
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

	var i sdbi.OrderItem
	i.BackOrdered = true

	sdb.MockOrderItem = &i

	var il []sdbi.OrderItem
	il = append(il, i)

	sdb.MockOrderItemList = &il

	fil := m.GetOrderItemList(4, 5)
	if len(*fil) != 0 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteOrderItem(t *testing.T) {
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

	var i sdbi.OrderItem
	i.BackOrdered = true

	sdb.MockOrderItem = &i

	sdb.MockDeleteOrderItemSuccess = true

	res := m.DeleteOrderItem(3, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteOrderItemFail(t *testing.T) {
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

	var i sdbi.OrderItem
	i.BackOrdered = true

	sdb.MockOrderItem = &i

	sdb.MockDeleteOrderItemSuccess = true

	res := m.DeleteOrderItem(3, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteOrderItemFail2(t *testing.T) {
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

	var i sdbi.OrderItem
	i.BackOrdered = true

	sdb.MockOrderItem = &i

	//sdb.MockDeleteOrderItemSuccess = true

	res := m.DeleteOrderItem(3, 5)
	if res.Success {
		t.Fail()
	}
}
