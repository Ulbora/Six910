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

// AddDistributor godoc
// @Summary Add a new distributor
// @Description Adds a new distributor to a store
// @Tags Distributor
// @Accept  json
// @Produce  json
// @Param distributor body six910-database-interface.Distributor true "distributor"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/distributor/add [post]
func (h *Six910Handler) AddDistributor(w http.ResponseWriter, r *http.Request) {
	var addDisURL = "/six910/rs/distributor/add"
	var adc jv.Claim
	adc.Role = storeAdmin
	adc.URL = addDisURL
	adc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &adc)
	h.Log.Debug("dist add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var dist sdbi.Distributor
			sdsuc, sderr := h.ProcessBody(r, &dist)
			h.Log.Debug("sdsuc: ", sdsuc)
			h.Log.Debug("dist: ", dist)
			h.Log.Debug("sderr: ", sderr)
			if !sdsuc && sderr != nil {
				http.Error(w, sderr.Error(), http.StatusBadRequest)
			} else {
				adres1 := h.Manager.AddDistributor(&dist)
				h.Log.Debug("acres: ", *adres1)
				if adres1.Success && adres1.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(adres1)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var adfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(adfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateDistributor godoc
// @Summary Update a distributor
// @Description Update distributor data
// @Tags Distributor
// @Accept  json
// @Produce  json
// @Param distributor body six910-database-interface.Distributor true "distributor"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/distributor/update [put]
func (h *Six910Handler) UpdateDistributor(w http.ResponseWriter, r *http.Request) {
	var upDisURL = "/six910/rs/distributor/update"
	var udc jv.Claim
	udc.Role = storeAdmin
	udc.URL = upDisURL
	udc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &udc)
	h.Log.Debug("dist update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var udis sdbi.Distributor
			udsuc, uderr := h.ProcessBody(r, &udis)
			h.Log.Debug("udsuc: ", udsuc)
			h.Log.Debug("udis: ", udis)
			h.Log.Debug("uderr: ", uderr)
			if !udsuc && uderr != nil {
				http.Error(w, uderr.Error(), http.StatusBadRequest)
			} else {
				udres := h.Manager.UpdateDistributor(&udis)
				h.Log.Debug("udres: ", *udres)
				if udres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(udres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var udfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(udfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetDistributor godoc
// @Summary Get details of a distributor by id
// @Description Get details of a distributor
// @Tags Distributor
// @Accept  json
// @Produce  json
// @Param id path string true "distributor id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.Distributor
// @Router /rs/distributor/get/id/{id}/{storeId} [get]
func (h *Six910Handler) GetDistributor(w http.ResponseWriter, r *http.Request) {
	var gDistURL = "/six910/rs/distributor/get"
	var gdc2 jv.Claim
	gdc2.Role = customerRole
	gdc2.URL = gDistURL
	gdc2.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gdc2)
	h.Log.Debug("dist get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var didStr = vars["id"]
			var storeIDStr = vars["storeId"]
			did, derr := strconv.ParseInt(didStr, 10, 64)
			storeID, serr := strconv.ParseInt(storeIDStr, 10, 64)
			var gdres *sdbi.Distributor
			if derr == nil && serr == nil {
				gdres = h.Manager.GetDistributor(did, storeID)
				h.Log.Debug("getDist: ", gdres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Distributor
				gdres = &nc
			}
			resJSON, _ := json.Marshal(gdres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetDistributorList godoc
// @Summary Get list of distributors
// @Description Get list of distributors for a store
// @Tags Distributor
// @Accept  json
// @Produce  json
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Distributor
// @Router /rs/distributor/get/list/{storeId} [get]
func (h *Six910Handler) GetDistributorList(w http.ResponseWriter, r *http.Request) {
	var gDislURL = "/six910/rs/distributor/list"
	var gdcl jv.Claim
	gdcl.Role = storeAdmin
	gdcl.URL = gDislURL
	gdcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gdcl)
	h.Log.Debug("dist get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var storeIDStr = vars["storeId"]
			storeID, serr := strconv.ParseInt(storeIDStr, 10, 64)
			var gdlres *[]sdbi.Distributor
			if serr == nil {
				gdlres = h.Manager.GetDistributorList(storeID)
				h.Log.Debug("getDist list: ", gdlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Distributor{}
				gdlres = &nc
			}
			resJSON, _ := json.Marshal(gdlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteDistributor godoc
// @Summary Delete a distributor
// @Description Delete a distributor from the store
// @Tags Distributor
// @Accept  json
// @Produce  json
// @Param id path string true "distributor id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/distributor/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeleteDistributor(w http.ResponseWriter, r *http.Request) {
	var dDisURL = "/six910/rs/distributor/delete"
	var dds jv.Claim
	dds.Role = storeAdmin
	dds.URL = dDisURL
	dds.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dds)
	h.Log.Debug("dist delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var ddidStr = vars["id"]
			var ddstoreIDStr = vars["storeId"]
			did, cerr := strconv.ParseInt(ddidStr, 10, 64)
			storeID, serr := strconv.ParseInt(ddstoreIDStr, 10, 64)
			var ddres *m.Response
			if cerr == nil && serr == nil {
				ddres = h.Manager.DeleteDistributor(did, storeID)
				h.Log.Debug("deleteCust: ", ddres)
				if ddres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				ddres = &nc
			}
			resJSON, _ := json.Marshal(ddres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
