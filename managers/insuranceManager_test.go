package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddInsurance(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.Cost = 10.00

	sdb.MockAddInsuranceSuccess = true
	sdb.MockInsuranceID = 4

	res := m.AddInsurance(&ins)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddInsuranceFail(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.Cost = 10.00

	sdb.MockAddInsuranceSuccess = true
	//sdb.MockInsuranceID = 4

	res := m.AddInsurance(&ins)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateInsurance(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.Cost = 10.00
	ins.StoreID = 6

	sdb.MockInsurance = &ins
	sdb.MockUpdateInsuranceSuccess = true

	res := m.UpdateInsurance(&ins)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateInsuranceFail(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.Cost = 10.00
	ins.StoreID = 6

	sdb.MockInsurance = &ins
	sdb.MockUpdateInsuranceSuccess = true

	var ins2 sdbi.Insurance
	ins2.Cost = 10.00
	ins2.StoreID = 66

	res := m.UpdateInsurance(&ins2)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateInsuranceFail2(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.Cost = 10.00
	ins.StoreID = 6

	sdb.MockInsurance = &ins
	//sdb.MockUpdateInsuranceSuccess = true

	res := m.UpdateInsurance(&ins)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetInsurance(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.Cost = 10.00
	ins.StoreID = 6

	sdb.MockInsurance = &ins

	fins := m.GetInsurance(4, 6)
	if fins.Cost != ins.Cost {
		t.Fail()
	}
}

func TestSix910Manager_GetInsuranceFail(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.Cost = 10.00
	ins.StoreID = 66

	sdb.MockInsurance = &ins

	fins := m.GetInsurance(4, 6)
	if fins.Cost == ins.Cost {
		t.Fail()
	}
}

func TestSix910Manager_GetInsuranceList(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.Cost = 10.00
	ins.StoreID = 6

	var ilst []sdbi.Insurance
	ilst = append(ilst, ins)

	sdb.MockInsuranceList = &ilst

	flst := m.GetInsuranceList(6)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteInsurance(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.Cost = 10.00
	ins.StoreID = 6

	sdb.MockInsurance = &ins

	sdb.MockDeleteInsuranceSuccess = true

	res := m.DeleteInsurance(4, 6)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteInsuranceFail(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.Cost = 10.00
	ins.StoreID = 66

	sdb.MockInsurance = &ins

	sdb.MockDeleteInsuranceSuccess = true

	res := m.DeleteInsurance(4, 6)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteInsuranceFail2(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.Cost = 10.00
	ins.StoreID = 6

	sdb.MockInsurance = &ins

	//sdb.MockDeleteInsuranceSuccess = true

	res := m.DeleteInsurance(4, 6)
	if res.Success {
		t.Fail()
	}
}
