package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func TestSix910Handler_AddShippingCarrier(t *testing.T) {
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

	sdb.MockAddShippingCarrierSuccess = true
	sdb.MockShippingCarrierID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"carrier":"test", "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddShippingCarrier(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AddShippingCarrierReq(t *testing.T) {
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

	sdb.MockAddShippingCarrierSuccess = true
	sdb.MockShippingCarrierID = 5

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"carrier":"test", "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddShippingCarrier(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_AddShippingCarrierMedia(t *testing.T) {
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

	sdb.MockAddShippingCarrierSuccess = true
	sdb.MockShippingCarrierID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"carrier":"test", "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddShippingCarrier(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_AddShippingCarrierFail(t *testing.T) {
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

	//sdb.MockAddShippingCarrierSuccess = true
	sdb.MockShippingCarrierID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"carrier":"test", "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddShippingCarrier(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_AddShippingCarrierAuth(t *testing.T) {
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

	sdb.MockAddShippingCarrierSuccess = true
	sdb.MockShippingCarrierID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"carrier":"test", "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddShippingCarrier(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateShippingCarrier(t *testing.T) {
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

	// var cus sdbi.StorePlugins
	// cus.StoreID = 5
	// cus.PluginName = "testPlugin"
	// sdb.MockStorePlugin = &cus

	sdb.MockUpdateShippingCarrierSuccess = true

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 1, "carrier":"test", "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateShippingCarrier(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateShippingCarrierReq(t *testing.T) {
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

	// var cus sdbi.StorePlugins
	// cus.StoreID = 5
	// cus.PluginName = "testPlugin"
	// sdb.MockStorePlugin = &cus

	sdb.MockUpdateShippingCarrierSuccess = true

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 1, "carrier":"test", "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", nil)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateShippingCarrier(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateShippingCarrierMedia(t *testing.T) {
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

	// var cus sdbi.StorePlugins
	// cus.StoreID = 5
	// cus.PluginName = "testPlugin"
	// sdb.MockStorePlugin = &cus

	sdb.MockUpdateShippingCarrierSuccess = true

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 1, "carrier":"test", "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateShippingCarrier(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateShippingCarrierFail(t *testing.T) {
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

	// var cus sdbi.StorePlugins
	// cus.StoreID = 5
	// cus.PluginName = "testPlugin"
	// sdb.MockStorePlugin = &cus

	//sdb.MockUpdateShippingCarrierSuccess = true

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 1, "carrier":"test", "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateShippingCarrier(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateShippingCarrierAuth(t *testing.T) {
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

	// var cus sdbi.StorePlugins
	// cus.StoreID = 5
	// cus.PluginName = "testPlugin"
	// sdb.MockStorePlugin = &cus

	sdb.MockUpdateShippingCarrierSuccess = true

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 1, "carrier":"test", "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateShippingCarrier(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetShippingCarrier(t *testing.T) {
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

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

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
		"id":      "1",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetShippingCarrier(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetShippingCarrierReq(t *testing.T) {
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

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

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
		"id":      "1f",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetShippingCarrier(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetShippingCarrierReq2(t *testing.T) {
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

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

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
		"id": "1",
		//"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetShippingCarrier(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetShippingCarrierAuth(t *testing.T) {
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

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

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
		"id":      "1",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetShippingCarrier(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetShippingCarrierList(t *testing.T) {
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

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	var lst []sdbi.ShippingCarrier
	lst = append(lst, shc)

	sdb.MockShippingCarrierList = &lst

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
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetShippingCarrierList(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetShippingCarrierListReq(t *testing.T) {
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

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	var lst []sdbi.ShippingCarrier
	lst = append(lst, shc)

	sdb.MockShippingCarrierList = &lst

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
		"storeId": "5h",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetShippingCarrierList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetShippingCarrierListReq2(t *testing.T) {
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

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	var lst []sdbi.ShippingCarrier
	lst = append(lst, shc)

	sdb.MockShippingCarrierList = &lst

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
		//"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetShippingCarrierList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetShippingCarrierListAuth(t *testing.T) {
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

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	var lst []sdbi.ShippingCarrier
	lst = append(lst, shc)

	sdb.MockShippingCarrierList = &lst

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
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetShippingCarrierList(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteShippingCarrier(t *testing.T) {
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

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteShippingCarrierSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":      "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteShippingCarrier(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete store plugin in test: ", string(body))

	if w.Code != 200 || !bdy.Success {
		t.Fail()
	}
}

func TestSix910Handler_DeleteShippingCarrierReq(t *testing.T) {
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

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteShippingCarrierSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":      "3g",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteShippingCarrier(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete store plugin in test: ", string(body))

	if w.Code != 400 || bdy.Success {
		t.Fail()
	}
}

func TestSix910Handler_DeleteShippingCarrierReq2(t *testing.T) {
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

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteShippingCarrierSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id": "3",
		//"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteShippingCarrier(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete store plugin in test: ", string(body))

	if w.Code != 400 || bdy.Success {
		t.Fail()
	}
}

func TestSix910Handler_DeleteShippingCarrierFail(t *testing.T) {
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

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//sdb.MockDeleteShippingCarrierSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":      "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteShippingCarrier(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete store plugin in test: ", string(body))

	if w.Code != 500 || bdy.Success {
		t.Fail()
	}
}

func TestSix910Handler_DeleteShippingCarrierAuth(t *testing.T) {
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

	var shc sdbi.ShippingCarrier
	shc.Carrier = "ups"
	shc.ID = 9
	shc.StoreID = 5

	sdb.MockShippingCarrier = &shc

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteShippingCarrierSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":      "3",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteShippingCarrier(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete store plugin in test: ", string(body))

	if w.Code != 401 || bdy.Success {
		t.Fail()
	}
}
