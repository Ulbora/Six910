package handlers

import (
	"net/http"
	"strconv"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	m "github.com/Ulbora/Six910/managers"
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

func (h *Six910Handler) processSecurity(r *http.Request, c *jv.Claim) bool {
	var rtn bool
	storeName := r.Header.Get("storeName")
	localDomain := r.Header.Get("localDomain")

	sp := h.Manager.GetSecurityProfile(storeName, localDomain)
	h.Log.Debug("security profile ", *sp)
	// if local with no OAuth process against local store
	//---- this uses user creds as basic auth
	//---- also uses api key that is in memory in both the Six910 UI and Server
	if !sp.IsOAuthOn && sp.Store != nil {
		username, pw, ok := r.BasicAuth()
		h.Log.Debug("basic auth ok ", ok)
		if ok {
			//////tokenHeader := r.Header.Get("Authorization")
			var user m.User
			user.Username = username
			user.Password = pw
			user.StoreID = sp.Store.ID
			u := h.Manager.GetUser(&user)
			user.CustomerID = u.CustomerID
			user.Enabled = u.Enabled
			h.Log.Debug("user to validate", user)
			res := h.Manager.ValidateUser(&user)
			h.Log.Debug("user validated: ", *res)
			apiKey := r.Header.Get("apiKey")
			h.Log.Debug("apiKey: ", apiKey)
			h.Log.Debug("h.APIKey: ", h.APIKey)
			h.Log.Debug("u.Role: ", u.Role)
			if res.Success && apiKey == h.APIKey && (u.Role == superAdmin || u.Role == storeAdmin || u.Role == c.Role) {
				rtn = true
			}
		}
		//else if oauth the proxy to GoAuth2 and validate token
	} else if sp.Store != nil {
		//needs ------
		//tokenHeader := r.Header.Get("Authorization")
		//clientIDStr := r.Header.Get("clientId")
		//userID := r.Header.Get("userId")

		role := r.Header.Get("superAdminRole")
		if role != superAdmin {
			clientIDStr := strconv.FormatInt(sp.Store.OauthClientID, 10)
			var newRoll = sp.Store.StoreName + clientIDStr + c.Role
			c.Role = newRoll
		} else {
			c.Role = role
		}
		h.Log.Debug("claim: ", *c)
		rtn = h.ValidatorClient.Authorize(r, c, h.ValidationURL)
	}
	return rtn
}

func (h *Six910Handler) processAPIKeySecurity(r *http.Request) bool {
	var rtn bool
	apiKey := r.Header.Get("apiKey")
	h.Log.Debug("apiKey: ", apiKey)
	h.Log.Debug("h.APIKey: ", h.APIKey)
	if apiKey == h.APIKey {
		rtn = true
	}
	return rtn
}
