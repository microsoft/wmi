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

// CIM_SwapSpaceCheck struct
type CIM_SwapSpaceCheck struct {
	*CIM_Check

	//
	SwapSpaceSize uint64
}

func NewCIM_SwapSpaceCheckEx1(instance *cim.WmiInstance) (newInstance *CIM_SwapSpaceCheck, err error) {
	tmp, err := NewCIM_CheckEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SwapSpaceCheck{
		CIM_Check: tmp,
	}
	return
}

func NewCIM_SwapSpaceCheckEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SwapSpaceCheck, err error) {
	tmp, err := NewCIM_CheckEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SwapSpaceCheck{
		CIM_Check: tmp,
	}
	return
}

// SetSwapSpaceSize sets the value of SwapSpaceSize for the instance
func (instance *CIM_SwapSpaceCheck) SetPropertySwapSpaceSize(value uint64) (err error) {
	return instance.SetProperty("SwapSpaceSize", value)
}

// GetSwapSpaceSize gets the value of SwapSpaceSize for the instance
func (instance *CIM_SwapSpaceCheck) GetPropertySwapSpaceSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("SwapSpaceSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
