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

//AddCustomer AddCustomer
func (m *Six910Manager) AddCustomer(c *sdbi.Customer) *ResponseID {
	var rtn ResponseID
	cs := m.Db.GetCustomer(c.Email, c.StoreID)
	if cs.ID == 0 {
		suc, id := m.Db.AddCustomer(c)
		if suc && id != 0 {
			rtn.ID = id
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusBadRequest
		}
	} else {
		rtn.Code = http.StatusConflict
		rtn.Message = customerAlreadyExists
	}
	return &rtn
}

//UpdateCustomer UpdateCustomer
func (m *Six910Manager) UpdateCustomer(c *sdbi.Customer) *Response {
	var rtn Response
	fc := m.Db.GetCustomerID(c.ID)
	if fc.StoreID == c.StoreID {
		suc := m.Db.UpdateCustomer(c)
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

//GetCustomer GetCustomer
func (m *Six910Manager) GetCustomer(email string, storeID int64) *sdbi.Customer {
	return m.Db.GetCustomer(email, storeID)
}

//GetCustomerID GetCustomerID
func (m *Six910Manager) GetCustomerID(id int64, storeID int64) *sdbi.Customer {
	var rtn *sdbi.Customer
	cus := m.Db.GetCustomerID(id)
	if cus.StoreID == storeID {
		rtn = cus
	} else {
		var nc sdbi.Customer
		rtn = &nc
	}
	return rtn
}

//GetCustomerList GetCustomerList
func (m *Six910Manager) GetCustomerList(storeID int64, start int64, end int64) *[]sdbi.Customer {
	return m.Db.GetCustomerList(storeID, start, end)
}

//DeleteCustomer DeleteCustomer
func (m *Six910Manager) DeleteCustomer(id int64, storeID int64) *Response {
	var rtn Response
	fc := m.Db.GetCustomerID(id)
	if fc.StoreID == storeID {
		suc := m.Db.DeleteCustomer(id)
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
