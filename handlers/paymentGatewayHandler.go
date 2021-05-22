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

//PaymentGatewayReq PaymentGatewayReq
type PaymentGatewayReq struct {
	StoreID        int64               `json:"storeId"`
	PaymentGateway sdbi.PaymentGateway `json:"paymentGateway"`
}

// AddPaymentGateway godoc
// @Summary Add a new PaymentGateway
// @Description Adds a new PaymentGateway to a store
// @Tags PaymentGateway
// @Accept  json
// @Produce  json
// @Param paymentGateway body PaymentGatewayReq true "paymentGateway"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/paymentGateway/add [post]
func (h *Six910Handler) AddPaymentGateway(w http.ResponseWriter, r *http.Request) {
	var addpgwURL = "/six910/rs/paymentGateway/add"
	var apgwc jv.Claim
	apgwc.Role = storeAdmin
	apgwc.URL = addpgwURL
	apgwc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &apgwc)
	h.Log.Debug("payment gateway add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var apgw PaymentGatewayReq
			apgwsuc, apgwerr := h.ProcessBody(r, &apgw)
			h.Log.Debug("apgwsuc: ", apgwsuc)
			h.Log.Debug("apgw: ", apgw)
			h.Log.Debug("apgwerr: ", apgwerr)
			if !apgwsuc && apgwerr != nil {
				http.Error(w, apgwerr.Error(), http.StatusBadRequest)
			} else {
				apgwres := h.Manager.AddPaymentGateway(&apgw.PaymentGateway, apgw.StoreID)
				h.Log.Debug("apgwres: ", *apgwres)
				if apgwres.Success && apgwres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(apgwres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var apgwfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(apgwfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdatePaymentGateway godoc
// @Summary Update a PaymentGateway
// @Description Update PaymentGateway data
// @Tags PaymentGateway
// @Accept  json
// @Produce  json
// @Param paymentGateway body PaymentGatewayReq true "paymentGateway"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/paymentGateway/update [put]
func (h *Six910Handler) UpdatePaymentGateway(w http.ResponseWriter, r *http.Request) {
	var uppgwURL = "/six910/rs/paymentGateway/update"
	var upgwc jv.Claim
	upgwc.Role = storeAdmin
	upgwc.URL = uppgwURL
	upgwc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &upgwc)
	h.Log.Debug("payment gateway update authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var upgw PaymentGatewayReq
			upgwsuc, upgwerr := h.ProcessBody(r, &upgw)
			h.Log.Debug("upgwsuc: ", upgwsuc)
			h.Log.Debug("upgw: ", upgw)
			h.Log.Debug("upgwerr: ", upgwerr)
			if !upgwsuc && upgwerr != nil {
				http.Error(w, upgwerr.Error(), http.StatusBadRequest)
			} else {
				upgwres := h.Manager.UpdatePaymentGateway(&upgw.PaymentGateway, upgw.StoreID)
				h.Log.Debug("upgwres: ", *upgwres)
				if upgwres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(upgwres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var upgwfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(upgwfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetPaymentGateway godoc
// @Summary Get details of a PaymentGateway by id
// @Description Get details of a PaymentGateway
// @Tags PaymentGateway
// @Accept  json
// @Produce  json
// @Param id path string true "paymentGateway id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.PaymentGateway
// @Router /rs/paymentGateway/get/id/{id}/{storeId} [get]
func (h *Six910Handler) GetPaymentGateway(w http.ResponseWriter, r *http.Request) {
	var gpgwURL = "/six910/rs/paymentGateway/get"
	var gpgwc jv.Claim
	gpgwc.Role = customerRole
	gpgwc.URL = gpgwURL
	gpgwc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("payment gateway get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gpgwidStr = vars["id"]
			var gpgwstoreIDStr = vars["storeId"]
			id, gpgwiderr := strconv.ParseInt(gpgwidStr, 10, 64)
			storeID, gpgwsiderr := strconv.ParseInt(gpgwstoreIDStr, 10, 64)
			var gpgwres *sdbi.PaymentGateway
			if gpgwiderr == nil && gpgwsiderr == nil {
				gpgwres = h.Manager.GetPaymentGateway(id, storeID)
				h.Log.Debug("gpgwres: ", gpgwres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.PaymentGateway
				gpgwres = &nc
			}
			resJSON, _ := json.Marshal(gpgwres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetPaymentGatewayByName godoc
// @Summary Get details of a PaymentGateway by name
// @Description Get details of a PaymentGateway
// @Tags PaymentGateway
// @Accept  json
// @Produce  json
// @Param id path string true "paymentGateway name"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.PaymentGateway
// @Router /rs/paymentGateway/get/name/{name}/{storeId} [get]
func (h *Six910Handler) GetPaymentGatewayByName(w http.ResponseWriter, r *http.Request) {
	var gpgwURL = "/six910/rs/paymentGateway/get/name"
	var gpgwc jv.Claim
	gpgwc.Role = customerRole
	gpgwc.URL = gpgwURL
	gpgwc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("payment gateway get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gpgwname = vars["name"]
			var gpgwstoreIDStr = vars["storeId"]
			//id, gpgwiderr := strconv.ParseInt(gpgwidStr, 10, 64)
			storeID, gpgwsiderr := strconv.ParseInt(gpgwstoreIDStr, 10, 64)
			var gpgwnres *sdbi.PaymentGateway
			if gpgwsiderr == nil {
				gpgwnres = h.Manager.GetPaymentGatewayByName(gpgwname, storeID)
				h.Log.Debug("gpgwres: ", gpgwnres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.PaymentGateway
				gpgwnres = &nc
			}
			resJSON, _ := json.Marshal(gpgwnres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetPaymentGateways godoc
// @Summary Get list of PaymentGateway
// @Description Get list of PaymentGateway for a store
// @Tags PaymentGateway
// @Accept  json
// @Produce  json
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.PaymentGateway
// @Router /rs/paymentGateway/get/list/{storeId} [get]
func (h *Six910Handler) GetPaymentGateways(w http.ResponseWriter, r *http.Request) {
	var gpgwlURL = "/six910/rs/paymentGateway/list"
	var gpgwcl jv.Claim
	gpgwcl.Role = customerRole
	gpgwcl.URL = gpgwlURL
	gpgwcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("payment gateway get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var pgwlstoreIDStr = vars["storeId"]
			storeID, spgwlerr := strconv.ParseInt(pgwlstoreIDStr, 10, 64)
			var gpgwlres *[]sdbi.PaymentGateway
			if spgwlerr == nil {
				gpgwlres = h.Manager.GetPaymentGateways(storeID)
				h.Log.Debug("get store plugin list: ", gpgwlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.PaymentGateway{}
				gpgwlres = &nc
			}
			resJSON, _ := json.Marshal(gpgwlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeletePaymentGateway godoc
// @Summary Delete a PaymentGateway
// @Description Delete a PaymentGateway from the store
// @Tags PaymentGateway
// @Accept  json
// @Produce  json
// @Param id path string true "paymentGateway id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/paymentGateway/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeletePaymentGateway(w http.ResponseWriter, r *http.Request) {
	var dpgwURL = "/six910/rs/paymentGateway/delete"
	var dpgws jv.Claim
	dpgws.Role = storeAdmin
	dpgws.URL = dpgwURL
	dpgws.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dpgws)
	h.Log.Debug("payment gateway delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dpgwidStr = vars["id"]
			var dpgwstoreIDStr = vars["storeId"]
			id, dpgwiderr := strconv.ParseInt(dpgwidStr, 10, 64)
			storeID, dpgwidserr := strconv.ParseInt(dpgwstoreIDStr, 10, 64)
			var dpgwres *m.Response
			if dpgwiderr == nil && dpgwidserr == nil {
				dpgwres = h.Manager.DeletePaymentGateway(id, storeID)
				h.Log.Debug("delete payment gateway: ", dpgwres)
				if dpgwres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dpgwres = &nc
			}
			resJSON, _ := json.Marshal(dpgwres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
