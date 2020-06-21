package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddIncludedSubRegion(t *testing.T) {
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

	var sr sdbi.IncludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg
	sdb.MockAddIncludedSubRegionSuccess = true
	sdb.MockIncludedSubRegionID = 3

	res := m.AddIncludedSubRegion(&sr, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddIncludedSubRegionFail(t *testing.T) {
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

	var sr sdbi.IncludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg
	sdb.MockAddIncludedSubRegionSuccess = true
	sdb.MockIncludedSubRegionID = 3

	res := m.AddIncludedSubRegion(&sr, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddIncludedSubRegionFail2(t *testing.T) {
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

	var sr sdbi.IncludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg
	//sdb.MockAddIncludedSubRegionSuccess = true
	sdb.MockIncludedSubRegionID = 3

	res := m.AddIncludedSubRegion(&sr, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateIncludedSubRegion(t *testing.T) {
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

	var sr sdbi.IncludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	sdb.MockUpdateIncludedSubRegionSuccess = true

	res := m.UpdateIncludedSubRegion(&sr, 5)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateIncludedSubRegionFail(t *testing.T) {
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

	var sr sdbi.IncludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 53
	sdb.MockRegion = &rg

	sdb.MockUpdateIncludedSubRegionSuccess = true

	res := m.UpdateIncludedSubRegion(&sr, 5)

	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_UpdateIncludedSubRegionFail2(t *testing.T) {
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

	var sr sdbi.IncludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	//sdb.MockUpdateIncludedSubRegionSuccess = true

	res := m.UpdateIncludedSubRegion(&sr, 5)

	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetIncludedSubRegion(t *testing.T) {
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

	var sr sdbi.IncludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	sdb.MockIncludedSubRegion = &sr

	fesr := m.GetIncludedSubRegion(2, 5)
	if fesr.SubRegionID != sr.SubRegionID {
		t.Fail()
	}
}

func TestSix910Manager_GetIncludedSubRegionFail(t *testing.T) {
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

	var sr sdbi.IncludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 53
	sdb.MockRegion = &rg

	sdb.MockIncludedSubRegion = &sr

	fesr := m.GetIncludedSubRegion(2, 5)
	if fesr.SubRegionID == sr.SubRegionID {
		t.Fail()
	}
}

func TestSix910Manager_GetIncludedSubRegionList(t *testing.T) {
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

	var sr sdbi.IncludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	var srlst []sdbi.IncludedSubRegion
	srlst = append(srlst, sr)

	sdb.MockIncludedSubRegionList = &srlst

	flst := m.GetIncludedSubRegionList(2, 5)
	if len(*flst) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetIncludedSubRegionListFail(t *testing.T) {
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

	var sr sdbi.IncludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg

	var srlst []sdbi.IncludedSubRegion
	srlst = append(srlst, sr)

	sdb.MockIncludedSubRegionList = &srlst

	flst := m.GetIncludedSubRegionList(2, 5)
	if len(*flst) != 0 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteIncludedSubRegion(t *testing.T) {
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

	var sr sdbi.IncludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	sdb.MockIncludedSubRegion = &sr

	sdb.MockDeleteIncludedSubRegionSuccess = true

	res := m.DeleteIncludedSubRegion(3, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteIncludedSubRegionFail(t *testing.T) {
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

	var sr sdbi.IncludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg

	sdb.MockIncludedSubRegion = &sr

	sdb.MockDeleteIncludedSubRegionSuccess = true

	res := m.DeleteIncludedSubRegion(3, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteIncludedSubRegionfail2(t *testing.T) {
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

	var sr sdbi.IncludedSubRegion
	sr.SubRegionID = 2
	sr.RegionID = 2

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	sdb.MockIncludedSubRegion = &sr

	//sdb.MockDeleteIncludedSubRegionSuccess = true

	res := m.DeleteIncludedSubRegion(3, 5)
	if res.Success {
		t.Fail()
	}
}
