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

//AddDataStoreWriteLock AddDataStoreWriteLock
func (m *Six910Manager) AddDataStoreWriteLock(w *sdbi.DataStoreWriteLock) *Response {
	var rtn Response
	suc := m.Db.AddDataStoreWriteLock(w)
	if suc {
		rtn.Success = suc
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateDataStoreWriteLock UpdateDataStoreWriteLock
func (m *Six910Manager) UpdateDataStoreWriteLock(w *sdbi.DataStoreWriteLock) *Response {
	var rtn Response
	suc := m.Db.UpdateDataStoreWriteLock(w)
	if suc {
		rtn.Success = suc
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//GetDataStoreWriteLock GetDataStoreWriteLock
func (m *Six910Manager) GetDataStoreWriteLock(dataStore string, storeID int64) *sdbi.DataStoreWriteLock {
	return m.Db.GetDataStoreWriteLock(dataStore, storeID)
}
