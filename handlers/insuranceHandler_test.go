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

func TestSix910Handler_AddInsurance(t *testing.T) {
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

	sdb.MockAddInsuranceSuccess = true
	sdb.MockInsuranceID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cost":4.95, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddInsurance(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AddInsuranceReq(t *testing.T) {
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

	sdb.MockAddInsuranceSuccess = true
	sdb.MockInsuranceID = 5

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cost":4.95, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddInsurance(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_AddInsuranceFail(t *testing.T) {
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

	//sdb.MockAddInsuranceSuccess = true
	sdb.MockInsuranceID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cost":4.95, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddInsurance(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_AddInsuranceMedia(t *testing.T) {
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

	sdb.MockAddInsuranceSuccess = true
	sdb.MockInsuranceID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cost":4.95, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddInsurance(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_AddInsuranceAuth(t *testing.T) {
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

	sdb.MockAddInsuranceSuccess = true
	sdb.MockInsuranceID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"cost":4.95, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddInsurance(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateInsurance(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

	sdb.MockUpdateInsuranceSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 1, "cost":4.95 , "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateInsurance(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateInsuranceReq(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

	sdb.MockUpdateInsuranceSuccess = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 1, "cost":4.95 , "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", nil)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateInsurance(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateInsuranceFail(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

	//sdb.MockUpdateInsuranceSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 1, "cost":4.95 , "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateInsurance(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateInsuranceAuth(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

	sdb.MockUpdateInsuranceSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 1, "cost":4.95 , "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateInsurance(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateInsuranceMedia(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

	sdb.MockUpdateInsuranceSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 1, "cost":4.95 , "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateInsurance(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_GetInsurance(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

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

	h.GetInsurance(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetInsuranceReq(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

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
		"id":      "f1",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetInsurance(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetInsuranceReq2(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

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

	h.GetInsurance(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetInsuranceAuth(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

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

	h.GetInsurance(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetInsuranceList(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5

	var lst []sdbi.Insurance
	lst = append(lst, ins)

	sdb.MockInsuranceList = &lst

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

	h.GetInsuranceList(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetInsuranceListReq(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5

	var lst []sdbi.Insurance
	lst = append(lst, ins)

	sdb.MockInsuranceList = &lst

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
		"storeId": "5f",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetInsuranceList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetInsuranceListReq2(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5

	var lst []sdbi.Insurance
	lst = append(lst, ins)

	sdb.MockInsuranceList = &lst

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

	h.GetInsuranceList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetInsuranceListAuth(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5

	var lst []sdbi.Insurance
	lst = append(lst, ins)

	sdb.MockInsuranceList = &lst

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

	h.GetInsuranceList(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteInsurance(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteInsuranceSuccess = true

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

	h.DeleteInsurance(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete insurance in test: ", string(body))

	if w.Code != 200 || !bdy.Success {
		t.Fail()
	}
}

func TestSix910Handler_DeleteInsuranceReq(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteInsuranceSuccess = true

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

	h.DeleteInsurance(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete insurance in test: ", string(body))

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteInsuranceReq2(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteInsuranceSuccess = true

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

	h.DeleteInsurance(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete insurance in test: ", string(body))

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteInsuranceFail(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//sdb.MockDeleteInsuranceSuccess = true

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

	h.DeleteInsurance(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete insurance in test: ", string(body))

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteInsuranceAuth(t *testing.T) {
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

	var ins sdbi.Insurance
	ins.StoreID = 5
	sdb.MockInsurance = &ins

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteInsuranceSuccess = true

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

	h.DeleteInsurance(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete insurance in test: ", string(body))

	if w.Code != 401 {
		t.Fail()
	}
}
