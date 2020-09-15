package managers

import (
	"time"

	sdbi "github.com/Ulbora/six910-database-interface"
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

//dbmod "github.com/Ulbora/six910-mysql"
//sdbi "github.com/Ulbora/six910-database-interface"

//LocalStoreAdminUser LocalStoreAdminUser
type LocalStoreAdminUser struct {
	Username string
	Password string
}

//LocalStoreResponse LocalStoreResponse
type LocalStoreResponse struct {
	Success bool   `json:"success"`
	StoreID int64  `json:"storeId"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

//User User
type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	OldPassword string `json:"oldPassword"`
	Role        string `json:"role"`
	CustomerID  int64  `json:"customerId"`
	StoreID     int64  `json:"storeId"`
	Enabled     bool   `json:"enabled"`
}

//ResponseID ResponseID
type ResponseID struct {
	ID      int64  `json:"id"`
	Success bool   `json:"success"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

//Response Response
type Response struct {
	Success bool   `json:"success"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

//UserResponse UserResponse
type UserResponse struct {
	Username   string `json:"username"`
	Role       string `json:"role"`
	CustomerID int64  `json:"customerId"`
	StoreID    int64  `json:"storeId"`
	Enabled    bool   `json:"enabled"`
}

//OAuthUser OAuthUser
type OAuthUser struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Enabled      bool   `json:"enabled"`
	EmailAddress string `json:"emailAddress"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	RoleID       int64  `json:"roleId"`
	ClientID     int64  `json:"clientId"`
}

//OAuthUserUser OAuthUserUser
type OAuthUserUser struct {
	Username  string    `json:"username"`
	Enabled   bool      `json:"enabled"`
	Entered   time.Time `json:"dateEntered"`
	Email     string    `json:"emailAddress"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	RoleID    int64     `json:"roleId"`
	ClientID  int64     `json:"clientId"`
}

//OauthUserList OauthUserList
type OauthUserList struct {
	Username  string `json:"username"`
	Enabled   bool   `json:"enabled"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ClientID  int64  `json:"clientId"`
}

//Auth Auth
type Auth struct {
	Token    string
	ClientID string
}

//SecurityProfile SecurityProfile
type SecurityProfile struct {
	IsOAuthOn bool
	Store     *sdbi.Store
}

//Manager Manager
type Manager interface {
	GetSecurityProfile(storeName string, localDomain string) *SecurityProfile

	CreateLocalStore(auth *LocalStoreAdminUser) *LocalStoreResponse

	// // Store --------------------------
	// //Can only add a store when in external oauth mode
	AddStore(s *sdbi.Store) *ResponseID

	UpdateStore(s *sdbi.Store) *Response
	GetStore(sname string, localDomain string) *sdbi.Store

	// //can only delete a store when in the external oauth mode
	DeleteStore(sname string, localDomain string) *Response

	// //customer
	AddCustomer(c *sdbi.Customer) *ResponseID
	UpdateCustomer(c *sdbi.Customer) *Response
	GetCustomer(email string, storeID int64) *sdbi.Customer
	GetCustomerID(id int64, storeID int64) *sdbi.Customer
	GetCustomerList(storeID int64) *[]sdbi.Customer
	DeleteCustomer(id int64, storeID int64) *Response

	// only for local single store installations
	AddAdminUser(u *User) *Response

	AddCustomerUser(u *User) *Response
	UpdateUser(u *User) *Response
	GetUser(u *User) *UserResponse
	GetAdminUsers(storeID int64) *[]UserResponse
	GetCustomerUsers(storeID int64) *[]UserResponse
	ValidateUser(u *User) *Response

	//oauth users
	AddOAuthUser(user *OAuthUser, auth *Auth) *Response
	UpdateOAuthUser(user *OAuthUser, auth *Auth) *Response
	GetOAuthUser(username string, clientID string, auth *Auth) (*OAuthUserUser, int)
	GetOAuthUserList(clientID string, auth *Auth) (*[]OAuthUser, int)
	DeleteOAuthUser(username string, clientID string, auth *Auth) *Response

	// //distributors
	AddDistributor(d *sdbi.Distributor) *ResponseID
	UpdateDistributor(d *sdbi.Distributor) *Response
	GetDistributor(id int64, storeID int64) *sdbi.Distributor
	GetDistributorList(storeID int64) *[]sdbi.Distributor
	DeleteDistributor(id int64, storeID int64) *Response

	// //Cart
	AddCart(c *sdbi.Cart) *ResponseID
	UpdateCart(c *sdbi.Cart) *Response
	GetCart(cid int64, storeID int64) *sdbi.Cart
	DeleteCart(id int64, cid int64, storeID int64) *Response

	// //cart item
	AddCartItem(ci *sdbi.CartItem, cid int64, sid int64) *ResponseID
	UpdateCartItem(ci *sdbi.CartItem, cid int64, sid int64) *Response
	GetCarItem(cid int64, prodID int64, sid int64) *sdbi.CartItem
	GetCartItemList(cartID int64, cid int64, sid int64) *[]sdbi.CartItem
	DeleteCartItem(id int64, prodID int64, cartID int64) *Response

	// //address
	AddAddress(a *sdbi.Address, sid int64) *ResponseID
	UpdateAddress(a *sdbi.Address, sid int64) *Response
	GetAddress(id int64, cid int64, sid int64) *sdbi.Address
	GetAddressList(cid int64, sid int64) *[]sdbi.Address
	DeleteAddress(id int64, cid int64, sid int64) *Response

	// //category
	AddCategory(c *sdbi.Category) *ResponseID
	UpdateCategory(c *sdbi.Category) *Response
	GetCategory(id int64, sid int64) *sdbi.Category
	GetHierarchicalCategoryList(storeID int64) *[]sdbi.Category
	GetCategoryList(storeID int64) *[]sdbi.Category
	GetSubCategoryList(catID int64) *[]sdbi.Category
	DeleteCategory(id int64, sid int64) *Response

	// //shipping method
	AddShippingMethod(s *sdbi.ShippingMethod) *ResponseID
	UpdateShippingMethod(s *sdbi.ShippingMethod) *Response
	GetShippingMethod(id int64, sid int64) *sdbi.ShippingMethod
	GetShippingMethodList(storeID int64) *[]sdbi.ShippingMethod
	DeleteShippingMethod(id int64, sid int64) *Response

	// //shipping insurance
	AddInsurance(i *sdbi.Insurance) *ResponseID
	UpdateInsurance(i *sdbi.Insurance) *Response
	GetInsurance(id int64, sid int64) *sdbi.Insurance
	GetInsuranceList(storeID int64) *[]sdbi.Insurance
	DeleteInsurance(id int64, sid int64) *Response

	//tax rate
	AddTaxRate(t *sdbi.TaxRate) *ResponseID
	UpdateTaxRate(t *sdbi.TaxRate) *Response
	GetTaxRate(country string, state string, sid int64) *[]sdbi.TaxRate
	GetTaxRateList(storeID int64) *[]sdbi.TaxRate
	DeleteTaxRate(id int64, sid int64) *Response

	// //product
	AddProduct(p *sdbi.Product) *ResponseID
	UpdateProduct(p *sdbi.Product) *Response
	GetProductByID(id int64, sid int64) *sdbi.Product
	GetProductByBySku(sku string, distributorID int64, sid int64) *sdbi.Product
	GetProductsByPromoted(sid int64, start int64, end int64) *[]sdbi.Product
	GetProductsByName(name string, sid int64, start int64, end int64) *[]sdbi.Product
	GetProductsByCaterory(catID int64, sid int64, start int64, end int64) *[]sdbi.Product
	GetProductList(storeID int64, start int64, end int64) *[]sdbi.Product
	DeleteProduct(id int64, sid int64) *Response

	// //Geographic Regions
	AddRegion(r *sdbi.Region) *ResponseID
	UpdateRegion(r *sdbi.Region) *Response
	GetRegion(id int64, sid int64) *sdbi.Region
	GetRegionList(storeID int64) *[]sdbi.Region
	DeleteRegion(id int64, sid int64) *Response

	// //Geographic Sub Regions
	AddSubRegion(s *sdbi.SubRegion, sid int64) *ResponseID
	UpdateSubRegion(s *sdbi.SubRegion, sid int64) *Response
	GetSubRegion(id int64, sid int64) *sdbi.SubRegion
	GetSubRegionList(regionID int64, sid int64) *[]sdbi.SubRegion
	DeleteSubRegion(id int64, sid int64) *Response

	// //excluded sub regions
	AddExcludedSubRegion(e *sdbi.ExcludedSubRegion, sid int64) *ResponseID
	UpdateExcludedSubRegion(e *sdbi.ExcludedSubRegion, sid int64) *Response
	GetExcludedSubRegion(id int64, sid int64) *sdbi.ExcludedSubRegion
	GetExcludedSubRegionList(regionID int64, sid int64) *[]sdbi.ExcludedSubRegion
	DeleteExcludedSubRegion(id int64, regionID int64, sid int64) *Response

	// //included sub regions
	AddIncludedSubRegion(e *sdbi.IncludedSubRegion, sid int64) *ResponseID
	UpdateIncludedSubRegion(e *sdbi.IncludedSubRegion, sid int64) *Response
	GetIncludedSubRegion(id int64, sid int64) *sdbi.IncludedSubRegion
	GetIncludedSubRegionList(regionID int64, sid int64) *[]sdbi.IncludedSubRegion
	DeleteIncludedSubRegion(id int64, regionID int64, sid int64) *Response

	// //limit exclusions and inclusions to a zip code
	AddZoneZip(z *sdbi.ZoneZip, sid int64) *ResponseID
	GetZoneZipListByExclusion(exID int64, sid int64) *[]sdbi.ZoneZip
	GetZoneZipListByInclusion(incID int64, sid int64) *[]sdbi.ZoneZip
	DeleteZoneZip(id int64, incID int64, exID int64, sid int64) *Response

	// //product category
	AddProductCategory(pc *sdbi.ProductCategory, sid int64) *Response
	GetProductCategoryList(productID int64) *[]int64
	DeleteProductCategory(pc *sdbi.ProductCategory, sid int64) *Response

	// //Orders
	AddOrder(o *sdbi.Order) *ResponseID
	UpdateOrder(o *sdbi.Order) *Response
	GetOrder(id int64, sid int64) *sdbi.Order
	GetOrderList(cid int64, sid int64) *[]sdbi.Order
	GetStoreOrderList(sid int64) *[]sdbi.Order
	GetStoreOrderListByStatus(status string, sid int64) *[]sdbi.Order
	DeleteOrder(id int64, sid int64) *Response

	// //Order Items
	AddOrderItem(i *sdbi.OrderItem, sid int64) *ResponseID
	UpdateOrderItem(i *sdbi.OrderItem, sid int64) *Response
	GetOrderItem(id int64, sid int64) *sdbi.OrderItem
	GetOrderItemList(orderID int64, sid int64) *[]sdbi.OrderItem
	DeleteOrderItem(id int64, sid int64) *Response

	// //Order Comments
	AddOrderComments(c *sdbi.OrderComment, sid int64) *ResponseID
	GetOrderCommentList(orderID int64, sid int64) *[]sdbi.OrderComment

	// //Order Payment Transactions
	AddOrderTransaction(t *sdbi.OrderTransaction, sid int64) *ResponseID
	GetOrderTransactionList(orderID int64, sid int64) *[]sdbi.OrderTransaction

	// //shipment
	AddShipment(s *sdbi.Shipment, sid int64) *ResponseID
	UpdateShipment(s *sdbi.Shipment, sid int64) *Response
	GetShipment(id int64, sid int64) *sdbi.Shipment
	GetShipmentList(orderID int64, sid int64) *[]sdbi.Shipment
	DeleteShipment(id int64, sid int64) *Response

	// //shipment boxes
	AddShipmentBox(sb *sdbi.ShipmentBox, sid int64) *ResponseID
	UpdateShipmentBox(sb *sdbi.ShipmentBox, sid int64) *Response
	GetShipmentBox(id int64, sid int64) *sdbi.ShipmentBox
	GetShipmentBoxList(shipmentID int64, sid int64) *[]sdbi.ShipmentBox
	DeleteShipmentBox(id int64, sid int64) *Response

	// //Shipment Items in box
	AddShipmentItem(si *sdbi.ShipmentItem, sid int64) *ResponseID
	UpdateShipmentItem(si *sdbi.ShipmentItem, sid int64) *Response
	GetShipmentItem(id int64, sid int64) *sdbi.ShipmentItem
	GetShipmentItemList(shipmentID int64, sid int64) *[]sdbi.ShipmentItem
	GetShipmentItemListByBox(boxNumber int64, shipmentID int64, sid int64) *[]sdbi.ShipmentItem
	DeleteShipmentItem(id int64, sid int64) *Response

	// //Global Plugins
	AddPlugin(p *sdbi.Plugins) *ResponseID
	UpdatePlugin(p *sdbi.Plugins) *Response
	GetPlugin(id int64) *sdbi.Plugins
	GetPluginList(start int64, end int64) *[]sdbi.Plugins
	DeletePlugin(id int64) *Response

	// //store plugins installed
	AddStorePlugin(sp *sdbi.StorePlugins) *ResponseID
	UpdateStorePlugin(sp *sdbi.StorePlugins) *Response
	GetStorePlugin(id int64, sid int64) *sdbi.StorePlugins
	GetStorePluginList(storeID int64) *[]sdbi.StorePlugins
	DeleteStorePlugin(id int64, sid int64) *Response

	// //Plugins that are payment gateways
	AddPaymentGateway(pgw *sdbi.PaymentGateway, sid int64) *ResponseID
	UpdatePaymentGateway(pgw *sdbi.PaymentGateway, sid int64) *Response
	GetPaymentGateway(id int64, sid int64) *sdbi.PaymentGateway
	GetPaymentGateways(storeID int64) *[]sdbi.PaymentGateway
	DeletePaymentGateway(id int64, sid int64) *Response

	// //store shipment carrier like UPS and FEDex
	AddShippingCarrier(c *sdbi.ShippingCarrier) *ResponseID
	UpdateShippingCarrier(c *sdbi.ShippingCarrier) *Response
	GetShippingCarrier(id int64, sid int64) *sdbi.ShippingCarrier
	GetShippingCarrierList(storeID int64) *[]sdbi.ShippingCarrier
	DeleteShippingCarrier(id int64, sid int64) *Response

	// //----UI Cluster installation: this is only called if UI is running in a cluster---
	// //Handle the situation where clients are running in a cluster
	// //This gives a way to make sure the json datastores are update on each node in the cluster

	// //----------------start datastore------------------------------------
	// //this gets called when a node start up and add only if it doesn't already exist
	AddLocalDatastore(d *sdbi.LocalDataStore) *Response

	// //This get get called when a change is made to a datastore from a node in the cluster
	// //Or after all reloads have completed and then get set to Reload = false
	UpdateLocalDatastore(d *sdbi.LocalDataStore) *Response

	// //This gets call by cluster nodes to see if there are pending reload
	GetLocalDatastore(storeID int64, dataStoreName string) *sdbi.LocalDataStore

	// //---------------------start instance--------------------
	// // this gets called when each instance is started and added only if never added before
	// //The instance name is pulled from Docker or an manually entered env variable
	AddInstance(i *sdbi.Instances) *Response

	// //This gets called after instance gets reloaded
	UpdateInstance(i *sdbi.Instances) *Response

	// //Gets called before updating an instance
	GetInstance(name string, dataStoreName string, storeID int64) *sdbi.Instances

	// //Gets called before or after updating an instance and allows a single instance
	// to clear the datastore reload flag if all instances have reloaded
	// this give any single instance visibility into other instances
	GetInstanceList(dataStoreName string, storeID int64) *[]sdbi.Instances

	// //-------------------start write lock-------------
	// //gets called after UI makes changes to one of the datastores
	// //If the datastore already exists, the Update method is called from within add
	AddDataStoreWriteLock(w *sdbi.DataStoreWriteLock) *Response
	UpdateDataStoreWriteLock(w *sdbi.DataStoreWriteLock) *Response

	// //gets called from within the add method and by any node trying to update a datastore
	GetDataStoreWriteLock(dataStore string, storeID int64) *sdbi.DataStoreWriteLock
}
