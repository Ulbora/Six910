package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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

//AddUser AddUser
func (h *Six910Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	var addUsURL = "/six910/rs/user/add"
	var auc jv.Claim
	auc.Role = storeAdmin
	auc.URL = addUsURL
	auc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &auc)
	h.Log.Debug("user add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		auOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", auOk)
		if !auOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var user m.User
			ausuc, auerr := h.ProcessBody(r, &user)
			h.Log.Debug("ausuc: ", ausuc)
			h.Log.Debug("user: ", user)
			h.Log.Debug("auerr: ", auerr)
			if !ausuc && auerr != nil {
				http.Error(w, auerr.Error(), http.StatusBadRequest)
			} else {
				var aures *m.Response
				if user.Role == storeAdmin && user.CustomerID == 0 {
					aures = h.Manager.AddAdminUser(&user)
					h.Log.Debug("aures adminuser: ", *aures)
				} else if user.Role == customerRole && user.CustomerID != 0 {
					aures = h.Manager.AddCustomerUser(&user)
					h.Log.Debug("aures customeruser: ", *aures)
				} else {
					var nr m.Response
					aures = &nr
				}
				h.Log.Debug("aures: ", *aures)
				if aures.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aures)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var acfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(acfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateUser UpdateUser
func (h *Six910Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var upUsURL = "/six910/rs/user/update"
	var uuc jv.Claim
	uuc.Role = customerRole
	uuc.URL = upUsURL
	uuc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uuc)
	h.Log.Debug("user update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uus m.User
			ussuc, userr := h.ProcessBody(r, &uus)
			h.Log.Debug("ussuc: ", ussuc)
			h.Log.Debug("uus: ", uus)
			h.Log.Debug("userr: ", userr)
			if !ussuc && userr != nil {
				http.Error(w, userr.Error(), http.StatusBadRequest)
			} else {
				uures := h.Manager.UpdateUser(&uus)
				h.Log.Debug("uures: ", *uures)
				if uures.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uures)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uufl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uufl)
		fmt.Fprint(w, string(resJSON))
	}
}
