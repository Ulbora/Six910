package managers

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_CreateLocalStore(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "six910"
	var dbi db.Database = &mydb

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	sdb.DB = dbi
	dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var locacc LocalStoreAdminUser
	locacc.Username = "admin"
	locacc.Password = "admin"

	m := sm.GetNew()
	var sec sdbi.Security
	sdb.MockSecurity = &sec
	sdb.MockStoreCount = 0
	sdb.MockAddStoreSuccess = true
	sdb.MockStoreID = 5
	sdb.MockAddSecuritySuccess = true
	sdb.MockSecurityID = 4
	sdb.MockAddLocalAccountSuccess = true
	resp := m.CreateLocalStore(&locacc)
	fmt.Println("resp: ", resp)

	if !resp.Success {
		t.Fail()
	}

	// dsuc := sdb.DeleteStore(resp.StoreID)
	// fmt.Println("delete suc: ", dsuc)
	// if !dsuc {
	// 	t.Fail()
	// }

	// dsec := sdb.DeleteSecurity()
	// if !dsec {
	// 	t.Fail()
	// }

	// dbi.Close()
}

func TestSix910Manager_CreateLocalStoreStoreExists(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "six910"
	var dbi db.Database = &mydb

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	sdb.DB = dbi
	dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var locacc LocalStoreAdminUser
	locacc.Username = "admin"
	locacc.Password = "admin"

	m := sm.GetNew()
	var sec sdbi.Security
	sdb.MockSecurity = &sec
	sdb.MockStoreCount = 1
	sdb.MockAddStoreSuccess = true
	sdb.MockStoreID = 5
	sdb.MockAddSecuritySuccess = true
	sdb.MockSecurityID = 4
	sdb.MockAddLocalAccountSuccess = true
	resp := m.CreateLocalStore(&locacc)
	fmt.Println("resp: ", resp)

	if resp.Success {
		t.Fail()
	}

	// dsuc := sdb.DeleteStore(resp.StoreID)
	// fmt.Println("delete suc: ", dsuc)
	// if !dsuc {
	// 	t.Fail()
	// }

	// dsec := sdb.DeleteSecurity()
	// if !dsec {
	// 	t.Fail()
	// }

	// dbi.Close()
}

func TestSix910Manager_CreateLocalStoreAddStoreFail(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "six910"
	var dbi db.Database = &mydb

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	sdb.DB = dbi
	dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var locacc LocalStoreAdminUser
	locacc.Username = "admin"
	locacc.Password = "admin"

	m := sm.GetNew()
	var sec sdbi.Security
	sdb.MockSecurity = &sec
	sdb.MockStoreCount = 0
	//sdb.MockAddStoreSuccess = true
	//sdb.MockStoreID = 5
	sdb.MockAddSecuritySuccess = true
	sdb.MockSecurityID = 4
	sdb.MockAddLocalAccountSuccess = true
	resp := m.CreateLocalStore(&locacc)
	fmt.Println("resp: ", resp)

	if resp.Success {
		t.Fail()
	}

	// dsuc := sdb.DeleteStore(resp.StoreID)
	// fmt.Println("delete suc: ", dsuc)
	// if !dsuc {
	// 	t.Fail()
	// }

	// dsec := sdb.DeleteSecurity()
	// if !dsec {
	// 	t.Fail()
	// }

	// dbi.Close()
}

func TestSix910Manager_CreateLocalStoreSecFail(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "six910"
	var dbi db.Database = &mydb

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	sdb.DB = dbi
	dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var locacc LocalStoreAdminUser
	locacc.Username = "admin"
	locacc.Password = "admin"

	m := sm.GetNew()
	var sec sdbi.Security
	sdb.MockSecurity = &sec
	sdb.MockStoreCount = 0
	sdb.MockAddStoreSuccess = true
	sdb.MockStoreID = 5
	//sdb.MockAddSecuritySuccess = true
	//sdb.MockSecurityID = 4
	sdb.MockAddLocalAccountSuccess = true
	resp := m.CreateLocalStore(&locacc)
	fmt.Println("resp: ", resp)

	if resp.Success {
		t.Fail()
	}

	// dsuc := sdb.DeleteStore(resp.StoreID)
	// fmt.Println("delete suc: ", dsuc)
	// if !dsuc {
	// 	t.Fail()
	// }

	// dsec := sdb.DeleteSecurity()
	// if !dsec {
	// 	t.Fail()
	// }

	// dbi.Close()
}

func TestSix910Manager_CreateLocalStoreExistingSecLocalFail(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "six910"
	var dbi db.Database = &mydb

	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	sdb.DB = dbi
	dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var locacc LocalStoreAdminUser
	locacc.Username = "admin"
	locacc.Password = "admin"

	m := sm.GetNew()
	var sec sdbi.Security
	sec.ID = 1
	sdb.MockSecurity = &sec
	sdb.MockStoreCount = 0
	sdb.MockAddStoreSuccess = true
	sdb.MockStoreID = 5
	sdb.MockAddSecuritySuccess = true
	sdb.MockSecurityID = 4
	//sdb.MockAddLocalAccountSuccess = true
	resp := m.CreateLocalStore(&locacc)
	fmt.Println("resp: ", resp)

	if resp.Success {
		t.Fail()
	}

	// dsuc := sdb.DeleteStore(resp.StoreID)
	// fmt.Println("delete suc: ", dsuc)
	// if !dsuc {
	// 	t.Fail()
	// }

	// dsec := sdb.DeleteSecurity()
	// if !dsec {
	// 	t.Fail()
	// }

	// dbi.Close()
}
