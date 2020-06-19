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

//AddAddress AddAddress
func (m *Six910Manager) AddAddress(a *sdbi.Address, sid int64) *ResponseID {
	var rtn ResponseID
	cus := m.Db.GetCustomerID(a.CustomerID)
	if cus.StoreID == sid {
		suc, id := m.Db.AddAddress(a)
		if suc && id != 0 {
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

//UpdateAddress UpdateAddress
func (m *Six910Manager) UpdateAddress(a *sdbi.Address, sid int64) *Response {
	var rtn Response
	cus := m.Db.GetCustomerID(a.CustomerID)
	if cus.StoreID == sid {
		suc := m.Db.UpdateAddress(a)
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

//GetAddress GetAddress
func (m *Six910Manager) GetAddress(id int64, cid int64, sid int64) *sdbi.Address {
	var rtn *sdbi.Address
	add := m.Db.GetAddress(id)
	m.Log.Debug("address: ", *add)
	cus := m.Db.GetCustomerID(cid)
	m.Log.Debug("customer: ", *cus)
	if cus.ID == add.CustomerID && cus.StoreID == sid {
		rtn = add
		m.Log.Debug("address in if: ", *rtn)
	} else {
		var na sdbi.Address
		rtn = &na
		m.Log.Debug("address in fail if: ", *rtn)
	}
	return rtn
}

//GetAddressList GetAddressList
func (m *Six910Manager) GetAddressList(cid int64, sid int64) *[]sdbi.Address {
	var rtn *[]sdbi.Address
	cus := m.Db.GetCustomerID(cid)
	if cus.StoreID == sid {
		rtn = m.Db.GetAddressList(cid)
	} else {
		var na = []sdbi.Address{}
		rtn = &na
	}
	return rtn
}

//DeleteAddress DeleteAddress
func (m *Six910Manager) DeleteAddress(id int64, cid int64, sid int64) *Response {
	var rtn Response
	add := m.Db.GetAddress(id)
	cus := m.Db.GetCustomerID(cid)
	if add.CustomerID == cid && cus.StoreID == sid {
		suc := m.Db.DeleteAddress(id)
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
