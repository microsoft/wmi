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

// CIM_BIOSLoadedInNV struct
type CIM_BIOSLoadedInNV struct {
	*CIM_Dependency

	//
	EndingAddress uint64

	//
	StartingAddress uint64
}

func NewCIM_BIOSLoadedInNVEx1(instance *cim.WmiInstance) (newInstance *CIM_BIOSLoadedInNV, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_BIOSLoadedInNV{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_BIOSLoadedInNVEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_BIOSLoadedInNV, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_BIOSLoadedInNV{
		CIM_Dependency: tmp,
	}
	return
}

// SetEndingAddress sets the value of EndingAddress for the instance
func (instance *CIM_BIOSLoadedInNV) SetPropertyEndingAddress(value uint64) (err error) {
	return instance.SetProperty("EndingAddress", value)
}

// GetEndingAddress gets the value of EndingAddress for the instance
func (instance *CIM_BIOSLoadedInNV) GetPropertyEndingAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("EndingAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartingAddress sets the value of StartingAddress for the instance
func (instance *CIM_BIOSLoadedInNV) SetPropertyStartingAddress(value uint64) (err error) {
	return instance.SetProperty("StartingAddress", value)
}

// GetStartingAddress gets the value of StartingAddress for the instance
func (instance *CIM_BIOSLoadedInNV) GetPropertyStartingAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("StartingAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
