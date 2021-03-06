package managers

import (
	"math/rand"
	"strings"

	"net/http"
	"time"

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
	//this is only used when running without oauth2
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
			la.Role = storeAdmin
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
	m.Log.Debug("Add Customer Hash password success for "+u.Password+" :", hpwsuc)
	m.Log.Debug("Add Customer Hashed password for "+u.Password+" :", hpw)
	m.Log.Debug("Add Customer user with customerId :", u.CustomerID)
	if hpwsuc {
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
	if lu.CustomerID == u.CustomerID && lu.StoreID == u.StoreID {
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
		m.Log.Debug("updated success ", suc)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusInternalServerError
		}
	}
	return &rtn
}

//AdminUpdateUser AdminUpdateUser
func (m *Six910Manager) AdminUpdateUser(u *User) *Response {
	var rtn Response
	alu := m.Db.GetLocalAccount(u.Username, u.StoreID)
	if alu.CustomerID == u.CustomerID && alu.StoreID == u.StoreID {
		alu.Enabled = u.Enabled
		if u.Password != "" {
			hpwsuc, hpw := m.hashPassword(u.Password)
			if hpwsuc {
				m.Log.Debug("update password with new Hash password for " + u.Password)
				alu.Password = hpw
			}
		}
		suc := m.Db.UpdateLocalAccount(alu)
		m.Log.Debug("admin updated success ", suc)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusInternalServerError
		}
	}
	return &rtn
}

//GetUser GetUser
func (m *Six910Manager) GetUser(u *User) *UserResponse {
	var rtn UserResponse
	lu := m.Db.GetLocalAccount(u.Username, u.StoreID)
	if lu.UserName == u.Username {
		rtn.CustomerID = lu.CustomerID
		rtn.Role = lu.Role
		rtn.StoreID = lu.StoreID
		rtn.Username = lu.UserName
		rtn.Enabled = lu.Enabled
	}
	m.Log.Debug("UserResponse in GetUser ", rtn)
	return &rtn
}

//GetAdminUsers GetAdminUsers
func (m *Six910Manager) GetAdminUsers(storeID int64) *[]UserResponse {
	//this is only used when running without oauth2
	var rtn []UserResponse
	lulist := m.Db.GetLocalAccountList(storeID)
	for _, l := range *lulist {
		if l.Role == storeAdmin {
			var u UserResponse
			u.CustomerID = l.CustomerID
			u.Role = l.Role
			u.StoreID = l.StoreID
			u.Username = l.UserName
			u.Enabled = l.Enabled
			rtn = append(rtn, u)
		}
	}
	return &rtn
}

//GetCustomerUsers GetCustomerUsers
func (m *Six910Manager) GetCustomerUsers(storeID int64) *[]UserResponse {
	var rtn []UserResponse
	lulist := m.Db.GetLocalAccountList(storeID)
	for _, l := range *lulist {
		if l.Role == customerRole {
			var u UserResponse
			u.CustomerID = l.CustomerID
			u.Role = l.Role
			u.StoreID = l.StoreID
			u.Username = l.UserName
			u.Enabled = l.Enabled
			rtn = append(rtn, u)
		}
	}
	return &rtn
}

//GetUsersByCustomer GetUsersByCustomer
func (m *Six910Manager) GetUsersByCustomer(cid int64, storeID int64) *[]UserResponse {
	var rtn []UserResponse
	clulist := m.Db.GetCustomerUsers(cid, storeID)
	for _, l := range *clulist {
		if l.CustomerID == cid {
			var u UserResponse
			u.CustomerID = l.CustomerID
			u.Role = l.Role
			u.StoreID = l.StoreID
			u.Username = l.UserName
			u.Enabled = l.Enabled
			rtn = append(rtn, u)
		}
	}
	return &rtn
}

//ValidateUser ValidateUser
func (m *Six910Manager) ValidateUser(u *User) *Response {
	var rtn Response
	m.Log.Debug("use in validate: ", *u)
	lu := m.Db.GetLocalAccount(u.Username, u.StoreID)
	m.Log.Debug("local user account in validate: ", *lu)
	m.Log.Debug("lu.CustomerID: ", lu.CustomerID)
	m.Log.Debug("u.CustomerID: ", u.CustomerID)
	m.Log.Debug("lu.UserName: ", lu.UserName)
	m.Log.Debug("u.Username: ", u.Username)
	m.Log.Debug("u.Enabled: ", u.Enabled)
	if lu.CustomerID == u.CustomerID && lu.UserName == u.Username && u.Enabled {
		m.Log.Debug("lu.Password: ", lu.Password)
		m.Log.Debug("u.Password: ", u.Password)
		mtch := m.validatePassword(u.Password, lu.Password)
		m.Log.Debug("password validate in ValidateUser ", mtch)
		if mtch {
			rtn.Success = true
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusUnauthorized
		}
	} else {
		rtn.Code = http.StatusUnauthorized
	}
	return &rtn
}

//ResetCustomerPassword ResetCustomerPassword
func (m *Six910Manager) ResetCustomerPassword(u *User) *CustomerPasswordResponse {
	var rtn CustomerPasswordResponse
	lu := m.Db.GetLocalAccount(u.Username, u.StoreID)
	if lu.UserName == u.Username && lu.Enabled {
		npw := m.generateNewPassword()
		rhpwsuc, hpw := m.hashPassword(npw)
		if rhpwsuc {
			lu.Password = hpw
			suc := m.Db.UpdateLocalAccount(lu)
			m.Log.Debug("generateing now pw and updated success ", suc)
			if suc {
				rtn.Password = npw
				rtn.Success = suc
				rtn.Username = lu.UserName
			}
		}
	}
	return &rtn
}

func (m *Six910Manager) generateNewPassword() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 15
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}
