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
	suc, id := m.Db.AddCustomer(c)
	if suc && id != 0 {
		rtn.ID = id
		rtn.Success = suc
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateCustomer UpdateCustomer
func (m *Six910Manager) UpdateCustomer(c *sdbi.Customer) *Response {
	var rtn Response
	suc := m.Db.UpdateCustomer(c)
	if suc {
		rtn.Success = suc
		rtn.Code = http.StatusOK
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
func (m *Six910Manager) GetCustomerID(id int64) *sdbi.Customer {
	return m.Db.GetCustomerID(id)
}

//GetCustomerList GetCustomerList
func (m *Six910Manager) GetCustomerList(storeID int64) *[]sdbi.Customer {
	return m.Db.GetCustomerList(storeID)
}

//DeleteCustomer DeleteCustomer
func (m *Six910Manager) DeleteCustomer(id int64) *Response {
	var rtn Response
	suc := m.Db.DeleteCustomer(id)
	if suc {
		rtn.Success = suc
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}
