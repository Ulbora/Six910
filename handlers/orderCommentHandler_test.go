package handlers

import (
	"bytes"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"

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

func TestSix910Handler_AddOrderComments(t *testing.T) {
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

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	sdb.MockAddOrderCommentSuccess = true
	sdb.MockOrderCommentID = 5

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester1"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	h := sh.GetNew()
	aJSON := io.NopCloser(bytes.NewBufferString(`{"orderComment": {"comment": "test", "orderId": 3 }, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester1")
	w := httptest.NewRecorder()

	h.AddOrderComments(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AddOrderComments2(t *testing.T) {
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

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	sdb.MockAddOrderCommentSuccess = true
	sdb.MockOrderCommentID = 5

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester1"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	h := sh.GetNew()
	aJSON := io.NopCloser(bytes.NewBufferString(`{"orderComment": {"comment": "test", "orderId": 3 }, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.Header.Set("clientId", "123456")
	r.SetBasicAuth("tester", "tester1")
	w := httptest.NewRecorder()

	h.AddOrderComments(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AddOrderCommentsReq(t *testing.T) {
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

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	sdb.MockAddOrderCommentSuccess = true
	sdb.MockOrderCommentID = 5

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester1"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"orderComment": {"comment": "test", "orderId": 3 }, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester1")
	w := httptest.NewRecorder()

	h.AddOrderComments(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_AddOrderCommentsMedia(t *testing.T) {
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

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	sdb.MockAddOrderCommentSuccess = true
	sdb.MockOrderCommentID = 5

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester1"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	h := sh.GetNew()
	aJSON := io.NopCloser(bytes.NewBufferString(`{"orderComment": {"comment": "test", "orderId": 3 }, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester1")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddOrderComments(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_AddOrderCommentsFail(t *testing.T) {
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

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	//sdb.MockAddOrderCommentSuccess = true
	sdb.MockOrderCommentID = 5

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester1"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	h := sh.GetNew()
	aJSON := io.NopCloser(bytes.NewBufferString(`{"orderComment": {"comment": "test", "orderId": 3 }, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester1")
	w := httptest.NewRecorder()

	h.AddOrderComments(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_AddOrderCommentsAuth(t *testing.T) {
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

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 8
	oc.OrderID = 3

	var lst []sdbi.OrderComment
	lst = append(lst, oc)

	sdb.MockOrderCommentList = &lst

	sdb.MockAddOrderCommentSuccess = true
	sdb.MockOrderCommentID = 5

	h := sh.GetNew()
	aJSON := io.NopCloser(bytes.NewBufferString(`{"orderComment": {"comment": "test", "orderId": 3 }, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddOrderComments(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetOrderCommentList(t *testing.T) {
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

	var cat sdbi.Category
	cat.ID = 1
	cat.StoreID = 5
	sdb.MockCategory = &cat

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 8
	oc.OrderID = 3

	var lst []sdbi.OrderComment
	lst = append(lst, oc)

	sdb.MockOrderCommentList = &lst

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester1"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"orderId": "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester1")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetOrderCommentList(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetOrderCommentList2(t *testing.T) {
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

	var cat sdbi.Category
	cat.ID = 1
	cat.StoreID = 5
	sdb.MockCategory = &cat

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 8
	oc.OrderID = 3

	var lst []sdbi.OrderComment
	lst = append(lst, oc)

	sdb.MockOrderCommentList = &lst

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester1"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"orderId": "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.Header.Set("clientId", "123456")
	r.SetBasicAuth("tester", "tester1")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetOrderCommentList(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetOrderCommentListReq(t *testing.T) {
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

	var cat sdbi.Category
	cat.ID = 1
	cat.StoreID = 5
	sdb.MockCategory = &cat

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 8
	oc.OrderID = 3

	var lst []sdbi.OrderComment
	lst = append(lst, oc)

	sdb.MockOrderCommentList = &lst

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester1"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"orderId": "3d",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester1")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetOrderCommentList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetOrderCommentListReq2(t *testing.T) {
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

	var cat sdbi.Category
	cat.ID = 1
	cat.StoreID = 5
	sdb.MockCategory = &cat

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 8
	oc.OrderID = 3

	var lst []sdbi.OrderComment
	lst = append(lst, oc)

	sdb.MockOrderCommentList = &lst

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "customer"
	lu.StoreID = 4
	lu.UserName = "tester"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("tester1"), bcrypt.DefaultCost)
	if err == nil {
		lu.Password = string(hashedPw)
		fmt.Println("hpw: ", lu.Password)
	}
	//hpw := sm
	//lu.Password = "tester1"
	lu.Enabled = true

	sdb.MockLocalAccount = &lu

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"orderId": "3",
		//"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester1")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetOrderCommentList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetOrderCommentListAuth(t *testing.T) {
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

	var cat sdbi.Category
	cat.ID = 1
	cat.StoreID = 5
	sdb.MockCategory = &cat

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 8
	oc.OrderID = 3

	var lst []sdbi.OrderComment
	lst = append(lst, oc)

	sdb.MockOrderCommentList = &lst

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
		"orderId": "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetOrderCommentList(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}
