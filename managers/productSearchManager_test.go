package managers

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_GetProductManufacturerListByProductName(t *testing.T) {
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

	var p sdbi.Product
	p.Color = "blue"
	p.StoreID = 4

	var plst []sdbi.Product
	plst = append(plst, p)

	sdb.MockProductList = &plst

	sdb.MockManufacturerList = &[]string{"test\\\\1\\", "test2"}

	flst3 := m.GetProductManufacturerListByProductName("test", 4)
	fmt.Println("flst3: ", flst3)
	if len(*flst3) != 2 {
		t.Fail()
	}
}

func TestSix910Manager_GetProductByNameAndManufacturerName(t *testing.T) {
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

	var p sdbi.Product
	p.Color = "blue"
	p.StoreID = 4

	var plst []sdbi.Product
	plst = append(plst, p)

	sdb.MockProductList = &plst

	sdb.MockManufacturerList = &[]string{"test1", "test2"}

	flst3 := m.GetProductByNameAndManufacturerName("test", "test", 4, 0, 2)
	if len(*flst3) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetProductManufacturerListByCatID(t *testing.T) {
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

	var p sdbi.Product
	p.Color = "blue"
	p.StoreID = 4

	var plst []sdbi.Product
	plst = append(plst, p)

	sdb.MockProductList = &plst

	sdb.MockManufacturerList = &[]string{"test\\\\1", "test2"}

	flst3 := m.GetProductManufacturerListByCatID(5, 4)
	fmt.Println("flst3: ", flst3)
	if len(*flst3) != 2 {
		t.Fail()
	}
}

func TestSix910Manager_GetProductByCatAndManufacturer(t *testing.T) {
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

	var p sdbi.Product
	p.Color = "blue"
	p.StoreID = 4

	var plst []sdbi.Product
	plst = append(plst, p)

	sdb.MockProductList = &plst

	sdb.MockManufacturerList = &[]string{"test1", "test2"}

	flst3 := m.GetProductByCatAndManufacturer(5, "test", 4, 0, 2)
	if len(*flst3) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_ProductSearch(t *testing.T) {
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

	var p sdbi.Product
	p.Color = "blue"
	p.StoreID = 4

	var plst []sdbi.Product
	plst = append(plst, p)

	sdb.MockProductSearchList = &plst

	var ps sdbi.ProductSearch
	satb := []string{"test"}
	ps.DescAttributes = &satb

	flst3 := m.ProductSearch(&ps)
	if len(*flst3) != 1 {
		t.Fail()
	}
}


func TestSix910Manager_ProductSearchNoImput(t *testing.T) {
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

	var p sdbi.Product
	p.Color = "blue"
	p.StoreID = 4

	var plst []sdbi.Product
	plst = append(plst, p)

	sdb.MockProductSearchList = &plst

	var ps sdbi.ProductSearch
	//satb := []string{"test"}
	//ps.DescAttributes = &satb

	flst3 := m.ProductSearch(&ps)
	if len(*flst3) != 0 {
		t.Fail()
	}
}


func TestSix910Manager_GetProductManufacturerListByProductSearch(t *testing.T) {
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

	var p sdbi.Product
	p.Color = "blue"
	p.StoreID = 4

	var plst []sdbi.Product
	plst = append(plst, p)

	sdb.MockProductList = &plst

	sdb.MockManufacturerList = &[]string{"test\\\\1\\", "test2"}

	flst3 := m.GetProductManufacturerListByProductSearch("test thing", 4)
	fmt.Println("flst3: ", flst3)
	if len(*flst3) != 2 {
		t.Fail()
	}
}