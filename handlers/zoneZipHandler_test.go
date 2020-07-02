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

func TestSix910Handler_AddZoneZip(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.RegionID = 1

	sdb.MockIncludedSubRegion = &sr

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockAddZoneZipSuccess = true
	sdb.MockZoneZipID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"storeId": 5, "zoneZip": { "zipCode":"12345", "includedSubRegionId": 7}}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddZoneZip(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AddZoneZipReq(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.RegionID = 1

	sdb.MockIncludedSubRegion = &sr

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockAddZoneZipSuccess = true
	sdb.MockZoneZipID = 5

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"storeId": 5, "zoneZip": { "zipCode":"12345", "includedSubRegionId": 7}}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddZoneZip(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_AddZoneZipMedia(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.RegionID = 1

	sdb.MockIncludedSubRegion = &sr

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockAddZoneZipSuccess = true
	sdb.MockZoneZipID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"storeId": 5, "zoneZip": { "zipCode":"12345", "includedSubRegionId": 7}}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddZoneZip(w, r)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestSix910Handler_AddZoneZipFail(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.RegionID = 1

	sdb.MockIncludedSubRegion = &sr

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	//sdb.MockAddZoneZipSuccess = true
	sdb.MockZoneZipID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"storeId": 5, "zoneZip": { "zipCode":"12345", "includedSubRegionId": 7}}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddZoneZip(w, r)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestSix910Handler_AddZoneZipAuth(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.RegionID = 1

	sdb.MockIncludedSubRegion = &sr

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l
	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//h := sh.GetNew()

	sdb.MockAddZoneZipSuccess = true
	sdb.MockZoneZipID = 5

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"storeId": 5, "zoneZip": { "zipCode":"12345", "includedSubRegionId": 7}}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddZoneZip(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetZoneZipListByExclusion(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.ExcludedSubRegion
	sr.ID = 2
	sr.RegionID = 1
	sdb.MockExcludedSubRegion = &sr

	var zz sdbi.ZoneZip
	zz.ExcludedSubRegionID = 4
	zz.ID = 22
	zz.ZipCode = "12345"

	var lst []sdbi.ZoneZip
	lst = append(lst, zz)

	sdb.MockZoneZipList = &lst

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
		"exId":    "2",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetZoneZipListByExclusion(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetZoneZipListByExclusionReq(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.ExcludedSubRegion
	sr.ID = 2
	sr.RegionID = 1
	sdb.MockExcludedSubRegion = &sr

	var zz sdbi.ZoneZip
	zz.ExcludedSubRegionID = 4
	zz.ID = 22
	zz.ZipCode = "12345"

	var lst []sdbi.ZoneZip
	lst = append(lst, zz)

	sdb.MockZoneZipList = &lst

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
		"exId":    "2g",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetZoneZipListByExclusion(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetZoneZipListByExclusionReq2(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.ExcludedSubRegion
	sr.ID = 2
	sr.RegionID = 1
	sdb.MockExcludedSubRegion = &sr

	var zz sdbi.ZoneZip
	zz.ExcludedSubRegionID = 4
	zz.ID = 22
	zz.ZipCode = "12345"

	var lst []sdbi.ZoneZip
	lst = append(lst, zz)

	sdb.MockZoneZipList = &lst

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
		"exId": "2",
		//"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetZoneZipListByExclusion(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetZoneZipListByExclusionAuth(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.ExcludedSubRegion
	sr.ID = 2
	sr.RegionID = 1
	sdb.MockExcludedSubRegion = &sr

	var zz sdbi.ZoneZip
	zz.ExcludedSubRegionID = 4
	zz.ID = 22
	zz.ZipCode = "12345"

	var lst []sdbi.ZoneZip
	lst = append(lst, zz)

	sdb.MockZoneZipList = &lst

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
		"exId":    "2",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetZoneZipListByExclusion(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_GetZoneZipListByInclusion(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.ID = 2
	sr.RegionID = 1
	sdb.MockIncludedSubRegion = &sr

	var zz sdbi.ZoneZip
	zz.ExcludedSubRegionID = 4
	zz.ID = 22
	zz.ZipCode = "12345"

	var lst []sdbi.ZoneZip
	lst = append(lst, zz)

	sdb.MockZoneZipList = &lst

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
		"incId":   "2",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetZoneZipListByInclusion(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_GetZoneZipListByInclusionReq(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.ID = 2
	sr.RegionID = 1
	sdb.MockIncludedSubRegion = &sr

	var zz sdbi.ZoneZip
	zz.ExcludedSubRegionID = 4
	zz.ID = 22
	zz.ZipCode = "12345"

	var lst []sdbi.ZoneZip
	lst = append(lst, zz)

	sdb.MockZoneZipList = &lst

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
		"incId":   "2f",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetZoneZipListByInclusion(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetZoneZipListByInclusionReq2(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.ID = 2
	sr.RegionID = 1
	sdb.MockIncludedSubRegion = &sr

	var zz sdbi.ZoneZip
	zz.ExcludedSubRegionID = 4
	zz.ID = 22
	zz.ZipCode = "12345"

	var lst []sdbi.ZoneZip
	lst = append(lst, zz)

	sdb.MockZoneZipList = &lst

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
		"incId": "2",
		//"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetZoneZipListByInclusion(w, r)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestSix910Handler_GetZoneZipListByInclusionAuth(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.ID = 2
	sr.RegionID = 1
	sdb.MockIncludedSubRegion = &sr

	var zz sdbi.ZoneZip
	zz.ExcludedSubRegionID = 4
	zz.ID = 22
	zz.ZipCode = "12345"

	var lst []sdbi.ZoneZip
	lst = append(lst, zz)

	sdb.MockZoneZipList = &lst

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
		"incId":   "2",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.GetZoneZipListByInclusion(w, r)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteZoneZip(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.ID = 2
	sr.RegionID = 1
	sdb.MockIncludedSubRegion = &sr

	var zz sdbi.ZoneZip
	zz.ID = 2
	zz.IncludedSubRegionID = 2

	var lst []sdbi.ZoneZip
	lst = append(lst, zz)

	sdb.MockZoneZipList = &lst

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteZoneZipSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":      "2",
		"incId":   "2",
		"exId":    "0",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteZoneZip(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete zip zone in test: ", string(body))

	if w.Code != 200 || !bdy.Success {
		t.Fail()
	}
}

func TestSix910Handler_DeleteZoneZipReq(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.ID = 2
	sr.RegionID = 1
	sdb.MockIncludedSubRegion = &sr

	var zz sdbi.ZoneZip
	zz.ID = 2
	zz.IncludedSubRegionID = 2

	var lst []sdbi.ZoneZip
	lst = append(lst, zz)

	sdb.MockZoneZipList = &lst

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteZoneZipSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":      "2d",
		"incId":   "2",
		"exId":    "0",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteZoneZip(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete zip zone in test: ", string(body))

	if w.Code != 400 || bdy.Success {
		t.Fail()
	}
}

func TestSix910Handler_DeleteZoneZipReq2(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.ID = 2
	sr.RegionID = 1
	sdb.MockIncludedSubRegion = &sr

	var zz sdbi.ZoneZip
	zz.ID = 2
	zz.IncludedSubRegionID = 2

	var lst []sdbi.ZoneZip
	lst = append(lst, zz)

	sdb.MockZoneZipList = &lst

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteZoneZipSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":    "2",
		"incId": "2",
		"exId":  "0",
		//"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteZoneZip(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete zip zone in test: ", string(body))

	if w.Code != 400 || bdy.Success {
		t.Fail()
	}
}

func TestSix910Handler_DeleteZoneZipFail(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.ID = 2
	sr.RegionID = 1
	sdb.MockIncludedSubRegion = &sr

	var zz sdbi.ZoneZip
	zz.ID = 2
	zz.IncludedSubRegionID = 2

	var lst []sdbi.ZoneZip
	lst = append(lst, zz)

	sdb.MockZoneZipList = &lst

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	//sdb.MockDeleteZoneZipSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":      "2",
		"incId":   "2",
		"exId":    "0",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteZoneZip(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete zip zone in test: ", string(body))

	if w.Code != 500 || bdy.Success {
		t.Fail()
	}
}

func TestSix910Handler_DeleteZoneZipAuth(t *testing.T) {
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

	var reg sdbi.Region
	reg.ID = 1
	reg.StoreID = 5
	sdb.MockRegion = &reg

	var sr sdbi.IncludedSubRegion
	sr.ID = 2
	sr.RegionID = 1
	sdb.MockIncludedSubRegion = &sr

	var zz sdbi.ZoneZip
	zz.ID = 2
	zz.IncludedSubRegionID = 2

	var lst []sdbi.ZoneZip
	lst = append(lst, zz)

	sdb.MockZoneZipList = &lst

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	sdb.MockDeleteZoneZipSuccess = true

	//h := sh.GetNew()

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "storeName":"TestStore", "city":"atlanta", "OauthClientID": 2}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id":      "2",
		"incId":   "2",
		"exId":    "0",
		"storeId": "5",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")

	r.Header.Set("superAdminRole", "superAdmin")

	w := httptest.NewRecorder()

	h.DeleteZoneZip(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy man.Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body delete zip zone in test: ", string(body))

	if w.Code != 401 || bdy.Success {
		t.Fail()
	}
}
