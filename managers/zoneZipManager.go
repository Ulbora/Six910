package managers

import (
	"net/http"

	sdbi "github.com/Ulbora/six910-database-interface"
)

/*
 Six910 is a shopping cart and E-commerce system.

 Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2020 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

//AddZoneZip AddZoneZip
func (m *Six910Manager) AddZoneZip(z *sdbi.ZoneZip, sid int64) *ResponseID {
	var rtn ResponseID
	var r *sdbi.Region
	if z.IncludedSubRegionID != 0 {
		isr := m.Db.GetIncludedSubRegion(z.IncludedSubRegionID)
		m.Log.Debug("isr :", isr)
		r = m.Db.GetRegion(isr.RegionID)
	} else if z.ExcludedSubRegionID != 0 {
		esr := m.Db.GetExcludedSubRegion(z.ExcludedSubRegionID)
		r = m.Db.GetRegion(esr.RegionID)
	}
	if r != nil && r.StoreID == sid {
		suc, id := m.Db.AddZoneZip(z)
		if suc && id != 0 {
			rtn.Success = suc
			rtn.ID = id
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusBadRequest
		}
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//GetZoneZipListByExclusion GetZoneZipListByExclusion
func (m *Six910Manager) GetZoneZipListByExclusion(exID int64, sid int64) *[]sdbi.ZoneZip {
	var rtn *[]sdbi.ZoneZip
	esr := m.Db.GetExcludedSubRegion(exID)
	r := m.Db.GetRegion(esr.RegionID)
	if r.StoreID == sid {
		rtn = m.Db.GetZoneZipListByExclusion(exID)
	} else {
		var nz = []sdbi.ZoneZip{}
		rtn = &nz
	}
	return rtn
}

//GetZoneZipListByInclusion GetZoneZipListByInclusion
func (m *Six910Manager) GetZoneZipListByInclusion(incID int64, sid int64) *[]sdbi.ZoneZip {
	var rtn *[]sdbi.ZoneZip
	isr := m.Db.GetIncludedSubRegion(incID)
	r := m.Db.GetRegion(isr.RegionID)
	if r.StoreID == sid {
		rtn = m.Db.GetZoneZipListByInclusion(incID)
	} else {
		var nz = []sdbi.ZoneZip{}
		rtn = &nz
	}
	return rtn
}

//DeleteZoneZip DeleteZoneZip
func (m *Six910Manager) DeleteZoneZip(id int64, incID int64, exID int64, sid int64) *Response {
	var rtn Response
	var zl *[]sdbi.ZoneZip
	var r *sdbi.Region
	if incID != 0 {
		in := m.Db.GetIncludedSubRegion(incID)
		r = m.Db.GetRegion(in.RegionID)
		zl = m.Db.GetZoneZipListByInclusion(incID)
	} else if exID != 0 {
		ex := m.Db.GetExcludedSubRegion(exID)
		r = m.Db.GetRegion(ex.RegionID)
		zl = m.Db.GetZoneZipListByExclusion(exID)
	}
	if zl != nil && r != nil && r.StoreID == sid {
		var found bool
		for _, z := range *zl {
			if z.ID == id {
				found = true
				break
			}
		}
		if found {
			suc := m.Db.DeleteZoneZip(id)
			if suc {
				rtn.Success = true
				rtn.Code = http.StatusOK
			} else {
				rtn.Code = http.StatusBadRequest
			}
		} else {
			rtn.Code = http.StatusUnprocessableEntity
		}
	} else {
		rtn.Code = http.StatusInternalServerError
	}
	return &rtn
}
