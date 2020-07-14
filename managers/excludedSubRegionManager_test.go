package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddExcludedSubRegion(t *testing.T) {
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

	var sr sdbi.ExcludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg
	sdb.MockAddExcludedSubRegionSuccess = true
	sdb.MockExcludedSubRegionID = 3

	res := m.AddExcludedSubRegion(&sr, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddExcludedSubRegionFail(t *testing.T) {
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

	var sr sdbi.ExcludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg
	sdb.MockAddExcludedSubRegionSuccess = true
	sdb.MockExcludedSubRegionID = 3

	res := m.AddExcludedSubRegion(&sr, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddExcludedSubRegionFail2(t *testing.T) {
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

	var sr sdbi.ExcludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg
	//sdb.MockAddExcludedSubRegionSuccess = true
	sdb.MockExcludedSubRegionID = 3

	res := m.AddExcludedSubRegion(&sr, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateExcludedSubRegion(t *testing.T) {
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

	var sr sdbi.ExcludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	sdb.MockUpdateExcludedSubRegionSuccess = true

	res := m.UpdateExcludedSubRegion(&sr, 5)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateExcludedSubRegionFail(t *testing.T) {
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

	var sr sdbi.ExcludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg

	sdb.MockUpdateExcludedSubRegionSuccess = true

	res := m.UpdateExcludedSubRegion(&sr, 5)

	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateExcludedSubRegionfail2(t *testing.T) {
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

	var sr sdbi.ExcludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	//sdb.MockUpdateExcludedSubRegionSuccess = true

	res := m.UpdateExcludedSubRegion(&sr, 5)

	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetExcludedSubRegion(t *testing.T) {
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

	var sr sdbi.ExcludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	sdb.MockExcludedSubRegion = &sr

	fesr := m.GetExcludedSubRegion(2, 5)
	if fesr.SubRegionID != sr.SubRegionID {
		t.Fail()
	}
}

func TestSix910Manager_GetExcludedSubRegionFail(t *testing.T) {
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

	var sr sdbi.ExcludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg

	sdb.MockExcludedSubRegion = &sr

	fesr := m.GetExcludedSubRegion(2, 5)
	if fesr.SubRegionID == sr.SubRegionID {
		t.Fail()
	}
}

func TestSix910Manager_GetExcludedSubRegionList(t *testing.T) {
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

	var sr sdbi.ExcludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	var srlst []sdbi.ExcludedSubRegion
	srlst = append(srlst, sr)

	sdb.MockExcludedSubRegionList = &srlst

	flst := m.GetExcludedSubRegionList(2, 5)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetExcludedSubRegionListFail(t *testing.T) {
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

	var sr sdbi.ExcludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg

	var srlst []sdbi.ExcludedSubRegion
	srlst = append(srlst, sr)

	sdb.MockExcludedSubRegionList = &srlst

	flst := m.GetExcludedSubRegionList(2, 5)
	if len(*flst) != 0 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteExcludedSubRegion(t *testing.T) {
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

	// var sr sdbi.ExcludedSubRegion
	// sr.SubRegionID = 2
	// sr.RegionID = 2

	var sr sdbi.ExcludedSubRegion
	sr.ID = 3
	sr.SubRegionID = 2
	sr.RegionID = 2
	var lst []sdbi.ExcludedSubRegion
	lst = append(lst, sr)

	sdb.MockExcludedSubRegionList = &lst

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	sdb.MockExcludedSubRegion = &sr

	sdb.MockDeleteExcludedSubRegionSuccess = true

	res := m.DeleteExcludedSubRegion(3, 2, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteExcludedSubRegionFail(t *testing.T) {
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

	var sr sdbi.ExcludedSubRegion
	sr.ID = 3
	sr.SubRegionID = 2
	sr.RegionID = 2
	var lst []sdbi.ExcludedSubRegion
	lst = append(lst, sr)

	sdb.MockExcludedSubRegionList = &lst

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg

	sdb.MockExcludedSubRegion = &sr

	sdb.MockDeleteExcludedSubRegionSuccess = true

	res := m.DeleteExcludedSubRegion(3, 2, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteExcludedSubRegionFail2(t *testing.T) {
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

	var sr sdbi.ExcludedSubRegion
	sr.ID = 3
	sr.SubRegionID = 2
	sr.RegionID = 2
	var lst []sdbi.ExcludedSubRegion
	lst = append(lst, sr)

	sdb.MockExcludedSubRegionList = &lst

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	sdb.MockExcludedSubRegion = &sr

	//sdb.MockDeleteExcludedSubRegionSuccess = true

	res := m.DeleteExcludedSubRegion(3, 2, 5)
	if res.Success {
		t.Fail()
	}
}
