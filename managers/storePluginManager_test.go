package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddStorePlugin(t *testing.T) {
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

	sdb.MockAddStorePluginSuccess = true
	sdb.MockStorePluginID = 2

	res := m.AddStorePlugin(&sp)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddStorePluginfail(t *testing.T) {
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

	sdb.MockAddStorePluginSuccess = true
	//sdb.MockStorePluginID = 2

	res := m.AddStorePlugin(&sp)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateStorePlugin(t *testing.T) {
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

	sdb.MockStorePlugin = &sp

	sdb.MockUpdateStorePluginSuccess = true

	res := m.UpdateStorePlugin(&sp)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateStorePluginfail(t *testing.T) {
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
	sp.StoreID = 55

	var sp2 sdbi.StorePlugins
	sp2.APIKey = "1234"
	sp2.StoreID = 5

	sdb.MockStorePlugin = &sp

	sdb.MockUpdateStorePluginSuccess = true

	res := m.UpdateStorePlugin(&sp2)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateStorePluginFail2(t *testing.T) {
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

	sdb.MockStorePlugin = &sp

	//sdb.MockUpdateStorePluginSuccess = true

	res := m.UpdateStorePlugin(&sp)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetStorePlugin(t *testing.T) {
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
	fsp := m.GetStorePlugin(4, 5)
	if fsp.APIKey != sp.APIKey {
		t.Fail()
	}
}

func TestSix910Manager_GetStorePluginFAil(t *testing.T) {
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
	sp.StoreID = 55

	sdb.MockStorePlugin = &sp
	fsp := m.GetStorePlugin(4, 5)
	if fsp.APIKey == sp.APIKey {
		t.Fail()
	}
}

func TestSix910Manager_GetStorePluginList(t *testing.T) {
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

	var lst []sdbi.StorePlugins
	lst = append(lst, sp)

	sdb.MockStorePluginList = &lst

	flst := m.GetStorePluginList(5)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteStorePlugin(t *testing.T) {
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

	sdb.MockDeleteStorePluginSuccess = true

	res := m.DeleteStorePlugin(4, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteStorePluginFail(t *testing.T) {
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
	sp.StoreID = 55

	sdb.MockStorePlugin = &sp

	sdb.MockDeleteStorePluginSuccess = true

	res := m.DeleteStorePlugin(4, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteStorePluginFail2(t *testing.T) {
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

	//sdb.MockDeleteStorePluginSuccess = true

	res := m.DeleteStorePlugin(4, 5)
	if res.Success {
		t.Fail()
	}
}
