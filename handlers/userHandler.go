package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	m "github.com/Ulbora/Six910/managers"
	"github.com/gorilla/mux"
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

//GetUser GetUser
func (h *Six910Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	var gUsURL = "/six910/rs/user/get"
	var guc jv.Claim
	guc.Role = customerRole
	guc.URL = gUsURL
	guc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &guc)

	h.Log.Debug("user get authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var username = vars["username"]
			var storeIDStr = vars["storeId"]
			storeID, err := strconv.ParseInt(storeIDStr, 10, 64)
			var gures *m.UserResponse
			if err == nil {
				var gureq m.User
				gureq.Username = username
				gureq.StoreID = storeID
				gures = h.Manager.GetUser(&gureq)
				h.Log.Debug("gures: ", gures)
				w.WriteHeader(http.StatusOK)
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
				var nu m.UserResponse
				gures = &nu
			}
			resJSON, _ := json.Marshal(gures)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetAdminUserList GetAdminUserList
func (h *Six910Handler) GetAdminUserList(w http.ResponseWriter, r *http.Request) {
	var gUslURL = "/six910/rs/adminuser/list"
	var gucl jv.Claim
	gucl.Role = storeAdmin
	gucl.URL = gUslURL
	gucl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gucl)
	h.Log.Debug("admin user get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var storeIDStr = vars["storeId"]
			storeID, serr := strconv.ParseInt(storeIDStr, 10, 64)
			var gulres *[]m.UserResponse
			if serr == nil {
				gulres = h.Manager.GetAdminUsers(storeID)
				h.Log.Debug("gulres list: ", gulres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []m.UserResponse{}
				gulres = &nc
			}
			resJSON, _ := json.Marshal(gulres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetCustomerUserList GetCustomerUserList
func (h *Six910Handler) GetCustomerUserList(w http.ResponseWriter, r *http.Request) {
	var gcUslURL = "/six910/rs/customeruser/list"
	var gcucl jv.Claim
	gcucl.Role = storeAdmin
	gcucl.URL = gcUslURL
	gcucl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gcucl)
	h.Log.Debug("admin user get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var storeIDStr = vars["storeId"]
			storeID, serr := strconv.ParseInt(storeIDStr, 10, 64)
			var gculres *[]m.UserResponse
			if serr == nil {
				gculres = h.Manager.GetCustomerUsers(storeID)
				h.Log.Debug("gculres list: ", gculres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var ncc = []m.UserResponse{}
				gculres = &ncc
			}
			resJSON, _ := json.Marshal(gculres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
