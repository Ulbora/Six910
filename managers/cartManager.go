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

//AddCart AddCart
func (m *Six910Manager) AddCart(c *sdbi.Cart) *ResponseID {
	var rtn ResponseID
	m.Log.Debug("adding cart :", *c)
	var cus *sdbi.Customer
	var crt *sdbi.Cart
	if c.CustomerID != 0 {
		cus = m.Db.GetCustomerID(c.CustomerID)
		if cus.StoreID == c.StoreID {
			crt = m.Db.GetCart(cus.ID)
		}
	}
	m.Log.Debug("existing cart :", crt)
	if crt != nil && crt.ID != 0 {
		rtn.ID = crt.ID
		rtn.Success = true
		rtn.Code = http.StatusOK
	} else {
		suc, id := m.Db.AddCart(c)
		if suc && id != 0 {
			rtn.ID = id
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusBadRequest
		}
	}
	return &rtn
}

//UpdateCart UpdateCart
func (m *Six910Manager) UpdateCart(c *sdbi.Cart) *Response {
	var rtn Response
	var ok bool
	cart := m.Db.GetCart(c.CustomerID)
	m.Log.Debug("existing cart :", cart)
	if cart != nil && cart.StoreID == c.StoreID {
		ok = true
		m.Log.Debug("updating cart :", *cart)
	} else if cart == nil && c.ID != 0 {
		ok = true
	}
	if ok {
		m.Log.Debug("updating new customer cart cart :", *c)
		suc := m.Db.UpdateCart(c)
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

//GetCart GetCart
func (m *Six910Manager) GetCart(cid int64, storeID int64) *sdbi.Cart {
	var rtn *sdbi.Cart
	cart := m.Db.GetCart(cid)
	m.Log.Debug("cart in get :", cart)
	if cart != nil && cart.StoreID == storeID {
		rtn = cart
	} else {
		var nc sdbi.Cart
		rtn = &nc
	}
	return rtn
}

//DeleteCart DeleteCart
func (m *Six910Manager) DeleteCart(id int64, cid int64, storeID int64) *Response {
	var rtn Response
	cart := m.Db.GetCart(cid)
	if cart != nil && cart.ID == id && cart.CustomerID == cid && cart.StoreID == storeID {
		suc := m.Db.DeleteCart(id)
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
