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

// AddInstance godoc
// @Summary Add a new datastore UI instance in a cluster
// @Description Adds a new datastore UI instance in a cluster
// @Tags Instance (datastore UI instance in a cluster)
// @Accept  json
// @Produce  json
// @Param instance body six910-database-interface.Instances true "instance"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/instance/add [post]
func (h *Six910Handler) AddInstance(w http.ResponseWriter, r *http.Request) {
	var addinstURL = "/six910/rs/instance/add"
	var ainstc jv.Claim
	ainstc.Role = storeAdmin
	ainstc.URL = addinstURL
	ainstc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ainstc)
	h.Log.Debug("instance add authorized: ", auth)
	h.SetContentType(w)
	if auth {
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
		var ainstfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ainstfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateInstance godoc
// @Summary Update a datastore UI instance in a cluster
// @Description Update datastore UI instance in a cluster
// @Tags Instance (datastore UI instance in a cluster)
// @Accept  json
// @Produce  json
// @Param instance body six910-database-interface.Instances true "instance"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/instance/update [put]
func (h *Six910Handler) UpdateInstance(w http.ResponseWriter, r *http.Request) {
	var upinstURL = "/six910/rs/instance/update"
	var uinstc jv.Claim
	uinstc.Role = storeAdmin
	uinstc.URL = upinstURL
	uinstc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uinstc)
	h.Log.Debug("instance update authorized: ", auth)
	h.SetContentType(w)
	if auth {
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
		var uinstfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uinstfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetInstance godoc
// @Summary Get details of a datastore UI instance in a cluster
// @Description Get details of a datastore UI instance in a cluster
// @Tags Instance (datastore UI instance in a cluster)
// @Accept  json
// @Produce  json
// @Param name path string true "name"
// @Param dataStoreName path string true "dataStoreName"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.Instances
// @Router /rs/instance/get/name/{name}/{dataStoreName}/{storeId} [get]
func (h *Six910Handler) GetInstance(w http.ResponseWriter, r *http.Request) {
	var ginstURL = "/six910/rs/instance/get"
	var ginstc jv.Claim
	ginstc.Role = customerRole
	ginstc.URL = ginstURL
	ginstc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ginstc)
	h.Log.Debug("instance get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var gname = vars["name"]
			var gdname = vars["dataStoreName"]
			var ginststoreIDStr = vars["storeId"]
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

// GetInstanceList godoc
// @Summary Get list of datastore UI instance in a cluster
// @Description Get list of datastore UI instance in a cluster
// @Tags Instance (datastore UI instance in a cluster)
// @Accept  json
// @Produce  json
// @Param dataStoreName path string true "dataStoreName"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Instances
// @Router /rs/instance/get/list/{dataStoreName}/{storeId} [get]
func (h *Six910Handler) GetInstanceList(w http.ResponseWriter, r *http.Request) {
	var ginlURL = "/six910/rs/instance/list"
	var gincl jv.Claim
	gincl.Role = customerRole
	gincl.URL = ginlURL
	gincl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gincl)
	h.Log.Debug("instance get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dslname = vars["dataStoreName"]
			var dslstoreIDStr = vars["storeId"]
			//cID, sorlciderr := strconv.ParseInt(orlcidStr, 10, 64)
			storeID, sorlserr := strconv.ParseInt(dslstoreIDStr, 10, 64)
			var ginlres *[]sdbi.Instances
			if sorlserr == nil {
				ginlres = h.Manager.GetInstanceList(dslname, storeID)
				h.Log.Debug("get instances list: ", ginlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Instances{}
				ginlres = &nc
			}
			resJSON, _ := json.Marshal(ginlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
