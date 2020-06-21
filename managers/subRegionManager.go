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

//AddSubRegion AddSubRegion
func (m *Six910Manager) AddSubRegion(s *sdbi.SubRegion) *ResponseID {
	var rtn ResponseID
	suc, id := m.Db.AddSubRegion(s)
	if suc && id != 0 {
		rtn.Success = suc
		rtn.ID = id
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateSubRegion UpdateSubRegion
func (m *Six910Manager) UpdateSubRegion(s *sdbi.SubRegion, sid int64) *Response {
	var rtn Response
	r := m.Db.GetRegion(s.RegionID)
	if r.StoreID == sid {
		suc := m.Db.UpdateSubRegion(s)
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

//GetSubRegion GetSubRegion
func (m *Six910Manager) GetSubRegion(id int64, sid int64) *sdbi.SubRegion {
	var rtn *sdbi.SubRegion
	sr := m.Db.GetSubRegion(id)
	r := m.Db.GetRegion(sr.RegionID)
	if r.StoreID == sid {
		rtn = sr
	} else {
		var ns sdbi.SubRegion
		rtn = &ns
	}
	return rtn
}

//GetSubRegionList GetSubRegionList
func (m *Six910Manager) GetSubRegionList(regionID int64, sid int64) *[]sdbi.SubRegion {
	var rtn *[]sdbi.SubRegion
	r := m.Db.GetRegion(regionID)
	if r.StoreID == sid {
		rtn = m.Db.GetSubRegionList(regionID)
	} else {
		var nr = []sdbi.SubRegion{}
		rtn = &nr
	}
	return rtn
}

//DeleteSubRegion DeleteSubRegion
func (m *Six910Manager) DeleteSubRegion(id int64, sid int64) *Response {
	var rtn Response
	sr := m.Db.GetSubRegion(id)
	r := m.Db.GetRegion(sr.RegionID)
	if r.StoreID == sid {
		suc := m.Db.DeleteSubRegion(id)
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
