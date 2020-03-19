// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ProductSupport struct
type CIM_ProductSupport struct {
	*cim.WmiInstance

	//
	Product CIM_Product

	//
	Support CIM_SupportAccess
}

func NewCIM_ProductSupportEx1(instance *cim.WmiInstance) (newInstance *CIM_ProductSupport, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_ProductSupport{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_ProductSupportEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ProductSupport, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ProductSupport{
		WmiInstance: tmp,
	}
	return
}

// SetProduct sets the value of Product for the instance
func (instance *CIM_ProductSupport) SetPropertyProduct(value CIM_Product) (err error) {
	return instance.SetProperty("Product", value)
}

// GetProduct gets the value of Product for the instance
func (instance *CIM_ProductSupport) GetPropertyProduct() (value CIM_Product, err error) {
	retValue, err := instance.GetProperty("Product")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Product)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupport sets the value of Support for the instance
func (instance *CIM_ProductSupport) SetPropertySupport(value CIM_SupportAccess) (err error) {
	return instance.SetProperty("Support", value)
}

// GetSupport gets the value of Support for the instance
func (instance *CIM_ProductSupport) GetPropertySupport() (value CIM_SupportAccess, err error) {
	retValue, err := instance.GetProperty("Support")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_SupportAccess)
	if !ok {
		// TODO: Set an error
	}
	return
}
