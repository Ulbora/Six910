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

//AddDataStoreWriteLock AddDataStoreWriteLock
func (h *Six910Handler) AddDataStoreWriteLock(w http.ResponseWriter, r *http.Request) {
	var addrlkURL = "/six910/rs/dataStoreWriteLock/add"
	var arlkc jv.Claim
	arlkc.Role = storeAdmin
	arlkc.URL = addrlkURL
	arlkc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &arlkc)
	h.Log.Debug("write lock add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var arlk sdbi.DataStoreWriteLock
			arlksuc, arlkerr := h.ProcessBody(r, &arlk)
			h.Log.Debug("arlksuc: ", arlksuc)
			h.Log.Debug("arlk: ", arlk)
			h.Log.Debug("arlkerr: ", arlkerr)
			if !arlksuc && arlkerr != nil {
				http.Error(w, arlkerr.Error(), http.StatusBadRequest)
			} else {
				arlkres := h.Manager.AddDataStoreWriteLock(&arlk)
				h.Log.Debug("arlkres: ", *arlkres)
				if arlkres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(arlkres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var arlkfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(arlkfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateDataStoreWriteLock UpdateDataStoreWriteLock
func (h *Six910Handler) UpdateDataStoreWriteLock(w http.ResponseWriter, r *http.Request) {
	var uprlkURL = "/six910/rs/dataStoreWriteLock/update"
	var urlkc jv.Claim
	urlkc.Role = storeAdmin
	urlkc.URL = uprlkURL
	urlkc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &urlkc)
	h.Log.Debug("write lock update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var urlk sdbi.DataStoreWriteLock
			urlksuc, urlkerr := h.ProcessBody(r, &urlk)
			h.Log.Debug("urlksuc: ", urlksuc)
			h.Log.Debug("urlk: ", urlk)
			h.Log.Debug("urlkerr: ", urlkerr)
			if !urlksuc && urlkerr != nil {
				http.Error(w, urlkerr.Error(), http.StatusBadRequest)
			} else {
				urlkres := h.Manager.UpdateDataStoreWriteLock(&urlk)
				h.Log.Debug("urlkres: ", *urlkres)
				if urlkres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(urlkres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var urlkfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(urlkfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetDataStoreWriteLock GetDataStoreWriteLock
func (h *Six910Handler) GetDataStoreWriteLock(w http.ResponseWriter, r *http.Request) {
	var grlkURL = "/six910/rs/dataStoreWriteLock/get"
	var grlkc jv.Claim
	grlkc.Role = customerRole
	grlkc.URL = grlkURL
	grlkc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &grlkc)
	h.Log.Debug("write lock get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var datastore = vars["dataStore"]
			var grlkstoreIDStr = vars["storeId"]
			//id, gspiiderr := strconv.ParseInt(gspiidStr, 10, 64)
			storeID, grlksiderr := strconv.ParseInt(grlkstoreIDStr, 10, 64)
			var grlkres *sdbi.DataStoreWriteLock
			if grlksiderr == nil {
				grlkres = h.Manager.GetDataStoreWriteLock(datastore, storeID)
				h.Log.Debug("grlkres: ", grlkres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.DataStoreWriteLock
				grlkres = &nc
			}
			resJSON, _ := json.Marshal(grlkres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
