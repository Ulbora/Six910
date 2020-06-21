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

//AddOrder AddOrder
func (m *Six910Manager) AddOrder(o *sdbi.Order) *ResponseID {
	var rtn ResponseID
	suc, id := m.Db.AddOrder(o)
	if suc && id != 0 {
		rtn.ID = id
		rtn.Success = suc
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateOrder UpdateOrder
func (m *Six910Manager) UpdateOrder(o *sdbi.Order) *Response {
	var rtn Response
	fo := m.Db.GetOrder(o.ID)
	if fo.StoreID == o.StoreID {
		suc := m.Db.UpdateOrder(o)
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

//GetOrder GetOrder
func (m *Six910Manager) GetOrder(id int64, sid int64) *sdbi.Order {
	var rtn *sdbi.Order
	o := m.Db.GetOrder(id)
	if o.StoreID == sid {
		rtn = o
	} else {
		var no sdbi.Order
		rtn = &no
	}
	return rtn
}

//GetOrderList GetOrderList
func (m *Six910Manager) GetOrderList(cid int64, sid int64) *[]sdbi.Order {
	return m.Db.GetOrderList(cid, sid)
}

//DeleteOrder DeleteOrder
func (m *Six910Manager) DeleteOrder(id int64, sid int64) *Response {
	var rtn Response
	fo := m.Db.GetOrder(id)
	if fo.StoreID == sid {
		suc := m.Db.DeleteOrder(id)
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
