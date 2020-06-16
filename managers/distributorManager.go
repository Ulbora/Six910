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

//AddDistributor AddDistributor
func (m *Six910Manager) AddDistributor(d *sdbi.Distributor) *ResponseID {
	var rtn ResponseID
	suc, id := m.Db.AddDistributor(d)
	if suc && id != 0 {
		rtn.ID = id
		rtn.Success = suc
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateDistributor UpdateDistributor
func (m *Six910Manager) UpdateDistributor(d *sdbi.Distributor) *Response {
	var rtn Response
	dist := m.Db.GetDistributor(d.ID)
	if dist.StoreID == d.StoreID {
		suc := m.Db.UpdateDistributor(d)
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

//GetDistributor GetDistributor
func (m *Six910Manager) GetDistributor(id int64, storeID int64) *sdbi.Distributor {
	var rtn *sdbi.Distributor
	dst := m.Db.GetDistributor(id)
	if dst.StoreID == storeID {
		rtn = dst
	} else {
		var ns sdbi.Distributor
		rtn = &ns
	}
	return rtn
}

//GetDistributorList GetDistributorList
func (m *Six910Manager) GetDistributorList(storeID int64) *[]sdbi.Distributor {
	return m.Db.GetDistributorList(storeID)
}

//DeleteDistributor DeleteDistributor
func (m *Six910Manager) DeleteDistributor(id int64, storeID int64) *Response {
	var rtn Response
	dst := m.Db.GetDistributor(id)
	if dst.StoreID == storeID {
		suc := m.Db.DeleteDistributor(dst.ID)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusInternalServerError
		}
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}
