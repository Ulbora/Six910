package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddDistributor(t *testing.T) {
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

	var dis sdbi.Distributor
	dis.Company = "tester"
	dis.StoreID = 4

	sdb.MockAddDistributorSuccess = true
	sdb.MockDistributorID = 5

	res := m.AddDistributor(&dis)
	if !res.Success || res.ID != 5 {
		t.Fail()
	}
}

func TestSix910Manager_AddDistributorFail(t *testing.T) {
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

	var dis sdbi.Distributor
	dis.Company = "tester"
	dis.StoreID = 4

	//sdb.MockAddDistributorSuccess = true
	//sdb.MockDistributorID = 5

	res := m.AddDistributor(&dis)
	if res.Success || res.ID == 5 {
		t.Fail()
	}
}

func TestSix910Manager_UpdateDistributor(t *testing.T) {
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

	var dis sdbi.Distributor
	dis.ID = 2
	dis.Company = "tester"
	dis.StoreID = 4

	sdb.MockDistributor = &dis

	var dis2 sdbi.Distributor
	dis2.ID = 2
	dis2.Company = "tester"
	dis2.StoreID = 4

	sdb.MockUpdateDistributorSuccess = true

	res := m.UpdateDistributor(&dis2)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateDistributorFail2(t *testing.T) {
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

	var dis sdbi.Distributor
	dis.ID = 2
	dis.Company = "tester"
	dis.StoreID = 44

	sdb.MockDistributor = &dis

	var dis2 sdbi.Distributor
	dis2.ID = 2
	dis2.Company = "tester"
	dis2.StoreID = 4

	sdb.MockUpdateDistributorSuccess = true

	res := m.UpdateDistributor(&dis2)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateDistributorfail(t *testing.T) {
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

	var dis sdbi.Distributor
	dis.ID = 2
	dis.Company = "tester"
	dis.StoreID = 4

	sdb.MockDistributor = &dis

	var dis2 sdbi.Distributor
	dis2.ID = 2
	dis2.Company = "tester"
	dis2.StoreID = 4

	//sdb.MockUpdateDistributorSuccess = true

	res := m.UpdateDistributor(&dis2)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetDistributor(t *testing.T) {
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

	var dis sdbi.Distributor
	dis.ID = 2
	dis.Company = "tester"
	dis.StoreID = 4

	sdb.MockDistributor = &dis

	fdis := m.GetDistributor(2, 4)
	if fdis.ID != 2 {
		t.Fail()
	}

}

func TestSix910Manager_GetDistributorfail(t *testing.T) {
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

	var dis sdbi.Distributor
	dis.ID = 2
	dis.Company = "tester"
	dis.StoreID = 44

	sdb.MockDistributor = &dis

	fdis := m.GetDistributor(2, 4)
	if fdis.ID == 2 {
		t.Fail()
	}

}

func TestSix910Manager_GetDistributorList(t *testing.T) {
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

	var dis sdbi.Distributor
	dis.ID = 2
	dis.Company = "tester"
	dis.StoreID = 4

	var dlst []sdbi.Distributor
	dlst = append(dlst, dis)

	sdb.MockDistributorList = &dlst

	fdlst := m.GetDistributorList(4)
	if len(*fdlst) != 1 {
		t.Fail()
	}

}

func TestSix910Manager_DeleteDistributor(t *testing.T) {
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

	var dis sdbi.Distributor
	dis.ID = 2
	dis.Company = "tester"
	dis.StoreID = 4

	sdb.MockDistributor = &dis
	sdb.MockDeleteDistributorSuccess = true

	res := m.DeleteDistributor(2, 4)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteDistributorFail2(t *testing.T) {
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

	var dis sdbi.Distributor
	dis.ID = 2
	dis.Company = "tester"
	dis.StoreID = 44

	sdb.MockDistributor = &dis
	sdb.MockDeleteDistributorSuccess = true

	res := m.DeleteDistributor(2, 4)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteDistributorFail(t *testing.T) {
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

	var dis sdbi.Distributor
	dis.ID = 2
	dis.Company = "tester"
	dis.StoreID = 4

	sdb.MockDistributor = &dis
	//sdb.MockDeleteDistributorSuccess = true

	res := m.DeleteDistributor(2, 4)
	if res.Success {
		t.Fail()
	}
}
