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

//AddShippingMethod AddShippingMethod
func (m *Six910Manager) AddShippingMethod(s *sdbi.ShippingMethod) *ResponseID {
	var rtn ResponseID
	suc, id := m.Db.AddShippingMethod(s)
	if suc && id != 0 {
		rtn.Success = suc
		rtn.ID = id
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateShippingMethod UpdateShippingMethod
func (m *Six910Manager) UpdateShippingMethod(s *sdbi.ShippingMethod) *Response {
	var rtn Response
	sm := m.Db.GetShippingMethod(s.ID)
	if sm.StoreID == s.StoreID {
		suc := m.Db.UpdateShippingMethod(s)
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

//GetShippingMethod GetShippingMethod
func (m *Six910Manager) GetShippingMethod(id int64, sid int64) *sdbi.ShippingMethod {
	var rtn *sdbi.ShippingMethod
	sm := m.Db.GetShippingMethod(id)
	if sm.StoreID == sid {
		rtn = sm
	} else {
		var ns sdbi.ShippingMethod
		rtn = &ns
	}
	return rtn
}

//GetShippingMethodList GetShippingMethodList
func (m *Six910Manager) GetShippingMethodList(storeID int64) *[]sdbi.ShippingMethod {
	return m.Db.GetShippingMethodList(storeID)
}

//DeleteShippingMethod DeleteShippingMethod
func (m *Six910Manager) DeleteShippingMethod(id int64, sid int64) *Response {
	var rtn Response
	sm := m.Db.GetShippingMethod(id)
	if sm.StoreID == sid {
		suc := m.Db.DeleteShippingMethod(id)
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
