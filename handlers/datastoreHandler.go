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

// AddLocalDatastore godoc
// @Summary Add a new JSON datastore
// @Description Adds a new JSON datastore for store content, css and others used to sync UI
// @Tags Datastore (JSON Datastore used to sync UI when clustered: content or css or others)
// @Accept  json
// @Produce  json
// @Param datastore body six910-database-interface.LocalDataStore true "datastore"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/datastore/add [post]
func (h *Six910Handler) AddLocalDatastore(w http.ResponseWriter, r *http.Request) {
	var addldsURL = "/six910/rs/datastore/add"
	var aldsc jv.Claim
	aldsc.Role = storeAdmin
	aldsc.URL = addldsURL
	aldsc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &aldsc)
	h.Log.Debug("local datastore add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var alds sdbi.LocalDataStore
			aldssuc, aldserr := h.ProcessBody(r, &alds)
			h.Log.Debug("aldssuc: ", aldssuc)
			h.Log.Debug("alds: ", alds)
			h.Log.Debug("aldserr: ", aldserr)
			if !aldssuc && aldserr != nil {
				http.Error(w, aldserr.Error(), http.StatusBadRequest)
			} else {
				aldsres := h.Manager.AddLocalDatastore(&alds)
				h.Log.Debug("aldsres: ", *aldsres)
				if aldsres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aldsres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aldsfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aldsfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateLocalDatastore godoc
// @Summary Update a JSON datastore
// @Description Update a JSON datastore for store content, css and others used to sync UI
// @Tags Datastore (JSON Datastore used to sync UI when clustered: content or css or others)
// @Accept  json
// @Produce  json
// @Param datastore body six910-database-interface.LocalDataStore true "datastore"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/datastore/update [put]
func (h *Six910Handler) UpdateLocalDatastore(w http.ResponseWriter, r *http.Request) {
	var upldsURL = "/six910/rs/datastore/update"
	var uldsc jv.Claim
	uldsc.Role = storeAdmin
	uldsc.URL = upldsURL
	uldsc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uldsc)
	h.Log.Debug("local datastore update authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var ulds sdbi.LocalDataStore
			uldssuc, uldserr := h.ProcessBody(r, &ulds)
			h.Log.Debug("uldssuc: ", uldssuc)
			h.Log.Debug("ulds: ", ulds)
			h.Log.Debug("uldserr: ", uldserr)
			if !uldssuc && uldserr != nil {
				http.Error(w, uldserr.Error(), http.StatusBadRequest)
			} else {
				uldsres := h.Manager.UpdateLocalDatastore(&ulds)
				h.Log.Debug("uldsres: ", *uldsres)
				if uldsres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uldsres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uldsfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uldsfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetLocalDatastore godoc
// @Summary Get details of a JSON datastore by name
// @Description Get details of a JSON datastore for store content, css and others used to sync UI
// @Tags Datastore (JSON Datastore used to sync UI when clustered: content or css or others)
// @Accept  json
// @Produce  json
// @Param name path string true "name"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.LocalDataStore
// @Router /rs/datastore/get/{name}/{storeId} [get]
func (h *Six910Handler) GetLocalDatastore(w http.ResponseWriter, r *http.Request) {
	var gldsURL = "/six910/rs/datastore/get"
	var gldsc jv.Claim
	gldsc.Role = customerRole
	gldsc.URL = gldsURL
	gldsc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gldsc)
	h.Log.Debug("local datastore get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var name = vars["name"]
			var gldsstoreIDStr = vars["storeId"]
			storeID, gldssiderr := strconv.ParseInt(gldsstoreIDStr, 10, 64)
			var gldsres *sdbi.LocalDataStore
			if gldssiderr == nil {
				gldsres = h.Manager.GetLocalDatastore(storeID, name)
				h.Log.Debug("gldsres: ", gldsres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.LocalDataStore
				gldsres = &nc
			}
			resJSON, _ := json.Marshal(gldsres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
