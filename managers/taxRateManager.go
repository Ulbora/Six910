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

//AddTaxRate AddTaxRate
func (m *Six910Manager) AddTaxRate(t *sdbi.TaxRate) *ResponseID {
	var rtn ResponseID
	suc, id := m.Db.AddTaxRate(t)
	if suc && id != 0 {
		rtn.Success = suc
		rtn.ID = id
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateTaxRate UpdateTaxRate
func (m *Six910Manager) UpdateTaxRate(t *sdbi.TaxRate) *Response {
	var rtn Response
	suc := m.Db.UpdateTaxRate(t)
	if suc {
		rtn.Success = suc
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//GetTaxRate GetTaxRate
func (m *Six910Manager) GetTaxRate(country string, state string, sid int64) *[]sdbi.TaxRate {
	return m.Db.GetTaxRate(country, state, sid)
}

//GetTaxRateList GetTaxRateList
func (m *Six910Manager) GetTaxRateList(storeID int64) *[]sdbi.TaxRate {
	return m.Db.GetTaxRateList(storeID)
}

//DeleteTaxRate DeleteTaxRate
func (m *Six910Manager) DeleteTaxRate(id int64, sid int64) *Response {
	var rtn Response
	rtl := m.Db.GetTaxRateList(sid)
	for _, rt := range *rtl {
		if rt.ID == id {
			suc := m.Db.DeleteTaxRate(id)
			if suc {
				rtn.Success = suc
				rtn.Code = http.StatusOK
			} else {
				rtn.Code = http.StatusOK
			}
			break
		}
	}
	return &rtn
}
