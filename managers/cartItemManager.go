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
	crt := m.Db.GetCart(ci.CartID)
	if crt.StoreID == sid && crt.CustomerID == cid {
		suc, id := m.Db.AddCartItem(ci)
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

//UpdateCartItem UpdateCartItem
func (m *Six910Manager) UpdateCartItem(ci *sdbi.CartItem, cid int64, sid int64) *Response {
	var rtn Response
	crt := m.Db.GetCart(ci.CartID)
	if crt.StoreID == sid && crt.CustomerID == cid {
		suc := m.Db.UpdateCartItem(ci)
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

//GetCarItem GetCarItem
func (m *Six910Manager) GetCarItem(cartID int64, prodID int64, sid int64) *sdbi.CartItem {
	var rtn *sdbi.CartItem
	crt := m.Db.GetCart(cartID)
	if crt.StoreID == sid {
		rtn = m.Db.GetCarItem(cartID, prodID)
	} else {
		var nc sdbi.CartItem
		rtn = &nc
	}
	return rtn
}

//GetCartItemList GetCartItemList
func (m *Six910Manager) GetCartItemList(cartID int64, cid int64, sid int64) *[]sdbi.CartItem {
	var rtn *[]sdbi.CartItem
	crt := m.Db.GetCart(cartID)
	if crt.CustomerID == cid && crt.StoreID == sid {
		rtn = m.Db.GetCartItemList(cartID)
	} else {
		var nc = []sdbi.CartItem{}
		rtn = &nc
	}
	return rtn
}

//DeleteCartItem DeleteCartItem
func (m *Six910Manager) DeleteCartItem(id int64, prodID int64, cartID int64) *Response {
	var rtn Response
	ci := m.Db.GetCarItem(id, prodID)
	if ci.CartID == cartID {
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
