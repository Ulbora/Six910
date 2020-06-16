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

//AddStore AddStore
func (m *Six910Manager) AddStore(s *sdbi.Store) *ResponseID {
	var rtn ResponseID
	if m.Db.GetSecurity().OauthOn {
		suc, id := m.Db.AddStore(s)
		if suc && id != 0 {
			rtn.ID = id
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusBadRequest
			rtn.Message = failStoreMayAlreadyExist
		}
	}
	return &rtn
}

//UpdateStore UpdateStore
func (m *Six910Manager) UpdateStore(s *sdbi.Store) *Response {
	var rtn Response
	suc := m.Db.UpdateStore(s)
	if suc {
		rtn.Success = suc
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//GetStore GetStore
func (m *Six910Manager) GetStore(sname string) *sdbi.Store {
	return m.Db.GetStore(sname)
}

//GetStoreID GetStoreID
func (m *Six910Manager) GetStoreID(id int64) *sdbi.Store {
	return m.Db.GetStoreID(id)
}

//GetStoreByLocal GetStoreByLocal
func (m *Six910Manager) GetStoreByLocal(localDomain string) *sdbi.Store {
	return m.Db.GetStoreByLocal(localDomain)
}

//DeleteStore DeleteStore
func (m *Six910Manager) DeleteStore(id int64) *Response {
	var rtn Response
	if m.Db.GetSecurity().OauthOn {
		suc := m.Db.DeleteStore(id)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusInternalServerError
		}
	}
	return &rtn
}
