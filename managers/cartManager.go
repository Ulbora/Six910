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
	cus := m.Db.GetCustomerID(c.CustomerID)
	if cus.StoreID == c.StoreID {
		crt := m.Db.GetCart(cus.ID)
		if crt.ID != 0 {
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
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}
