package managers

import (
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

//GetProductManufacturerListByProductName GetProductManufacturerListByProductName
func (m *Six910Manager) GetProductManufacturerListByProductName(name string, storeID int64) *[]string {
	return m.Db.GetProductManufacturerListByProductName(name, storeID)
}

//GetProductByNameAndManufacturerName GetProductByNameAndManufacturerName
func (m *Six910Manager) GetProductByNameAndManufacturerName(manf string, name string, storeID int64,
	start int64, end int64) *[]sdbi.Product {
	return m.Db.GetProductByNameAndManufacturerName(manf, name, storeID, start, end)
}

//GetProductManufacturerListByCatID GetProductManufacturerListByCatID
func (m *Six910Manager) GetProductManufacturerListByCatID(catID int64, storeID int64) *[]string {
	return m.Db.GetProductManufacturerListByCatID(catID, storeID)
}

//GetProductByCatAndManufacturer GetProductByCatAndManufacturer
func (m *Six910Manager) GetProductByCatAndManufacturer(catID int64, manf string, storeID int64,
	start int64, end int64) *[]sdbi.Product {
	return m.Db.GetProductByCatAndManufacturer(catID, manf, storeID, start, end)
}
