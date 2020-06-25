package managers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddOAuthUser(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l
	sm.UserHost = "http://localhost:3001"
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "code": 200}`))
	p.MockResp = &ress
	p.MockRespCode = 200
	sm.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", sm.Proxy)

	m := sm.GetNew()

	var u OAuthUser
	u.ClientID = 5
	u.Enabled = true

	var ath Auth
	ath.ClientID = "5"
	ath.Token = "12345"

	res := m.AddOAuthUser(&u, &ath)
	fmt.Println("res in oauth add user: ", res)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateOAuthUser(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l
	sm.UserHost = "http://localhost:3001"
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "code": 200}`))
	p.MockResp = &ress
	p.MockRespCode = 200
	sm.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", sm.Proxy)

	m := sm.GetNew()

	var u OAuthUser
	u.ClientID = 5
	u.Enabled = true

	var ath Auth
	ath.ClientID = "5"
	ath.Token = "12345"

	res := m.UpdateOAuthUser(&u, &ath)
	fmt.Println("res in oauth update user: ", res)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetOAuthUser(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l
	sm.UserHost = "http://localhost:3001"
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "roleId":1, "enabled": true}`))
	p.MockResp = &ress
	p.MockRespCode = 200
	sm.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", sm.Proxy)

	m := sm.GetNew()

	var u OAuthUser
	u.ClientID = 5
	u.Enabled = true

	var ath Auth
	ath.ClientID = "5"
	ath.Token = "12345"

	fu, code := m.GetOAuthUser("test", "4", &ath)
	fmt.Println("res in oauth get user: ", fu)
	if code != 200 || !fu.Enabled {
		t.Fail()
	}
}

func TestSix910Manager_GetOAuthUserList(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l
	sm.UserHost = "http://localhost:3001"
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"username":"tester", "roleId":2, "enabled": true}]`))
	p.MockResp = &ress
	p.MockRespCode = 200
	sm.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", sm.Proxy)

	m := sm.GetNew()

	var u OAuthUser
	u.ClientID = 5
	u.Enabled = true

	var ath Auth
	ath.ClientID = "5"
	ath.Token = "12345"

	fls, code := m.GetOAuthUserList("3", &ath)
	fmt.Println("res in oauth get list user: ", fls)
	if code != 200 || len(*fls) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteOAuthUser(t *testing.T) {
	var sdb sixmdb.MockSix910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	//sdb.DB = dbi
	//dbi.Connect()

	var sm Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l
	sm.UserHost = "http://localhost:3001"
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "code": 200}`))
	p.MockResp = &ress
	p.MockRespCode = 200
	sm.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", sm.Proxy)

	m := sm.GetNew()

	var u OAuthUser
	u.ClientID = 5
	u.Enabled = true

	var ath Auth
	ath.ClientID = "5"
	ath.Token = "12345"

	res := m.DeleteOAuthUser("tester", "3", &ath)
	if !res.Success {
		t.Fail()
	}
}
