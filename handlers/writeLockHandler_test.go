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

func TestSix910Handler_AddDataStoreWriteLock(t *testing.T) {
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

	sdb.MockAddDataStoreWriteLockSuccess = true
	sdb.MockDataStoreWriteLockID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{ "dataStoreName":"test", "locked": true, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddDataStoreWriteLock(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AddDataStoreWriteLockReq(t *testing.T) {
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

	sdb.MockAddDataStoreWriteLockSuccess = true
	sdb.MockDataStoreWriteLockID = 5

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{ "dataStoreName":"test", "locked": true, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddDataStoreWriteLock(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_AddDataStoreWriteLockMedia(t *testing.T) {
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

	sdb.MockAddDataStoreWriteLockSuccess = true
	sdb.MockDataStoreWriteLockID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{ "dataStoreName":"test", "locked": true, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddDataStoreWriteLock(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_AddDataStoreWriteLockFail(t *testing.T) {
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

	//sdb.MockAddDataStoreWriteLockSuccess = true
	sdb.MockDataStoreWriteLockID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{ "dataStoreName":"test", "locked": true, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddDataStoreWriteLock(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_AddDataStoreWriteLockAuth(t *testing.T) {
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

	sdb.MockAddDataStoreWriteLockSuccess = true
	sdb.MockDataStoreWriteLockID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{ "dataStoreName":"test", "locked": true, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddDataStoreWriteLock(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateDataStoreWriteLock(t *testing.T) {
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

	var cus sdbi.StorePlugins
	cus.StoreID = 5
	cus.PluginName = "testPlugin"
	sdb.MockStorePlugin = &cus

	sdb.MockUpdateDataStoreWriteLockSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{ "id": 4, "dataStoreName":"test", "locked": true, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateDataStoreWriteLock(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateDataStoreWriteLockReq(t *testing.T) {
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

	var cus sdbi.StorePlugins
	cus.StoreID = 5
	cus.PluginName = "testPlugin"
	sdb.MockStorePlugin = &cus

	sdb.MockUpdateDataStoreWriteLockSuccess = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{ "id": 4, "dataStoreName":"test", "locked": true, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", nil)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateDataStoreWriteLock(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateDataStoreWriteLockMedia(t *testing.T) {
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

	var cus sdbi.StorePlugins
	cus.StoreID = 5
	cus.PluginName = "testPlugin"
	sdb.MockStorePlugin = &cus

	sdb.MockUpdateDataStoreWriteLockSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{ "id": 4, "dataStoreName":"test", "locked": true, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateDataStoreWriteLock(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateDataStoreWriteLockFail(t *testing.T) {
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

	var cus sdbi.StorePlugins
	cus.StoreID = 5
	cus.PluginName = "testPlugin"
	sdb.MockStorePlugin = &cus

	//sdb.MockUpdateDataStoreWriteLockSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{ "id": 4, "dataStoreName":"test", "locked": true, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateDataStoreWriteLock(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateDataStoreWriteLockAuth(t *testing.T) {
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

	var cus sdbi.StorePlugins
	cus.StoreID = 5
	cus.PluginName = "testPlugin"
	sdb.MockStorePlugin = &cus

	sdb.MockUpdateDataStoreWriteLockSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{ "id": 4, "dataStoreName":"test", "locked": true, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateDataStoreWriteLock(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetDataStoreWriteLock(t *testing.T) {
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

	var spi sdbi.DataStoreWriteLock
	spi.StoreID = 5
	spi.DataStoreName = "content"
	sdb.MockDataStoreWriteLock = &spi

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
		"dataStore": "content",
		"storeId":   "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetDataStoreWriteLock(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetDataStoreWriteLockReq(t *testing.T) {
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

	var spi sdbi.DataStoreWriteLock
	spi.StoreID = 5
	spi.DataStoreName = "content"
	sdb.MockDataStoreWriteLock = &spi

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
		"dataStore": "content",
		"storeId":   "5f",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetDataStoreWriteLock(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetDataStoreWriteLockReq2(t *testing.T) {
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

	var spi sdbi.DataStoreWriteLock
	spi.StoreID = 5
	spi.DataStoreName = "content"
	sdb.MockDataStoreWriteLock = &spi

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
		"dataStore": "content",
		//"storeId":   "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetDataStoreWriteLock(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetDataStoreWriteLockAuth(t *testing.T) {
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

	var spi sdbi.DataStoreWriteLock
	spi.StoreID = 5
	spi.DataStoreName = "content"
	sdb.MockDataStoreWriteLock = &spi

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
		"dataStore": "content",
		"storeId":   "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetDataStoreWriteLock(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}
