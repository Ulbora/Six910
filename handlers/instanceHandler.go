package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	m "github.com/Ulbora/Six910/managers"
	sdbi "github.com/Ulbora/six910-database-interface"
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

//AddInstance AddInstance
func (h *Six910Handler) AddInstance(w http.ResponseWriter, r *http.Request) {
	var addinstURL = "/six910/rs/instance/add"
	var ainstc jv.Claim
	ainstc.Role = storeAdmin
	ainstc.URL = addinstURL
	ainstc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ainstc)
	h.Log.Debug("instance add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var ainst sdbi.Instances
			ainstsuc, ainsterr := h.ProcessBody(r, &ainst)
			h.Log.Debug("ainstsuc: ", ainstsuc)
			h.Log.Debug("ainst: ", ainst)
			h.Log.Debug("ainsterr: ", ainsterr)
			if !ainstsuc && ainsterr != nil {
				http.Error(w, ainsterr.Error(), http.StatusBadRequest)
			} else {
				ainstres := h.Manager.AddInstance(&ainst)
				h.Log.Debug("ainstres: ", *ainstres)
				if ainstres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ainstres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var ainstfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ainstfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateInstance UpdateInstance
func (h *Six910Handler) UpdateInstance(w http.ResponseWriter, r *http.Request) {
	var upinstURL = "/six910/rs/instance/update"
	var uinstc jv.Claim
	uinstc.Role = storeAdmin
	uinstc.URL = upinstURL
	uinstc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uinstc)
	h.Log.Debug("instance update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uinst sdbi.Instances
			uinstsuc, uinsterr := h.ProcessBody(r, &uinst)
			h.Log.Debug("uinstsuc: ", uinstsuc)
			h.Log.Debug("uspi: ", uinst)
			h.Log.Debug("uinsterr: ", uinsterr)
			if !uinstsuc && uinsterr != nil {
				http.Error(w, uinsterr.Error(), http.StatusBadRequest)
			} else {
				uinstres := h.Manager.UpdateInstance(&uinst)
				h.Log.Debug("uinstres: ", *uinstres)
				if uinstres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uinstres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uinstfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uinstfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetInstance GetInstance
func (h *Six910Handler) GetInstance(w http.ResponseWriter, r *http.Request) {
	var ginstURL = "/six910/rs/instance/get"
	var ginstc jv.Claim
	ginstc.Role = customerRole
	ginstc.URL = ginstURL
	ginstc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ginstc)
	h.Log.Debug("instance get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var gname = vars["name"]
			var gdname = vars["dataStoreName"]
			var ginststoreIDStr = vars["storeId"]
			//id, gspiiderr := strconv.ParseInt(gspiidStr, 10, 64)
			storeID, ginstsiderr := strconv.ParseInt(ginststoreIDStr, 10, 64)
			var ginstres *sdbi.Instances
			if ginstsiderr == nil {
				ginstres = h.Manager.GetInstance(gname, gdname, storeID)
				h.Log.Debug("ginstres: ", *ginstres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Instances
				ginstres = &nc
			}
			resJSON, _ := json.Marshal(ginstres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
