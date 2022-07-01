package managers

import (
	"fmt"
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

// func TestSix910Manager_AddProduct_fail_subsku(t *testing.T) {
// 	var sdb sixmdb.MockSix910Mysql
// 	var l lg.Logger
// 	l.LogLevel = lg.AllLevel
// 	sdb.Log = &l
// 	//sdb.DB = dbi
// 	//dbi.Connect()

// 	var sm Six910Manager
// 	sm.Db = sdb.GetNew()
// 	sm.Log = &l

// 	m := sm.GetNew()

// 	var p sdbi.Product
// 	p.Color = "blue"
// 	p.StoreID = 4
// 	//p.SubSku = true
// 	//p.ParentProductID = 4

// 	sdb.MockAddProductSuccess = true
// 	sdb.MockProductID = 7

// 	res := m.AddProduct(&p)
// 	if res.Success {
// 		t.Fail()
// 	}
// }

func TestSix910Manager_AddProduct2(t *testing.T) {
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
	//p.SubSku = true
	p.ParentProductID = 5

	sdb.MockAddProductSuccess = true
	sdb.MockProductID = 7

	res := m.AddProduct(&p)
	if !res.Success {
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

func TestSix910Manager_UpdateProduct2(t *testing.T) {
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
	//p.SubSku = true
	p.ParentProductID = 5

	sdb.MockProduct = &p

	sdb.MockUpdateProductSuccess = true

	res := m.UpdateProduct(&p)
	if !res.Success {
		t.Fail()
	}
}

// func TestSix910Manager_UpdateProduct2_fail(t *testing.T) {
// 	var sdb sixmdb.MockSix910Mysql
// 	var l lg.Logger
// 	l.LogLevel = lg.AllLevel
// 	sdb.Log = &l
// 	//sdb.DB = dbi
// 	//dbi.Connect()

// 	var sm Six910Manager
// 	sm.Db = sdb.GetNew()
// 	sm.Log = &l

// 	m := sm.GetNew()

// 	var p sdbi.Product
// 	p.Color = "blue"
// 	p.StoreID = 4
// 	//p.SubSku = true

// 	sdb.MockProduct = &p

// 	sdb.MockUpdateProductSuccess = true

// 	res := m.UpdateProduct(&p)
// 	if res.Success {
// 		t.Fail()
// 	}
// }

func TestSix910Manager_UpdateProductQuantity(t *testing.T) {
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

	sdb.MockUpdateProductQuantitySuccess = true

	res := m.UpdateProductQuantity(&p)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateProductQuantityFail(t *testing.T) {
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

	var p2 sdbi.Product
	p2.Color = "blue"
	p2.StoreID = 42

	sdb.MockUpdateProductQuantitySuccess = true

	res := m.UpdateProductQuantity(&p2)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateProductQuantityFail2(t *testing.T) {
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

	//sdb.MockUpdateProductQuantitySuccess = true

	res := m.UpdateProductQuantity(&p)
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

	var mpl []sdbi.Product
	var sp sdbi.Product
	sp.Cost = 1.10
	mpl = append(mpl, sp)
	fmt.Println("mpl: ", mpl)
	sdb.MockProductSubSkuList = &mpl

	fp := m.GetProductByID(4, 4)
	fmt.Println("product: ", fp)
	if fp.Color != p.Color || (*fp.SubSkuList)[0].Cost != 1.10 {
		t.Fail()
	}
}

func TestSix910Manager_GetProductByIDNoSub(t *testing.T) {
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

	var mpl []sdbi.Product
	// var sp sdbi.Product
	// sp.Cost = 1.10
	// mpl = append(mpl, sp)
	fmt.Println("mpl: ", mpl)
	sdb.MockProductSubSkuList = &mpl

	fp := m.GetProductByID(4, 4)
	fmt.Println("product: ", fp)
	if fp.Color != p.Color || len(*fp.SubSkuList) != 0 {
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

	var mpl []sdbi.Product
	var sp sdbi.Product
	sp.Cost = 1.10
	mpl = append(mpl, sp)
	fmt.Println("mpl: ", mpl)
	sdb.MockProductSubSkuList = &mpl

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

func TestSix910Manager_GetProductSubSkuList(t *testing.T) {
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

	sdb.MockProductSubSkuList = &plst

	flst3 := m.GetProductSubSkuList(2, 4)
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
	p.ID = 1
	p.Color = "blue"
	p.StoreID = 4

	sdb.MockProduct = &p

	var mpl []sdbi.Product
	var sp sdbi.Product
	sp.Cost = 1.10
	mpl = append(mpl, sp)
	fmt.Println("mpl: ", mpl)
	sdb.MockProductSubSkuList = &mpl

	fp := m.GetProductByBySku("345", 4, 4)
	if fp.Color != p.Color || (*fp.SubSkuList)[0].Cost != 1.10 {
		t.Fail()
	}
}

func TestSix910Manager_GetProductByBySkuNoSub(t *testing.T) {
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

	var mpl []sdbi.Product
	// var sp sdbi.Product
	// sp.Cost = 1.10
	// mpl = append(mpl, sp)
	fmt.Println("mpl: ", mpl)
	sdb.MockProductSubSkuList = &mpl

	fp := m.GetProductByBySku("345", 4, 4)
	if fp.Color != p.Color || len(*fp.SubSkuList) != 0 {
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

	var mpl []sdbi.Product
	// var sp sdbi.Product
	// sp.Cost = 1.10
	// mpl = append(mpl, sp)
	fmt.Println("mpl: ", mpl)
	sdb.MockProductSubSkuList = &mpl

	fp := m.GetProductByBySku("345", 4, 4)
	if fp.Color == p.Color {
		t.Fail()
	}
}

func TestSix910Manager_GetProductIDList(t *testing.T) {
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

	var idlst []int64
	idlst = append(idlst, 4)
	sdb.MockProductIDList = &idlst
	fidlst := m.GetProductIDList(4)
	if len(*fidlst) == 0 {
		t.Fail()
	}
}

func TestSix910Manager_GetProductIDListByCategories(t *testing.T) {
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

	var idlst []int64
	idlst = append(idlst, 4)
	sdb.MockProductIDList = &idlst

	var catlst []int64
	catlst = append(catlst, 5)

	fidlst := m.GetProductIDListByCategories(2, &catlst)
	if len(*fidlst) == 0 {
		t.Fail()
	}
}
