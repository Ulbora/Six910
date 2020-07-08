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
	str.ID = 5
	str.City = "Atlanta"
	str.FirstName = "tester"
	str.LastName = "tester"
	str.StoreName = "teststore"

	sdb.MockStore = &str

	sdb.MockUpdateStoreSuccess = true
	res := m.UpdateStore(&str)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateStoreFail2(t *testing.T) {

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
	str.StoreName = "teststore"

	sdb.MockStore = &str

	var str2 sdbi.Store
	str2.ID = 55
	str2.City = "Atlanta"
	str2.FirstName = "tester"
	str2.LastName = "tester"
	str2.StoreName = "teststore"

	sdb.MockUpdateStoreSuccess = true
	res := m.UpdateStore(&str2)
	if res.Success {
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
	sdb.MockStore = &str

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
	str.LocalDomain = "localtester"
	sdb.MockStore = &str

	fstr := m.GetStore("test", "localtester")
	if fstr.ID != 5 {
		t.Fail()
	}
}

func TestSix910Manager_GetStoreFail(t *testing.T) {

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
	str.LocalDomain = "localtester2"
	sdb.MockStore = &str

	fstr := m.GetStore("test", "localtester")
	if fstr.ID == 5 {
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

	var str sdbi.Store
	str.ID = 5
	str.City = "Atlanta"
	str.FirstName = "tester"
	str.LastName = "tester"
	str.LocalDomain = "localtester"
	str.StoreName = "teststore"
	sdb.MockStore = &str

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec
	sdb.MockDeleteStoreSuccess = true
	res := m.DeleteStore("teststore", "localtester")
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteStoreFail2(t *testing.T) {

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
	str.LocalDomain = "localtester1"
	str.StoreName = "teststore"
	sdb.MockStore = &str

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec
	sdb.MockDeleteStoreSuccess = true
	res := m.DeleteStore("teststore", "localtester")
	if res.Success {
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

	var str sdbi.Store
	str.ID = 5
	str.City = "Atlanta"
	str.FirstName = "tester"
	str.LastName = "tester"
	str.LocalDomain = "localtester"
	str.StoreName = "teststore"
	sdb.MockStore = &str

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec
	//sdb.MockDeleteStoreSuccess = true
	res := m.DeleteStore("teststore", "localtester")
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteStoreNoOauth(t *testing.T) {

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
	str.LocalDomain = "localtester"
	str.StoreName = "teststore"
	sdb.MockStore = &str

	var sec sdbi.Security
	//sec.OauthOn = true
	sdb.MockSecurity = &sec
	sdb.MockDeleteStoreSuccess = true
	res := m.DeleteStore("teststore", "localtester")
	if res.Success || res.Code != 401 {
		t.Fail()
	}
}
