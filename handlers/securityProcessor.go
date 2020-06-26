package handlers

import (
	"net/http"
	"strings"

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

//Claim Claim
type Claim struct {
	Role  string
	URL   string
	Scope string
}

func (h *Six910Handler) processSecurity(r *http.Request, c *Claim) bool {
	var rtn bool
	storeName := r.Header.Get("storeName")
	localDomain := r.Header.Get("localDomain")
	sp := h.Manager.GetSecurityProfile(storeName, localDomain)
	// if local with no OAuth process against local store
	//---- this uses user creds as basic auth
	//---- also uses access token that is in memory in both the Six910 UI and Server
	if !sp.IsOAuthOn && sp.Store != nil {
		username, pw, ok := r.BasicAuth()
		if ok {
			//////tokenHeader := r.Header.Get("Authorization")
			var user m.User
			user.Username = username
			user.Password = pw
			user.StoreID = sp.Store.ID
			u := h.Manager.GetUser(&user)
			user.CustomerID = u.CustomerID
			res := h.Manager.ValidateUser(&user)
			apiKey := r.Header.Get("apiKey")
			if res.Success && u.Role == c.Role && apiKey == h.APIKey {
				rtn = true
			}
		}
	}

	//else if oauth the proxy to GoAuth2 and validate token

	return rtn
}

func (h *Six910Handler) getLocalToken(token string) {
	tokenArray := strings.Split(token, " ")
	//fmt.Println("tokenArray", tokenArray)
	if len(tokenArray) == 2 {

	}
}
