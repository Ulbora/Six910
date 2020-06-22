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

//AddShipmentBox AddShipmentBox
func (m *Six910Manager) AddShipmentBox(sb *sdbi.ShipmentBox, sid int64) *ResponseID {
	var rtn ResponseID
	sp := m.Db.GetShipment(sb.ShipmentID)
	o := m.Db.GetOrder(sp.OrderID)
	if o.StoreID == sid {
		suc, id := m.Db.AddShipmentBox(sb)
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

//UpdateShipmentBox UpdateShipmentBox
func (m *Six910Manager) UpdateShipmentBox(sb *sdbi.ShipmentBox, sid int64) *Response {
	var rtn Response
	sp := m.Db.GetShipment(sb.ShipmentID)
	o := m.Db.GetOrder(sp.OrderID)
	if o.StoreID == sid {
		suc := m.Db.UpdateShipmentBox(sb)
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

//GetShipmentBox GetShipmentBox
func (m *Six910Manager) GetShipmentBox(id int64, sid int64) *sdbi.ShipmentBox {
	var rtn *sdbi.ShipmentBox
	sb := m.Db.GetShipmentBox(id)
	sh := m.Db.GetShipment(sb.ShipmentID)
	o := m.Db.GetOrder(sh.OrderID)
	if o.StoreID == sid {
		rtn = sb
	} else {
		var nb sdbi.ShipmentBox
		rtn = &nb
	}
	return rtn
}

//GetShipmentBoxList GetShipmentBoxList
func (m *Six910Manager) GetShipmentBoxList(shipmentID int64, sid int64) *[]sdbi.ShipmentBox {
	var rtn *[]sdbi.ShipmentBox
	sh := m.Db.GetShipment(shipmentID)
	o := m.Db.GetOrder(sh.OrderID)
	if o.StoreID == sid {
		rtn = m.Db.GetShipmentBoxList(shipmentID)
	} else {
		var ns = []sdbi.ShipmentBox{}
		rtn = &ns
	}
	return rtn
}

//DeleteShipmentBox DeleteShipmentBox
func (m *Six910Manager) DeleteShipmentBox(id int64, sid int64) *Response {
	var rtn Response
	sb := m.Db.GetShipmentBox(id)
	sh := m.Db.GetShipment(sb.ShipmentID)
	o := m.Db.GetOrder(sh.OrderID)
	if o.StoreID == sid {
		suc := m.Db.DeleteShipmentBox(id)
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
