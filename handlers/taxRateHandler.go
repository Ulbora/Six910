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

// AddTaxRate godoc
// @Summary Add tax rate
// @Description Adds tax rate to a store
// @Tags TaxRate
// @Accept  json
// @Produce  json
// @Param tax rate body six910-database-interface.TaxRate true "tax rate"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/taxRate/add [post]
func (h *Six910Handler) AddTaxRate(w http.ResponseWriter, r *http.Request) {
	var addtrURL = "/six910/rs/taxRate/add"
	var atrc jv.Claim
	atrc.Role = storeAdmin
	atrc.URL = addtrURL
	atrc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &atrc)
	h.Log.Debug("tax rate add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var atr sdbi.TaxRate
			atrsuc, atrerr := h.ProcessBody(r, &atr)
			h.Log.Debug("atrsuc: ", atrsuc)
			h.Log.Debug("ains: ", atr)
			h.Log.Debug("atrerr: ", atrerr)
			if !atrsuc && atrerr != nil {
				http.Error(w, atrerr.Error(), http.StatusBadRequest)
			} else {
				atrres := h.Manager.AddTaxRate(&atr)
				h.Log.Debug("atrres: ", *atrres)
				if atrres.Success && atrres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(atrres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var atrfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(atrfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateTaxRate godoc
// @Summary Update TaxRate
// @Description Update TaxRate data
// @Tags TaxRate
// @Accept  json
// @Produce  json
// @Param TaxRate body six910-database-interface.TaxRate true "tax rate"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/taxRate/update [put]
func (h *Six910Handler) UpdateTaxRate(w http.ResponseWriter, r *http.Request) {
	var uptrURL = "/six910/rs/taxRate/update"
	var utrc jv.Claim
	utrc.Role = storeAdmin
	utrc.URL = uptrURL
	utrc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &utrc)
	h.Log.Debug("tax rate update authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var utr sdbi.TaxRate
			utrsuc, utrerr := h.ProcessBody(r, &utr)
			h.Log.Debug("uinssuc: ", utrsuc)
			h.Log.Debug("uins: ", utr)
			h.Log.Debug("uinserr: ", utrerr)
			if !utrsuc && utrerr != nil {
				http.Error(w, utrerr.Error(), http.StatusBadRequest)
			} else {
				utrres := h.Manager.UpdateTaxRate(&utr)
				h.Log.Debug("utrres: ", *utrres)
				if utrres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(utrres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var utrfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(utrfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetTaxRate godoc
// @Summary Get list of TaxRates
// @Description Get list of TaxRate by country and state for a store
// @Tags TaxRate
// @Accept  json
// @Produce  json
// @Param country path string true "country"
// @Param state path string true "state"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.TaxRate
// @Router /rs/taxRate/get/country/{country}/{state}/{storeId} [get]
func (h *Six910Handler) GetTaxRate(w http.ResponseWriter, r *http.Request) {
	var gtrURL = "/six910/rs/taxRate/get"
	var gtrc jv.Claim
	gtrc.Role = customerRole
	gtrc.URL = gtrURL
	gtrc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("tax rate get authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var gtrcntry = vars["country"]
			var gtrstate = vars["state"]
			var trstoreIDStr = vars["storeId"]
			storeID, strerr := strconv.ParseInt(trstoreIDStr, 10, 64)
			var gtrres *[]sdbi.TaxRate
			if strerr == nil {
				gtrres = h.Manager.GetTaxRate(gtrcntry, gtrstate, storeID)
				h.Log.Debug("get tar rate: ", *gtrres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.TaxRate{}
				gtrres = &nc
			}
			resJSON, _ := json.Marshal(gtrres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetTaxRateList godoc
// @Summary Get list of TaxRates
// @Description Get list of TaxRate for a store
// @Tags TaxRate
// @Accept  json
// @Produce  json
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.TaxRate
// @Router /rs/taxRate/get/list/{storeId} [get]
func (h *Six910Handler) GetTaxRateList(w http.ResponseWriter, r *http.Request) {
	var gtrlURL = "/six910/rs/taxRate/list"
	var gtrcl jv.Claim
	gtrcl.Role = customerRole
	gtrcl.URL = gtrlURL
	gtrcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gtrcl)
	h.Log.Debug("tax rate get authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var trlstoreIDStr = vars["storeId"]
			storeID, strerr := strconv.ParseInt(trlstoreIDStr, 10, 64)
			var gtrlres *[]sdbi.TaxRate
			if strerr == nil {
				gtrlres = h.Manager.GetTaxRateList(storeID)
				h.Log.Debug("get tar rate: ", *gtrlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.TaxRate{}
				gtrlres = &nc
			}
			resJSON, _ := json.Marshal(gtrlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteTaxRate godoc
// @Summary Delete a TaxRate
// @Description Delete a TaxRate from the store
// @Tags TaxRate
// @Accept  json
// @Produce  json
// @Param id path string true "tax rate id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/taxRate/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeleteTaxRate(w http.ResponseWriter, r *http.Request) {
	var dtrURL = "/six910/rs/tarRate/delete"
	var dtrs jv.Claim
	dtrs.Role = storeAdmin
	dtrs.URL = dtrURL
	dtrs.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dtrs)
	h.Log.Debug("tax rate delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dtridStr = vars["id"]
			var dtrstoreIDStr = vars["storeId"]
			id, dtriderr := strconv.ParseInt(dtridStr, 10, 64)
			storeID, dtridserr := strconv.ParseInt(dtrstoreIDStr, 10, 64)
			var dtrres *m.Response
			if dtriderr == nil && dtridserr == nil {
				dtrres = h.Manager.DeleteTaxRate(id, storeID)
				h.Log.Debug("dinsres: ", dtrres)
				if dtrres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dtrres = &nc
			}
			resJSON, _ := json.Marshal(dtrres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
