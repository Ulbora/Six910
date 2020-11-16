package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	m "github.com/Ulbora/Six910/managers"
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

//AddVisit AddVisit
func (h *Six910Handler) AddVisit(w http.ResponseWriter, r *http.Request) {
	var addvURL = "/six910/rs/visit/add"
	var avc jv.Claim
	avc.Role = customerRole
	avc.URL = addvURL
	avc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("visit add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var av sdbi.Visitor
			avsuc, averr := h.ProcessBody(r, &av)
			h.Log.Debug("avsuc: ", avsuc)
			h.Log.Debug("av: ", av)
			h.Log.Debug("averr: ", averr)
			if !avsuc && averr != nil {
				http.Error(w, averr.Error(), http.StatusBadRequest)
			} else {
				avres := h.Manager.AddVisit(&av)
				var resp m.Response
				resp.Success = avres
				h.Log.Debug("avres: ", avres)
				if avres {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(resp)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aorf m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aorf)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetVisitorData GetVisitorData
func (h *Six910Handler) GetVisitorData(w http.ResponseWriter, r *http.Request) {
	var gvdlURL = "/six910/rs/visitor/data/list"
	var gvdcl jv.Claim
	gvdcl.Role = customerRole
	gvdcl.URL = gvdlURL
	gvdcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gvdcl)
	h.Log.Debug("visitor get  list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var vdlstoreIDStr = vars["storeId"]
			storeID, vdlerr := strconv.ParseInt(vdlstoreIDStr, 10, 64)
			var vdList *[]sdbi.VisitorData
			if vdlerr == nil {
				vdList = h.Manager.GetVisitorData(storeID)
				h.Log.Debug("get visitor data list: ", vdList)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.VisitorData{}
				vdList = &nc
			}
			resJSON, _ := json.Marshal(vdList)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
