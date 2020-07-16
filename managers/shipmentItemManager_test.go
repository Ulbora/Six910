package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddShipmentItem(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	sdb.MockAddShipmentItemSuccess = true
	sdb.MockShipmentItemID = 3

	var si sdbi.ShipmentItem
	si.Quantity = 3

	res := m.AddShipmentItem(&si, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddShipmentItemFAil(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	sdb.MockAddShipmentItemSuccess = true
	sdb.MockShipmentItemID = 3

	var si sdbi.ShipmentItem
	si.Quantity = 3

	res := m.AddShipmentItem(&si, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddShipmentItemFail2(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	//sdb.MockAddShipmentItemSuccess = true
	sdb.MockShipmentItemID = 3

	var si sdbi.ShipmentItem
	si.Quantity = 3

	res := m.AddShipmentItem(&si, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShipmentItem(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	sdb.MockUpdateShipmentItemSuccess = true

	var si sdbi.ShipmentItem
	si.Quantity = 3

	res := m.UpdateShipmentItem(&si, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShipmentItemFail(t *testing.T) {
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
	o.StoreID = 57

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	sdb.MockUpdateShipmentItemSuccess = true

	var si sdbi.ShipmentItem
	si.Quantity = 3

	res := m.UpdateShipmentItem(&si, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShipmentItemFail2(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	//sdb.MockUpdateShipmentItemSuccess = true

	var si sdbi.ShipmentItem
	si.Quantity = 3

	res := m.UpdateShipmentItem(&si, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetShipmentItem(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	var si sdbi.ShipmentItem
	si.Quantity = 3

	sdb.MockShipmentItem = &si

	fsi := m.GetShipmentItem(4, 5)
	if fsi.Quantity != si.Quantity {
		t.Fail()
	}
}

func TestSix910Manager_GetShipmentItemFAil(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	var si sdbi.ShipmentItem
	si.Quantity = 3

	sdb.MockShipmentItem = &si

	fsi := m.GetShipmentItem(4, 5)
	if fsi.Quantity == si.Quantity {
		t.Fail()
	}
}

func TestSix910Manager_GetShipmentItemList(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	var si sdbi.ShipmentItem
	si.Quantity = 3

	var sblst []sdbi.ShipmentItem
	sblst = append(sblst, si)

	sdb.MockShipmentItemList = &sblst

	flst := m.GetShipmentItemList(4, 5)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetShipmentItemListFail(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	var si sdbi.ShipmentItem
	si.Quantity = 3

	var sblst []sdbi.ShipmentItem
	sblst = append(sblst, si)

	sdb.MockShipmentItemList = &sblst

	flst := m.GetShipmentItemList(4, 5)
	if len(*flst) != 0 {
		t.Fail()
	}
}

func TestSix910Manager_GetShipmentItemListByBox(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	var si sdbi.ShipmentItem
	si.Quantity = 3
	var si2 sdbi.ShipmentItem
	si2.Quantity = 3

	var sblst []sdbi.ShipmentItem
	sblst = append(sblst, si)
	sblst = append(sblst, si2)

	sdb.MockShipmentItemList = &sblst

	flst := m.GetShipmentItemListByBox(4, 3, 5)
	if len(*flst) != 2 {
		t.Fail()
	}
}

func TestSix910Manager_GetShipmentItemListByBoxFail(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	var si sdbi.ShipmentItem
	si.Quantity = 3
	var si2 sdbi.ShipmentItem
	si2.Quantity = 3

	var sblst []sdbi.ShipmentItem
	sblst = append(sblst, si)
	sblst = append(sblst, si2)

	sdb.MockShipmentItemList = &sblst

	flst := m.GetShipmentItemListByBox(4, 3, 5)
	if len(*flst) != 0 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShipmentItem(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	var si sdbi.ShipmentItem
	si.Quantity = 4

	sdb.MockShipmentItem = &si

	sdb.MockDeleteShipmentItemSuccess = true

	res := m.DeleteShipmentItem(3, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShipmentItemFail(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	var si sdbi.ShipmentItem
	si.Quantity = 4

	sdb.MockShipmentItem = &si

	sdb.MockDeleteShipmentItemSuccess = true

	res := m.DeleteShipmentItem(3, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShipmentItemFail2(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	var si sdbi.ShipmentItem
	si.Quantity = 4

	sdb.MockShipmentItem = &si

	//sdb.MockDeleteShipmentItemSuccess = true

	res := m.DeleteShipmentItem(3, 5)
	if res.Success {
		t.Fail()
	}
}
