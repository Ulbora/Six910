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

// AddRegion godoc
// @Summary Add a new region
// @Description Adds a new region to a store
// @Tags Region (Geographic Sales Region)
// @Accept  json
// @Produce  json
// @Param region body six910-database-interface.Region true "region"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/region/add [post]
func (h *Six910Handler) AddRegion(w http.ResponseWriter, r *http.Request) {
	var addregURL = "/six910/rs/region/add"
	var aregc jv.Claim
	aregc.Role = storeAdmin
	aregc.URL = addregURL
	aregc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &aregc)
	h.Log.Debug("region add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var areg sdbi.Region
			aregsuc, aregerr := h.ProcessBody(r, &areg)
			h.Log.Debug("aregsuc: ", aregsuc)
			h.Log.Debug("areg: ", areg)
			h.Log.Debug("aregerr: ", aregerr)
			if !aregsuc && aregerr != nil {
				http.Error(w, aregerr.Error(), http.StatusBadRequest)
			} else {
				aregres := h.Manager.AddRegion(&areg)
				h.Log.Debug("aregres: ", *aregres)
				if aregres.Success && aregres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aregres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aregfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aregfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateRegion godoc
// @Summary Update a region
// @Description Update region data
// @Tags Region (Geographic Sales Region)
// @Accept  json
// @Produce  json
// @Param region body six910-database-interface.Region true "region"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/region/update [put]
func (h *Six910Handler) UpdateRegion(w http.ResponseWriter, r *http.Request) {
	var upregURL = "/six910/rs/region/update"
	var uregc jv.Claim
	uregc.Role = storeAdmin
	uregc.URL = upregURL
	uregc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uregc)
	h.Log.Debug("region update authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var ureg sdbi.Region
			uregsuc, uregerr := h.ProcessBody(r, &ureg)
			h.Log.Debug("uregsuc: ", uregsuc)
			h.Log.Debug("ureg: ", ureg)
			h.Log.Debug("uregerr: ", uregerr)
			if !uregsuc && uregerr != nil {
				http.Error(w, uregerr.Error(), http.StatusBadRequest)
			} else {
				uregres := h.Manager.UpdateRegion(&ureg)
				h.Log.Debug("uregres: ", *uregres)
				if uregres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uregres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uregfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uregfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetRegion godoc
// @Summary Get details of a region by id
// @Description Get details of a region
// @Tags Region (Geographic Sales Region)
// @Accept  json
// @Produce  json
// @Param id path string true "region id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.Region
// @Router /rs/region/get/id/{id}/{storeId} [get]
func (h *Six910Handler) GetRegion(w http.ResponseWriter, r *http.Request) {
	var gregURL = "/six910/rs/region/get"
	var gregc jv.Claim
	gregc.Role = customerRole
	gregc.URL = gregURL
	gregc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gregc)
	h.Log.Debug("region get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gregidStr = vars["id"]
			var gregstoreIDStr = vars["storeId"]
			id, gregiderr := strconv.ParseInt(gregidStr, 10, 64)
			storeID, gregsiderr := strconv.ParseInt(gregstoreIDStr, 10, 64)
			var gregres *sdbi.Region
			if gregiderr == nil && gregsiderr == nil {
				gregres = h.Manager.GetRegion(id, storeID)
				h.Log.Debug("gregres: ", gregres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Region
				gregres = &nc
			}
			resJSON, _ := json.Marshal(gregres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetRegionList godoc
// @Summary Get list of regions
// @Description Get list of regions for a store
// @Tags Region (Geographic Sales Region)
// @Accept  json
// @Produce  json
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Region
// @Router /rs/region/get/list/{storeId} [get]
func (h *Six910Handler) GetRegionList(w http.ResponseWriter, r *http.Request) {
	var greglURL = "/six910/rs/region/list"
	var gregcl jv.Claim
	gregcl.Role = customerRole
	gregcl.URL = greglURL
	gregcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gregcl)
	h.Log.Debug("region get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var reglstoreIDStr = vars["storeId"]
			storeID, sreglerr := strconv.ParseInt(reglstoreIDStr, 10, 64)
			var greglres *[]sdbi.Region
			if sreglerr == nil {
				greglres = h.Manager.GetRegionList(storeID)
				h.Log.Debug("get region list: ", greglres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Region{}
				greglres = &nc
			}
			resJSON, _ := json.Marshal(greglres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteRegion godoc
// @Summary Delete a region
// @Description Delete a region from the store
// @Tags Region (Geographic Sales Region)
// @Accept  json
// @Produce  json
// @Param id path string true "region id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/region/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeleteRegion(w http.ResponseWriter, r *http.Request) {
	var dregURL = "/six910/rs/region/delete"
	var dregs jv.Claim
	dregs.Role = storeAdmin
	dregs.URL = dregURL
	dregs.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dregs)
	h.Log.Debug("region delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dregidStr = vars["id"]
			var dregstoreIDStr = vars["storeId"]
			id, dregiderr := strconv.ParseInt(dregidStr, 10, 64)
			storeID, dregidserr := strconv.ParseInt(dregstoreIDStr, 10, 64)
			var dregres *m.Response
			if dregiderr == nil && dregidserr == nil {
				dregres = h.Manager.DeleteRegion(id, storeID)
				h.Log.Debug("dregres: ", dregres)
				if dregres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dregres = &nc
			}
			resJSON, _ := json.Marshal(dregres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
