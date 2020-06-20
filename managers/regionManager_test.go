package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddRegion(t *testing.T) {
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

	var reg sdbi.Region
	reg.StoreID = 3

	sdb.MockAddRegionSuccess = true
	sdb.MockRegionID = 5

	res := m.AddRegion(&reg)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddRegionFail(t *testing.T) {
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

	var reg sdbi.Region
	reg.StoreID = 3

	sdb.MockAddRegionSuccess = true
	//sdb.MockRegionID = 5

	res := m.AddRegion(&reg)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateRegion(t *testing.T) {
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

	var reg sdbi.Region
	reg.StoreID = 3

	sdb.MockRegion = &reg

	sdb.MockUpdateRegionSuccess = true

	res := m.UpdateRegion(&reg)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateRegionFail(t *testing.T) {
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

	var reg sdbi.Region
	reg.StoreID = 3

	sdb.MockRegion = &reg

	sdb.MockUpdateRegionSuccess = true

	var reg2 sdbi.Region
	reg2.StoreID = 32

	res := m.UpdateRegion(&reg2)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateRegionFail2(t *testing.T) {
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

	var reg sdbi.Region
	reg.StoreID = 3

	sdb.MockRegion = &reg

	//sdb.MockUpdateRegionSuccess = true

	res := m.UpdateRegion(&reg)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetRegion(t *testing.T) {
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

	var reg sdbi.Region
	reg.Name = "test"
	reg.StoreID = 3

	sdb.MockRegion = &reg

	freg := m.GetRegion(3, 3)
	if freg.Name != reg.Name {
		t.Fail()
	}
}

func TestSix910Manager_GetRegionFail(t *testing.T) {
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

	var reg sdbi.Region
	reg.Name = "test"
	reg.StoreID = 33

	sdb.MockRegion = &reg

	freg := m.GetRegion(3, 3)
	if freg.Name == reg.Name {
		t.Fail()
	}
}

func TestSix910Manager_GetRegionList(t *testing.T) {
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

	var reg sdbi.Region
	reg.Name = "test"
	reg.StoreID = 33

	var rlst []sdbi.Region
	rlst = append(rlst, reg)

	sdb.MockRegionList = &rlst

	flst := m.GetRegionList(33)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteRegion(t *testing.T) {
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

	var reg sdbi.Region
	reg.Name = "test"
	reg.StoreID = 3

	sdb.MockRegion = &reg
	sdb.MockDeleteRegionSuccess = true

	res := m.DeleteRegion(2, 3)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteRegionFail(t *testing.T) {
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

	var reg sdbi.Region
	reg.Name = "test"
	reg.StoreID = 34

	sdb.MockRegion = &reg
	sdb.MockDeleteRegionSuccess = true

	res := m.DeleteRegion(2, 3)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteRegionFail2(t *testing.T) {
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

	var reg sdbi.Region
	reg.Name = "test"
	reg.StoreID = 3

	sdb.MockRegion = &reg
	//sdb.MockDeleteRegionSuccess = true

	res := m.DeleteRegion(2, 3)
	if res.Success {
		t.Fail()
	}
}
