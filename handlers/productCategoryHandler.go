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

//ProductCategory ProductCategory
type ProductCategory struct {
	CategoryID int64 `json:"categoryId"`
	ProductID  int64 `json:"productId"`
}

//ProductCategoryReq ProductCategoryReq
type ProductCategoryReq struct {
	StoreID         int64           `json:"storeId"`
	ProductCategory ProductCategory `json:"productCategory"`
}

// AddProductCategory godoc
// @Summary Add a product to a category
// @Description Adds a product to a category in a store
// @Tags ProductCategory
// @Accept  json
// @Produce  json
// @Param productCategory body ProductCategoryReq true "Product to Category"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/productCategory/add [post]
func (h *Six910Handler) AddProductCategory(w http.ResponseWriter, r *http.Request) {
	var addpcURL = "/six910/rs/productCategory/add"
	var apcc jv.Claim
	apcc.Role = storeAdmin
	apcc.URL = addpcURL
	apcc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &apcc)
	h.Log.Debug("product category add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var apcr ProductCategoryReq
			apcsuc, apcerr := h.ProcessBody(r, &apcr)
			h.Log.Debug("apcsuc: ", apcsuc)
			h.Log.Debug("apcr: ", apcr)
			h.Log.Debug("apcerr: ", apcerr)
			if !apcsuc && apcerr != nil {
				http.Error(w, apcerr.Error(), http.StatusBadRequest)
			} else {
				var pc sdbi.ProductCategory
				pc.CategoryID = apcr.ProductCategory.CategoryID
				pc.ProductID = apcr.ProductCategory.ProductID
				apcres := h.Manager.AddProductCategory(&pc, apcr.StoreID)
				h.Log.Debug("apcres: ", *apcres)
				if apcres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(apcres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var apcf m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(apcf)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetProductCategoryList godoc
// @Summary Get list of category IDs for a product
// @Description Get list of category IDs for a product
// @Tags ProductCategory
// @Accept  json
// @Produce  json
// @Param productId path string true "productId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} int64
// @Router /rs/productCategory/list/{productId} [get]
func (h *Six910Handler) GetProductCategoryList(w http.ResponseWriter, r *http.Request) {
	var gcatlURL = "/six910/rs/productCategory/list"
	var gcatcl jv.Claim
	gcatcl.Role = customerRole
	gcatcl.URL = gcatlURL
	gcatcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("prod cat get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var ctlprodIDStr = vars["productId"]
			pID, splerr := strconv.ParseInt(ctlprodIDStr, 10, 64)
			var gpcatlres = []int64{}
			if splerr == nil {
				pcids := h.Manager.GetProductCategoryList(pID)
				if pcids != nil && len(*pcids) > 0 {
					gpcatlres = *pcids
				}
				h.Log.Debug("get prod cat list: ", gpcatlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []int64{}
				gpcatlres = nc
			}
			resJSON, _ := json.Marshal(gpcatlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteProductCategory godoc
// @Summary Delete a product from a category
// @Description Delete a product from a category in a store
// @Tags ProductCategory
// @Accept  json
// @Produce  json
// @Param categoryId path string true "category id"
// @Param productId path string true "product id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/productCategory/delete/{categoryId}/{productId}/{storeId} [delete]
func (h *Six910Handler) DeleteProductCategory(w http.ResponseWriter, r *http.Request) {
	var dpcURL = "/six910/rs/productCategory/delete"
	var dpcc jv.Claim
	dpcc.Role = storeAdmin
	dpcc.URL = dpcURL
	dpcc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dpcc)
	h.Log.Debug("product category delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var dpccatidStr = vars["categoryId"]
			var dpcprodidStr = vars["productId"]
			var dprodstoreIDStr = vars["storeId"]
			cid, dpccatiderr := strconv.ParseInt(dpccatidStr, 10, 64)
			pid, dpcprodiderr := strconv.ParseInt(dpcprodidStr, 10, 64)
			storeID, dpcsiderr := strconv.ParseInt(dprodstoreIDStr, 10, 64)
			var dpcres *m.Response
			if dpccatiderr == nil && dpcprodiderr == nil && dpcsiderr == nil {
				var pc sdbi.ProductCategory
				pc.CategoryID = cid
				pc.ProductID = pid
				dpcres = h.Manager.DeleteProductCategory(&pc, storeID)
				h.Log.Debug("dpcres: ", *dpcres)
				if dpcres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dpcres = &nc
			}
			resJSON, _ := json.Marshal(dpcres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
