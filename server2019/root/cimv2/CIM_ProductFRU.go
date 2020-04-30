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

// CIM_ProductFRU struct
type CIM_ProductFRU struct {
	*cim.WmiInstance

	//
	FRU CIM_FRU

	//
	Product CIM_Product
}

func NewCIM_ProductFRUEx1(instance *cim.WmiInstance) (newInstance *CIM_ProductFRU, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_ProductFRU{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_ProductFRUEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ProductFRU, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ProductFRU{
		WmiInstance: tmp,
	}
	return
}

// SetFRU sets the value of FRU for the instance
func (instance *CIM_ProductFRU) SetPropertyFRU(value CIM_FRU) (err error) {
	return instance.SetProperty("FRU", value)
}

// GetFRU gets the value of FRU for the instance
func (instance *CIM_ProductFRU) GetPropertyFRU() (value CIM_FRU, err error) {
	retValue, err := instance.GetProperty("FRU")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_FRU)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProduct sets the value of Product for the instance
func (instance *CIM_ProductFRU) SetPropertyProduct(value CIM_Product) (err error) {
	return instance.SetProperty("Product", value)
}

// GetProduct gets the value of Product for the instance
func (instance *CIM_ProductFRU) GetPropertyProduct() (value CIM_Product, err error) {
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