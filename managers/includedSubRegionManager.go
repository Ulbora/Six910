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

//AddIncludedSubRegion AddIncludedSubRegion
func (m *Six910Manager) AddIncludedSubRegion(e *sdbi.IncludedSubRegion, sid int64) *ResponseID {
	var rtn ResponseID
	r := m.Db.GetRegion(e.RegionID)
	if r.StoreID == sid {
		suc, id := m.Db.AddIncludedSubRegion(e)
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

//UpdateIncludedSubRegion UpdateIncludedSubRegion
func (m *Six910Manager) UpdateIncludedSubRegion(e *sdbi.IncludedSubRegion, sid int64) *Response {
	var rtn Response
	r := m.Db.GetRegion(e.RegionID)
	if r.StoreID == sid {
		suc := m.Db.UpdateIncludedSubRegion(e)
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

//GetIncludedSubRegion GetIncludedSubRegion
func (m *Six910Manager) GetIncludedSubRegion(id int64, sid int64) *sdbi.IncludedSubRegion {
	var rtn *sdbi.IncludedSubRegion
	sr := m.Db.GetIncludedSubRegion(id)
	r := m.Db.GetRegion(sr.RegionID)
	if r.StoreID == sid {
		rtn = sr
	} else {
		var ns sdbi.IncludedSubRegion
		rtn = &ns
	}
	return rtn
}

//GetIncludedSubRegionList GetIncludedSubRegionList
func (m *Six910Manager) GetIncludedSubRegionList(regionID int64, sid int64) *[]sdbi.IncludedSubRegion {
	var rtn *[]sdbi.IncludedSubRegion
	r := m.Db.GetRegion(regionID)
	if r.StoreID == sid {
		rtn = m.Db.GetIncludedSubRegionList(regionID)
	} else {
		var nr = []sdbi.IncludedSubRegion{}
		rtn = &nr
	}
	return rtn
}

//DeleteIncludedSubRegion DeleteIncludedSubRegion
func (m *Six910Manager) DeleteIncludedSubRegion(id int64, sid int64) *Response {
	var rtn Response
	sr := m.Db.GetIncludedSubRegion(id)
	r := m.Db.GetRegion(sr.RegionID)
	if r.StoreID == sid {
		suc := m.Db.DeleteIncludedSubRegion(id)
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
