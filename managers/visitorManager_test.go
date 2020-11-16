package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	sixmdb "github.com/Ulbora/six910-mysql"
)

func TestSix910Manager_AddVisit(t *testing.T) {
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

	var v sdbi.Visitor
	v.Origin = "test"

	sdb.MockAddVisitorResp = true

	suc := m.AddVisit(&v)
	if !suc {
		t.Fail()
	}
}

func TestSix910Manager_GetVisitorData(t *testing.T) {
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

	var vd sdbi.VisitorData
	vd.VisitCount = 5
	var vdlst []sdbi.VisitorData
	vdlst = append(vdlst, vd)
	sdb.MockVisitorData = &vdlst

	fvlst := m.GetVisitorData(5)
	if len(*fvlst) == 0 {
		t.Fail()
	}
}
