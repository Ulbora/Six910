package handlers

import "net/http"

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

const (
	validationServiceLocal = "http://localhost:3000/rs/token/validate"

	//roles
	superAdmin   = "superAdmin"
	storeAdmin   = "StoreAdmin"
	customerRole = "customer"

	//messages
	storeAddMessage    = "Store can not be added."
	storeDeleteMessage = "Store can not be deleted."
)

//Handler Handler
type Handler interface {
	//////security
	//////processSecurity(r *http.Request, c *jv.Claim) bool

	//store
	AddStore(w http.ResponseWriter, r *http.Request)
	UpdateStore(w http.ResponseWriter, r *http.Request)
	GetStore(w http.ResponseWriter, r *http.Request)
	DeleteStore(w http.ResponseWriter, r *http.Request)

	//customer
	AddCustomer(w http.ResponseWriter, r *http.Request)
	UpdateCustomer(w http.ResponseWriter, r *http.Request)
	GetCustomer(w http.ResponseWriter, r *http.Request)
	GetCustomerID(w http.ResponseWriter, r *http.Request)
	GetCustomerList(w http.ResponseWriter, r *http.Request)
	DeleteCustomer(w http.ResponseWriter, r *http.Request)

	//users
	AddUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)

	//only for non oauth stores
	GetAdminUserList(w http.ResponseWriter, r *http.Request)
	GetCustomerUserList(w http.ResponseWriter, r *http.Request)
	// DeleteUser(w http.ResponseWriter, r *http.Request)

	//distributors
	AddDistributor(w http.ResponseWriter, r *http.Request)
	UpdateDistributor(w http.ResponseWriter, r *http.Request)
	GetDistributor(w http.ResponseWriter, r *http.Request)
	GetDistributorList(w http.ResponseWriter, r *http.Request)
	DeleteDistributor(w http.ResponseWriter, r *http.Request)

	//Cart
	AddCart(w http.ResponseWriter, r *http.Request)
	UpdateCart(w http.ResponseWriter, r *http.Request)
	GetCart(w http.ResponseWriter, r *http.Request)
	DeleteCart(w http.ResponseWriter, r *http.Request)

	//cart item
	AddCartItem(w http.ResponseWriter, r *http.Request)
	UpdateCartItem(w http.ResponseWriter, r *http.Request)
	GetCartItem(w http.ResponseWriter, r *http.Request)
	GetCartItemList(w http.ResponseWriter, r *http.Request)
	DeleteCartItem(w http.ResponseWriter, r *http.Request)

	//address
	AddAddress(w http.ResponseWriter, r *http.Request)
	UpdateAddress(w http.ResponseWriter, r *http.Request)
	GetAddress(w http.ResponseWriter, r *http.Request)
	GetAddressList(w http.ResponseWriter, r *http.Request)
	DeleteAddress(w http.ResponseWriter, r *http.Request)

	//category
	AddCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	GetCategory(w http.ResponseWriter, r *http.Request)
	GetCategoryList(w http.ResponseWriter, r *http.Request)
	GetSubCategoryList(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)

	//shipping method
	AddShippingMethod(w http.ResponseWriter, r *http.Request)
	UpdateShippingMethod(w http.ResponseWriter, r *http.Request)
	GetShippingMethod(w http.ResponseWriter, r *http.Request)
	GetShippingMethodList(w http.ResponseWriter, r *http.Request)
	DeleteShippingMethod(w http.ResponseWriter, r *http.Request)

	//shipping insurance
	AddInsurance(w http.ResponseWriter, r *http.Request)
	UpdateInsurance(w http.ResponseWriter, r *http.Request)
	GetInsurance(w http.ResponseWriter, r *http.Request)
	GetInsuranceList(w http.ResponseWriter, r *http.Request)
	DeleteInsurance(w http.ResponseWriter, r *http.Request)

	//product
	AddProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	GetProductByID(w http.ResponseWriter, r *http.Request)
	GetProductBySku(w http.ResponseWriter, r *http.Request)
	GetProductsByName(w http.ResponseWriter, r *http.Request)
	GetProductsByPromoted(w http.ResponseWriter, r *http.Request)
	GetProductsByCaterory(w http.ResponseWriter, r *http.Request)
	GetProductList(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)

	//Geographic Regions
	AddRegion(w http.ResponseWriter, r *http.Request)
	UpdateRegion(w http.ResponseWriter, r *http.Request)
	GetRegion(w http.ResponseWriter, r *http.Request)
	GetRegionList(w http.ResponseWriter, r *http.Request)
	DeleteRegion(w http.ResponseWriter, r *http.Request)

	//Geographic Sub Regions
	AddSubRegion(w http.ResponseWriter, r *http.Request)
	UpdateSubRegion(w http.ResponseWriter, r *http.Request)
	GetSubRegion(w http.ResponseWriter, r *http.Request)
	GetSubRegionList(w http.ResponseWriter, r *http.Request)
	DeleteSubRegion(w http.ResponseWriter, r *http.Request)

	//excluded sub regions
	AddExcludedSubRegion(w http.ResponseWriter, r *http.Request)

	//-----currently not used
	UpdateExcludedSubRegion(w http.ResponseWriter, r *http.Request)
	//-----currently not used
	GetExcludedSubRegion(w http.ResponseWriter, r *http.Request)

	GetExcludedSubRegionList(w http.ResponseWriter, r *http.Request)
	DeleteExcludedSubRegion(w http.ResponseWriter, r *http.Request)

	//included sub regions
	AddIncludedSubRegion(w http.ResponseWriter, r *http.Request)

	//-----currently not used
	UpdateIncludedSubRegion(w http.ResponseWriter, r *http.Request)
	//-----currently not used
	GetIncludedSubRegion(w http.ResponseWriter, r *http.Request)

	GetIncludedSubRegionList(w http.ResponseWriter, r *http.Request)
	DeleteIncludedSubRegion(w http.ResponseWriter, r *http.Request)

	//limit exclusions and inclusions to a zip code
	AddZoneZip(w http.ResponseWriter, r *http.Request)
	GetZoneZipListByExclusion(w http.ResponseWriter, r *http.Request)
	GetZoneZipListByInclusion(w http.ResponseWriter, r *http.Request)
	DeleteZoneZip(w http.ResponseWriter, r *http.Request)

	//product category
	AddProductCategory(w http.ResponseWriter, r *http.Request)
	DeleteProductCategory(w http.ResponseWriter, r *http.Request)

	//Orders
	AddOrder(w http.ResponseWriter, r *http.Request)
	UpdateOrder(w http.ResponseWriter, r *http.Request)
	GetOrder(w http.ResponseWriter, r *http.Request)
	GetOrderList(w http.ResponseWriter, r *http.Request)
	GetStoreOrderList(w http.ResponseWriter, r *http.Request)
	GetStoreOrderListByStatus(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)

	//Order Items
	AddOrderItem(w http.ResponseWriter, r *http.Request)
	UpdateOrderItem(w http.ResponseWriter, r *http.Request)
	GetOrderItem(w http.ResponseWriter, r *http.Request)
	GetOrderItemList(w http.ResponseWriter, r *http.Request)
	DeleteOrderItem(w http.ResponseWriter, r *http.Request)

	//Order Comments
	AddOrderComments(w http.ResponseWriter, r *http.Request)
	GetOrderCommentList(w http.ResponseWriter, r *http.Request)

	//Order Payment Transactions
	AddOrderTransaction(w http.ResponseWriter, r *http.Request)
	GetOrderTransactionList(w http.ResponseWriter, r *http.Request)

	//shipment
	AddShipment(w http.ResponseWriter, r *http.Request)
	UpdateShipment(w http.ResponseWriter, r *http.Request)
	GetShipment(w http.ResponseWriter, r *http.Request)
	GetShipmentList(w http.ResponseWriter, r *http.Request)
	DeleteShipment(w http.ResponseWriter, r *http.Request)

	//shipment boxes
	AddShipmentBox(w http.ResponseWriter, r *http.Request)
	UpdateShipmentBox(w http.ResponseWriter, r *http.Request)
	GetShipmentBox(w http.ResponseWriter, r *http.Request)
	GetShipmentBoxList(w http.ResponseWriter, r *http.Request)
	DeleteShipmentBox(w http.ResponseWriter, r *http.Request)

	//Shipment Items in box
	AddShipmentItem(w http.ResponseWriter, r *http.Request)
	UpdateShipmentItem(w http.ResponseWriter, r *http.Request)
	GetShipmentItem(w http.ResponseWriter, r *http.Request)
	GetShipmentItemList(w http.ResponseWriter, r *http.Request)
	GetShipmentItemListByBox(w http.ResponseWriter, r *http.Request)
	DeleteShipmentItem(w http.ResponseWriter, r *http.Request)

	//Global Plugins
	AddPlugin(w http.ResponseWriter, r *http.Request)
	UpdatePlugin(w http.ResponseWriter, r *http.Request)
	GetPlugin(w http.ResponseWriter, r *http.Request)
	GetPluginList(w http.ResponseWriter, r *http.Request)
	DeletePlugin(w http.ResponseWriter, r *http.Request)

	//store plugins installed
	AddStorePlugin(w http.ResponseWriter, r *http.Request)
	UpdateStorePlugin(w http.ResponseWriter, r *http.Request)
	GetStorePlugin(w http.ResponseWriter, r *http.Request)
	GetStorePluginList(w http.ResponseWriter, r *http.Request)
	DeleteStorePlugin(w http.ResponseWriter, r *http.Request)

	//Plugins that are payment gateways
	AddPaymentGateway(w http.ResponseWriter, r *http.Request)
	UpdatePaymentGateway(w http.ResponseWriter, r *http.Request)
	GetPaymentGateway(w http.ResponseWriter, r *http.Request)
	GetPaymentGateways(w http.ResponseWriter, r *http.Request)
	DeletePaymentGateway(w http.ResponseWriter, r *http.Request)

	//store shipment carrier like UPS and FEDex
	AddShippingCarrier(w http.ResponseWriter, r *http.Request)
	UpdateShippingCarrier(w http.ResponseWriter, r *http.Request)
	GetShippingCarrier(w http.ResponseWriter, r *http.Request)
	GetShippingCarrierList(w http.ResponseWriter, r *http.Request)
	DeleteShippingCarrier(w http.ResponseWriter, r *http.Request)

	//----UI Cluster installation: this is only called if UI is running in a cluster---
	//Handles the situation where clients are running in a cluster
	//This gives a way to make sure the json datastores are update on each node in the cluster

	//----------------start datastore------------------------------------
	//this gets called when a node starts up and add only if it doesn't already exist
	AddLocalDatastore(w http.ResponseWriter, r *http.Request)

	//This gets called when a change is made to a datastore from a node in the cluster
	//Or after all reloads have completed and then gets set to Reload = false
	UpdateLocalDatastore(w http.ResponseWriter, r *http.Request)

	//This gets call by cluster nodes to see if there are pending reload
	GetLocalDatastore(w http.ResponseWriter, r *http.Request)

	//---------------------start instance--------------------
	// this gets called when each instance is started and added only if never added before
	//The instance name is pulled from Docker or an manually entered env variable
	AddInstance(w http.ResponseWriter, r *http.Request)

	//This gets called after instance gets reloaded
	UpdateInstance(w http.ResponseWriter, r *http.Request)

	//Gets called before updating an instance reloaded
	GetInstance(w http.ResponseWriter, r *http.Request)

	//Gets called before updating or after an instance is reloaded
	GetInstanceList(w http.ResponseWriter, r *http.Request)

	//-------------------start write lock-------------
	//gets called after UI makes changes to one of the datastores
	//If the datastore already exists, the Update method is called from within add
	AddDataStoreWriteLock(w http.ResponseWriter, r *http.Request)
	UpdateDataStoreWriteLock(w http.ResponseWriter, r *http.Request)

	//gets called from within the add method and by any node trying to update a datastore
	GetDataStoreWriteLock(w http.ResponseWriter, r *http.Request)
}
