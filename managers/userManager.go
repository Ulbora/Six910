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

//AddAdminUser AddAdminUser
func (m *Six910Manager) AddAdminUser(u *User) *Response {
	var rtn Response
	var oauth = m.Db.GetSecurity().OauthOn
	m.Log.Debug("Creating admin user with oauth :", oauth)
	if !oauth {
		hpwsuc, hpw := m.hashPassword(u.Password)
		m.Log.Debug("Add Admin Hash password for "+u.Password+" :", hpwsuc)
		if hpwsuc {
			var la sdbi.LocalAccount
			la.Enabled = true
			la.UserName = u.Username
			la.Password = hpw
			la.StoreID = u.StoreID
			la.Role = u.Role
			suc := m.Db.AddLocalAccount(&la)
			if suc {
				rtn.Success = suc
				rtn.Code = http.StatusOK
			} else {
				rtn.Code = http.StatusBadRequest
			}
		}
	}
	return &rtn
}

//AddCustomerUser AddCustomerUser
func (m *Six910Manager) AddCustomerUser(u *User) *Response {
	//this is secured in handler by api key entered at startup in env var
	//also api key gets passed in header
	var rtn Response
	hpwsuc, hpw := m.hashPassword(u.Password)
	m.Log.Debug("Add Customer Hash password for "+u.Password+" :", hpwsuc)
	m.Log.Debug("Add Customer user with customerId :", u.CustomerID)
	if hpwsuc && u.CustomerID != 0 {
		var la sdbi.LocalAccount
		la.Enabled = true
		la.UserName = u.Username
		la.Password = hpw
		la.StoreID = u.StoreID
		la.CustomerID = u.CustomerID
		la.Role = customerRole
		suc := m.Db.AddLocalAccount(&la)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusBadRequest
		}
	}
	return &rtn
}

//UpdateUser UpdateUser
func (m *Six910Manager) UpdateUser(u *User) *Response {
	var rtn Response
	lu := m.Db.GetLocalAccount(u.Username, u.StoreID)
	if lu.CustomerID == u.CustomerID {
		lu.Enabled = u.Enabled
		if u.Password != "" && u.OldPassword != "" {
			m.Log.Debug("update user Hash password for " + lu.Password)
			mtch := m.validatePassword(u.OldPassword, lu.Password)
			m.Log.Debug("password validate ", mtch)
			if mtch {
				hpwsuc, hpw := m.hashPassword(u.Password)
				if hpwsuc {
					m.Log.Debug("update password with new Hash password for " + u.Password)
					lu.Password = hpw
				}
			}
		}
		suc := m.Db.UpdateLocalAccount(lu)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusInternalServerError
		}
	}

	return &rtn
}
