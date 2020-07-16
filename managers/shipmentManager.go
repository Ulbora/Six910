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

//AddShipment AddShipment
func (m *Six910Manager) AddShipment(s *sdbi.Shipment, sid int64) *ResponseID {
	var rtn ResponseID
	fo := m.Db.GetOrder(s.OrderID)
	m.Log.Debug("Order in add shipment: ", *fo)
	m.Log.Debug("OrderId in add shipment: ", s.OrderID)
	if fo.StoreID == sid {
		suc, id := m.Db.AddShipment(s)
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

//UpdateShipment UpdateShipment
func (m *Six910Manager) UpdateShipment(s *sdbi.Shipment, sid int64) *Response {
	var rtn Response
	fo := m.Db.GetOrder(s.OrderID)
	if fo.StoreID == sid {
		sp := m.Db.GetShipment(s.ID)
		if sp.OrderID == fo.ID {
			suc := m.Db.UpdateShipment(s)
			if suc {
				rtn.Success = suc
				rtn.Code = http.StatusOK
			} else {
				rtn.Code = http.StatusBadRequest
			}
		} else {
			rtn.Code = http.StatusInternalServerError
		}
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//GetShipment GetShipment
func (m *Six910Manager) GetShipment(id int64, sid int64) *sdbi.Shipment {
	var rtn *sdbi.Shipment
	sp := m.Db.GetShipment(id)
	o := m.Db.GetOrder(sp.OrderID)
	if o.StoreID == sid {
		rtn = sp
	} else {
		var ns sdbi.Shipment
		rtn = &ns
	}
	return rtn
}

//GetShipmentList GetShipmentList
func (m *Six910Manager) GetShipmentList(orderID int64, sid int64) *[]sdbi.Shipment {
	var rtn *[]sdbi.Shipment
	o := m.Db.GetOrder(orderID)
	if o.StoreID == sid {
		rtn = m.Db.GetShipmentList(orderID)
	} else {
		var ns = []sdbi.Shipment{}
		rtn = &ns
	}
	return rtn
}

//DeleteShipment DeleteShipment
func (m *Six910Manager) DeleteShipment(id int64, sid int64) *Response {
	var rtn Response
	sp := m.Db.GetShipment(id)
	o := m.Db.GetOrder(sp.OrderID)
	if o.StoreID == sid {
		suc := m.Db.DeleteShipment(id)
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
