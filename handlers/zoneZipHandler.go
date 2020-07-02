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

//ZoneZipReq ZoneZipReq
type ZoneZipReq struct {
	StoreID int64        `json:"storeId"`
	ZoneZip sdbi.ZoneZip `json:"zoneZip"`
}

//AddZoneZip AddZoneZip
func (h *Six910Handler) AddZoneZip(w http.ResponseWriter, r *http.Request) {
	var addzzURL = "/six910/rs/zoneZip/add"
	var azzc jv.Claim
	azzc.Role = storeAdmin
	azzc.URL = addzzURL
	azzc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &azzc)
	h.Log.Debug("zone zip add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var azzr ZoneZipReq
			azzsuc, azzerr := h.ProcessBody(r, &azzr)
			h.Log.Debug("azzsuc: ", azzsuc)
			h.Log.Debug("azzr: ", azzr)
			h.Log.Debug("azzerr: ", azzerr)
			if !azzsuc && azzerr != nil {
				http.Error(w, azzerr.Error(), http.StatusBadRequest)
			} else {
				azzres := h.Manager.AddZoneZip(&azzr.ZoneZip, azzr.StoreID)
				h.Log.Debug("azzres: ", *azzres)
				if azzres.Success && azzres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(azzres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var azzf m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(azzf)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetZoneZipListByExclusion GetZoneZipListByExclusion
func (h *Six910Handler) GetZoneZipListByExclusion(w http.ResponseWriter, r *http.Request) {
	var gzzURL = "/six910/rs/exZoneZip/list"
	var gzzc jv.Claim
	gzzc.Role = customerRole
	gzzc.URL = gzzURL
	gzzc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gzzc)
	h.Log.Debug("excluded zip zone get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gzzexidStr = vars["exId"]
			var gzzstoreIDStr = vars["storeId"]
			id, gzzexiderr := strconv.ParseInt(gzzexidStr, 10, 64)
			storeID, gzzsiderr := strconv.ParseInt(gzzstoreIDStr, 10, 64)
			var gzzlres *[]sdbi.ZoneZip
			if gzzexiderr == nil && gzzsiderr == nil {
				gzzlres = h.Manager.GetZoneZipListByExclusion(id, storeID)
				h.Log.Debug("gzzlres: ", gzzlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc []sdbi.ZoneZip
				gzzlres = &nc
			}
			resJSON, _ := json.Marshal(gzzlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetZoneZipListByInclusion GetZoneZipListByInclusion
func (h *Six910Handler) GetZoneZipListByInclusion(w http.ResponseWriter, r *http.Request) {
	var gizzURL = "/six910/rs/incZoneZip/list"
	var gizzc jv.Claim
	gizzc.Role = customerRole
	gizzc.URL = gizzURL
	gizzc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gizzc)
	h.Log.Debug("included zip zone get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gizzexidStr = vars["incId"]
			var gizzstoreIDStr = vars["storeId"]
			id, gizzexiderr := strconv.ParseInt(gizzexidStr, 10, 64)
			storeID, gizzsiderr := strconv.ParseInt(gizzstoreIDStr, 10, 64)
			var gizzlres *[]sdbi.ZoneZip
			if gizzexiderr == nil && gizzsiderr == nil {
				gizzlres = h.Manager.GetZoneZipListByInclusion(id, storeID)
				h.Log.Debug("gizzlres: ", gizzlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc []sdbi.ZoneZip
				gizzlres = &nc
			}
			resJSON, _ := json.Marshal(gizzlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeleteZoneZip DeleteZoneZip
func (h *Six910Handler) DeleteZoneZip(w http.ResponseWriter, r *http.Request) {
	var dzzURL = "/six910/rs/zoneZip/delete"
	var dzzc jv.Claim
	dzzc.Role = storeAdmin
	dzzc.URL = dzzURL
	dzzc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dzzc)
	h.Log.Debug("zonezip delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 4 {
			h.Log.Debug("vars: ", vars)
			var dzzidStr = vars["id"]
			var dzziidStr = vars["incId"]
			var dzzeidStr = vars["exId"]
			var dzzstoreIDStr = vars["storeId"]
			id, dzziderr := strconv.ParseInt(dzzidStr, 10, 64)
			incID, dzziiderr := strconv.ParseInt(dzziidStr, 10, 64)
			exID, dzzeiderr := strconv.ParseInt(dzzeidStr, 10, 64)
			storeID, dzzidserr := strconv.ParseInt(dzzstoreIDStr, 10, 64)
			var dzzres *m.Response
			if dzziderr == nil && dzziiderr == nil && dzzeiderr == nil && dzzidserr == nil {
				dzzres = h.Manager.DeleteZoneZip(id, incID, exID, storeID)
				h.Log.Debug("dzzres: ", dzzres)
				if dzzres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dzzres = &nc
			}
			resJSON, _ := json.Marshal(dzzres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
