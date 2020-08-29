package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	lg "github.com/Ulbora/Level_Logger"
	man "github.com/Ulbora/Six910/managers"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
	"github.com/gorilla/mux"
)

func TestSix910Handler_AddCartItem(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 22
	var lst []sdbi.CartItem
	lst = append(lst, ci)

	sdb.MockCartItemList = &lst

	sdb.MockAddCartItemSuccess = true
	sdb.MockCartItemID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cartItem": {"cartId": 9}, "customerId":3, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.AddCartItem(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AddCartItemAuth(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l
	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 22
	var lst []sdbi.CartItem
	lst = append(lst, ci)

	sdb.MockCartItemList = &lst

	sdb.MockAddCartItemSuccess = true
	sdb.MockCartItemID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cartItem": {"cartId": 4}, "customerId":3, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddCartItem(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_AddCartItemFail(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 22
	var lst []sdbi.CartItem
	lst = append(lst, ci)

	sdb.MockCartItemList = &lst

	//sdb.MockAddCartItemSuccess = true
	sdb.MockCartItemID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cartItem": {"cartId": 4}, "customerId":3, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.AddCartItem(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_AddCartItemMedia(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 22
	var lst []sdbi.CartItem
	lst = append(lst, ci)

	sdb.MockCartItemList = &lst

	sdb.MockAddCartItemSuccess = true
	sdb.MockCartItemID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cartItem": {"cartId": 4}, "customerId":3, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("apiKey", "123456")

	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddCartItem(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_AddCartItemReq(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.CartID = 4
	ci.ProductID = 22
	var lst []sdbi.CartItem
	lst = append(lst, ci)

	sdb.MockCartItemList = &lst

	sdb.MockAddCartItemSuccess = true
	sdb.MockCartItemID = 5

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cartItem": {"cartId": 4}, "customerId":3, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.AddCartItem(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateCartItem(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	sdb.MockUpdateCartItemSuccess = true

	var cus sdbi.Customer
	cus.StoreID = 4
	sdb.MockCustomer = &cus

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cartItem": {"id": 5, "cartId": 4}, "customerId":3, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateCartItem(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateCartItemAuth(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	sdb.MockUpdateCartItemSuccess = true

	var cus sdbi.Customer
	cus.StoreID = 4
	sdb.MockCustomer = &cus

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cartItem": {"cartId": 4}, "customerId":3, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateCartItem(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateCartItemFail(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	//sdb.MockUpdateCartItemSuccess = true

	var cus sdbi.Customer
	cus.StoreID = 4
	sdb.MockCustomer = &cus

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cartItem": {"cartId": 4}, "customerId":3, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateCartItem(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateCartItemMedia(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	sdb.MockUpdateCartItemSuccess = true

	var cus sdbi.Customer
	cus.StoreID = 4
	sdb.MockCustomer = &cus

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cartItem": {"cartId": 4}, "customerId":3, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")
	r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.UpdateCartItem(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateCartItemReq(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	sdb.MockUpdateCartItemSuccess = true

	var cus sdbi.Customer
	cus.StoreID = 4
	sdb.MockCustomer = &cus

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cartItem": {"cartId": 4}, "customerId":3, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", nil)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")
	r.Header.Set("apiKey", "123456")

	w := httptest.NewRecorder()

	h.UpdateCartItem(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetCarItem(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.ID = 66
	ci.CartID = 9
	sdb.MockCartItem = &ci

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"cid":     "9",
		"prodId":  "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetCartItem(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetCarItemAuth(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.ID = 66
	ci.CartID = 9
	sdb.MockCartItem = &ci

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"cid":     "9",
		"prodId":  "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetCartItem(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetCarItemReq2(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.ID = 66
	ci.CartID = 9
	sdb.MockCartItem = &ci

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"cid":    "9",
		"prodId": "3",
		//"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetCartItem(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetCarItemReq(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.ID = 66
	ci.CartID = 9
	sdb.MockCartItem = &ci

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"cid":     "9d",
		"prodId":  "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetCartItem(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetCartItemList(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.ID = 66
	ci.CartID = 9
	//sdb.MockCartItem = &ci

	var lst []sdbi.CartItem
	lst = append(lst, ci)
	sdb.MockCartItemList = &lst

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"cartId":  "9",
		"cid":     "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetCartItemList(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetCartItemListAuth(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.ID = 66
	ci.CartID = 9
	//sdb.MockCartItem = &ci

	var lst []sdbi.CartItem
	lst = append(lst, ci)
	sdb.MockCartItemList = &lst

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"cartId":  "9",
		"cid":     "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetCartItemList(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetCartItemListReq(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.ID = 66
	ci.CartID = 9
	//sdb.MockCartItem = &ci

	var lst []sdbi.CartItem
	lst = append(lst, ci)
	sdb.MockCartItemList = &lst

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"cartId": "9",
		"cid":    "3",
		//"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetCartItemList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetCartItemListReq2(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.ID = 66
	ci.CartID = 9
	//sdb.MockCartItem = &ci

	var lst []sdbi.CartItem
	lst = append(lst, ci)
	sdb.MockCartItemList = &lst

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"cartId":  "9g",
		"cid":     "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetCartItemList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteCartItem(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.ID = 5
	ci.CartID = 9
	sdb.MockCartItem = &ci

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteCartItemSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":     "5",
		"prodId": "3",
		"cartId": "9",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteCartItem(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteCartItemAuth(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.ID = 66
	ci.CartID = 9
	sdb.MockCartItem = &ci

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteCartItemSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":     "5",
		"prodId": "3",
		"cartId": "9",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteCartItem(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteCartItemFail(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.ID = 66
	ci.CartID = 9
	sdb.MockCartItem = &ci

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//sdb.MockDeleteCartItemSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":     "5",
		"prodId": "3",
		"cartId": "9",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteCartItem(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteCartItemReq(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.ID = 66
	ci.CartID = 9
	sdb.MockCartItem = &ci

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteCartItemSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":     "5",
		"prodId": "3",
		//"cartId": "9",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteCartItem(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteCartItemReq2(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l

	var sec sdbi.Security
	sec.OauthOn = true
	sdb.MockSecurity = &sec

	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	str.OauthClientID = 5
	sdb.MockStore = &str

	var cart sdbi.Cart
	cart.ID = 9
	cart.CustomerID = 3
	cart.StoreID = 5

	sdb.MockCart = &cart

	var ci sdbi.CartItem
	ci.ID = 66
	ci.CartID = 9
	sdb.MockCartItem = &ci

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteCartItemSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":     "5g",
		"prodId": "3",
		"cartId": "9",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteCartItem(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}
