// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_BIOSElement struct
type CIM_BIOSElement struct {
	*CIM_SoftwareElement

	//
	PrimaryBIOS bool
}

func NewCIM_BIOSElementEx1(instance *cim.WmiInstance) (newInstance *CIM_BIOSElement, err error) {
	tmp, err := NewCIM_SoftwareElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_BIOSElement{
		CIM_SoftwareElement: tmp,
	}
	return
}

func NewCIM_BIOSElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_BIOSElement, err error) {
	tmp, err := NewCIM_SoftwareElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_BIOSElement{
		CIM_SoftwareElement: tmp,
	}
	return
}

// SetPrimaryBIOS sets the value of PrimaryBIOS for the instance
func (instance *CIM_BIOSElement) SetPropertyPrimaryBIOS(value bool) (err error) {
	return instance.SetProperty("PrimaryBIOS", value)
}

// GetPrimaryBIOS gets the value of PrimaryBIOS for the instance
func (instance *CIM_BIOSElement) GetPropertyPrimaryBIOS() (value bool, err error) {
	retValue, err := instance.GetProperty("PrimaryBIOS")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
