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

//ExcludedSubRegionReq ExcludedSubRegionReq
type ExcludedSubRegionReq struct {
	StoreID           int64                  `json:"storeId"`
	ExcludedSubRegion sdbi.ExcludedSubRegion `json:"excludedSubRegion"`
}

// AddExcludedSubRegion godoc
// @Summary Add a new ExcludedSubRegion
// @Description Adds a new ExcludedSubRegion to a store
// @Tags Excluded Sub Regions (Excluded Geographic Sales Sub Regions)
// @Accept  json
// @Produce  json
// @Param excludedSubRegion body ExcludedSubRegionReq true "excludedSubRegion"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/excludedSubRegion/add [post]
func (h *Six910Handler) AddExcludedSubRegion(w http.ResponseWriter, r *http.Request) {
	var addesregURL = "/six910/rs/excludedSubRegion/add"
	var aesregc jv.Claim
	aesregc.Role = storeAdmin
	aesregc.URL = addesregURL
	aesregc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &aesregc)
	h.Log.Debug("excluded sub region add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var aesreg ExcludedSubRegionReq
			aesregsuc, aesregerr := h.ProcessBody(r, &aesreg)
			h.Log.Debug("aesregsuc: ", aesregsuc)
			h.Log.Debug("asreg: ", aesreg)
			h.Log.Debug("asregerr: ", aesregerr)
			if !aesregsuc && aesregerr != nil {
				http.Error(w, aesregerr.Error(), http.StatusBadRequest)
			} else {
				aesregres := h.Manager.AddExcludedSubRegion(&aesreg.ExcludedSubRegion, aesreg.StoreID)
				h.Log.Debug("aesregres: ", *aesregres)
				if aesregres.Success && aesregres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aesregres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aesrfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aesrfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateExcludedSubRegion UpdateExcludedSubRegion
func (h *Six910Handler) UpdateExcludedSubRegion(w http.ResponseWriter, r *http.Request) {
	var upesregURL = "/six910/rs/excludedSubRegion/update"
	var uesregc jv.Claim
	uesregc.Role = storeAdmin
	uesregc.URL = upesregURL
	uesregc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uesregc)
	h.Log.Debug("excluded sub region update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uesreg ExcludedSubRegionReq
			uesregsuc, uesregerr := h.ProcessBody(r, &uesreg)
			h.Log.Debug("uesregsuc: ", uesregsuc)
			h.Log.Debug("uesreg: ", uesreg)
			h.Log.Debug("uesregerr: ", uesregerr)
			if !uesregsuc && uesregerr != nil {
				http.Error(w, uesregerr.Error(), http.StatusBadRequest)
			} else {
				uesregres := h.Manager.UpdateExcludedSubRegion(&uesreg.ExcludedSubRegion, uesreg.StoreID)
				h.Log.Debug("uesregres: ", *uesregres)
				if uesregres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uesregres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uesregfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uesregfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetExcludedSubRegion GetExcludedSubRegion
func (h *Six910Handler) GetExcludedSubRegion(w http.ResponseWriter, r *http.Request) {
	var gesregURL = "/six910/rs/excludedSubRegion/get"
	var gesregc jv.Claim
	gesregc.Role = customerRole
	gesregc.URL = gesregURL
	gesregc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gesregc)
	h.Log.Debug("excluded sub region get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gesregidStr = vars["id"]
			var gesregstoreIDStr = vars["storeId"]
			id, gesregiderr := strconv.ParseInt(gesregidStr, 10, 64)
			storeID, gesregsiderr := strconv.ParseInt(gesregstoreIDStr, 10, 64)
			var gesregres *sdbi.ExcludedSubRegion
			if gesregiderr == nil && gesregsiderr == nil {
				gesregres = h.Manager.GetExcludedSubRegion(id, storeID)
				h.Log.Debug("gesregres: ", gesregres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.ExcludedSubRegion
				gesregres = &nc
			}
			resJSON, _ := json.Marshal(gesregres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetExcludedSubRegionList godoc
// @Summary Get list of ExcludedSubRegion
// @Description Get list of ExcludedSubRegion for a store
// @Tags Excluded Sub Regions (Excluded Geographic Sales Sub Regions)
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
// @Success 200 {array} six910-database-interface.ExcludedSubRegion
// @Router /rs/excludedSubRegion/get/list/{regionId}/{storeId} [get]
func (h *Six910Handler) GetExcludedSubRegionList(w http.ResponseWriter, r *http.Request) {
	var gesreglURL = "/six910/rs/excludedSubRegion/list"
	var gesregcl jv.Claim
	gesregcl.Role = customerRole
	gesregcl.URL = gesreglURL
	gesregcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gesregcl)
	h.Log.Debug("excluded sub region get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var eregIDlStr = vars["regionId"]
			var esreglstoreIDStr = vars["storeId"]
			regionID, eregIDlerr := strconv.ParseInt(eregIDlStr, 10, 64)
			storeID, esreglerr := strconv.ParseInt(esreglstoreIDStr, 10, 64)
			var gesreglres *[]sdbi.ExcludedSubRegion
			if eregIDlerr == nil && esreglerr == nil {
				gesreglres = h.Manager.GetExcludedSubRegionList(regionID, storeID)
				h.Log.Debug("get excluded sub region list: ", gesreglres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.ExcludedSubRegion{}
				gesreglres = &nc
			}
			resJSON, _ := json.Marshal(gesreglres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteExcludedSubRegion godoc
// @Summary Delete a ExcludedSubRegion
// @Description Delete a ExcludedSubRegion from the store
// @Tags Excluded Sub Regions (Excluded Geographic Sales Sub Regions)
// @Accept  json
// @Produce  json
// @Param id path string true "excludedSubRegion id"
// @Param regionId path string true "region id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/excludedSubRegion/delete/{id}/{regionId}/{storeId} [delete]
func (h *Six910Handler) DeleteExcludedSubRegion(w http.ResponseWriter, r *http.Request) {
	var desregURL = "/six910/rs/excludedSubRegion/delete"
	var desregs jv.Claim
	desregs.Role = storeAdmin
	desregs.URL = desregURL
	desregs.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &desregs)
	h.Log.Debug("excluded sub region delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var desregidStr = vars["id"]
			var dregidStr = vars["regionId"]
			var desregstoreIDStr = vars["storeId"]
			id, desregiderr := strconv.ParseInt(desregidStr, 10, 64)
			regionID, dregiderr := strconv.ParseInt(dregidStr, 10, 64)
			storeID, desregidserr := strconv.ParseInt(desregstoreIDStr, 10, 64)
			var desregres *m.Response
			if desregiderr == nil && dregiderr == nil && desregidserr == nil {
				desregres = h.Manager.DeleteExcludedSubRegion(id, regionID, storeID)
				h.Log.Debug("desregres: ", *desregres)
				if desregres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				desregres = &nc
			}
			resJSON, _ := json.Marshal(desregres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
