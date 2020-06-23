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

//AddShippingCarrier AddShippingCarrier
func (m *Six910Manager) AddShippingCarrier(c *sdbi.ShippingCarrier) *ResponseID {
	var rtn ResponseID
	suc, id := m.Db.AddShippingCarrier(c)
	if suc && id != 0 {
		rtn.Success = suc
		rtn.ID = id
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateShippingCarrier UpdateShippingCarrier
func (m *Six910Manager) UpdateShippingCarrier(c *sdbi.ShippingCarrier) *Response {
	var rtn Response
	sc := m.Db.GetShippingCarrier(c.ID)
	if sc.StoreID == c.StoreID {
		suc := m.Db.UpdateShippingCarrier(c)
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

//GetShippingCarrier GetShippingCarrier
func (m *Six910Manager) GetShippingCarrier(id int64, sid int64) *sdbi.ShippingCarrier {
	var rtn *sdbi.ShippingCarrier
	sc := m.Db.GetShippingCarrier(id)
	if sc.StoreID == sid {
		rtn = sc
	} else {
		var nc sdbi.ShippingCarrier
		rtn = &nc
	}
	return rtn
}

//GetShippingCarrierList GetShippingCarrierList
func (m *Six910Manager) GetShippingCarrierList(storeID int64) *[]sdbi.ShippingCarrier {
	return m.Db.GetShippingCarrierList(storeID)
}

//DeleteShippingCarrier DeleteShippingCarrier
func (m *Six910Manager) DeleteShippingCarrier(id int64, sid int64) *Response {
	var rtn Response
	sc := m.Db.GetShippingCarrier(id)
	if sc.StoreID == sid {
		suc := m.Db.DeleteShippingCarrier(id)
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
