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

//IncludedSubRegionReq IncludedSubRegionReq
type IncludedSubRegionReq struct {
	StoreID           int64                  `json:"storeId"`
	IncludedSubRegion sdbi.IncludedSubRegion `json:"includedSubRegion"`
}

// AddIncludedSubRegion godoc
// @Summary Add new IncludedSubRegion
// @Description Adds new IncludedSubRegion to a store
// @Tags Included Sub Regions (Included Geographic Sales Sub Regions)
// @Accept  json
// @Produce  json
// @Param includedSubRegion body IncludedSubRegionReq true "includedSubRegion"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/includedSubRegion/add [post]
func (h *Six910Handler) AddIncludedSubRegion(w http.ResponseWriter, r *http.Request) {
	var addisregURL = "/six910/rs/includedSubRegion/add"
	var aisregc jv.Claim
	aisregc.Role = storeAdmin
	aisregc.URL = addisregURL
	aisregc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &aisregc)
	h.Log.Debug("included sub region add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var aisreg IncludedSubRegionReq
			aisregsuc, aisregerr := h.ProcessBody(r, &aisreg)
			h.Log.Debug("aisregsuc: ", aisregsuc)
			h.Log.Debug("aisreg: ", aisreg)
			h.Log.Debug("aisregerr: ", aisregerr)
			if !aisregsuc && aisregerr != nil {
				http.Error(w, aisregerr.Error(), http.StatusBadRequest)
			} else {
				aisregres := h.Manager.AddIncludedSubRegion(&aisreg.IncludedSubRegion, aisreg.StoreID)
				h.Log.Debug("aisregres: ", *aisregres)
				if aisregres.Success && aisregres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aisregres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aisrfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aisrfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateIncludedSubRegion UpdateIncludedSubRegion
func (h *Six910Handler) UpdateIncludedSubRegion(w http.ResponseWriter, r *http.Request) {
	var upisregURL = "/six910/rs/includedSubRegion/update"
	var uisregc jv.Claim
	uisregc.Role = storeAdmin
	uisregc.URL = upisregURL
	uisregc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uisregc)
	h.Log.Debug("included sub region update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uisreg IncludedSubRegionReq
			uisregsuc, uisregerr := h.ProcessBody(r, &uisreg)
			h.Log.Debug("uisregsuc: ", uisregsuc)
			h.Log.Debug("uisreg: ", uisreg)
			h.Log.Debug("uisregerr: ", uisregerr)
			if !uisregsuc && uisregerr != nil {
				http.Error(w, uisregerr.Error(), http.StatusBadRequest)
			} else {
				uisregres := h.Manager.UpdateIncludedSubRegion(&uisreg.IncludedSubRegion, uisreg.StoreID)
				h.Log.Debug("uisregres: ", *uisregres)
				if uisregres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uisregres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uisregfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uisregfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetIncludedSubRegion GetIncludedSubRegion
func (h *Six910Handler) GetIncludedSubRegion(w http.ResponseWriter, r *http.Request) {
	var gisregURL = "/six910/rs/includedSubRegion/get"
	var gisregc jv.Claim
	gisregc.Role = customerRole
	gisregc.URL = gisregURL
	gisregc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gisregc)
	h.Log.Debug("included sub region get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gisregidStr = vars["id"]
			var gisregstoreIDStr = vars["storeId"]
			id, gisregiderr := strconv.ParseInt(gisregidStr, 10, 64)
			storeID, gisregsiderr := strconv.ParseInt(gisregstoreIDStr, 10, 64)
			var gisregres *sdbi.IncludedSubRegion
			if gisregiderr == nil && gisregsiderr == nil {
				gisregres = h.Manager.GetIncludedSubRegion(id, storeID)
				h.Log.Debug("gisregres: ", gisregres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.IncludedSubRegion
				gisregres = &nc
			}
			resJSON, _ := json.Marshal(gisregres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetIncludedSubRegionList godoc
// @Summary Get list of IncludedSubRegion
// @Description Get list of IncludedSubRegion for a store
// @Tags Included Sub Regions (Included Geographic Sales Sub Regions)
// @Accept  json
// @Produce  json
// @Param regionId path string true "region id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.IncludedSubRegion
// @Router /rs/includedSubRegion/get/list/{regionId}/{storeId} [get]
func (h *Six910Handler) GetIncludedSubRegionList(w http.ResponseWriter, r *http.Request) {
	var gisreglURL = "/six910/rs/includedSubRegion/list"
	var gisregcl jv.Claim
	gisregcl.Role = customerRole
	gisregcl.URL = gisreglURL
	gisregcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gisregcl)
	h.Log.Debug("included sub region get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var iregIDlStr = vars["regionId"]
			var isreglstoreIDStr = vars["storeId"]
			regionID, iregIDlerr := strconv.ParseInt(iregIDlStr, 10, 64)
			storeID, isreglerr := strconv.ParseInt(isreglstoreIDStr, 10, 64)
			var gisreglres *[]sdbi.IncludedSubRegion
			if iregIDlerr == nil && isreglerr == nil {
				gisreglres = h.Manager.GetIncludedSubRegionList(regionID, storeID)
				h.Log.Debug("get included sub region list: ", gisreglres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.IncludedSubRegion{}
				gisreglres = &nc
			}
			resJSON, _ := json.Marshal(gisreglres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteIncludedSubRegion godoc
// @Summary Delete a IncludedSubRegion
// @Description Delete IncludedSubRegion from the store
// @Tags Included Sub Regions (Included Geographic Sales Sub Regions)
// @Accept  json
// @Produce  json
// @Param id path string true "includedSubRegion id"
// @Param regionId path string true "region id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/includedSubRegion/delete/{id}/{regionId}/{storeId} [delete]
func (h *Six910Handler) DeleteIncludedSubRegion(w http.ResponseWriter, r *http.Request) {
	var disregURL = "/six910/rs/includedSubRegion/delete"
	var disregs jv.Claim
	disregs.Role = storeAdmin
	disregs.URL = disregURL
	disregs.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &disregs)
	h.Log.Debug("included sub region delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var disregidStr = vars["id"]
			var disaaregidStr = vars["regionId"]
			var disregstoreIDStr = vars["storeId"]
			id, disregiderr := strconv.ParseInt(disregidStr, 10, 64)
			regionID, daaregiderr := strconv.ParseInt(disaaregidStr, 10, 64)
			storeID, disregidserr := strconv.ParseInt(disregstoreIDStr, 10, 64)
			var disregres *m.Response
			if disregiderr == nil && daaregiderr == nil && disregidserr == nil {
				disregres = h.Manager.DeleteIncludedSubRegion(id, regionID, storeID)
				h.Log.Debug("disregres: ", disregres)
				if disregres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				disregres = &nc
			}
			resJSON, _ := json.Marshal(disregres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
