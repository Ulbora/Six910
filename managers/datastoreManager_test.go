package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddLocalDatastore(t *testing.T) {
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

	var ds sdbi.LocalDataStore
	ds.DataStoreName = "test"
	ds.StoreID = 5

	sdb.MockAddLocalDataStoreSuccess = true

	res := m.AddLocalDatastore(&ds)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddLocalDatastoreFail(t *testing.T) {
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

	var ds sdbi.LocalDataStore
	ds.DataStoreName = "test"
	ds.StoreID = 5

	//sdb.MockAddLocalDataStoreSuccess = true

	res := m.AddLocalDatastore(&ds)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateLocalDatastore(t *testing.T) {
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

	var ds sdbi.LocalDataStore
	ds.DataStoreName = "test"
	ds.StoreID = 5

	sdb.MockLocalDataStore = &ds

	sdb.MockUpdateLocalDataStoreSuccess = true

	res := m.UpdateLocalDatastore(&ds)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateLocalDatastoreFail(t *testing.T) {
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

	var ds sdbi.LocalDataStore
	ds.DataStoreName = "test"
	ds.StoreID = 5

	sdb.MockLocalDataStore = &ds

	sdb.MockUpdateLocalDataStoreSuccess = true

	var ds2 sdbi.LocalDataStore
	ds2.DataStoreName = "test"
	ds2.StoreID = 54

	res := m.UpdateLocalDatastore(&ds2)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateLocalDatastoreFail2(t *testing.T) {
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

	var ds sdbi.LocalDataStore
	ds.DataStoreName = "test"
	ds.StoreID = 5

	sdb.MockLocalDataStore = &ds

	//sdb.MockUpdateLocalDataStoreSuccess = true

	res := m.UpdateLocalDatastore(&ds)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetLocalDatastore(t *testing.T) {
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

	var ds sdbi.LocalDataStore
	ds.DataStoreName = "test"
	ds.StoreID = 5
	ds.Reload = true

	sdb.MockLocalDataStore = &ds

	fds := m.GetLocalDatastore(5, "test")
	if fds.Reload != ds.Reload {
		t.Fail()
	}
}
