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

//AddCartItem AddCartItem
func (m *Six910Manager) AddCartItem(ci *sdbi.CartItem, cid int64, sid int64) *ResponseID {
	var rtn ResponseID
	var ok bool
	var ctid int64
	if cid != 0 {
		crt := m.Db.GetCart(cid)
		m.Log.Debug("found cart in add: ", crt)
		if crt != nil && crt.StoreID == sid && crt.ID == ci.CartID {
			ctid = crt.ID
			ok = true
		}
	} else {
		ctid = ci.CartID
		ok = true
	}
	if ok {
		crtLst := m.Db.GetCartItemList(ctid)
		var foundItem bool
		for _, fc := range *crtLst {
			if fc.ProductID == ci.ProductID {
				ci.ID = fc.ID
				ci.Quantity += fc.Quantity
				foundItem = true
			}
		}
		if foundItem {
			suc := m.Db.UpdateCartItem(ci)
			rtn.Success = suc
			rtn.ID = ci.ID
			rtn.Code = http.StatusOK
		} else {
			suc, id := m.Db.AddCartItem(ci)
			if suc && id != 0 {
				rtn.Success = suc
				rtn.ID = id
				rtn.Code = http.StatusOK
			}
		}
	}

	if !rtn.Success {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateCartItem UpdateCartItem
func (m *Six910Manager) UpdateCartItem(ci *sdbi.CartItem, cid int64, sid int64) *Response {
	var rtn Response
	var ok bool
	if cid != 0 {
		crt := m.Db.GetCart(cid)
		if crt != nil && crt.StoreID == sid && crt.CustomerID == cid {
			ok = true
		}
	} else {
		ok = true
	}
	if ok {
		suc := m.Db.UpdateCartItem(ci)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusBadRequest
		}
	}
	return &rtn
}

//GetCarItem GetCarItem
func (m *Six910Manager) GetCarItem(cid int64, prodID int64, sid int64) *sdbi.CartItem {
	var rtn *sdbi.CartItem
	crt := m.Db.GetCart(cid)
	if crt != nil && crt.StoreID == sid {
		rtn = m.Db.GetCarItem(crt.ID, prodID)
	} else {
		var nc sdbi.CartItem
		rtn = &nc
	}
	return rtn
}

//GetCartItemList GetCartItemList
func (m *Six910Manager) GetCartItemList(cartID int64, cid int64, sid int64) *[]sdbi.CartItem {
	var rtn *[]sdbi.CartItem
	if cid != 0 {
		crt := m.Db.GetCart(cid)
		m.Log.Debug("cart in get list: ", crt)
		if crt != nil && crt.CustomerID == cid && crt.StoreID == sid && cartID == crt.ID {
			rtn = m.Db.GetCartItemList(cartID)
		} else {
			var nc = []sdbi.CartItem{}
			rtn = &nc
		}
	} else {
		rtn = m.Db.GetCartItemList(cartID)
	}
	return rtn
}

//DeleteCartItem DeleteCartItem
func (m *Six910Manager) DeleteCartItem(id int64, prodID int64, cartID int64) *Response {
	var rtn Response
	ci := m.Db.GetCarItem(cartID, prodID)
	m.Log.Debug("cartItem in get delete: ", *ci)
	m.Log.Debug("cartItem ID in get delete: ", ci.ID)
	if ci.CartID == cartID && ci.ID == id {
		suc := m.Db.DeleteCartItem(id)
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
