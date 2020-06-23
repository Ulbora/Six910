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

//AddPlugin AddPlugin
func (m *Six910Manager) AddPlugin(p *sdbi.Plugins) *ResponseID {
	var rtn ResponseID
	suc, id := m.Db.AddPlugin(p)
	if suc && id != 0 {
		rtn.Success = suc
		rtn.ID = id
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdatePlugin UpdatePlugin
func (m *Six910Manager) UpdatePlugin(p *sdbi.Plugins) *Response {
	var rtn Response
	suc := m.Db.UpdatePlugin(p)
	if suc {
		rtn.Success = suc
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//GetPlugin GetPlugin
func (m *Six910Manager) GetPlugin(id int64) *sdbi.Plugins {
	return m.Db.GetPlugin(id)
}

//GetPluginList GetPluginList
func (m *Six910Manager) GetPluginList(start int64, end int64) *[]sdbi.Plugins {
	return m.Db.GetPluginList(start, end)
}

//DeletePlugin DeletePlugin
func (m *Six910Manager) DeletePlugin(id int64) *Response {
	var rtn Response
	suc := m.Db.DeletePlugin(id)
	if suc {
		rtn.Success = suc
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}
