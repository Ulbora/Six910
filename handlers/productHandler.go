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

//ProdIDReq ProdIDReq
type ProdIDReq struct {
	StoreID      int64    `json:"storeId"`
	CategoryList *[]int64 `json:"categoryList"`
}

// AddProduct godoc
// @Summary Add a new product
// @Description Adds a new product to a store
// @Tags Product
// @Accept  json
// @Produce  json
// @Param product body six910-database-interface.Product true "product"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/product/add [post]
func (h *Six910Handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var addprodURL = "/six910/rs/product/add"
	var aprodc jv.Claim
	aprodc.Role = storeAdmin
	aprodc.URL = addprodURL
	aprodc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &aprodc)
	h.Log.Debug("product add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var aprod sdbi.Product
			aprodsuc, aproderr := h.ProcessBody(r, &aprod)
			h.Log.Debug("aprodsuc: ", aprodsuc)
			h.Log.Debug("aprod: ", aprod)
			h.Log.Debug("aproderr: ", aproderr)
			if !aprodsuc && aproderr != nil {
				http.Error(w, aproderr.Error(), http.StatusBadRequest)
			} else {
				aprodres := h.Manager.AddProduct(&aprod)
				h.Log.Debug("aprodres: ", *aprodres)
				if aprodres.Success && aprodres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aprodres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aprodfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aprodfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update product data
// @Tags Product
// @Accept  json
// @Produce  json
// @Param product body six910-database-interface.Product true "product"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/product/update [put]
func (h *Six910Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var upprodURL = "/six910/rs/product/update"
	var uprodc jv.Claim
	uprodc.Role = storeAdmin
	uprodc.URL = upprodURL
	uprodc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uprodc)
	h.Log.Debug("product update authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uprod sdbi.Product
			uprodsuc, uproderr := h.ProcessBody(r, &uprod)
			h.Log.Debug("uprodsuc: ", uprodsuc)
			h.Log.Debug("uprod: ", uprod)
			h.Log.Debug("uproderr: ", uproderr)
			if !uprodsuc && uproderr != nil {
				http.Error(w, uproderr.Error(), http.StatusBadRequest)
			} else {
				uprodres := h.Manager.UpdateProduct(&uprod)
				h.Log.Debug("uprodres: ", *uprodres)
				if uprodres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uprodres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uprodfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uprodfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateProductQuantity godoc
// @Summary Update a product Quantity
// @Description Update product data
// @Tags Product
// @Accept  json
// @Produce  json
// @Param product body six910-database-interface.Product true "product"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/product/update/quantity [put]
func (h *Six910Handler) UpdateProductQuantity(w http.ResponseWriter, r *http.Request) {
	var upqprodURL = "/six910/rs/product/update/quantity"
	var upqrodc jv.Claim
	upqrodc.Role = storeAdmin
	upqrodc.URL = upqprodURL
	upqrodc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("product update quantity authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uqprod sdbi.Product
			uqprodsuc, uqproderr := h.ProcessBody(r, &uqprod)
			h.Log.Debug("uqprodsuc: ", uqprodsuc)
			h.Log.Debug("uqprod: ", uqprod)
			h.Log.Debug("uqproderr: ", uqproderr)
			if !uqprodsuc && uqproderr != nil {
				http.Error(w, uqproderr.Error(), http.StatusBadRequest)
			} else {
				uqprodres := h.Manager.UpdateProductQuantity(&uqprod)
				h.Log.Debug("uqprodres: ", *uqprodres)
				if uqprodres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uqprodres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uqprodfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uqprodfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetProductByID godoc
// @Summary Get details of a product by id
// @Description Get details of a product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "product id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string true "apiKey required"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.Product
// @Router /rs/product/get/id/{id}/{storeId} [get]
func (h *Six910Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	var gprodURL = "/six910/rs/product/get"
	var gprodc jv.Claim
	gprodc.Role = customerRole
	gprodc.URL = gprodURL
	gprodc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("product get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gprodidStr = vars["id"]
			var gprodstoreIDStr = vars["storeId"]
			id, gprodiderr := strconv.ParseInt(gprodidStr, 10, 64)
			storeID, gprodsiderr := strconv.ParseInt(gprodstoreIDStr, 10, 64)
			var gprodres *sdbi.Product
			if gprodiderr == nil && gprodsiderr == nil {
				gprodres = h.Manager.GetProductByID(id, storeID)
				h.Log.Debug("gprodres: ", gprodres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Product
				gprodres = &nc
			}
			resJSON, _ := json.Marshal(gprodres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetProductBySku godoc
// @Summary Get details of a product by sku
// @Description Get details of a product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param sku path string true "product sku"
// @Param did path string true "product distributor id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string true "apiKey required"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.Product
// @Router /rs/product/get/sku/{sku}/{did}/{storeId} [get]
func (h *Six910Handler) GetProductBySku(w http.ResponseWriter, r *http.Request) {
	var gprodURL = "/six910/rs/product/get/sku"
	var gprodc jv.Claim
	gprodc.Role = customerRole
	gprodc.URL = gprodURL
	gprodc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("product get sku authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var sku = vars["sku"]
			var gproddidStr = vars["did"]
			var gprodstoreIDStr = vars["storeId"]
			did, gproddiderr := strconv.ParseInt(gproddidStr, 10, 64)
			storeID, gprodsiderr := strconv.ParseInt(gprodstoreIDStr, 10, 64)
			var gprodsres *sdbi.Product
			if gproddiderr == nil && gprodsiderr == nil {
				gprodsres = h.Manager.GetProductByBySku(sku, did, storeID)
				h.Log.Debug("gprodsres: ", gprodsres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Product
				gprodsres = &nc
			}
			resJSON, _ := json.Marshal(gprodsres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetProductsByPromoted godoc
// @Summary Get list of products by product name
// @Description Get list of products for a store
// @Tags Product
// @Accept  json
// @Produce  json
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
// @Router /rs/product/get/promoted/{storeId}/{start}/{end} [get]
func (h *Six910Handler) GetProductsByPromoted(w http.ResponseWriter, r *http.Request) {
	var gpprodnlURL = "/six910/rs/product/name/list/promoted"
	var gpprodncl jv.Claim
	gpprodncl.Role = customerRole
	gpprodncl.URL = gpprodnlURL
	gpprodncl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("producr promoted get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var pprodnlstoreIDStr = vars["storeId"]
			var pprodnstartStr = vars["start"]
			var pprodnendStr = vars["end"]
			storeID, spprodnlerr := strconv.ParseInt(pprodnlstoreIDStr, 10, 64)
			prodnStart, pprodnstarterr := strconv.ParseInt(pprodnstartStr, 10, 64)
			prodnEnd, pprodnenderr := strconv.ParseInt(pprodnendStr, 10, 64)
			var gpprodnlres *[]sdbi.Product
			if spprodnlerr == nil && pprodnstarterr == nil && pprodnenderr == nil {
				gpprodnlres = h.Manager.GetProductsByPromoted(storeID, prodnStart, prodnEnd)
				h.Log.Debug("get product prpomoted list: ", gpprodnlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Product{}
				gpprodnlres = &nc
			}
			resJSON, _ := json.Marshal(gpprodnlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetProductsByName godoc
// @Summary Get list of products by product name
// @Description Get list of products for a store
// @Tags Product
// @Accept  json
// @Produce  json
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
// @Router /rs/product/get/name/{name}/{storeId}/{start}/{end} [get]
func (h *Six910Handler) GetProductsByName(w http.ResponseWriter, r *http.Request) {
	var gprodnlURL = "/six910/rs/product/name/list"
	var gprodncl jv.Claim
	gprodncl.Role = customerRole
	gprodncl.URL = gprodnlURL
	gprodncl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("producr name get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 4 {
			h.Log.Debug("vars: ", vars)
			var prodName = vars["name"]
			var prodnlstoreIDStr = vars["storeId"]
			var prodnstartStr = vars["start"]
			var prodnendStr = vars["end"]
			storeID, sprodnlerr := strconv.ParseInt(prodnlstoreIDStr, 10, 64)
			prodnStart, prodnstarterr := strconv.ParseInt(prodnstartStr, 10, 64)
			prodnEnd, prodnenderr := strconv.ParseInt(prodnendStr, 10, 64)
			var gprodnlres *[]sdbi.Product
			if sprodnlerr == nil && prodnstarterr == nil && prodnenderr == nil {
				gprodnlres = h.Manager.GetProductsByName(prodName, storeID, prodnStart, prodnEnd)
				h.Log.Debug("get product name list: ", gprodnlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Product{}
				gprodnlres = &nc
			}
			resJSON, _ := json.Marshal(gprodnlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetProductsByCaterory godoc
// @Summary Get list of products by category
// @Description Get list of products for a store
// @Tags Product
// @Accept  json
// @Produce  json
// @Param catId path string true "category id"
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
// @Router /rs/product/get/category/{catId}/{storeId}/{start}/{end} [get]
func (h *Six910Handler) GetProductsByCaterory(w http.ResponseWriter, r *http.Request) {
	var gprodclURL = "/six910/rs/product/category/list"
	var gprodccl jv.Claim
	gprodccl.Role = customerRole
	gprodccl.URL = gprodclURL
	gprodccl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("producr cat get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 4 {
			h.Log.Debug("vars: ", vars)
			var prodCatIDStr = vars["catId"]
			var prodclstoreIDStr = vars["storeId"]
			var prodcstartStr = vars["start"]
			var prodcendStr = vars["end"]
			catID, prodclerr := strconv.ParseInt(prodCatIDStr, 10, 64)
			storeID, sprodclerr := strconv.ParseInt(prodclstoreIDStr, 10, 64)
			prodcStart, prodcstarterr := strconv.ParseInt(prodcstartStr, 10, 64)
			prodcEnd, prodcenderr := strconv.ParseInt(prodcendStr, 10, 64)
			var gprodclres *[]sdbi.Product
			if prodclerr == nil && sprodclerr == nil && prodcstarterr == nil && prodcenderr == nil {
				gprodclres = h.Manager.GetProductsByCaterory(catID, storeID, prodcStart, prodcEnd)
				h.Log.Debug("get product cat list: ", gprodclres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Product{}
				gprodclres = &nc
			}
			resJSON, _ := json.Marshal(gprodclres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetProductList godoc
// @Summary Get list of products
// @Description Get list of products for a store
// @Tags Product
// @Accept  json
// @Produce  json
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
// @Router /rs/product/get/list/{storeId}/{start}/{end} [get]
func (h *Six910Handler) GetProductList(w http.ResponseWriter, r *http.Request) {
	var gprodlURL = "/six910/rs/product/list"
	var gprodcl jv.Claim
	gprodcl.Role = customerRole
	gprodcl.URL = gprodlURL
	gprodcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("producr get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var prodlstoreIDStr = vars["storeId"]
			var prodstartStr = vars["start"]
			var prodendStr = vars["end"]
			storeID, sprodlerr := strconv.ParseInt(prodlstoreIDStr, 10, 64)
			prodStart, prodstarterr := strconv.ParseInt(prodstartStr, 10, 64)
			prodEnd, prodenderr := strconv.ParseInt(prodendStr, 10, 64)
			var gprodlres *[]sdbi.Product
			if sprodlerr == nil && prodstarterr == nil && prodenderr == nil {
				gprodlres = h.Manager.GetProductList(storeID, prodStart, prodEnd)
				h.Log.Debug("get product list: ", gprodlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Product{}
				gprodlres = &nc
			}
			resJSON, _ := json.Marshal(gprodlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetProductSubSkuList godoc
// @Summary Get list of products
// @Description Get list of products for a store
// @Tags Product
// @Accept  json
// @Produce  json
// @Param storeId path string true "store storeId"
// @Param parentProdID path string true "parent Product ID"
// @Param apiKey header string true "apiKey required"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Product
// @Router /rs/product/get/list/{storeId}/{start}/{end} [get]
func (h *Six910Handler) GetProductSubSkuList(w http.ResponseWriter, r *http.Request) {
	var gprodlURL = "/six910/rs/product/subskus"
	var gprodcl jv.Claim
	gprodcl.Role = customerRole
	gprodcl.URL = gprodlURL
	gprodcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("produce get subsku list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var prodlstoreIDStr = vars["storeId"]
			var parentProdIDStr = vars["parentProdID"]
			storeID, sprodlerr := strconv.ParseInt(prodlstoreIDStr, 10, 64)
			parentProdID, parentProdIDerr := strconv.ParseInt(parentProdIDStr, 10, 64)
			var gprodsslres *[]sdbi.Product
			if sprodlerr == nil && parentProdIDerr == nil {
				gprodsslres = h.Manager.GetProductSubSkuList(storeID, parentProdID)
				h.Log.Debug("get product subsku list: ", gprodsslres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Product{}
				gprodsslres = &nc
			}
			resJSON, _ := json.Marshal(gprodsslres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetProductIDList GetProductIDList
func (h *Six910Handler) GetProductIDList(w http.ResponseWriter, r *http.Request) {
	var gprodilURL = "/six910/rs/product/id/list"
	var gprodicl jv.Claim
	gprodicl.Role = customerRole
	gprodicl.URL = gprodilURL
	gprodicl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("produce get id list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var prodilstoreIDStr = vars["storeId"]
			storeID, sprodilerr := strconv.ParseInt(prodilstoreIDStr, 10, 64)
			var pidList *[]int64
			if sprodilerr == nil {
				pidList = h.Manager.GetProductIDList(storeID)
				h.Log.Debug("get product id list: ", pidList)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []int64{}
				pidList = &nc
			}
			resJSON, _ := json.Marshal(pidList)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetProductIDListByCategories GetProductIDListByCategories
func (h *Six910Handler) GetProductIDListByCategories(w http.ResponseWriter, r *http.Request) {
	var gprodiclURL = "/six910/rs/product/id/list/cat"
	var gprodiccl jv.Claim
	gprodiccl.Role = customerRole
	gprodiccl.URL = gprodiclURL
	gprodiccl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("produce get id list by cat authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var prodic ProdIDReq
			prodicsuc, prodicerr := h.ProcessBody(r, &prodic)
			h.Log.Debug("prodicsuc: ", prodicsuc)
			h.Log.Debug("prodic: ", prodic)
			h.Log.Debug("prodicerr: ", prodicerr)
			if !prodicsuc && prodicerr != nil {
				http.Error(w, prodicerr.Error(), http.StatusBadRequest)
			} else {
				prodicres := h.Manager.GetProductIDListByCategories(prodic.StoreID, prodic.CategoryList)
				h.Log.Debug("prodicres: ", *prodicres)
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(prodicres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteProduct godoc
// @Summary Delete a products
// @Description Delete a products from the store
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "products id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/product/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var dprodURL = "/six910/rs/product/delete"
	var dprodc jv.Claim
	dprodc.Role = storeAdmin
	dprodc.URL = dprodURL
	dprodc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dprodc)
	h.Log.Debug("product delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dprodidStr = vars["id"]
			var dprodstoreIDStr = vars["storeId"]
			id, dprodiderr := strconv.ParseInt(dprodidStr, 10, 64)
			storeID, dprodidserr := strconv.ParseInt(dprodstoreIDStr, 10, 64)
			var dprodres *m.Response
			if dprodiderr == nil && dprodidserr == nil {
				dprodres = h.Manager.DeleteProduct(id, storeID)
				h.Log.Debug("dprodres: ", *dprodres)
				if dprodres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dprodres = &nc
			}
			resJSON, _ := json.Marshal(dprodres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
