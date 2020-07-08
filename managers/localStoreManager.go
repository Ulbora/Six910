package managers

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

import (
	"net/http"

	sdbi "github.com/Ulbora/six910-database-interface"
)

//CreateLocalStore CreateLocalStore
func (m *Six910Manager) CreateLocalStore(auth *LocalStoreAdminUser) *LocalStoreResponse {
	var rtn LocalStoreResponse
	sec := m.Db.GetSecurity()
	if !sec.OauthOn {
		strCnt := m.Db.GetStoreCount()
		if strCnt == 0 {
			var str sdbi.Store
			str.City = "Atlanta"
			str.State = "GA"
			str.Zip = "12345"
			str.FirstName = "default"
			str.LastName = "store"
			str.StoreName = "defaultLocalStore"
			str.LocalDomain = "defaultLocalStore.mydomain.com"
			str.Enabled = true
			m.Log.Debug("Creating local store")
			suc, sid := m.Db.AddStore(&str)
			if suc {
				var ssuc bool
				var ssid int64
				if sec.ID == 0 {
					m.Log.Debug("Adding new security record for local store")
					ssuc, ssid = m.Db.AddSecurity(sec)
				} else {
					ssuc = true
					ssid = sec.ID
				}
				if suc && ssuc && ssid != 0 {
					hpwsuc, hpw := m.hashPassword(auth.Password)
					var accsuc bool
					if hpwsuc {
						var lacc sdbi.LocalAccount
						lacc.UserName = auth.Username
						lacc.Password = hpw
						lacc.StoreID = sid
						lacc.Role = storeAdmin
						lacc.Enabled = true
						accsuc = m.Db.AddLocalAccount(&lacc)
					}
					if accsuc {
						rtn.Success = suc
						rtn.StoreID = sid
						rtn.Code = http.StatusOK
					} else {
						rtn.Code = http.StatusInternalServerError
						rtn.Message = failToAddDefaultUserAccount
					}
				} else {
					rtn.Code = http.StatusInternalServerError
					rtn.Message = failToAddSecurity
				}
				//}
			} else {
				rtn.Code = http.StatusInternalServerError
				rtn.Message = failToAddDefaultStore
			}
		} else {
			rtn.Code = http.StatusOK
			rtn.Message = storeAlreadyExists
		}
	}
	return &rtn
}
