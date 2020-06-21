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

//AddProductCategory AddProductCategory
func (m *Six910Manager) AddProductCategory(pc *sdbi.ProductCategory, sid int64) *Response {
	var rtn Response
	cat := m.Db.GetCategory(pc.CategoryID)
	if cat.StoreID == sid {
		p := m.Db.GetProductByID(pc.ProductID)
		if p.StoreID == sid {
			suc := m.Db.AddProductCategory(pc)
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
		rtn.Code = http.StatusInternalServerError
	}
	return &rtn
}

//DeleteProductCategory DeleteProductCategory
func (m *Six910Manager) DeleteProductCategory(pc *sdbi.ProductCategory, sid int64) *Response {
	var rtn Response
	cat := m.Db.GetCategory(pc.CategoryID)
	if cat.StoreID == sid {
		p := m.Db.GetProductByID(pc.ProductID)
		if p.StoreID == sid {
			suc := m.Db.DeleteProductCategory(pc)
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
		rtn.Code = http.StatusInternalServerError
	}
	return &rtn
}
