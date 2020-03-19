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

// CIM_RealizesDiskPartition struct
type CIM_RealizesDiskPartition struct {
	*CIM_Realizes

	//
	StartingAddress uint64
}

func NewCIM_RealizesDiskPartitionEx1(instance *cim.WmiInstance) (newInstance *CIM_RealizesDiskPartition, err error) {
	tmp, err := NewCIM_RealizesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_RealizesDiskPartition{
		CIM_Realizes: tmp,
	}
	return
}

func NewCIM_RealizesDiskPartitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_RealizesDiskPartition, err error) {
	tmp, err := NewCIM_RealizesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_RealizesDiskPartition{
		CIM_Realizes: tmp,
	}
	return
}

// SetStartingAddress sets the value of StartingAddress for the instance
func (instance *CIM_RealizesDiskPartition) SetPropertyStartingAddress(value uint64) (err error) {
	return instance.SetProperty("StartingAddress", value)
}

// GetStartingAddress gets the value of StartingAddress for the instance
func (instance *CIM_RealizesDiskPartition) GetPropertyStartingAddress() (value uint64, err error) {
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
