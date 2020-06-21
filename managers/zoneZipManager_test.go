package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddZoneZipIn(t *testing.T) {
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
	sdb.MockIncludedSubRegion = &sr

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	var z sdbi.ZoneZip
	z.IncludedSubRegionID = 4
	z.ZipCode = "12345"

	sdb.MockZoneZipID = 1
	sdb.MockAddZoneZipSuccess = true

	res := m.AddZoneZip(&z, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddZoneZipEx(t *testing.T) {
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
	sdb.MockExcludedSubRegion = &sr

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	var z sdbi.ZoneZip
	z.ExcludedSubRegionID = 4
	z.ZipCode = "12345"

	sdb.MockZoneZipID = 1
	sdb.MockAddZoneZipSuccess = true

	res := m.AddZoneZip(&z, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddZoneZipExFail(t *testing.T) {
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
	sdb.MockExcludedSubRegion = &sr

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg

	var z sdbi.ZoneZip
	z.ExcludedSubRegionID = 4
	z.ZipCode = "12345"

	sdb.MockZoneZipID = 1
	sdb.MockAddZoneZipSuccess = true

	res := m.AddZoneZip(&z, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_AddZoneZipExFail2(t *testing.T) {
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
	sdb.MockExcludedSubRegion = &sr

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	var z sdbi.ZoneZip
	z.ExcludedSubRegionID = 4
	z.ZipCode = "12345"

	sdb.MockZoneZipID = 1
	//sdb.MockAddZoneZipSuccess = true

	res := m.AddZoneZip(&z, 5)
	if res.Success {
		t.Fail()
	}
}

func TestSix910Manager_GetZoneZipListByExclusion(t *testing.T) {
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
	sdb.MockExcludedSubRegion = &sr

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	var z sdbi.ZoneZip
	z.ExcludedSubRegionID = 4
	z.ZipCode = "12345"

	var zlst []sdbi.ZoneZip
	zlst = append(zlst, z)

	sdb.MockZoneZipList = &zlst

	fezl := m.GetZoneZipListByExclusion(2, 5)
	if len(*fezl) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetZoneZipListByExclusionFail(t *testing.T) {
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
	sdb.MockExcludedSubRegion = &sr

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg

	var z sdbi.ZoneZip
	z.ExcludedSubRegionID = 4
	z.ZipCode = "12345"

	var zlst []sdbi.ZoneZip
	zlst = append(zlst, z)

	sdb.MockZoneZipList = &zlst

	fezl := m.GetZoneZipListByExclusion(2, 5)
	if len(*fezl) != 0 {
		t.Fail()
	}
}

func TestSix910Manager_GetZoneZipListByInclusion(t *testing.T) {
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
	sdb.MockIncludedSubRegion = &sr

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	var z sdbi.ZoneZip
	z.ExcludedSubRegionID = 4
	z.ZipCode = "12345"

	var zlst []sdbi.ZoneZip
	zlst = append(zlst, z)

	sdb.MockZoneZipList = &zlst

	fezl := m.GetZoneZipListByInclusion(2, 5)
	if len(*fezl) != 1 {
		t.Fail()
	}
}

func TestSix910Manager_GetZoneZipListByInclusionFail(t *testing.T) {
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
	sdb.MockIncludedSubRegion = &sr

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg

	var z sdbi.ZoneZip
	z.ExcludedSubRegionID = 4
	z.ZipCode = "12345"

	var zlst []sdbi.ZoneZip
	zlst = append(zlst, z)

	sdb.MockZoneZipList = &zlst

	fezl := m.GetZoneZipListByInclusion(2, 5)
	if len(*fezl) != 0 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteZoneZipIn(t *testing.T) {
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
	sdb.MockIncludedSubRegion = &sr

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	var z sdbi.ZoneZip
	z.ID = 3
	z.IncludedSubRegionID = 4
	z.ZipCode = "12345"

	var zlst []sdbi.ZoneZip
	zlst = append(zlst, z)
	sdb.MockZoneZipList = &zlst

	sdb.MockDeleteZoneZipSuccess = true

	res := m.DeleteZoneZip(3, 4, 4, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteZoneZipEx(t *testing.T) {
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
	sdb.MockExcludedSubRegion = &sr

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	var z sdbi.ZoneZip
	z.ID = 3
	z.ExcludedSubRegionID = 4
	z.ZipCode = "12345"

	var zlst []sdbi.ZoneZip
	zlst = append(zlst, z)
	sdb.MockZoneZipList = &zlst

	sdb.MockDeleteZoneZipSuccess = true

	res := m.DeleteZoneZip(3, 0, 4, 5)
	if !res.Success {
		t.Fail()
	}
}

func TestSix910Manager_DeleteZoneZipExFail1(t *testing.T) {
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
	sdb.MockExcludedSubRegion = &sr

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	var z sdbi.ZoneZip
	z.ID = 3
	z.ExcludedSubRegionID = 4
	z.ZipCode = "12345"

	var zlst []sdbi.ZoneZip
	zlst = append(zlst, z)
	sdb.MockZoneZipList = &zlst

	//sdb.MockDeleteZoneZipSuccess = true

	res := m.DeleteZoneZip(3, 0, 4, 5)
	if res.Success || res.Code != 400 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteZoneZipExFail2(t *testing.T) {
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
	sdb.MockExcludedSubRegion = &sr

	var rg sdbi.Region
	rg.StoreID = 5
	sdb.MockRegion = &rg

	var z sdbi.ZoneZip
	z.ID = 4
	z.ExcludedSubRegionID = 4
	z.ZipCode = "12345"

	var zlst []sdbi.ZoneZip
	zlst = append(zlst, z)
	sdb.MockZoneZipList = &zlst

	sdb.MockDeleteZoneZipSuccess = true

	res := m.DeleteZoneZip(3, 0, 4, 5)
	if res.Success || res.Code != 422 {
		t.Fail()
	}
}

func TestSix910Manager_DeleteZoneZipExFail3(t *testing.T) {
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
	sdb.MockExcludedSubRegion = &sr

	var rg sdbi.Region
	rg.StoreID = 55
	sdb.MockRegion = &rg

	var z sdbi.ZoneZip
	z.ID = 3
	z.ExcludedSubRegionID = 4
	z.ZipCode = "12345"

	var zlst []sdbi.ZoneZip
	zlst = append(zlst, z)
	sdb.MockZoneZipList = &zlst

	sdb.MockDeleteZoneZipSuccess = true

	res := m.DeleteZoneZip(3, 0, 4, 5)
	if res.Success || res.Code != 500 {
		t.Fail()
	}
}
