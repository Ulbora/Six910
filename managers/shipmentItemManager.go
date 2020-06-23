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

//AddShipmentItem AddShipmentItem
func (m *Six910Manager) AddShipmentItem(si *sdbi.ShipmentItem, sid int64) *ResponseID {
	var rtn ResponseID
	sp := m.Db.GetShipment(si.ShipmentID)
	o := m.Db.GetOrder(sp.OrderID)
	if o.StoreID == sid {
		suc, id := m.Db.AddShipmentItem(si)
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

//UpdateShipmentItem UpdateShipmentItem
func (m *Six910Manager) UpdateShipmentItem(si *sdbi.ShipmentItem, sid int64) *Response {
	var rtn Response
	sp := m.Db.GetShipment(si.ShipmentID)
	o := m.Db.GetOrder(sp.OrderID)
	if o.StoreID == sid {
		suc := m.Db.UpdateShipmentItem(si)
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

//GetShipmentItem GetShipmentItem
func (m *Six910Manager) GetShipmentItem(id int64, sid int64) *sdbi.ShipmentItem {
	var rtn *sdbi.ShipmentItem
	si := m.Db.GetShipmentItem(id)
	sh := m.Db.GetShipment(si.ShipmentID)
	o := m.Db.GetOrder(sh.OrderID)
	if o.StoreID == sid {
		rtn = si
	} else {
		var nb sdbi.ShipmentItem
		rtn = &nb
	}
	return rtn
}

//GetShipmentItemList GetShipmentItemList
func (m *Six910Manager) GetShipmentItemList(shipmentID int64, sid int64) *[]sdbi.ShipmentItem {
	var rtn *[]sdbi.ShipmentItem
	sh := m.Db.GetShipment(shipmentID)
	o := m.Db.GetOrder(sh.OrderID)
	if o.StoreID == sid {
		rtn = m.Db.GetShipmentItemList(shipmentID)
	} else {
		var ns = []sdbi.ShipmentItem{}
		rtn = &ns
	}
	return rtn
}

//GetShipmentItemListByBox GetShipmentItemListByBox
func (m *Six910Manager) GetShipmentItemListByBox(boxNumber int64, sid int64) *[]sdbi.ShipmentItem {
	var rtn *[]sdbi.ShipmentItem
	si := m.Db.GetShipmentItemListByBox(boxNumber)
	var ok bool
	for _, i := range *si {
		sh := m.Db.GetShipment(i.ShipmentID)
		o := m.Db.GetOrder(sh.OrderID)
		if o.StoreID == sid {
			ok = true
		}
		break
	}
	if ok {
		rtn = si
	} else {
		var ni = []sdbi.ShipmentItem{}
		rtn = &ni
	}
	return rtn
}

//DeleteShipmentItem DeleteShipmentItem
func (m *Six910Manager) DeleteShipmentItem(id int64, sid int64) *Response {
	var rtn Response
	si := m.Db.GetShipmentItem(id)
	sh := m.Db.GetShipment(si.ShipmentID)
	o := m.Db.GetOrder(sh.OrderID)
	if o.StoreID == sid {
		suc := m.Db.DeleteShipmentItem(id)
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
