package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddProductCategory(t *testing.T) {
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

	var c sdbi.Category
	c.StoreID = 5

	var p sdbi.Product
	p.StoreID = 5

	sdb.MockProduct = &p
	sdb.MockCategory = &c

	var pc sdbi.ProductCategory
	pc.CategoryID = 1
	pc.ProductID = 2

	sdb.MockAddProductCategorySuccess = true

	res := m.AddProductCategory(&pc, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddProductCategoryFail1(t *testing.T) {
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

	var c sdbi.Category
	c.StoreID = 5

	var p sdbi.Product
	p.StoreID = 5

	sdb.MockProduct = &p
	sdb.MockCategory = &c

	var pc sdbi.ProductCategory
	pc.CategoryID = 1
	pc.ProductID = 2

	//sdb.MockAddProductCategorySuccess = true

	res := m.AddProductCategory(&pc, 5)
	if res.Success || res.Code != 400 {
		t.Fail()
	}
}

func TestSix910Manager_AddProductCategoryFail2(t *testing.T) {
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

	var c sdbi.Category
	c.StoreID = 55

	var p sdbi.Product
	p.StoreID = 5

	sdb.MockProduct = &p
	sdb.MockCategory = &c

	var pc sdbi.ProductCategory
	pc.CategoryID = 1
	pc.ProductID = 2

	sdb.MockAddProductCategorySuccess = true

	res := m.AddProductCategory(&pc, 5)
	if res.Success || res.Code != 500 {
		t.Fail()
	}
}

func TestSix910Manager_AddProductCategoryFail3(t *testing.T) {
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

	var c sdbi.Category
	c.StoreID = 5

	var p sdbi.Product
	p.StoreID = 55

	sdb.MockProduct = &p
	sdb.MockCategory = &c

	var pc sdbi.ProductCategory
	pc.CategoryID = 1
	pc.ProductID = 2

	sdb.MockAddProductCategorySuccess = true

	res := m.AddProductCategory(&pc, 5)
	if res.Success || res.Code != 500 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteProductCategory(t *testing.T) {
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

	var c sdbi.Category
	c.StoreID = 5

	var p sdbi.Product
	p.StoreID = 5

	sdb.MockProduct = &p
	sdb.MockCategory = &c

	var pc sdbi.ProductCategory
	pc.CategoryID = 1
	pc.ProductID = 2

	sdb.MockDeleteProductCategorySuccess = true

	res := m.DeleteProductCategory(&pc, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteProductCategoryFail1(t *testing.T) {
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

	var c sdbi.Category
	c.StoreID = 5

	var p sdbi.Product
	p.StoreID = 5

	sdb.MockProduct = &p
	sdb.MockCategory = &c

	var pc sdbi.ProductCategory
	pc.CategoryID = 1
	pc.ProductID = 2

	//sdb.MockDeleteProductCategorySuccess = true

	res := m.DeleteProductCategory(&pc, 5)
	if res.Success || res.Code != 400 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteProductCategoryFail2(t *testing.T) {
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

	var c sdbi.Category
	c.StoreID = 55

	var p sdbi.Product
	p.StoreID = 5

	sdb.MockProduct = &p
	sdb.MockCategory = &c

	var pc sdbi.ProductCategory
	pc.CategoryID = 1
	pc.ProductID = 2

	sdb.MockDeleteProductCategorySuccess = true

	res := m.DeleteProductCategory(&pc, 5)
	if res.Success || res.Code != 500 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteProductCategoryFail3(t *testing.T) {
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

	var c sdbi.Category
	c.StoreID = 5

	var p sdbi.Product
	p.StoreID = 55

	sdb.MockProduct = &p
	sdb.MockCategory = &c

	var pc sdbi.ProductCategory
	pc.CategoryID = 1
	pc.ProductID = 2

	sdb.MockDeleteProductCategorySuccess = true

	res := m.DeleteProductCategory(&pc, 5)
	if res.Success || res.Code != 500 {
		t.Fail()
	}
}
