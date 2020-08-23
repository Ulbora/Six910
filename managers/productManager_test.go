package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddProduct(t *testing.T) {
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

	sdb.MockAddProductSuccess = true
	sdb.MockProductID = 7

	res := m.AddProduct(&p)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddProductFail(t *testing.T) {
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

	sdb.MockAddProductSuccess = true
	//sdb.MockProductID = 7

	res := m.AddProduct(&p)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateProduct(t *testing.T) {
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

	sdb.MockProduct = &p

	sdb.MockUpdateProductSuccess = true

	res := m.UpdateProduct(&p)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateProductFail(t *testing.T) {
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

	sdb.MockProduct = &p

	sdb.MockUpdateProductSuccess = true

	var p2 sdbi.Product
	p2.Color = "blue"
	p2.StoreID = 48

	res := m.UpdateProduct(&p2)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateProductFail2(t *testing.T) {
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

	sdb.MockProduct = &p

	//sdb.MockUpdateProductSuccess = true

	res := m.UpdateProduct(&p)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetProductByID(t *testing.T) {
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

	sdb.MockProduct = &p

	fp := m.GetProductByID(4, 4)
	if fp.Color != p.Color {
		t.Fail()
	}
}

func TestSix910Manager_GetProductByIDFail(t *testing.T) {
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
	p.StoreID = 44

	sdb.MockProduct = &p

	fp := m.GetProductByID(4, 4)
	if fp.Color == p.Color {
		t.Fail()
	}
}

func TestSix910Manager_GetProductsByPromoted(t *testing.T) {
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

	flst := m.GetProductsByPromoted(4, 0, 10)
	if len(*flst) != 1 {
		t.Fail()
	}

}

func TestSix910Manager_GetProductsByName(t *testing.T) {
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

	flst := m.GetProductsByName("test", 4, 0, 10)
	if len(*flst) != 1 {
		t.Fail()
	}

}

func TestSix910Manager_GetProductsByCaterory(t *testing.T) {
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

	var cat sdbi.Category
	cat.StoreID = 4

	sdb.MockCategory = &cat

	flst2 := m.GetProductsByCaterory(2, 4, 0, 5)
	if len(*flst2) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetProductsByCateroryFail(t *testing.T) {
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

	var cat sdbi.Category
	cat.StoreID = 44

	sdb.MockCategory = &cat

	flst2 := m.GetProductsByCaterory(2, 4, 0, 5)
	if len(*flst2) != 0 {
		t.Fail()
	}
}

func TestSix910Manager_GetProductList(t *testing.T) {
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

	flst3 := m.GetProductList(4, 0, 10)
	if len(*flst3) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteProduct(t *testing.T) {
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

	sdb.MockProduct = &p

	sdb.MockDeleteProductSuccess = true

	res := m.DeleteProduct(5, 4)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteProductFail(t *testing.T) {
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
	p.StoreID = 44

	sdb.MockProduct = &p

	sdb.MockDeleteProductSuccess = true

	res := m.DeleteProduct(5, 4)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetProductByBySku(t *testing.T) {
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

	sdb.MockProduct = &p

	fp := m.GetProductByBySku("345", 4, 4)
	if fp.Color != p.Color {
		t.Fail()
	}
}

func TestSix910Manager_GetProductByBySkuFail(t *testing.T) {
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
	p.StoreID = 44

	sdb.MockProduct = &p

	fp := m.GetProductByBySku("345", 4, 4)
	if fp.Color == p.Color {
		t.Fail()
	}
}
