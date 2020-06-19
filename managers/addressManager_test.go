package managers

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddAddress(t *testing.T) {
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
	cus.StoreID = 5

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.CustomerID = 6
	sdb.MockAddAddressSuccess = true
	sdb.MockAddressID = 3
	res := m.AddAddress(&add, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddAddressFail1(t *testing.T) {
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
	cus.StoreID = 55

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.CustomerID = 6
	sdb.MockAddAddressSuccess = true
	sdb.MockAddressID = 3
	res := m.AddAddress(&add, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddAddressFail2(t *testing.T) {
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
	cus.StoreID = 5

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.CustomerID = 6
	//sdb.MockAddAddressSuccess = true
	sdb.MockAddressID = 3
	res := m.AddAddress(&add, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateAddress(t *testing.T) {
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
	cus.StoreID = 5

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.CustomerID = 6
	sdb.MockUpdateAddressSuccess = true

	res := m.UpdateAddress(&add, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateAddressFail1(t *testing.T) {
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
	cus.StoreID = 55

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.CustomerID = 6
	sdb.MockUpdateAddressSuccess = true

	res := m.UpdateAddress(&add, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateAddressFail2(t *testing.T) {
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
	cus.StoreID = 5

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.CustomerID = 6
	//sdb.MockUpdateAddressSuccess = true

	res := m.UpdateAddress(&add, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetAddress(t *testing.T) {
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
	cus.ID = 6
	cus.StoreID = 5

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.ID = 2
	add.CustomerID = 6
	add.City = "Atlanta"

	sdb.MockAddress = &add

	fadd := m.GetAddress(2, 6, 5)
	if fadd.City != add.City {
		t.Fail()
	}
}

func TestSix910Manager_GetAddressfail(t *testing.T) {
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
	cus.ID = 6
	cus.StoreID = 5

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.ID = 2
	add.CustomerID = 66
	add.City = "Atlanta"

	sdb.MockAddress = &add

	fadd := m.GetAddress(2, 6, 5)
	fmt.Println("add: ", add)
	fmt.Println("fadd: ", fadd)
	if fadd.City == add.City {
		t.Fail()
	}
}

func TestSix910Manager_GetAddressList(t *testing.T) {
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
	cus.ID = 6
	cus.StoreID = 5

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.ID = 2
	add.CustomerID = 6
	add.City = "Atlanta"

	var addlst []sdbi.Address
	addlst = append(addlst, add)

	sdb.MockAddressList = &addlst

	falst := m.GetAddressList(6, 5)
	if len(*falst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetAddressListFail(t *testing.T) {
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
	cus.ID = 6
	cus.StoreID = 55

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.ID = 2
	add.CustomerID = 6
	add.City = "Atlanta"

	var addlst []sdbi.Address
	addlst = append(addlst, add)

	sdb.MockAddressList = &addlst

	falst := m.GetAddressList(6, 5)
	if len(*falst) != 0 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteAddress(t *testing.T) {
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
	cus.ID = 6
	cus.StoreID = 5

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.ID = 2
	add.CustomerID = 6
	add.City = "Atlanta"

	sdb.MockAddress = &add
	sdb.MockDeleteAddressSuccess = true
	res := m.DeleteAddress(2, 6, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteAddressFail1(t *testing.T) {
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
	cus.ID = 6
	cus.StoreID = 55

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.ID = 2
	add.CustomerID = 6
	add.City = "Atlanta"

	sdb.MockAddress = &add
	sdb.MockDeleteAddressSuccess = true
	res := m.DeleteAddress(2, 6, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteAddressFail2(t *testing.T) {
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
	cus.ID = 6
	cus.StoreID = 5

	sdb.MockCustomer = &cus

	var add sdbi.Address
	add.ID = 2
	add.CustomerID = 6
	add.City = "Atlanta"

	sdb.MockAddress = &add
	//sdb.MockDeleteAddressSuccess = true
	res := m.DeleteAddress(2, 6, 5)
	if res.Success {
		t.Fail()
	}
}
