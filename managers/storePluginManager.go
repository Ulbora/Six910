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

//AddStorePlugin AddStorePlugin
func (m *Six910Manager) AddStorePlugin(sp *sdbi.StorePlugins) *ResponseID {
	var rtn ResponseID
	suc, id := m.Db.AddStorePlugin(sp)
	if suc && id != 0 {
		rtn.Success = suc
		rtn.ID = id
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateStorePlugin UpdateStorePlugin
func (m *Six910Manager) UpdateStorePlugin(sp *sdbi.StorePlugins) *Response {
	var rtn Response
	fsp := m.Db.GetStorePlugin(sp.ID)
	if fsp.StoreID == sp.StoreID {
		suc := m.Db.UpdateStorePlugin(sp)
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

//GetStorePlugin GetStorePlugin
func (m *Six910Manager) GetStorePlugin(id int64, sid int64) *sdbi.StorePlugins {
	var rtn *sdbi.StorePlugins
	sp := m.Db.GetStorePlugin(id)
	if sp.StoreID == sid {
		rtn = sp
	} else {
		var np sdbi.StorePlugins
		rtn = &np
	}
	return rtn
}

//GetStorePluginList GetStorePluginList
func (m *Six910Manager) GetStorePluginList(storeID int64) *[]sdbi.StorePlugins {
	return m.Db.GetStorePluginList(storeID)
}

//DeleteStorePlugin DeleteStorePlugin
func (m *Six910Manager) DeleteStorePlugin(id int64, sid int64) *Response {
	var rtn Response
	sp := m.Db.GetStorePlugin(id)
	if sp.StoreID == sid {
		suc := m.Db.DeleteStorePlugin(id)
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
