package managers

import (
	"fmt"

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
	lc.StoreID = 1
	//_, hpw := sm.hashPassword("tester")
	//lc.Password = hpw
	lc.Password = "$2a$10$jQeE.2i3EuggHY0ckMT7OuRW5zZeo.M2OTpNEU92skDVKoEmugBqe"
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
	lc.StoreID = 1
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

func TestSix910Manager_AdminUpdateUser(t *testing.T) {
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
	lc.StoreID = 1
	//_, hpw := sm.hashPassword("tester")
	//lc.Password = hpw
	lc.Password = "$2a$10$jQeE.2i3EuggHY0ckMT7OuRW5zZeo.M2OTpNEU92skDVKoEmugBqe"
	sdb.MockLocalAccount = &lc

	var u User
	u.Enabled = true
	u.StoreID = 1
	u.Username = "tester"
	u.OldPassword = "tester"
	u.Password = "tester2"
	u.Role = "admin"
	u.CustomerID = 2

	res := m.AdminUpdateUser(&u)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AdminUpdateUserFail(t *testing.T) {
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
	lc.StoreID = 1
	//_, hpw := sm.hashPassword("tester")
	//lc.Password = hpw
	lc.Password = "$2a$10$jQeE.2i3EuggHY0ckMT7OuRW5zZeo.M2OTpNEU92skDVKoEmugBqe"
	sdb.MockLocalAccount = &lc

	var u User
	u.Enabled = true
	u.StoreID = 1
	u.Username = "tester"
	u.OldPassword = "tester"
	u.Password = "tester2"
	u.Role = "admin"
	u.CustomerID = 2

	res := m.AdminUpdateUser(&u)
	if res.Success {
		t.Fail()
	}
}
func TestSix910Manager_GetUser(t *testing.T) {
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
	lc.UserName = "tester"
	lc.Role = customerRole
	lc.Enabled = true
	lc.StoreID = 5
	sdb.MockLocalAccount = &lc

	var u User
	u.Enabled = true
	u.StoreID = 1
	u.Username = "tester"
	u.OldPassword = "tester"
	u.Password = "tester2"
	u.Role = "admin"
	u.CustomerID = 2

	fu := m.GetUser(&u)
	if fu.Username != u.Username {
		t.Fail()
	}
}

func TestSix910Manager_GetAdminUsers(t *testing.T) {
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
	var lclist []sdbi.LocalAccount
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	_, hpw := sm.hashPassword("tester")
	lc.Password = hpw
	lc.UserName = "tester"
	lc.Role = customerRole
	lc.Enabled = true
	lc.StoreID = 5
	lclist = append(lclist, lc)

	var lc2 sdbi.LocalAccount
	//lc2.CustomerID = 0
	_, hpw2 := sm.hashPassword("tester2")
	lc2.Password = hpw2
	lc2.UserName = "tester2"
	lc2.Role = storeAdmin
	lc2.Enabled = true
	lc2.StoreID = 5
	lclist = append(lclist, lc2)

	sdb.MockLocalAccountList = &lclist

	fu := m.GetAdminUsers(5)
	fmt.Println("found admin users: ", fu)
	if len(*fu) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetCustomerUsers(t *testing.T) {
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
	var lclist []sdbi.LocalAccount
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	_, hpw := sm.hashPassword("tester")
	lc.Password = hpw
	lc.UserName = "tester"
	lc.Role = customerRole
	lc.Enabled = true
	lc.StoreID = 5
	lclist = append(lclist, lc)

	var lc2 sdbi.LocalAccount
	//lc2.CustomerID = 0
	_, hpw2 := sm.hashPassword("tester2")
	lc2.Password = hpw2
	lc2.UserName = "tester2"
	lc2.Role = storeAdmin
	lc2.Enabled = true
	lc2.StoreID = 5
	lclist = append(lclist, lc2)

	sdb.MockLocalAccountList = &lclist

	fu := m.GetCustomerUsers(5)
	fmt.Println("found customer users: ", fu)
	if len(*fu) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetUsersByCustomer(t *testing.T) {
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
	var lclist []sdbi.LocalAccount
	var lc sdbi.LocalAccount
	lc.CustomerID = 2
	_, hpw := sm.hashPassword("tester")
	lc.Password = hpw
	lc.UserName = "tester"
	lc.Role = customerRole
	lc.Enabled = true
	lc.StoreID = 5
	lclist = append(lclist, lc)

	var lc2 sdbi.LocalAccount
	//lc2.CustomerID = 0
	_, hpw2 := sm.hashPassword("tester2")
	lc2.Password = hpw2
	lc2.UserName = "tester2"
	lc2.Role = storeAdmin
	lc2.Enabled = true
	lc2.StoreID = 5
	lclist = append(lclist, lc2)

	sdb.MockLocalAccountList = &lclist

	fu := m.GetUsersByCustomer(2, 5)
	fmt.Println("found users by customer: ", fu)
	if len(*fu) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_ValidateUser(t *testing.T) {
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
	lc.UserName = "tester"
	lc.Role = customerRole
	lc.Enabled = true
	lc.StoreID = 5
	sdb.MockLocalAccount = &lc

	var u User
	u.Enabled = true
	u.StoreID = 1
	u.Username = "tester"
	//u.OldPassword = "tester"
	u.Password = "tester"
	u.Role = customerRole
	u.CustomerID = 2

	res := m.ValidateUser(&u)
	fmt.Println("validate user: ", res)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_ValidateUserFail(t *testing.T) {
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
	lc.UserName = "tester"
	lc.Role = customerRole
	lc.Enabled = true
	lc.StoreID = 5
	sdb.MockLocalAccount = &lc

	var u User
	u.Enabled = true
	u.StoreID = 1
	u.Username = "tester2"
	//u.OldPassword = "tester"
	u.Password = "tester"
	u.Role = customerRole
	u.CustomerID = 2

	res := m.ValidateUser(&u)
	fmt.Println("validate user: ", res)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_ValidateUserFail2(t *testing.T) {
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
	lc.UserName = "tester"
	lc.Role = customerRole
	lc.Enabled = true
	lc.StoreID = 5
	sdb.MockLocalAccount = &lc

	var u User
	u.Enabled = true
	u.StoreID = 1
	u.Username = "tester"
	//u.OldPassword = "tester"
	u.Password = "tester2"
	u.Role = customerRole
	u.CustomerID = 2

	res := m.ValidateUser(&u)
	fmt.Println("validate user: ", res)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_generateNewPassword(t *testing.T) {
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	var sm Six910Manager
	sm.Log = &l

	pw := sm.generateNewPassword()
	fmt.Println("new pw: ", pw)
	if pw == "" {
		t.Fail()
	}
}

func TestSix910Manager_ResetCustomerPassword(t *testing.T) {
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
	lc.StoreID = 1
	//_, hpw := sm.hashPassword("tester")
	//lc.Password = hpw
	lc.Enabled = true
	lc.UserName = "tester"
	lc.Password = "$2a$10$jQeE.2i3EuggHY0ckMT7OuRW5zZeo.M2OTpNEU92skDVKoEmugBqe"
	sdb.MockLocalAccount = &lc

	var u User
	u.Enabled = true
	u.StoreID = 1
	u.Username = "tester"
	u.OldPassword = "tester"
	u.Password = "tester2"
	u.Role = "admin"
	u.CustomerID = 2

	res := m.ResetCustomerPassword(&u)
	fmt.Println("res: ", *res)
	if !res.Success {
		t.Fail()
	}
}
