package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddAdminUser(t *testing.T) {
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
	sdb.MockSecurity = &sec
	sdb.MockAddLocalAccountSuccess = true

	var u User
	u.Enabled = true
	u.StoreID = 1
	u.Username = "tester"
	u.Password = "tester"
	u.Role = "admin"

	res := m.AddAdminUser(&u)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddAdminUserfail(t *testing.T) {
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
	sdb.MockSecurity = &sec
	//sdb.MockAddLocalAccountSuccess = true

	var u User
	u.Enabled = true
	u.StoreID = 1
	u.Username = "tester"
	u.Password = "tester"
	u.Role = "admin"

	res := m.AddAdminUser(&u)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddCustomerUser(t *testing.T) {
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

	//var sec sdbi.Security
	//sdb.MockSecurity = &sec
	sdb.MockAddLocalAccountSuccess = true

	var u User
	u.Enabled = true
	u.StoreID = 1
	u.Username = "tester"
	u.Password = "tester"
	u.Role = "admin"
	u.CustomerID = 2

	res := m.AddCustomerUser(&u)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddCustomerUserFail(t *testing.T) {
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

	//var sec sdbi.Security
	//sdb.MockSecurity = &sec
	//sdb.MockAddLocalAccountSuccess = true

	var u User
	u.Enabled = true
	u.StoreID = 1
	u.Username = "tester"
	u.Password = "tester"
	u.Role = "admin"
	u.CustomerID = 2

	res := m.AddCustomerUser(&u)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateUser(t *testing.T) {
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

	//var sec sdbi.Security
	//sdb.MockSecurity = &sec
	sdb.MockUpdateLocalAccountSuccess = true
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	_, hpw := sm.hashPassword("tester")
	lc.Password = hpw
	sdb.MockLocalAccount = &lc

	var u User
	u.Enabled = true
	u.StoreID = 1
	u.Username = "tester"
	u.OldPassword = "tester"
	u.Password = "tester2"
	u.Role = "admin"
	u.CustomerID = 2

	res := m.UpdateUser(&u)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateUserFail(t *testing.T) {
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

	//var sec sdbi.Security
	//sdb.MockSecurity = &sec
	//sdb.MockUpdateLocalAccountSuccess = true
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	_, hpw := sm.hashPassword("tester")
	lc.Password = hpw
	sdb.MockLocalAccount = &lc

	var u User
	u.Enabled = true
	u.StoreID = 1
	u.Username = "tester"
	u.OldPassword = "tester"
	u.Password = "tester2"
	u.Role = "admin"
	u.CustomerID = 2

	res := m.UpdateUser(&u)
	if res.Success {
		t.Fail()
	}
}
