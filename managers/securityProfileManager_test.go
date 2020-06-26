package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_GetSecurityProfile(t *testing.T) {
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

	var str sdbi.Store
	str.ID = 4
	str.LocalDomain = "testdomain"
	str.StoreName = "test"

	sdb.MockSecurity = &sec
	sdb.MockStore = &str

	prof := m.GetSecurityProfile("test", "testdomain")
	if prof.Store.ID != 4 {
		t.Fail()
	}
}

func TestSix910Manager_GetSecurityProfile2(t *testing.T) {
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
	//sec.OauthOn = true

	var str sdbi.Store
	str.ID = 4
	str.LocalDomain = "testdomain"
	str.StoreName = "test"

	sdb.MockSecurity = &sec
	sdb.MockStore = &str

	prof := m.GetSecurityProfile("test", "testdomain")
	if prof.Store.ID != 4 {
		t.Fail()
	}
}
