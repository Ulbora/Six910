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

//AddProduct AddProduct
func (h *Six910Handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var addprodURL = "/six910/rs/product/add"
	var aprodc jv.Claim
	aprodc.Role = storeAdmin
	aprodc.URL = addprodURL
	aprodc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &aprodc)
	h.Log.Debug("product add authorized: ", auth)
	if auth {
		h.SetContentType(w)
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

//UpdateProduct UpdateProduct
func (h *Six910Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var upprodURL = "/six910/rs/product/update"
	var uprodc jv.Claim
	uprodc.Role = storeAdmin
	uprodc.URL = upprodURL
	uprodc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uprodc)
	h.Log.Debug("product update authorized: ", auth)
	if auth {
		h.SetContentType(w)
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
		var uprodfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uprodfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetProductByID GetProductByID
func (h *Six910Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	var gprodURL = "/six910/rs/product/get"
	var gprodc jv.Claim
	gprodc.Role = customerRole
	gprodc.URL = gprodURL
	gprodc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gprodc)
	h.Log.Debug("product get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
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

//GetProductsByName GetProductsByName
func (h *Six910Handler) GetProductsByName(w http.ResponseWriter, r *http.Request) {
	var gprodnlURL = "/six910/rs/product/name/list"
	var gprodncl jv.Claim
	gprodncl.Role = customerRole
	gprodncl.URL = gprodnlURL
	gprodncl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gprodncl)
	h.Log.Debug("producr name get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
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

//GetProductsByCaterory GetProductsByCaterory
func (h *Six910Handler) GetProductsByCaterory(w http.ResponseWriter, r *http.Request) {
	var gprodclURL = "/six910/rs/product/category/list"
	var gprodccl jv.Claim
	gprodccl.Role = customerRole
	gprodccl.URL = gprodclURL
	gprodccl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gprodccl)
	h.Log.Debug("producr cat get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
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

//GetProductList GetProductList
func (h *Six910Handler) GetProductList(w http.ResponseWriter, r *http.Request) {
	var gprodlURL = "/six910/rs/product/list"
	var gprodcl jv.Claim
	gprodcl.Role = customerRole
	gprodcl.URL = gprodlURL
	gprodcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gprodcl)
	h.Log.Debug("producr get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
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

//DeleteProduct DeleteProduct
func (h *Six910Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var dprodURL = "/six910/rs/product/delete"
	var dprodc jv.Claim
	dprodc.Role = storeAdmin
	dprodc.URL = dprodURL
	dprodc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dprodc)
	h.Log.Debug("product delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
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
