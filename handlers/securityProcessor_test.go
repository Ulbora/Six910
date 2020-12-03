package handlers

import (
	"fmt"
	"net/http"
	"testing"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	lg "github.com/Ulbora/Level_Logger"
	man "github.com/Ulbora/Six910/managers"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
	"golang.org/x/crypto/bcrypt"
)

func TestSix910Handler_processSecurity(t *testing.T) {
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
	sdb.MockSecurity = &sec
	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	sdb.MockStore = &str

	var lu sdbi.LocalAccount
	lu.CustomerID = 2
	lu.Role = "StoreAdmin"
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

	// var us man.User
	// //us.CustomerID = 2
	// //us.Role = "StoreAdmin"
	// us.StoreID = 4
	// us.Username = "tester"
	//sdbi.mock

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	//h := sh.GetNew()

	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester1")

	var c jv.Claim
	c.Role = "StoreAdmin"

	auth := sh.processSecurity(r, &c)
	if !auth {
		t.Fail()
	}
}

func TestSix910Handler_processSecurityOauth(t *testing.T) {
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

	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")

	var c jv.Claim
	c.Role = "StoreAdmin"

	auth := sh.processSecurity(r, &c)
	if !auth {
		t.Fail()
	}
}

func TestSix910Handler_processBasicSecurity(t *testing.T) {
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
	sdb.MockSecurity = &sec
	m := sm.GetNew()

	var str sdbi.Store
	str.ID = 4
	str.StoreName = "TestStore"
	str.LocalDomain = "test.domain"
	sdb.MockStore = &str

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

	// var us man.User
	// //us.CustomerID = 2
	// //us.Role = "StoreAdmin"
	// us.StoreID = 4
	// us.Username = "tester"
	//sdbi.mock

	var sh Six910Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	//h := sh.GetNew()

	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("apiKey", "123456")
	r.SetBasicAuth("tester", "tester1")

	var c jv.Claim
	c.Role = "StoreAdmin"

	auth := sh.processBasicSecurity(r, &c)
	if !auth {
		t.Fail()
	}
}
