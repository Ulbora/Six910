package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddCustomer(t *testing.T) {

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

	var fcus sdbi.Customer
	fcus.StoreID = 1
	sdb.MockCustomer = &fcus

	var cus sdbi.Customer
	cus.StoreID = 1

	sdb.MockAddCustomerSuccess = true
	sdb.MockCustomerID = 2
	resp := m.AddCustomer(&cus)
	if !resp.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddCustomerExists(t *testing.T) {

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

	var fcus sdbi.Customer
	fcus.ID = 2
	fcus.StoreID = 1
	sdb.MockCustomer = &fcus

	var cus sdbi.Customer
	cus.StoreID = 1

	sdb.MockAddCustomerSuccess = true
	sdb.MockCustomerID = 2
	resp := m.AddCustomer(&cus)
	if resp.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddCustomerFail(t *testing.T) {

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

	var fcus sdbi.Customer
	fcus.StoreID = 1
	sdb.MockCustomer = &fcus

	var cus sdbi.Customer
	cus.StoreID = 1

	//sdb.MockAddCustomerSuccess = true
	sdb.MockCustomerID = 2
	resp := m.AddCustomer(&cus)
	if resp.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCustomer(t *testing.T) {

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

	var cus sdbi.Customer
	cus.StoreID = 1

	sdb.MockCustomer = &cus

	sdb.MockUpdateCustomerSuccess = true
	res := m.UpdateCustomer(&cus)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCustomerfail1(t *testing.T) {

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

	var cus sdbi.Customer
	cus.StoreID = 2

	sdb.MockCustomer = &cus

	var cus2 sdbi.Customer
	cus2.StoreID = 4

	sdb.MockUpdateCustomerSuccess = true
	res := m.UpdateCustomer(&cus2)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCustomerFail(t *testing.T) {

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

	var cus sdbi.Customer
	cus.StoreID = 1
	sdb.MockCustomer = &cus

	//sdb.MockUpdateCustomerSuccess = true
	res := m.UpdateCustomer(&cus)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetCustomer(t *testing.T) {

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

	var cus sdbi.Customer
	cus.StoreID = 1
	cus.ID = 6

	sdb.MockCustomer = &cus
	fcus := m.GetCustomer("test", 2)
	if fcus.ID != cus.ID {
		t.Fail()
	}
}

func TestSix910Manager_GetCustomerID(t *testing.T) {

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

	var cus sdbi.Customer
	cus.StoreID = 1
	cus.ID = 6

	sdb.MockCustomer = &cus
	fcus := m.GetCustomerID(4, 1)
	if fcus.ID != cus.ID {
		t.Fail()
	}
}

func TestSix910Manager_GetCustomerIDFail(t *testing.T) {

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

	var cus sdbi.Customer
	cus.StoreID = 11
	cus.ID = 6

	sdb.MockCustomer = &cus
	fcus := m.GetCustomerID(4, 1)
	if fcus.ID == cus.ID {
		t.Fail()
	}
}

func TestSix910Manager_GetCustomerList(t *testing.T) {

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

	var cus sdbi.Customer
	cus.StoreID = 1
	cus.ID = 6

	var cuslst []sdbi.Customer
	cuslst = append(cuslst, cus)

	sdb.MockCustomerList = &cuslst
	fcuslst := m.GetCustomerList(3)
	if (*fcuslst)[0].ID != cus.ID {
		t.Fail()
	}
}

func TestSix910Manager_DeleteCustomer(t *testing.T) {

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

	var cus sdbi.Customer
	cus.StoreID = 1
	cus.ID = 6

	sdb.MockCustomer = &cus

	sdb.MockDeleteCustomerSuccess = true
	res := m.DeleteCustomer(5, 1)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteCustomerFail(t *testing.T) {

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

	var cus sdbi.Customer
	cus.StoreID = 1
	cus.ID = 6
	sdb.MockCustomer = &cus

	//sdb.MockDeleteCustomerSuccess = true
	res := m.DeleteCustomer(5, 1)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteCustomerfail2(t *testing.T) {

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

	var cus sdbi.Customer
	cus.StoreID = 11
	cus.ID = 6

	sdb.MockCustomer = &cus

	sdb.MockDeleteCustomerSuccess = true
	res := m.DeleteCustomer(5, 1)
	if res.Success {
		t.Fail()
	}
}
