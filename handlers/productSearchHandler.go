package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	jv "github.com/Ulbora/GoAuth2JwtValidator"

	//m "github.com/Ulbora/Six910/managers"
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

// GetProductManufacturerListByProductName godoc
// @Summary Get list of Manufacturers for a product name
// @Description Get list of Manufacturers for a product name for a store
// @Tags Manufacturer
// @Accept  json
// @Produce  json
// @Param name path string true "product name"
// @Param storeId path string true "store storeId"
// @Param apiKey header string true "apiKey required"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} string
// @Router /rs/manufacturer/get/product/name/{name}/{storeId} [get]
func (h *Six910Handler) GetProductManufacturerListByProductName(w http.ResponseWriter, r *http.Request) {
	var gprodmnlURL = "/six910/rs/manufacturer/search/product/name/list"
	var gprodmncl jv.Claim
	gprodmncl.Role = customerRole
	gprodmncl.URL = gprodmnlURL
	gprodmncl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("produce manf list prod name get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var prodmName = vars["name"]
			var prodmnlstoreIDStr = vars["storeId"]
			storeID, sprodnlerr := strconv.ParseInt(prodmnlstoreIDStr, 10, 64)
			var gprodmnlres *[]string
			if sprodnlerr == nil {
				gprodmnlres = h.Manager.GetProductManufacturerListByProductName(prodmName, storeID)
				h.Log.Debug("get product manf name list: ", gprodmnlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []string{}
				gprodmnlres = &nc
			}
			resJSON, _ := json.Marshal(gprodmnlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetProductManufacturerListByProductSearch product desc attrs
// @Description Get list of Manufacturers for a product search attributes for a store
// @Tags Manufacturer
// @Accept  json
// @Produce  json
// @Param name path string true "product desc attrs"
// @Param storeId path string true "store storeId"
// @Param apiKey header string true "apiKey required"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} string
// @Router /rs/manufacturer/get/product/desc/{search}/{storeId} [get]
func (h *Six910Handler) GetProductManufacturerListByProductSearch(w http.ResponseWriter, r *http.Request) {
	var gprprodmnlURL = "/six910/rs/manufacturer/search/product/desc/list"
	var gprprodmncl jv.Claim
	gprprodmncl.Role = customerRole
	gprprodmncl.URL = gprprodmnlURL
	gprprodmncl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("produce manf list prod desc attr get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var prprodmattr = vars["search"]
			var prprodmnlstoreIDStr = vars["storeId"]
			storeID, sprodnlerr := strconv.ParseInt(prprodmnlstoreIDStr, 10, 64)
			var gprprodmnlres *[]string
			if sprodnlerr == nil {
				gprprodmnlres = h.Manager.GetProductManufacturerListByProductSearch(prprodmattr, storeID)
				h.Log.Debug("get product manf desc attrs list: ", gprprodmnlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []string{}
				gprprodmnlres = &nc
			}
			resJSON, _ := json.Marshal(gprprodmnlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetProductByNameAndManufacturerName godoc
// @Summary Get list of products by product name and manufacturer
// @Description Get list of products by name and manufacturerfor a store
// @Tags Product
// @Accept  json
// @Produce  json
// @Param manufacturer path string true "manufacturer name"
// @Param name path string true "product name"
// @Param storeId path string true "store storeId"
// @Param start path string true "start index 0 based"
// @Param end path string true "end index"
// @Param apiKey header string true "apiKey required"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Product
// @Router /rs/product/get/manufacturer/name/{manufacturer}/{name}/{storeId}/{start}/{end} [get]
func (h *Six910Handler) GetProductByNameAndManufacturerName(w http.ResponseWriter, r *http.Request) {
	var gprodnmlURL = "/six910/rs/product/name/manufacturer/list"
	var gprodnmcl jv.Claim
	gprodnmcl.Role = customerRole
	gprodnmcl.URL = gprodnmlURL
	gprodnmcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("produce by name and manf get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 5 {
			h.Log.Debug("vars: ", vars)
			var manfName = vars["manufacturer"]
			var prodmName = vars["name"]
			var prodnmlstoreIDStr = vars["storeId"]
			var prodnmstartStr = vars["start"]
			var prodnmendStr = vars["end"]
			storeID, sprodnlerr := strconv.ParseInt(prodnmlstoreIDStr, 10, 64)
			prodnStart, prodnstarterr := strconv.ParseInt(prodnmstartStr, 10, 64)
			prodnEnd, prodnenderr := strconv.ParseInt(prodnmendStr, 10, 64)
			var gprodnmlres *[]sdbi.Product
			if sprodnlerr == nil && prodnstarterr == nil && prodnenderr == nil {
				gprodnmlres = h.Manager.GetProductByNameAndManufacturerName(manfName, prodmName, storeID, prodnStart, prodnEnd)
				h.Log.Debug("get product name list: ", gprodnmlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Product{}
				gprodnmlres = &nc
			}
			resJSON, _ := json.Marshal(gprodnmlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetProductManufacturerListByCatID godoc
// @Summary Get list of Manufacturers for a category
// @Description Get list of Manufacturers for for a category for a store
// @Tags Manufacturer
// @Accept  json
// @Produce  json
// @Param catId path string true "category ID"
// @Param storeId path string true "store storeId"
// @Param apiKey header string true "apiKey required"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} string
// @Router /rs/manufacturer/get/category/{catId}/{storeId} [get]
func (h *Six910Handler) GetProductManufacturerListByCatID(w http.ResponseWriter, r *http.Request) {
	var gprodmcnlURL = "/six910/rs/manufacturer/search/cat/list"
	var gprodmcncl jv.Claim
	gprodmcncl.Role = customerRole
	gprodmcncl.URL = gprodmcnlURL
	gprodmcncl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("produce manf list cat get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var catm = vars["catId"]
			var prodmcnlstoreIDStr = vars["storeId"]
			catID, mcerr := strconv.ParseInt(catm, 10, 64)
			storeID, sprodnlerr := strconv.ParseInt(prodmcnlstoreIDStr, 10, 64)
			var gprodmcnlres *[]string
			if sprodnlerr == nil && mcerr == nil {
				gprodmcnlres = h.Manager.GetProductManufacturerListByCatID(catID, storeID)
				h.Log.Debug("get product manf cat list: ", gprodmcnlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []string{}
				gprodmcnlres = &nc
			}
			resJSON, _ := json.Marshal(gprodmcnlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetProductByCatAndManufacturer godoc
// @Summary Get list of products by category and manufacturer
// @Description Get list of products by category and manufacturerfor a store
// @Tags Product
// @Accept  json
// @Produce  json
// @Param catId path string true "category ID"
// @Param manufacturer path string true "manufacturer name"
// @Param storeId path string true "store storeId"
// @Param start path string true "start index 0 based"
// @Param end path string true "end index"
// @Param apiKey header string true "apiKey required"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Product
// @Router /rs/product/get/category/manufacturer/{catId}/{manufacturer}/{storeId}/{start}/{end} [get]
func (h *Six910Handler) GetProductByCatAndManufacturer(w http.ResponseWriter, r *http.Request) {
	var gprodnmclURL = "/six910/rs/product/cat/manufacturer/list"
	var gprodnmccl jv.Claim
	gprodnmccl.Role = customerRole
	gprodnmccl.URL = gprodnmclURL
	gprodnmccl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("produce by cat and manf get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 5 {
			h.Log.Debug("vars: ", vars)
			var catmID = vars["catId"]
			var manfName = vars["manufacturer"]
			var prodnmclstoreIDStr = vars["storeId"]
			var prodnmcstartStr = vars["start"]
			var prodnmcendStr = vars["end"]
			catID, catmIDerr := strconv.ParseInt(catmID, 10, 64)
			storeID, sprodnlerr := strconv.ParseInt(prodnmclstoreIDStr, 10, 64)
			prodnStart, prodnstarterr := strconv.ParseInt(prodnmcstartStr, 10, 64)
			prodnEnd, prodnenderr := strconv.ParseInt(prodnmcendStr, 10, 64)
			var gprodnmclres *[]sdbi.Product
			if catmIDerr == nil && sprodnlerr == nil && prodnstarterr == nil && prodnenderr == nil {
				gprodnmclres = h.Manager.GetProductByCatAndManufacturer(catID, manfName, storeID, prodnStart, prodnEnd)
				h.Log.Debug("get product cat and manf list: ", gprodnmclres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Product{}
				gprodnmclres = &nc
			}
			resJSON, _ := json.Marshal(gprodnmclres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// ProductSearch godoc
// @Summary Get Product List
// @Description Get Product List for a store by attributes
// @Tags Product
// @Accept  json
// @Produce  json
// @Param product body six910-database-interface.Product true "ProductSearch"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/product/search [post]
func (h *Six910Handler) ProductSearch(w http.ResponseWriter, r *http.Request) {
	var srprodURL = "/six910/rs/product/search"
	var srprodc jv.Claim
	srprodc.Role = customerRole
	srprodc.URL = srprodURL
	srprodc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	//auth := h.processSecurity(r, &srprodc)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("product search authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var srprod sdbi.ProductSearch
			srprodsuc, srproderr := h.ProcessBody(r, &srprod)
			h.Log.Debug("aprodsuc: ", srprodsuc)
			h.Log.Debug("aprod: ", srprod)
			h.Log.Debug("aproderr: ", srproderr)
			if !srprodsuc && srproderr != nil {
				http.Error(w, srproderr.Error(), http.StatusBadRequest)
			} else {
				srprodres := h.Manager.ProductSearch(&srprod)
				h.Log.Debug("srprodres: ", *srprodres)
				//if srprodres.Success && aprodres.ID != 0 {
				w.WriteHeader(http.StatusOK)
				//} else {
				//w.WriteHeader(http.StatusInternalServerError)
				//}
				resJSON, _ := json.Marshal(srprodres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
