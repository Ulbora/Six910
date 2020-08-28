package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddTaxRate(t *testing.T) {
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

	var cat sdbi.TaxRate
	cat.Country = "usa"

	sdb.MockAddTaxRateSuccess = true
	sdb.MockTaxRateID = 3
	res := m.AddTaxRate(&cat)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddTaxRateFail(t *testing.T) {
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

	var cat sdbi.TaxRate
	cat.Country = "usa"

	//sdb.MockAddTaxRateSuccess = true
	sdb.MockTaxRateID = 3
	res := m.AddTaxRate(&cat)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateTaxRate(t *testing.T) {
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

	var cat sdbi.TaxRate
	cat.Country = "usa"
	cat.StoreID = 6

	sdb.MockTaxRate = &cat
	sdb.MockUpdateTaxRateSuccess = true

	res := m.UpdateTaxRate(&cat)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateTaxRateFail(t *testing.T) {
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

	var cat sdbi.TaxRate
	cat.Country = "usa"
	cat.StoreID = 6

	sdb.MockTaxRate = &cat
	//sdb.MockUpdateTaxRateSuccess = true

	res := m.UpdateTaxRate(&cat)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetTaxRate(t *testing.T) {
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

	var cat sdbi.TaxRate
	cat.Country = "USA"
	cat.StoreID = 6

	var catlst []sdbi.TaxRate
	catlst = append(catlst, cat)
	sdb.MockTaxRateList = &catlst
	//sdb.MockUpdateCategorySuccess = true

	flst := m.GetTaxRate("usa", "GA", 6)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetTaxRateList(t *testing.T) {
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

	var cat sdbi.TaxRate
	cat.Country = "USA"
	cat.StoreID = 6

	var catlst []sdbi.TaxRate
	catlst = append(catlst, cat)
	sdb.MockTaxRateList = &catlst
	//sdb.MockUpdateCategorySuccess = true

	flst := m.GetTaxRateList(6)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteTaxRate(t *testing.T) {
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

	var cat sdbi.TaxRate
	cat.ID = 4
	cat.Country = "USA"
	cat.StoreID = 6

	var catlst []sdbi.TaxRate
	catlst = append(catlst, cat)
	sdb.MockTaxRateList = &catlst
	sdb.MockDeleteTaxRateSuccess = true
	//sdb.MockUpdateCategorySuccess = true

	res := m.DeleteTaxRate(4, 6)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteTaxRateFail(t *testing.T) {
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

	var cat sdbi.TaxRate
	cat.ID = 4
	cat.Country = "USA"
	cat.StoreID = 6

	var catlst []sdbi.TaxRate
	catlst = append(catlst, cat)
	sdb.MockTaxRateList = &catlst
	//sdb.MockDeleteTaxRateSuccess = true
	//sdb.MockUpdateCategorySuccess = true

	res := m.DeleteTaxRate(4, 6)
	if res.Success {
		t.Fail()
	}
}
