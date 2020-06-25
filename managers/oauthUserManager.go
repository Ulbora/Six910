package managers

import (
	"bytes"
	"encoding/json"
	"net/http"
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

//AddOAuthUser AddOAuthUser
func (m *Six910Manager) AddOAuthUser(user *OAuthUser, auth *Auth) *Response {
	var rtn Response
	var addURL = m.UserHost + "/rs/user/add"
	aJSON, err := json.Marshal(user)
	m.Log.Debug("Add user: ", err)
	if err == nil {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		m.Log.Debug("Add user req: ", rErr)
		if rErr == nil {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+auth.Token)
			req.Header.Set("clientId", auth.ClientID)
			//req.Header.Set("apiKey", u.APIKey)
			_, code := m.Proxy.Do(req, &rtn)
			rtn.Code = int64(code)
		}
	}
	return &rtn
}

//UpdateOAuthUser UpdateOAuthUser
func (m *Six910Manager) UpdateOAuthUser(user *OAuthUser, auth *Auth) *Response {
	var rtn Response
	var upURL = m.UserHost + "/rs/user/update"
	aJSON, err := json.Marshal(user)
	m.Log.Debug("update user: ", err)
	if err == nil {
		req, rErr := http.NewRequest("PUT", upURL, bytes.NewBuffer(aJSON))
		m.Log.Debug("update user req: ", rErr)
		if rErr == nil {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+auth.Token)
			req.Header.Set("clientId", auth.ClientID)
			//req.Header.Set("apiKey", u.APIKey)
			_, code := m.Proxy.Do(req, &rtn)
			rtn.Code = int64(code)
		}
	}
	return &rtn
}

//GetOAuthUser GetOAuthUser
func (m *Six910Manager) GetOAuthUser(username string, clientID string, auth *Auth) (*OAuthUserUser, int) {
	var rtn = new(OAuthUserUser)
	var code int
	var gURL = m.UserHost + "/rs/user/get/" + username + "/" + clientID
	req, rErr := http.NewRequest("GET", gURL, nil)
	m.Log.Debug("get user req: ", rErr)
	if rErr == nil {
		req.Header.Set("clientId", auth.ClientID)
		req.Header.Set("Authorization", "Bearer "+auth.Token)
		//req.Header.Set("apiKey", u.APIKey)
		_, code = m.Proxy.Do(req, &rtn)
	}
	return rtn, code
}

//GetOAuthUserList GetOAuthUserList
func (m *Six910Manager) GetOAuthUserList(clientID string, auth *Auth) (*[]OAuthUser, int) {
	var rtn = make([]OAuthUser, 0)
	var code int
	var gURL = m.UserHost + "/rs/user/search/" + clientID
	//fmt.Println(gURL)
	m.Log.Debug("gURL: ", gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	m.Log.Debug("search user list rErr: ", rErr)
	if rErr == nil {
		req.Header.Set("clientId", auth.ClientID)
		req.Header.Set("Authorization", "Bearer "+auth.Token)
		//req.Header.Set("apiKey", u.APIKey)
		_, code = m.Proxy.Do(req, &rtn)
		m.Log.Debug("search user list code: ", code)
	}
	return &rtn, code
}

//DeleteOAuthUser DeleteOAuthUser
func (m *Six910Manager) DeleteOAuthUser(username string, clientID string, auth *Auth) *Response {
	var rtn = new(Response)
	var gURL = m.UserHost + "/rs/user/delete/" + username + "/" + clientID
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	m.Log.Debug("delete user list req: ", rErr)
	if rErr == nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+auth.Token)
		req.Header.Set("clientId", auth.ClientID)
		//req.Header.Set("apiKey", u.APIKey)
		_, code := m.Proxy.Do(req, &rtn)
		rtn.Code = int64(code)
	}
	return rtn
}
