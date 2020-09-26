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

// AddCustomer godoc
// @Summary Add a new customer
// @Description Adds a new customer to a store
// @Tags Customer
// @Accept  json
// @Produce  json
// @Param customer body six910-database-interface.Customer true "customer"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/customer/add [post]
func (h *Six910Handler) AddCustomer(w http.ResponseWriter, r *http.Request) {
	var addCusURL = "/six910/rs/customer/add"
	var acc jv.Claim
	acc.Role = storeAdmin
	acc.URL = addCusURL
	acc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("cus add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var cust sdbi.Customer
			scsuc, scerr := h.ProcessBody(r, &cust)
			h.Log.Debug("scsuc: ", scsuc)
			h.Log.Debug("cust: ", cust)
			h.Log.Debug("scerr: ", scerr)
			if !scsuc && scerr != nil {
				http.Error(w, scerr.Error(), http.StatusBadRequest)
			} else {
				acres := h.Manager.AddCustomer(&cust)
				h.Log.Debug("acres: ", *acres)
				if acres.Success && acres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(acres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var acfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(acfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateCustomer godoc
// @Summary Update a customer
// @Description Update customer data
// @Tags Customer
// @Accept  json
// @Produce  json
// @Param customer body six910-database-interface.Customer true "customer"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/customer/update [put]
func (h *Six910Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var upCusURL = "/six910/rs/customer/update"
	var ucc jv.Claim
	ucc.Role = customerRole
	ucc.URL = upCusURL
	ucc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ucc)
	h.Log.Debug("customer update authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var ucus sdbi.Customer
			ucsuc, ucerr := h.ProcessBody(r, &ucus)
			h.Log.Debug("ucsuc: ", ucsuc)
			h.Log.Debug("ucus: ", ucus)
			h.Log.Debug("ucerr: ", ucerr)
			if !ucsuc && ucerr != nil {
				http.Error(w, ucerr.Error(), http.StatusBadRequest)
			} else {
				ucres := h.Manager.UpdateCustomer(&ucus)
				h.Log.Debug("ucres: ", *ucres)
				if ucres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ucres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var ucfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ucfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetCustomer godoc
// @Summary Get details of a customer by email
// @Description Get details of a customer
// @Tags Customer
// @Accept  json
// @Produce  json
// @Param email path string true "customer email"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.Customer
// @Router /rs/customer/get/email/{email}/{storeId} [get]
func (h *Six910Handler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	var gCusURL = "/six910/rs/customer/get"
	var gcc jv.Claim
	gcc.Role = customerRole
	gcc.URL = gCusURL
	gcc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)

	h.Log.Debug("cus get authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var storeName = vars["email"]
			var storeIDStr = vars["storeId"]
			storeID, err := strconv.ParseInt(storeIDStr, 10, 64)
			var gcres *sdbi.Customer
			if err == nil {
				gcres = h.Manager.GetCustomer(storeName, storeID)
				h.Log.Debug("getCust: ", gcres)
				w.WriteHeader(http.StatusOK)
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
				var nc sdbi.Customer
				gcres = &nc
			}
			resJSON, _ := json.Marshal(gcres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetCustomerID godoc
// @Summary Get details of a customer by id
// @Description Get details of a customer
// @Tags Customer
// @Accept  json
// @Produce  json
// @Param id path string true "customer id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.Customer
// @Router /rs/customer/get/id/{id}/{storeId} [get]
func (h *Six910Handler) GetCustomerID(w http.ResponseWriter, r *http.Request) {
	var gCus2URL = "/six910/rs/customer/get/id"
	var gcc2 jv.Claim
	gcc2.Role = customerRole
	gcc2.URL = gCus2URL
	gcc2.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("cus get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var cidStr = vars["id"]
			var storeIDStr = vars["storeId"]
			cid, cerr := strconv.ParseInt(cidStr, 10, 64)
			storeID, serr := strconv.ParseInt(storeIDStr, 10, 64)
			var gcres *sdbi.Customer
			if cerr == nil && serr == nil {
				gcres = h.Manager.GetCustomerID(cid, storeID)
				h.Log.Debug("getCustId: ", gcres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Customer
				gcres = &nc
			}
			resJSON, _ := json.Marshal(gcres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetCustomerList godoc
// @Summary Get list of a customers
// @Description Get list of a customers for a store
// @Tags Customer
// @Accept  json
// @Produce  json
// @Param storeId path string true "store storeId"
// @Param start path string true "start index 0 based"
// @Param end path string true "end index"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Customer
// @Router /rs/customer/get/list/{storeId}/{start}/{end} [get]
func (h *Six910Handler) GetCustomerList(w http.ResponseWriter, r *http.Request) {
	var gCuslURL = "/six910/rs/customer/list"
	var gccl jv.Claim
	gccl.Role = storeAdmin
	gccl.URL = gCuslURL
	gccl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gccl)
	h.Log.Debug("cus get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var storeIDStr = vars["storeId"]
			var cusstartStr = vars["start"]
			var cusendStr = vars["end"]
			storeID, serr := strconv.ParseInt(storeIDStr, 10, 64)
			cusStart, cusstarterr := strconv.ParseInt(cusstartStr, 10, 64)
			cusEnd, cusenderr := strconv.ParseInt(cusendStr, 10, 64)
			var gclres *[]sdbi.Customer
			if serr == nil && cusstarterr == nil && cusenderr == nil {
				gclres = h.Manager.GetCustomerList(storeID, cusStart, cusEnd)
				h.Log.Debug("getCust list: ", gclres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Customer{}
				gclres = &nc
			}
			resJSON, _ := json.Marshal(gclres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteCustomer godoc
// @Summary Delete a customer
// @Description Delete a customer from the store
// @Tags Customer
// @Accept  json
// @Produce  json
// @Param id path string true "customer id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/customer/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	var dCusURL = "/six910/rs/customer/delete"
	var dcs jv.Claim
	dcs.Role = storeAdmin
	dcs.URL = dCusURL
	dcs.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dcs)
	h.Log.Debug("cust delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dcidStr = vars["id"]
			var dstoreIDStr = vars["storeId"]
			cid, cerr := strconv.ParseInt(dcidStr, 10, 64)
			storeID, serr := strconv.ParseInt(dstoreIDStr, 10, 64)
			var dcres *m.Response
			if cerr == nil && serr == nil {
				dcres = h.Manager.DeleteCustomer(cid, storeID)
				h.Log.Debug("deleteCust: ", dcres)
				if dcres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dcres = &nc
			}
			resJSON, _ := json.Marshal(dcres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
