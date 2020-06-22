package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddShipmentBox(t *testing.T) {
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

	sdb.MockAddShipmentBoxSuccess = true
	sdb.MockShipmentBoxID = 3

	var sb sdbi.ShipmentBox
	sb.BoxNumber = 2

	res := m.AddShipmentBox(&sb, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddShipmentBoxFail(t *testing.T) {
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

	sdb.MockAddShipmentBoxSuccess = true
	sdb.MockShipmentBoxID = 3

	var sb sdbi.ShipmentBox
	sb.BoxNumber = 2

	res := m.AddShipmentBox(&sb, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddShipmentBoxfail2(t *testing.T) {
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

	//sdb.MockAddShipmentBoxSuccess = true
	sdb.MockShipmentBoxID = 3

	var sb sdbi.ShipmentBox
	sb.BoxNumber = 2

	res := m.AddShipmentBox(&sb, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShipmentBox(t *testing.T) {
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

	sdb.MockUpdateShipmentBoxSuccess = true

	var sb sdbi.ShipmentBox
	sb.BoxNumber = 2

	res := m.UpdateShipmentBox(&sb, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShipmentBoxFail(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	sdb.MockUpdateShipmentBoxSuccess = true

	var sb sdbi.ShipmentBox
	sb.BoxNumber = 2

	res := m.UpdateShipmentBox(&sb, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShipmentBoxFail2(t *testing.T) {
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

	//sdb.MockUpdateShipmentBoxSuccess = true

	var sb sdbi.ShipmentBox
	sb.BoxNumber = 2

	res := m.UpdateShipmentBox(&sb, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetShipmentBox(t *testing.T) {
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

	var sb sdbi.ShipmentBox
	sb.BoxNumber = 2

	sdb.MockShipmentBox = &sb

	fsb := m.GetShipmentBox(4, 5)
	if fsb.BoxNumber != sb.BoxNumber {
		t.Fail()
	}
}

func TestSix910Manager_GetShipmentBoxFail(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	var sb sdbi.ShipmentBox
	sb.BoxNumber = 2

	sdb.MockShipmentBox = &sb

	fsb := m.GetShipmentBox(4, 5)
	if fsb.BoxNumber == sb.BoxNumber {
		t.Fail()
	}
}

func TestSix910Manager_GetShipmentBoxList(t *testing.T) {
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

	var sb sdbi.ShipmentBox
	sb.BoxNumber = 2

	var sblst []sdbi.ShipmentBox
	sblst = append(sblst, sb)

	sdb.MockShipmentBoxList = &sblst

	flst := m.GetShipmentBoxList(3, 5)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetShipmentBoxListFAil(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	var sb sdbi.ShipmentBox
	sb.BoxNumber = 2

	var sblst []sdbi.ShipmentBox
	sblst = append(sblst, sb)

	sdb.MockShipmentBoxList = &sblst

	flst := m.GetShipmentBoxList(3, 5)
	if len(*flst) != 0 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShipmentBox(t *testing.T) {
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

	var sb sdbi.ShipmentBox
	sb.BoxNumber = 2

	sdb.MockShipmentBox = &sb

	sdb.MockDeleteShipmentBoxSuccess = true

	res := m.DeleteShipmentBox(3, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShipmentBoxFail(t *testing.T) {
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

	var sb sdbi.ShipmentBox
	sb.BoxNumber = 2

	sdb.MockShipmentBox = &sb

	//sdb.MockDeleteShipmentBoxSuccess = true

	res := m.DeleteShipmentBox(3, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShipmentBoxFAil2(t *testing.T) {
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

	var sh sdbi.Shipment
	sh.Boxes = 2
	sh.OrderID = 7

	sdb.MockOrder = &o
	sdb.MockShipment = &sh

	var sb sdbi.ShipmentBox
	sb.BoxNumber = 2

	sdb.MockShipmentBox = &sb

	sdb.MockDeleteShipmentBoxSuccess = true

	res := m.DeleteShipmentBox(3, 5)
	if res.Success {
		t.Fail()
	}
}
