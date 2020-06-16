package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddStore(t *testing.T) {

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

	var str sdbi.Store
	str.City = "Atlanta"
	str.FirstName = "tester"
	str.LastName = "tester"

	sdb.MockAddStoreSuccess = true
	sdb.MockStoreID = 5
	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	res := m.AddStore(&str)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddStorefail(t *testing.T) {

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

	var str sdbi.Store
	str.City = "Atlanta"
	str.FirstName = "tester"
	str.LastName = "tester"

	//sdb.MockAddStoreSuccess = true
	//sdb.MockStoreID = 5
	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	res := m.AddStore(&str)
	if res.Success || res.Message != failStoreMayAlreadyExist {
		t.Fail()
	}
}

func TestSix910Manager_UpdateStore(t *testing.T) {

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

	var str sdbi.Store
	str.City = "Atlanta"
	str.FirstName = "tester"
	str.LastName = "tester"

	sdb.MockUpdateStoreSuccess = true
	res := m.UpdateStore(&str)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateStorefail(t *testing.T) {

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

	var str sdbi.Store
	str.City = "Atlanta"
	str.FirstName = "tester"
	str.LastName = "tester"

	//sdb.MockUpdateStoreSuccess = true
	res := m.UpdateStore(&str)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetStore(t *testing.T) {

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

	var str sdbi.Store
	str.ID = 5
	str.City = "Atlanta"
	str.FirstName = "tester"
	str.LastName = "tester"
	sdb.MockStore = &str

	fstr := m.GetStore("test")
	if fstr.ID != 5 {
		t.Fail()
	}
}

func TestSix910Manager_GetStoreId(t *testing.T) {

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

	var str sdbi.Store
	str.ID = 5
	str.City = "Atlanta"
	str.FirstName = "tester"
	str.LastName = "tester"
	sdb.MockStore = &str

	fstr := m.GetStoreID(5)
	if fstr.ID != 5 {
		t.Fail()
	}
}

func TestSix910Manager_GetStoredomain(t *testing.T) {

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

	var str sdbi.Store
	str.ID = 5
	str.City = "Atlanta"
	str.FirstName = "tester"
	str.LastName = "tester"
	sdb.MockStore = &str

	fstr := m.GetStoreByLocal("domain")
	if fstr.ID != 5 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteStore(t *testing.T) {

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

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec
	sdb.MockDeleteStoreSuccess = true
	res := m.DeleteStore(5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteStorefail(t *testing.T) {

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

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec
	//sdb.MockDeleteStoreSuccess = true
	res := m.DeleteStore(5)
	if res.Success {
		t.Fail()
	}
}
