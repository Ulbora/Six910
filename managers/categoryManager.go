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

//AddCategory AddCategory
func (m *Six910Manager) AddCategory(c *sdbi.Category) *ResponseID {
	var rtn ResponseID
	suc, id := m.Db.AddCategory(c)
	if suc && id != 0 {
		rtn.Success = suc
		rtn.ID = id
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateCategory UpdateCategory
func (m *Six910Manager) UpdateCategory(c *sdbi.Category) *Response {
	var rtn Response
	cat := m.Db.GetCategory(c.ID)
	if cat.StoreID == c.StoreID {
		suc := m.Db.UpdateCategory(c)
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

//GetCategory GetCategory
func (m *Six910Manager) GetCategory(id int64, sid int64) *sdbi.Category {
	var rtn *sdbi.Category
	cat := m.Db.GetCategory(id)
	if cat.StoreID == sid {
		rtn = cat
	} else {
		var nc sdbi.Category
		rtn = &nc
	}
	return rtn
}

//GetCategoryList GetCategoryList
func (m *Six910Manager) GetCategoryList(storeID int64) *[]sdbi.Category {
	return m.Db.GetCategoryList(storeID)
}

//GetSubCategoryList GetSubCategoryList
func (m *Six910Manager) GetSubCategoryList(catID int64) *[]sdbi.Category {
	return m.Db.GetSubCategoryList(catID)
}

//DeleteCategory DeleteCategory
func (m *Six910Manager) DeleteCategory(id int64, sid int64) *Response {
	var rtn Response
	cat := m.Db.GetCategory(id)
	if cat.StoreID == sid {
		suc := m.Db.DeleteCategory(id)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusOK
		}
	} else {
		rtn.Code = http.StatusOK
	}
	return &rtn
}
