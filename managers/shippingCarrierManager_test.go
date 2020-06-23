package managers

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddShippingCarrier(t *testing.T) {
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

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"
	sc.StoreID = 5

	sdb.MockAddShippingCarrierSuccess = true
	sdb.MockShippingCarrierID = 1

	res := m.AddShippingCarrier(&sc)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddShippingCarrierFail(t *testing.T) {
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

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"
	sc.StoreID = 5

	//sdb.MockAddShippingCarrierSuccess = true
	sdb.MockShippingCarrierID = 1

	res := m.AddShippingCarrier(&sc)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShippingCarrier(t *testing.T) {
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

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"
	sc.StoreID = 5
	sdb.MockShippingCarrier = &sc

	sdb.MockUpdateShippingCarrierSuccess = true

	res := m.UpdateShippingCarrier(&sc)
	fmt.Println("s carrier res: ", res)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShippingCarrierFail2(t *testing.T) {
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

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"
	sc.StoreID = 5
	sdb.MockShippingCarrier = &sc

	//sdb.MockUpdateShippingCarrierSuccess = true

	res := m.UpdateShippingCarrier(&sc)
	fmt.Println("s carrier res: ", res)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateShippingCarrierFAil(t *testing.T) {
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

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"
	sc.StoreID = 5
	sdb.MockShippingCarrier = &sc

	sdb.MockUpdateShippingCarrierSuccess = true

	var sc2 sdbi.ShippingCarrier
	sc2.Carrier = "UPS"
	sc2.StoreID = 55

	res := m.UpdateShippingCarrier(&sc2)
	fmt.Println("s carrier res: ", res)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetShippingCarrier(t *testing.T) {
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

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"
	sc.StoreID = 5
	sdb.MockShippingCarrier = &sc

	fsc := m.GetShippingCarrier(4, 5)
	if fsc.Carrier != sc.Carrier {
		t.Fail()
	}
}

func TestSix910Manager_GetShippingCarrierfail(t *testing.T) {
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

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"
	sc.StoreID = 55
	sdb.MockShippingCarrier = &sc

	fsc := m.GetShippingCarrier(4, 5)
	if fsc.Carrier == sc.Carrier {
		t.Fail()
	}
}

func TestSix910Manager_GetShippingCarrierList(t *testing.T) {
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

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"
	sc.StoreID = 55

	var lst []sdbi.ShippingCarrier
	lst = append(lst, sc)

	sdb.MockShippingCarrierList = &lst

	flst := m.GetShippingCarrierList(4)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShippingCarrier(t *testing.T) {
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

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"
	sc.StoreID = 5
	sdb.MockShippingCarrier = &sc

	sdb.MockDeleteShippingCarrierSuccess = true

	res := m.DeleteShippingCarrier(4, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShippingCarrierFail(t *testing.T) {
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

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"
	sc.StoreID = 56
	sdb.MockShippingCarrier = &sc

	sdb.MockDeleteShippingCarrierSuccess = true

	res := m.DeleteShippingCarrier(4, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteShippingCarrierFail2(t *testing.T) {
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

	var sc sdbi.ShippingCarrier
	sc.Carrier = "UPS"
	sc.StoreID = 5
	sdb.MockShippingCarrier = &sc

	//sdb.MockDeleteShippingCarrierSuccess = true

	res := m.DeleteShippingCarrier(4, 5)
	if res.Success {
		t.Fail()
	}
}
