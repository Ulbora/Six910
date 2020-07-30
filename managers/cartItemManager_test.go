package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddCartItem(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	cart.CustomerID = 4
	sdb.MockCart = &cart

	var prod sdbi.Product
	prod.ID = 2
	prod.StoreID = 3
	sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 2

	var ci2 sdbi.CartItem
	ci2.CartID = 4
	ci2.ProductID = 22

	var lst []sdbi.CartItem
	lst = append(lst, ci2)

	sdb.MockCartItemList = &lst

	sdb.MockAddCartItemSuccess = true
	sdb.MockCartItemID = 2
	res := m.AddCartItem(&ci, 4, 3)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddCartItemNoCID(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	cart.CustomerID = 4
	sdb.MockCart = &cart

	var prod sdbi.Product
	prod.ID = 2
	prod.StoreID = 3
	sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 2

	var ci2 sdbi.CartItem
	ci2.CartID = 4
	ci2.ProductID = 22

	var lst []sdbi.CartItem
	lst = append(lst, ci2)

	sdb.MockCartItemList = &lst

	sdb.MockAddCartItemSuccess = true
	sdb.MockCartItemID = 2
	res := m.AddCartItem(&ci, 0, 3)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddCartItemExisting(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	cart.CustomerID = 4
	sdb.MockCart = &cart

	var prod sdbi.Product
	prod.ID = 2
	prod.StoreID = 3
	sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 2

	var ci2 sdbi.CartItem
	ci2.CartID = 4
	ci2.ProductID = 22

	var lst []sdbi.CartItem
	lst = append(lst, ci)

	sdb.MockCartItemList = &lst

	sdb.MockUpdateCartItemSuccess = true
	sdb.MockCartItemID = 2
	res := m.AddCartItem(&ci, 4, 3)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddCartItemFail1(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	cart.CustomerID = 4
	sdb.MockCart = &cart

	var prod sdbi.Product
	prod.ID = 2
	prod.StoreID = 3
	sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 2

	var lst []sdbi.CartItem
	lst = append(lst, ci)
	sdb.MockCartItemList = &lst

	//sdb.MockAddCartItemSuccess = true
	//sdb.MockCartItemID = 2
	res := m.AddCartItem(&ci, 4, 3)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddCartItemFail2(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	sdb.MockCart = &cart

	var prod sdbi.Product
	prod.ID = 2
	prod.StoreID = 33
	sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 2

	var lst []sdbi.CartItem
	lst = append(lst, ci)
	sdb.MockCartItemList = &lst

	sdb.MockAddCartItemSuccess = true
	sdb.MockCartItemID = 2
	res := m.AddCartItem(&ci, 4, 3)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCartItem(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	cart.CustomerID = 4
	sdb.MockCart = &cart

	var prod sdbi.Product
	prod.ID = 2
	prod.StoreID = 3
	sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 2

	sdb.MockUpdateCartItemSuccess = true

	res := m.UpdateCartItem(&ci, 4, 3)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCartItemNoCID(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	cart.CustomerID = 4
	sdb.MockCart = &cart

	var prod sdbi.Product
	prod.ID = 2
	prod.StoreID = 3
	sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 2

	sdb.MockUpdateCartItemSuccess = true

	res := m.UpdateCartItem(&ci, 0, 3)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCartItemFail1(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 33
	cart.CustomerID = 4
	sdb.MockCart = &cart

	var prod sdbi.Product
	prod.ID = 2
	prod.StoreID = 3
	sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 2

	sdb.MockUpdateCartItemSuccess = true

	res := m.UpdateCartItem(&ci, 4, 3)
	if res.Success {
		t.Fail()
	}
}
func TestSix910Manager_UpdateCartItemFail2(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	cart.CustomerID = 4
	sdb.MockCart = &cart

	var prod sdbi.Product
	prod.ID = 2
	prod.StoreID = 3
	sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 2

	//sdb.MockUpdateCartItemSuccess = true

	res := m.UpdateCartItem(&ci, 4, 3)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetCarItem(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	sdb.MockCart = &cart

	// var prod sdbi.Product
	// prod.ID = 2
	// prod.StoreID = 3
	// sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.ID = 6
	ci.CartID = 4
	ci.ProductID = 2
	ci.Quantity = 10

	sdb.MockCartItem = &ci

	fcart := m.GetCarItem(6, 2, 3)
	if fcart.Quantity != ci.Quantity {
		t.Fail()
	}
}

func TestSix910Manager_GetCarItemFail(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 33
	sdb.MockCart = &cart

	// var prod sdbi.Product
	// prod.ID = 2
	// prod.StoreID = 33
	// sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.ID = 6
	ci.CartID = 4
	ci.ProductID = 2
	ci.Quantity = 10

	sdb.MockCartItem = &ci

	fcart := m.GetCarItem(6, 2, 3)
	if fcart.Quantity == ci.Quantity {
		t.Fail()
	}
}

func TestSix910Manager_GetCartItemList(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	cart.CustomerID = 4
	sdb.MockCart = &cart

	// var prod sdbi.Product
	// prod.ID = 2
	// prod.StoreID = 3
	// sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.ID = 6
	ci.CartID = 4
	ci.ProductID = 2
	ci.Quantity = 10
	var cilst []sdbi.CartItem
	cilst = append(cilst, ci)

	sdb.MockCartItemList = &cilst

	fcilst := m.GetCartItemList(4, 4, 3)
	if len(*fcilst) != 1 {
		t.Fail()
	}

}

func TestSix910Manager_GetCartItemListNoCID(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	cart.CustomerID = 4
	sdb.MockCart = &cart

	// var prod sdbi.Product
	// prod.ID = 2
	// prod.StoreID = 3
	// sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.ID = 6
	ci.CartID = 4
	ci.ProductID = 2
	ci.Quantity = 10
	var cilst []sdbi.CartItem
	cilst = append(cilst, ci)

	sdb.MockCartItemList = &cilst

	fcilst := m.GetCartItemList(4, 0, 3)
	if len(*fcilst) != 1 {
		t.Fail()
	}

}

func TestSix910Manager_GetCartItemListFail(t *testing.T) {

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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 33
	cart.CustomerID = 4
	sdb.MockCart = &cart

	// var prod sdbi.Product
	// prod.ID = 2
	// prod.StoreID = 3
	// sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.ID = 6
	ci.CartID = 4
	ci.ProductID = 2
	ci.Quantity = 10
	var cilst []sdbi.CartItem
	cilst = append(cilst, ci)

	sdb.MockCartItemList = &cilst

	fcilst := m.GetCartItemList(4, 4, 3)
	if len(*fcilst) == 1 {
		t.Fail()
	}

}

func TestSix910Manager_DeleteCartItem(t *testing.T) {
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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	sdb.MockCart = &cart

	// var prod sdbi.Product
	// prod.ID = 2
	// prod.StoreID = 3
	// sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.ID = 6
	ci.CartID = 4
	ci.ProductID = 2
	ci.Quantity = 10

	sdb.MockCartItem = &ci
	sdb.MockDeleteCartItemSuccess = true

	res := m.DeleteCartItem(6, 2, 4)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteCartItemFail1(t *testing.T) {
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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	sdb.MockCart = &cart

	// var prod sdbi.Product
	// prod.ID = 2
	// prod.StoreID = 3
	// sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.ID = 6
	ci.CartID = 44
	ci.ProductID = 2
	ci.Quantity = 10

	sdb.MockCartItem = &ci
	sdb.MockDeleteCartItemSuccess = true

	res := m.DeleteCartItem(6, 2, 4)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteCartItemFail2(t *testing.T) {
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

	var cart sdbi.Cart
	cart.ID = 4
	cart.StoreID = 3
	sdb.MockCart = &cart

	// var prod sdbi.Product
	// prod.ID = 2
	// prod.StoreID = 3
	// sdb.MockProduct = &prod

	var ci sdbi.CartItem
	ci.ID = 6
	ci.CartID = 4
	ci.ProductID = 2
	ci.Quantity = 10

	sdb.MockCartItem = &ci
	//sdb.MockDeleteCartItemSuccess = true

	res := m.DeleteCartItem(6, 2, 4)
	if res.Success {
		t.Fail()
	}
}
