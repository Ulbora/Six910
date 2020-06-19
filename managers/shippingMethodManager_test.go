package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddShippingMethod(t *testing.T) {
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

	var smth sdbi.ShippingMethod

	sdb.MockAddShippingMethodSuccess = true
	sdb.MockShippingMethodID = 4

	res := m.AddShippingMethod(&smth)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddShippingMethodFail(t *testing.T) {
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

	var smth sdbi.ShippingMethod

	sdb.MockAddShippingMethodSuccess = true
	//sdb.MockShippingMethodID = 4

	res := m.AddShippingMethod(&smth)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShippingMethod(t *testing.T) {
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

	var smth sdbi.ShippingMethod
	sdb.MockShippingMethod = &smth

	sdb.MockUpdateShippingMethodSuccess = true

	res := m.UpdateShippingMethod(&smth)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShippingMethodFail(t *testing.T) {
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

	var smth2 sdbi.ShippingMethod
	smth2.StoreID = 5
	sdb.MockShippingMethod = &smth2

	var smth sdbi.ShippingMethod
	smth.StoreID = 2

	sdb.MockUpdateShippingMethodSuccess = true

	res := m.UpdateShippingMethod(&smth)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShippingMethodFail2(t *testing.T) {
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

	var smth sdbi.ShippingMethod
	sdb.MockShippingMethod = &smth

	//sdb.MockUpdateShippingMethodSuccess = true

	res := m.UpdateShippingMethod(&smth)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetShippingMethod(t *testing.T) {
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

	var smth2 sdbi.ShippingMethod
	smth2.Cost = 10.00
	smth2.StoreID = 5
	sdb.MockShippingMethod = &smth2

	fsm := m.GetShippingMethod(4, 5)
	if fsm.Cost != smth2.Cost {
		t.Fail()
	}
}

func TestSix910Manager_GetShippingMethodFail(t *testing.T) {
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

	var smth2 sdbi.ShippingMethod
	smth2.Cost = 10.00
	smth2.StoreID = 55
	sdb.MockShippingMethod = &smth2

	fsm := m.GetShippingMethod(4, 5)
	if fsm.Cost == smth2.Cost {
		t.Fail()
	}
}

func TestSix910Manager_GetShippingMethodList(t *testing.T) {
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

	var smth2 sdbi.ShippingMethod
	smth2.Cost = 10.00
	smth2.StoreID = 5

	var smlst []sdbi.ShippingMethod
	smlst = append(smlst, smth2)

	sdb.MockShippingMethodList = &smlst

	flst := m.GetShippingMethodList(5)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShippingMethod(t *testing.T) {
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

	var smth2 sdbi.ShippingMethod
	smth2.Cost = 10.00
	smth2.StoreID = 5
	sdb.MockShippingMethod = &smth2

	sdb.MockDeleteShippingMethodSuccess = true

	res := m.DeleteShippingMethod(3, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShippingMethodFail(t *testing.T) {
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

	var smth2 sdbi.ShippingMethod
	smth2.Cost = 10.00
	smth2.StoreID = 55
	sdb.MockShippingMethod = &smth2

	sdb.MockDeleteShippingMethodSuccess = true

	res := m.DeleteShippingMethod(3, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShippingMethodFail2(t *testing.T) {
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

	var smth2 sdbi.ShippingMethod
	smth2.Cost = 10.00
	smth2.StoreID = 5
	sdb.MockShippingMethod = &smth2

	//sdb.MockDeleteShippingMethodSuccess = true

	res := m.DeleteShippingMethod(3, 5)
	if res.Success {
		t.Fail()
	}
}
