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

//AddOrderItem AddOrderItem
func (m *Six910Manager) AddOrderItem(i *sdbi.OrderItem, sid int64) *ResponseID {
	var rtn ResponseID
	fo := m.Db.GetOrder(i.OrderID)
	if fo.StoreID == sid {
		suc, id := m.Db.AddOrderItem(i)
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

//UpdateOrderItem UpdateOrderItem
func (m *Six910Manager) UpdateOrderItem(i *sdbi.OrderItem, sid int64) *Response {
	var rtn Response
	fo := m.Db.GetOrder(i.OrderID)
	if fo.StoreID == sid {
		suc := m.Db.UpdateOrderItem(i)
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

//GetOrderItem GetOrderItem
func (m *Six910Manager) GetOrderItem(id int64, sid int64) *sdbi.OrderItem {
	var rtn *sdbi.OrderItem
	oi := m.Db.GetOrderItem(id)
	o := m.Db.GetOrder(oi.OrderID)
	if o.StoreID == sid {
		rtn = oi
	} else {
		var no sdbi.OrderItem
		rtn = &no
	}
	return rtn
}

//GetOrderItemList GetOrderItemList
func (m *Six910Manager) GetOrderItemList(orderID int64, sid int64) *[]sdbi.OrderItem {
	var rtn *[]sdbi.OrderItem
	o := m.Db.GetOrder(orderID)
	if o.StoreID == sid {
		rtn = m.Db.GetOrderItemList(orderID)
	} else {
		var nl = []sdbi.OrderItem{}
		rtn = &nl
	}
	return rtn
}

//DeleteOrderItem DeleteOrderItem
func (m *Six910Manager) DeleteOrderItem(id int64, sid int64) *Response {
	var rtn Response
	fi := m.Db.GetOrderItem(id)
	o := m.Db.GetOrder(fi.OrderID)
	if o.StoreID == sid {
		suc := m.Db.DeleteOrderItem(id)
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
