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

//AddInsurance AddInsurance
func (m *Six910Manager) AddInsurance(i *sdbi.Insurance) *ResponseID {
	var rtn ResponseID
	suc, id := m.Db.AddInsurance(i)
	if suc && id != 0 {
		rtn.Success = suc
		rtn.ID = id
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateInsurance UpdateInsurance
func (m *Six910Manager) UpdateInsurance(i *sdbi.Insurance) *Response {
	var rtn Response
	ins := m.Db.GetInsurance(i.ID)
	if ins.StoreID == i.StoreID {
		suc := m.Db.UpdateInsurance(i)
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

//GetInsurance GetInsurance
func (m *Six910Manager) GetInsurance(id int64, sid int64) *sdbi.Insurance {
	var rtn *sdbi.Insurance
	ins := m.Db.GetInsurance(id)
	if ins.StoreID == sid {
		rtn = ins
	} else {
		var ni sdbi.Insurance
		rtn = &ni
	}
	return rtn
}

//GetInsuranceList GetInsuranceList
func (m *Six910Manager) GetInsuranceList(storeID int64) *[]sdbi.Insurance {
	return m.Db.GetInsuranceList(storeID)
}

//DeleteInsurance DeleteInsurance
func (m *Six910Manager) DeleteInsurance(id int64, sid int64) *Response {
	var rtn Response
	ins := m.Db.GetInsurance(id)
	if ins.StoreID == sid {
		suc := m.Db.DeleteInsurance(id)
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
