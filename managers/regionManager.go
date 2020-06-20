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

//AddRegion AddRegion
func (m *Six910Manager) AddRegion(r *sdbi.Region) *ResponseID {
	var rtn ResponseID
	suc, id := m.Db.AddRegion(r)
	if suc && id != 0 {
		rtn.Success = suc
		rtn.ID = id
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateRegion UpdateRegion
func (m *Six910Manager) UpdateRegion(r *sdbi.Region) *Response {
	var rtn Response
	reg := m.Db.GetRegion(r.ID)
	if reg.StoreID == r.StoreID {
		suc := m.Db.UpdateRegion(r)
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

//GetRegion GetRegion
func (m *Six910Manager) GetRegion(id int64, sid int64) *sdbi.Region {
	var rtn *sdbi.Region
	reg := m.Db.GetRegion(id)
	if reg.StoreID == sid {
		rtn = reg
	} else {
		var nr sdbi.Region
		rtn = &nr
	}
	return rtn
}

//GetRegionList GetRegionList
func (m *Six910Manager) GetRegionList(storeID int64) *[]sdbi.Region {
	return m.Db.GetRegionList(storeID)
}

//DeleteRegion DeleteRegion
func (m *Six910Manager) DeleteRegion(id int64, sid int64) *Response {
	var rtn Response
	reg := m.Db.GetRegion(id)
	if reg.StoreID == sid {
		suc := m.Db.DeleteRegion(id)
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
