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

// AddShippingCarrier godoc
// @Summary Add a new Shipping Carrier
// @Description Adds a new Shipping Carrier to a store
// @Tags ShippingCarrier
// @Accept  json
// @Produce  json
// @Param shippingCarrier body six910-database-interface.ShippingCarrier true "shippingCarrier"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/shippingCarrier/add [post]
func (h *Six910Handler) AddShippingCarrier(w http.ResponseWriter, r *http.Request) {
	var addscrURL = "/six910/rs/shippingCarrier/add"
	var ascrc jv.Claim
	ascrc.Role = storeAdmin
	ascrc.URL = addscrURL
	ascrc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ascrc)
	h.Log.Debug("shipping carrier add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var ascr sdbi.ShippingCarrier
			ascrsuc, ascrerr := h.ProcessBody(r, &ascr)
			h.Log.Debug("ascrsuc: ", ascrsuc)
			h.Log.Debug("ascr: ", ascr)
			h.Log.Debug("ascrerr: ", ascrerr)
			if !ascrsuc && ascrerr != nil {
				http.Error(w, ascrerr.Error(), http.StatusBadRequest)
			} else {
				ascrres := h.Manager.AddShippingCarrier(&ascr)
				h.Log.Debug("ascrres: ", *ascrres)
				if ascrres.Success && ascrres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ascrres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var ascrfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ascrfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateShippingCarrier godoc
// @Summary Update a Shipping Carrier
// @Description Update Shipping Carrier data
// @Tags ShippingCarrier
// @Accept  json
// @Produce  json
// @Param shippingCarrier body six910-database-interface.ShippingCarrier true "shippingCarrier"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/shippingCarrier/update [put]
func (h *Six910Handler) UpdateShippingCarrier(w http.ResponseWriter, r *http.Request) {
	var upscrURL = "/six910/rs/shippingCarrier/update"
	var uscrc jv.Claim
	uscrc.Role = storeAdmin
	uscrc.URL = upscrURL
	uscrc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uscrc)
	h.Log.Debug("shipping carrier update authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uscr sdbi.ShippingCarrier
			uscrsuc, uscrerr := h.ProcessBody(r, &uscr)
			h.Log.Debug("uscrsuc: ", uscrsuc)
			h.Log.Debug("uscr: ", uscr)
			h.Log.Debug("uscrerr: ", uscrerr)
			if !uscrsuc && uscrerr != nil {
				http.Error(w, uscrerr.Error(), http.StatusBadRequest)
			} else {
				uscrres := h.Manager.UpdateShippingCarrier(&uscr)
				h.Log.Debug("uscrres: ", *uscrres)
				if uscrres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uscrres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uscrfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uscrfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetShippingCarrier godoc
// @Summary Get details of a Shipping Carrier by id
// @Description Get details of a Shipping Carrier
// @Tags ShippingCarrier
// @Accept  json
// @Produce  json
// @Param id path string true "shippingCarrier id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.ShippingCarrier
// @Router /rs/shippingCarrier/get/id/{id}/{storeId} [get]
func (h *Six910Handler) GetShippingCarrier(w http.ResponseWriter, r *http.Request) {
	var gscrURL = "/six910/rs/shippingCarrier/get"
	var gscrc jv.Claim
	gscrc.Role = customerRole
	gscrc.URL = gscrURL
	gscrc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gscrc)
	h.Log.Debug("shipping carrier get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gscridStr = vars["id"]
			var gscrstoreIDStr = vars["storeId"]
			id, gscriderr := strconv.ParseInt(gscridStr, 10, 64)
			storeID, gscrsiderr := strconv.ParseInt(gscrstoreIDStr, 10, 64)
			var gscrres *sdbi.ShippingCarrier
			if gscriderr == nil && gscrsiderr == nil {
				gscrres = h.Manager.GetShippingCarrier(id, storeID)
				h.Log.Debug("gscrres: ", gscrres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.ShippingCarrier
				gscrres = &nc
			}
			resJSON, _ := json.Marshal(gscrres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetShippingCarrierList godoc
// @Summary Get list of Shipping Carrier
// @Description Get list of Shipping Carrier for a store
// @Tags ShippingCarrier
// @Accept  json
// @Produce  json
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.ShippingCarrier
// @Router /rs/shippingCarrier/get/list/{storeId} [get]
func (h *Six910Handler) GetShippingCarrierList(w http.ResponseWriter, r *http.Request) {
	var gscrlURL = "/six910/rs/shippingCarrier/list"
	var gscrcl jv.Claim
	gscrcl.Role = customerRole
	gscrcl.URL = gscrlURL
	gscrcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gscrcl)
	h.Log.Debug("shipping carrier get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var scrlstoreIDStr = vars["storeId"]
			storeID, sscrlerr := strconv.ParseInt(scrlstoreIDStr, 10, 64)
			var gscrlres *[]sdbi.ShippingCarrier
			if sscrlerr == nil {
				gscrlres = h.Manager.GetShippingCarrierList(storeID)
				h.Log.Debug("get shipping carrier list: ", gscrlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.ShippingCarrier{}
				gscrlres = &nc
			}
			resJSON, _ := json.Marshal(gscrlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteShippingCarrier godoc
// @Summary Delete a Shipping Carrier
// @Description Delete a Shipping Carrier from the store
// @Tags ShippingCarrier
// @Accept  json
// @Produce  json
// @Param id path string true "shippingCarrier id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/shippingCarrier/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeleteShippingCarrier(w http.ResponseWriter, r *http.Request) {
	var dscrURL = "/six910/rs/shippingCarrier/delete"
	var dscrs jv.Claim
	dscrs.Role = storeAdmin
	dscrs.URL = dscrURL
	dscrs.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dscrs)
	h.Log.Debug("shipping carrier delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dscridStr = vars["id"]
			var dscrstoreIDStr = vars["storeId"]
			id, dscriderr := strconv.ParseInt(dscridStr, 10, 64)
			storeID, dscridserr := strconv.ParseInt(dscrstoreIDStr, 10, 64)
			var dscrres *m.Response
			if dscriderr == nil && dscridserr == nil {
				dscrres = h.Manager.DeleteShippingCarrier(id, storeID)
				h.Log.Debug("delete shipping carrier: ", dscrres)
				if dscrres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dscrres = &nc
			}
			resJSON, _ := json.Marshal(dscrres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
