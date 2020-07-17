package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddInstance(t *testing.T) {
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

	var in sdbi.Instances
	in.DataStoreName = "test"
	in.StoreID = 5

	sdb.MockAddInstancesSuccess = true

	res := m.AddInstance(&in)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddInstanceFail(t *testing.T) {
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

	var in sdbi.Instances
	in.DataStoreName = "test"
	in.StoreID = 5

	//sdb.MockAddInstancesSuccess = true

	res := m.AddInstance(&in)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateInstance(t *testing.T) {
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

	var in sdbi.Instances
	in.DataStoreName = "test"
	in.StoreID = 5

	sdb.MockUpdateInstancesSuccess = true

	res := m.UpdateInstance(&in)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateInstanceFail(t *testing.T) {
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

	var in sdbi.Instances
	in.DataStoreName = "test"
	in.StoreID = 5

	//sdb.MockUpdateInstancesSuccess = true

	res := m.UpdateInstance(&in)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetInstance(t *testing.T) {
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

	var in sdbi.Instances
	in.DataStoreName = "test"
	in.StoreID = 5

	sdb.MockInstances = &in

	fin := m.GetInstance("test", "test", 5)
	if fin.DataStoreName != in.DataStoreName {
		t.Fail()
	}
}

func TestSix910Manager_GetInstanceList(t *testing.T) {
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

	var in sdbi.Instances
	in.DataStoreName = "test"
	in.StoreID = 5

	var lst []sdbi.Instances
	lst = append(lst, in)

	sdb.MockInstancesList = &lst

	fin := m.GetInstanceList("test", 5)
	if (*fin)[0].DataStoreName != in.DataStoreName {
		t.Fail()
	}
}
