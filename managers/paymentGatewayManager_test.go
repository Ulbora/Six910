package managers

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddPaymentGateway(t *testing.T) {
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

	var sp sdbi.StorePlugins
	sp.APIKey = "1234"
	sp.StoreID = 5

	sdb.MockStorePlugin = &sp

	var gw sdbi.PaymentGateway
	gw.CheckoutURL = "test"
	gw.StorePluginsID = 3

	sdb.MockAddPaymentGatewaySuccess = true
	sdb.MockPaymentGatewayID = 2

	res := m.AddPaymentGateway(&gw, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddPaymentGatewayFail(t *testing.T) {
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

	var sp sdbi.StorePlugins
	sp.APIKey = "1234"
	sp.StoreID = 57

	sdb.MockStorePlugin = &sp

	var gw sdbi.PaymentGateway
	gw.CheckoutURL = "test"
	gw.StorePluginsID = 3

	sdb.MockAddPaymentGatewaySuccess = true
	sdb.MockPaymentGatewayID = 2

	res := m.AddPaymentGateway(&gw, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddPaymentGatewayFail2(t *testing.T) {
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

	var sp sdbi.StorePlugins
	sp.APIKey = "1234"
	sp.StoreID = 5

	sdb.MockStorePlugin = &sp

	var gw sdbi.PaymentGateway
	gw.CheckoutURL = "test"
	gw.StorePluginsID = 3

	//sdb.MockAddPaymentGatewaySuccess = true
	sdb.MockPaymentGatewayID = 2

	res := m.AddPaymentGateway(&gw, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdatePaymentGateway(t *testing.T) {
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

	var sp sdbi.StorePlugins
	sp.APIKey = "1234"
	sp.StoreID = 5

	sdb.MockStorePlugin = &sp

	var gw sdbi.PaymentGateway
	gw.CheckoutURL = "test"
	gw.StorePluginsID = 3

	sdb.MockUpdatePaymentGatewaySuccess = true

	res := m.UpdatePaymentGateway(&gw, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdatePaymentGatewayFail(t *testing.T) {
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

	var sp sdbi.StorePlugins
	sp.APIKey = "1234"
	sp.StoreID = 56

	sdb.MockStorePlugin = &sp

	var gw sdbi.PaymentGateway
	gw.CheckoutURL = "test"
	gw.StorePluginsID = 3

	sdb.MockUpdatePaymentGatewaySuccess = true

	res := m.UpdatePaymentGateway(&gw, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdatePaymentGatewayFail2(t *testing.T) {
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

	var sp sdbi.StorePlugins
	sp.APIKey = "1234"
	sp.StoreID = 5

	sdb.MockStorePlugin = &sp

	var gw sdbi.PaymentGateway
	gw.CheckoutURL = "test"
	gw.StorePluginsID = 3

	//sdb.MockUpdatePaymentGatewaySuccess = true

	res := m.UpdatePaymentGateway(&gw, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetPaymentGateway(t *testing.T) {
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

	var sp sdbi.StorePlugins
	sp.APIKey = "1234"
	sp.StoreID = 5

	sdb.MockStorePlugin = &sp

	var gw sdbi.PaymentGateway
	gw.CheckoutURL = "test"
	gw.StorePluginsID = 3

	sdb.MockPaymentGateway = &gw

	fg := m.GetPaymentGateway(3, 5)
	if fg.CheckoutURL != gw.CheckoutURL {
		t.Fail()
	}
}

func TestSix910Manager_GetPaymentGatewayFail(t *testing.T) {
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

	var sp sdbi.StorePlugins
	sp.APIKey = "1234"
	sp.StoreID = 56

	sdb.MockStorePlugin = &sp

	var gw sdbi.PaymentGateway
	gw.CheckoutURL = "test"
	gw.StorePluginsID = 3

	sdb.MockPaymentGateway = &gw

	fg := m.GetPaymentGateway(3, 5)
	if fg.CheckoutURL == gw.CheckoutURL {
		t.Fail()
	}
}

func TestSix910Manager_GetPaymentGatewayByName(t *testing.T) {
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

	var sp sdbi.StorePlugins
	sp.APIKey = "1234"
	sp.StoreID = 5

	sdb.MockStorePlugin = &sp

	var gw sdbi.PaymentGateway
	gw.CheckoutURL = "test"
	gw.StorePluginsID = 3
	gw.Name = "BTC"

	sdb.MockPaymentGateway = &gw

	fg := m.GetPaymentGatewayByName("BTC", 5)
	fmt.Println("fg: ", *fg)
	if fg.CheckoutURL != gw.CheckoutURL || fg.Name != "BTC" {
		t.Fail()
	}
	// t.Fail()
}

func TestSix910Manager_GetPaymentGateways(t *testing.T) {
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

	var sp sdbi.StorePlugins
	sp.APIKey = "1234"
	sp.StoreID = 56

	sdb.MockStorePlugin = &sp

	var gw sdbi.PaymentGateway
	gw.CheckoutURL = "test"
	gw.StorePluginsID = 3

	var lst []sdbi.PaymentGateway
	lst = append(lst, gw)

	sdb.MockPaymentGatewayList = &lst

	flst := m.GetPaymentGateways(5)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_DeletePaymentGateway(t *testing.T) {
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

	var sp sdbi.StorePlugins
	sp.APIKey = "1234"
	sp.StoreID = 5

	sdb.MockStorePlugin = &sp

	var gw sdbi.PaymentGateway
	gw.CheckoutURL = "test"
	gw.StorePluginsID = 3

	sdb.MockPaymentGateway = &gw

	sdb.MockDeletePaymentGatewaySuccess = true

	res := m.DeletePaymentGateway(4, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeletePaymentGatewayfail(t *testing.T) {
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

	var sp sdbi.StorePlugins
	sp.APIKey = "1234"
	sp.StoreID = 56

	sdb.MockStorePlugin = &sp

	var gw sdbi.PaymentGateway
	gw.CheckoutURL = "test"
	gw.StorePluginsID = 3

	sdb.MockPaymentGateway = &gw

	sdb.MockDeletePaymentGatewaySuccess = true

	res := m.DeletePaymentGateway(4, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeletePaymentGatewayFail2(t *testing.T) {
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

	var sp sdbi.StorePlugins
	sp.APIKey = "1234"
	sp.StoreID = 5

	sdb.MockStorePlugin = &sp

	var gw sdbi.PaymentGateway
	gw.CheckoutURL = "test"
	gw.StorePluginsID = 3

	sdb.MockPaymentGateway = &gw

	//sdb.MockDeletePaymentGatewaySuccess = true

	res := m.DeletePaymentGateway(4, 5)
	if res.Success {
		t.Fail()
	}
}
