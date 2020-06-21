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

//AddExcludedSubRegion AddExcludedSubRegion
func (m *Six910Manager) AddExcludedSubRegion(e *sdbi.ExcludedSubRegion, sid int64) *ResponseID {
	var rtn ResponseID
	r := m.Db.GetRegion(e.RegionID)
	if r.StoreID == sid {
		suc, id := m.Db.AddExcludedSubRegion(e)
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

//UpdateExcludedSubRegion UpdateExcludedSubRegion
func (m *Six910Manager) UpdateExcludedSubRegion(e *sdbi.ExcludedSubRegion, sid int64) *Response {
	var rtn Response
	r := m.Db.GetRegion(e.RegionID)
	if r.StoreID == sid {
		suc := m.Db.UpdateExcludedSubRegion(e)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusBadRequest
		}
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//GetExcludedSubRegion GetExcludedSubRegion
func (m *Six910Manager) GetExcludedSubRegion(id int64, sid int64) *sdbi.ExcludedSubRegion {
	var rtn *sdbi.ExcludedSubRegion
	sr := m.Db.GetExcludedSubRegion(id)
	r := m.Db.GetRegion(sr.RegionID)
	if r.StoreID == sid {
		rtn = sr
	} else {
		var ns sdbi.ExcludedSubRegion
		rtn = &ns
	}
	return rtn
}

//GetExcludedSubRegionList GetExcludedSubRegionList
func (m *Six910Manager) GetExcludedSubRegionList(regionID int64, sid int64) *[]sdbi.ExcludedSubRegion {
	var rtn *[]sdbi.ExcludedSubRegion
	r := m.Db.GetRegion(regionID)
	if r.StoreID == sid {
		rtn = m.Db.GetExcludedSubRegionList(regionID)
	} else {
		var nr = []sdbi.ExcludedSubRegion{}
		rtn = &nr
	}
	return rtn
}

//DeleteExcludedSubRegion DeleteExcludedSubRegion
func (m *Six910Manager) DeleteExcludedSubRegion(id int64, sid int64) *Response {
	var rtn Response
	sr := m.Db.GetExcludedSubRegion(id)
	r := m.Db.GetRegion(sr.RegionID)
	if r.StoreID == sid {
		suc := m.Db.DeleteExcludedSubRegion(id)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusBadRequest
		}
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}
