package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddPlugin(t *testing.T) {
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

	var p sdbi.Plugins
	p.ActivateURL = "test"

	sdb.MockAddPluginSuccess = true
	sdb.MockPluginID = 4

	res := m.AddPlugin(&p)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddPluginFail(t *testing.T) {
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

	var p sdbi.Plugins
	p.ActivateURL = "test"

	//sdb.MockAddPluginSuccess = true
	sdb.MockPluginID = 4

	res := m.AddPlugin(&p)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdatePlugin(t *testing.T) {
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

	var p sdbi.Plugins
	p.ActivateURL = "test"

	sdb.MockUpdatePluginSuccess = true
	sdb.MockPluginID = 4

	res := m.UpdatePlugin(&p)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdatePluginFail(t *testing.T) {
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

	var p sdbi.Plugins
	p.ActivateURL = "test"

	//sdb.MockUpdatePluginSuccess = true
	sdb.MockPluginID = 4

	res := m.UpdatePlugin(&p)

	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetPlugin(t *testing.T) {
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

	var p sdbi.Plugins
	p.ActivateURL = "test"

	sdb.MockPlugin = &p

	fp := m.GetPlugin(8)
	if fp.ActivateURL != p.ActivateURL {
		t.Fail()
	}
}

func TestSix910Manager_GetPluginList(t *testing.T) {
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

	var p sdbi.Plugins
	p.ActivateURL = "test"

	var lst []sdbi.Plugins
	lst = append(lst, p)

	sdb.MockPluginList = &lst

	flst := m.GetPluginList(2, 4)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_DeletePlugin(t *testing.T) {
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

	sdb.MockDeletePluginSuccess = true

	res := m.DeletePlugin(4)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeletePluginFail(t *testing.T) {
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

	//sdb.MockDeletePluginSuccess = true

	res := m.DeletePlugin(4)
	if res.Success {
		t.Fail()
	}
}
