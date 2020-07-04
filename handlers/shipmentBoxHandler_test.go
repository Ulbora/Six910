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

func TestSix910Handler_AddShipmentBox(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	sdb.MockAddShipmentBoxSuccess = true
	sdb.MockShipmentBoxID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"shipmentBox":{"boxNumber": 3, "shipmentId": 3}, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddShipmentBox(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AddShipmentBoxReq(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	sdb.MockAddShipmentBoxSuccess = true
	sdb.MockShipmentBoxID = 5

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"shipmentBox":{"boxNumber": 3, "shipmentId": 3}, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddShipmentBox(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_AddShipmentBoxMedia(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	sdb.MockAddShipmentBoxSuccess = true
	sdb.MockShipmentBoxID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"shipmentBox":{"boxNumber": 3, "shipmentId": 3}, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddShipmentBox(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_AddShipmentBoxFail(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	//sdb.MockAddShipmentBoxSuccess = true
	sdb.MockShipmentBoxID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"shipmentBox":{"boxNumber": 3, "shipmentId": 3}, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddShipmentBox(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_AddShipmentBoxAuth(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	sdb.MockAddShipmentBoxSuccess = true
	sdb.MockShipmentBoxID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"shipmentBox":{"boxNumber": 3, "shipmentId": 3}, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddShipmentBox(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateShipmentBox(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	sdb.MockUpdateShipmentBoxSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"shipmentBox":{"id": 33, "boxNumber": 3, "shipmentId": 3}, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateShipmentBox(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateShipmentBoxReq(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	sdb.MockUpdateShipmentBoxSuccess = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"shipmentBox":{"id": 33, "boxNumber": 3, "shipmentId": 3}, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", nil)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateShipmentBox(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateShipmentBoxMedia(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	sdb.MockUpdateShipmentBoxSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"shipmentBox":{"id": 33, "boxNumber": 3, "shipmentId": 3}, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateShipmentBox(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateShipmentBoxFail(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	//sdb.MockUpdateShipmentBoxSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"shipmentBox":{"id": 33, "boxNumber": 3, "shipmentId": 3}, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateShipmentBox(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateShipmentBoxAuth(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	sdb.MockUpdateShipmentBoxSuccess = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"shipmentBox":{"id": 33, "boxNumber": 3, "shipmentId": 3}, "storeId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("storeName2", "TestStore2")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.UpdateShipmentBox(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetShipmentBox(t *testing.T) {
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

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	var sb sdbi.ShipmentBox
	sb.ID = 33
	sb.BoxNumber = 5
	sb.ShipmentID = 3

	sdb.MockShipmentBox = &sb

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

	h.GetShipmentBox(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetShipmentBoxReq(t *testing.T) {
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

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	var sb sdbi.ShipmentBox
	sb.ID = 33
	sb.BoxNumber = 5
	sb.ShipmentID = 3

	sdb.MockShipmentBox = &sb

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
		"id":      "1g",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetShipmentBox(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetShipmentBoxReq2(t *testing.T) {
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

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	var sb sdbi.ShipmentBox
	sb.ID = 33
	sb.BoxNumber = 5
	sb.ShipmentID = 3

	sdb.MockShipmentBox = &sb

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

	h.GetShipmentBox(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetShipmentBoxAuth(t *testing.T) {
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

	var ord sdbi.Order
	ord.ID = 3
	ord.StoreID = 5

	sdb.MockOrder = &ord

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	var sb sdbi.ShipmentBox
	sb.ID = 33
	sb.BoxNumber = 5
	sb.ShipmentID = 3

	sdb.MockShipmentBox = &sb

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

	h.GetShipmentBox(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetShipmentBoxList(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	var sb sdbi.ShipmentBox
	sb.ID = 33
	sb.BoxNumber = 5
	sb.ShipmentID = 3

	var lst []sdbi.ShipmentBox
	lst = append(lst, sb)

	sdb.MockShipmentBoxList = &lst

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
		"shipmentId": "3",
		"storeId":    "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetShipmentBoxList(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetShipmentBoxListReq(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	var sb sdbi.ShipmentBox
	sb.ID = 33
	sb.BoxNumber = 5
	sb.ShipmentID = 3

	var lst []sdbi.ShipmentBox
	lst = append(lst, sb)

	sdb.MockShipmentBoxList = &lst

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
		"shipmentId": "3g",
		"storeId":    "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetShipmentBoxList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetShipmentBoxListReq2(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	var sb sdbi.ShipmentBox
	sb.ID = 33
	sb.BoxNumber = 5
	sb.ShipmentID = 3

	var lst []sdbi.ShipmentBox
	lst = append(lst, sb)

	sdb.MockShipmentBoxList = &lst

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
		"shipmentId": "3",
		//"storeId":    "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetShipmentBoxList(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetShipmentBoxListAuth(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	var sb sdbi.ShipmentBox
	sb.ID = 33
	sb.BoxNumber = 5
	sb.ShipmentID = 3

	var lst []sdbi.ShipmentBox
	lst = append(lst, sb)

	sdb.MockShipmentBoxList = &lst

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
		"shipmentId": "3",
		"storeId":    "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetShipmentBoxList(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteShipmentBox(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	var sb sdbi.ShipmentBox
	sb.ID = 33
	sb.BoxNumber = 5
	sb.ShipmentID = 3

	sdb.MockShipmentBox = &sb

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteShipmentBoxSuccess = true

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

	h.DeleteShipmentBox(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete shipment box in test: ", string(body))

	if w.Code != 200 || !bdy.Success {
		t.Fail()
	}
}

func TestSix910Handler_DeleteShipmentBoxReq(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	var sb sdbi.ShipmentBox
	sb.ID = 33
	sb.BoxNumber = 5
	sb.ShipmentID = 3

	sdb.MockShipmentBox = &sb

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteShipmentBoxSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":      "3f",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteShipmentBox(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete shipment box in test: ", string(body))

	if w.Code != 400 || bdy.Success {
		t.Fail()
	}
}

func TestSix910Handler_DeleteShipmentBoxReq2(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	var sb sdbi.ShipmentBox
	sb.ID = 33
	sb.BoxNumber = 5
	sb.ShipmentID = 3

	sdb.MockShipmentBox = &sb

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteShipmentBoxSuccess = true

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

	h.DeleteShipmentBox(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete shipment box in test: ", string(body))

	if w.Code != 400 || bdy.Success {
		t.Fail()
	}
}

func TestSix910Handler_DeleteShipmentBoxFail(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	var sb sdbi.ShipmentBox
	sb.ID = 33
	sb.BoxNumber = 5
	sb.ShipmentID = 3

	sdb.MockShipmentBox = &sb

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//sdb.MockDeleteShipmentBoxSuccess = true

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

	h.DeleteShipmentBox(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete shipment box in test: ", string(body))

	if w.Code != 500 || bdy.Success {
		t.Fail()
	}
}

func TestSix910Handler_DeleteShipmentBoxAuth(t *testing.T) {
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

	var shp sdbi.Shipment
	shp.ID = 4
	shp.Boxes = 5
	shp.OrderID = 3

	sdb.MockShipment = &shp

	var sb sdbi.ShipmentBox
	sb.ID = 33
	sb.BoxNumber = 5
	sb.ShipmentID = 3

	sdb.MockShipmentBox = &sb

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteShipmentBoxSuccess = true

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

	h.DeleteShipmentBox(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete shipment box in test: ", string(body))

	if w.Code != 401 || bdy.Success {
		t.Fail()
	}
}
