package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddSubRegion(t *testing.T) {
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

	var sr sdbi.SubRegion
	sr.Name = "sub"
	sr.RegionID = 2

	sdb.MockSubRegionID = 6
	sdb.MockAddSubRegionSuccess = true

	res := m.AddSubRegion(&sr)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddSubRegionFail(t *testing.T) {
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

	var sr sdbi.SubRegion
	sr.Name = "sub"
	sr.RegionID = 2

	sdb.MockSubRegionID = 6
	//sdb.MockAddSubRegionSuccess = true

	res := m.AddSubRegion(&sr)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateSubRegion(t *testing.T) {
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

	var sr sdbi.SubRegion
	sr.Name = "sub"
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	sdb.MockUpdateSubRegionSuccess = true

	res := m.UpdateSubRegion(&sr, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateSubRegionFail(t *testing.T) {
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

	var sr sdbi.SubRegion
	sr.Name = "sub"
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg

	sdb.MockUpdateSubRegionSuccess = true

	res := m.UpdateSubRegion(&sr, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateSubRegionFail2(t *testing.T) {
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

	var sr sdbi.SubRegion
	sr.Name = "sub"
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	//sdb.MockUpdateSubRegionSuccess = true

	res := m.UpdateSubRegion(&sr, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetSubRegion(t *testing.T) {
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

	var sr sdbi.SubRegion
	sr.Name = "sub"
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	sdb.MockSubRegion = &sr

	fsr := m.GetSubRegion(2, 5)
	if fsr.Name != sr.Name {
		t.Fail()
	}
}

func TestSix910Manager_GetSubRegionFail(t *testing.T) {
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

	var sr sdbi.SubRegion
	sr.Name = "sub"
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg

	sdb.MockSubRegion = &sr

	fsr := m.GetSubRegion(2, 5)
	if fsr.Name == sr.Name {
		t.Fail()
	}
}

func TestSix910Manager_GetSubRegionList(t *testing.T) {
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

	var sr sdbi.SubRegion
	sr.Name = "sub"
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	var rlst []sdbi.SubRegion
	rlst = append(rlst, sr)
	sdb.MockSubRegionList = &rlst

	flst := m.GetSubRegionList(4, 5)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetSubRegionListFail(t *testing.T) {
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

	var sr sdbi.SubRegion
	sr.Name = "sub"
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg

	var rlst []sdbi.SubRegion
	rlst = append(rlst, sr)
	sdb.MockSubRegionList = &rlst

	flst := m.GetSubRegionList(4, 5)
	if len(*flst) != 0 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteSubRegion(t *testing.T) {
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

	var sr sdbi.SubRegion
	sr.Name = "sub"
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg
	sdb.MockSubRegion = &sr

	sdb.MockDeleteSubRegionSuccess = true

	res := m.DeleteSubRegion(3, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteSubRegionFail(t *testing.T) {
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

	var sr sdbi.SubRegion
	sr.Name = "sub"
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg
	sdb.MockSubRegion = &sr

	sdb.MockDeleteSubRegionSuccess = true

	res := m.DeleteSubRegion(3, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteSubRegionFail2(t *testing.T) {
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

	var sr sdbi.SubRegion
	sr.Name = "sub"
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg
	sdb.MockSubRegion = &sr

	//sdb.MockDeleteSubRegionSuccess = true

	res := m.DeleteSubRegion(3, 5)
	if res.Success {
		t.Fail()
	}
}
