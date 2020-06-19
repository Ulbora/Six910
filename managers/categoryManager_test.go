package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddCategory(t *testing.T) {
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

	var cat sdbi.Category
	cat.Description = "stuff"

	sdb.MockAddCategorySuccess = true
	sdb.MockCategoryID = 3
	res := m.AddCategory(&cat)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddCategoryFail(t *testing.T) {
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

	var cat sdbi.Category
	cat.Description = "stuff"

	//sdb.MockAddCategorySuccess = true
	sdb.MockCategoryID = 3
	res := m.AddCategory(&cat)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCategory(t *testing.T) {
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

	var cat sdbi.Category
	cat.Description = "stuff"
	cat.StoreID = 6

	sdb.MockCategory = &cat
	sdb.MockUpdateCategorySuccess = true

	res := m.UpdateCategory(&cat)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCategoryFail1(t *testing.T) {
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

	var cat sdbi.Category
	cat.Description = "stuff"
	cat.StoreID = 6

	sdb.MockCategory = &cat
	sdb.MockUpdateCategorySuccess = true

	var cat2 sdbi.Category
	cat2.Description = "stuff"
	cat2.StoreID = 66

	res := m.UpdateCategory(&cat2)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCategoryFail(t *testing.T) {
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

	var cat sdbi.Category
	cat.Description = "stuff"
	cat.StoreID = 6

	sdb.MockCategory = &cat
	//sdb.MockUpdateCategorySuccess = true

	res := m.UpdateCategory(&cat)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetCategory(t *testing.T) {
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

	var cat sdbi.Category
	cat.Description = "stuff"
	cat.StoreID = 6

	sdb.MockCategory = &cat
	//sdb.MockUpdateCategorySuccess = true

	fcat := m.GetCategory(3, 6)
	if fcat.Description != cat.Description {
		t.Fail()
	}
}

func TestSix910Manager_GetCategoryFail(t *testing.T) {
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

	var cat sdbi.Category
	cat.Description = "stuff"
	cat.StoreID = 66

	sdb.MockCategory = &cat
	//sdb.MockUpdateCategorySuccess = true

	fcat := m.GetCategory(3, 6)
	if fcat.Description == cat.Description {
		t.Fail()
	}
}

func TestSix910Manager_GetCategoryList(t *testing.T) {
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

	var cat sdbi.Category
	cat.Description = "stuff"
	cat.StoreID = 6

	sdb.MockCategory = &cat
	var catlst []sdbi.Category
	catlst = append(catlst, cat)
	sdb.MockCategoryList = &catlst
	//sdb.MockUpdateCategorySuccess = true

	flst := m.GetCategoryList(6)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetSubCategoryList(t *testing.T) {
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

	var cat sdbi.Category
	cat.Description = "stuff"
	cat.StoreID = 6

	sdb.MockCategory = &cat
	var catlst []sdbi.Category
	catlst = append(catlst, cat)
	sdb.MockCategoryList = &catlst
	//sdb.MockUpdateCategorySuccess = true

	flst := m.GetSubCategoryList(64)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteCategory(t *testing.T) {
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

	var cat sdbi.Category
	cat.Description = "stuff"
	cat.StoreID = 6

	sdb.MockCategory = &cat
	sdb.MockDeleteCategorySuccess = true
	//sdb.MockUpdateCategorySuccess = true

	res := m.DeleteCategory(4, 6)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteCategoryFail(t *testing.T) {
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

	var cat sdbi.Category
	cat.Description = "stuff"
	cat.StoreID = 66

	sdb.MockCategory = &cat
	sdb.MockDeleteCategorySuccess = true
	//sdb.MockUpdateCategorySuccess = true

	res := m.DeleteCategory(4, 6)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteCategoryFail2(t *testing.T) {
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

	var cat sdbi.Category
	cat.Description = "stuff"
	cat.StoreID = 6

	sdb.MockCategory = &cat
	//sdb.MockDeleteCategorySuccess = true
	//sdb.MockUpdateCategorySuccess = true

	res := m.DeleteCategory(4, 6)
	if res.Success {
		t.Fail()
	}
}
