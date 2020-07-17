package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddDataStoreWriteLock(t *testing.T) {
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

	var dl sdbi.DataStoreWriteLock
	dl.DataStoreName = "test"

	sdb.MockAddDataStoreWriteLockSuccess = true
	sdb.MockDataStoreWriteLockID = 3

	res := m.AddDataStoreWriteLock(&dl)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddDataStoreWriteLockFail(t *testing.T) {
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

	var dl sdbi.DataStoreWriteLock
	dl.DataStoreName = "test"

	//sdb.MockAddDataStoreWriteLockSuccess = true
	//sdb.MockDataStoreWriteLockID = 3

	res := m.AddDataStoreWriteLock(&dl)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateDataStoreWriteLock(t *testing.T) {
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

	var dl sdbi.DataStoreWriteLock
	dl.DataStoreName = "test"

	sdb.MockUpdateDataStoreWriteLockSuccess = true

	res := m.UpdateDataStoreWriteLock(&dl)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateDataStoreWriteLockFail(t *testing.T) {
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

	var dl sdbi.DataStoreWriteLock
	dl.DataStoreName = "test"

	//sdb.MockUpdateDataStoreWriteLockSuccess = true

	res := m.UpdateDataStoreWriteLock(&dl)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetDataStoreWriteLock(t *testing.T) {
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

	var dl sdbi.DataStoreWriteLock
	dl.DataStoreName = "test"

	sdb.MockDataStoreWriteLock = &dl

	fld := m.GetDataStoreWriteLock("test", 5)
	if fld.DataStoreName != dl.DataStoreName {
		t.Fail()
	}
}
