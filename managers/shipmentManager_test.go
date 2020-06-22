package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddShipment(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockAddShipmentSuccess = true
	sdb.MockShipmentID = 2

	res := m.AddShipment(&sh, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddShipmentFail(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockAddShipmentSuccess = true
	sdb.MockShipmentID = 2

	res := m.AddShipment(&sh, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddShipmentFail2(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	//sdb.MockAddShipmentSuccess = true
	sdb.MockShipmentID = 2

	res := m.AddShipment(&sh, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShipment(t *testing.T) {
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
	o.ID = 7
	o.BillingAddress = "123"
	o.StoreID = 5

	sdb.MockOrder = &o

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockShipment = &sh

	sdb.MockUpdateShipmentSuccess = true

	res := m.UpdateShipment(&sh, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShipmentFail(t *testing.T) {
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
	o.ID = 7
	o.BillingAddress = "123"
	o.StoreID = 58

	sdb.MockOrder = &o

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockShipment = &sh

	sdb.MockUpdateShipmentSuccess = true

	res := m.UpdateShipment(&sh, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShipmentFail2(t *testing.T) {
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
	o.ID = 7
	o.BillingAddress = "123"
	o.StoreID = 5

	sdb.MockOrder = &o

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 77

	sdb.MockShipment = &sh

	sdb.MockUpdateShipmentSuccess = true

	res := m.UpdateShipment(&sh, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShipmentFail3(t *testing.T) {
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
	o.ID = 7
	o.BillingAddress = "123"
	o.StoreID = 5

	sdb.MockOrder = &o

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockShipment = &sh

	//sdb.MockUpdateShipmentSuccess = true

	res := m.UpdateShipment(&sh, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetShipment(t *testing.T) {
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
	o.ID = 7
	o.BillingAddress = "123"
	o.StoreID = 5

	sdb.MockOrder = &o

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7
	sdb.MockShipment = &sh

	fsh := m.GetShipment(4, 5)
	if fsh.Boxes != sh.Boxes {
		t.Fail()
	}
}

func TestSix910Manager_GetShipmentFail(t *testing.T) {
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
	o.ID = 7
	o.BillingAddress = "123"
	o.StoreID = 55

	sdb.MockOrder = &o

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7
	sdb.MockShipment = &sh

	fsh := m.GetShipment(4, 5)
	if fsh.Boxes == sh.Boxes {
		t.Fail()
	}
}

func TestSix910Manager_GetShipmentList(t *testing.T) {
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
	o.ID = 7
	o.BillingAddress = "123"
	o.StoreID = 5

	sdb.MockOrder = &o

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	var shl []sdbi.Shipment
	shl = append(shl, sh)
	sdb.MockShipmentList = &shl

	fshl := m.GetShipmentList(3, 5)
	if len(*fshl) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetShipmentListFail(t *testing.T) {
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
	o.ID = 7
	o.BillingAddress = "123"
	o.StoreID = 55

	sdb.MockOrder = &o

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	var shl []sdbi.Shipment
	shl = append(shl, sh)
	sdb.MockShipmentList = &shl

	fshl := m.GetShipmentList(3, 5)
	if len(*fshl) != 0 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShipment(t *testing.T) {
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
	o.ID = 7
	o.BillingAddress = "123"
	o.StoreID = 5

	sdb.MockOrder = &o

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockShipment = &sh

	sdb.MockDeleteShipmentSuccess = true

	res := m.DeleteShipment(4, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShipmentFail(t *testing.T) {
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
	o.ID = 7
	o.BillingAddress = "123"
	o.StoreID = 55

	sdb.MockOrder = &o

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockShipment = &sh

	sdb.MockDeleteShipmentSuccess = true

	res := m.DeleteShipment(4, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShipmentFAil2(t *testing.T) {
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
	o.ID = 7
	o.BillingAddress = "123"
	o.StoreID = 5

	sdb.MockOrder = &o

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockShipment = &sh

	//sdb.MockDeleteShipmentSuccess = true

	res := m.DeleteShipment(4, 5)
	if res.Success {
		t.Fail()
	}
}
