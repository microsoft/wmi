// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_CollectionStatistics struct
type Win32_CollectionStatistics struct {
	cim.WmiInstance

	//
	Collection CIM_CollectionOfMSEs

	//
	Stats CIM_StatisticalInformation
}

// SetCollection sets the value of Collection for the instance
func (instance *Win32_CollectionStatistics) SetPropertyCollection(value CIM_CollectionOfMSEs) (err error) {
	return instance.SetProperty("Collection", value)
}

// GetCollection gets the value of Collection for the instance
func (instance *Win32_CollectionStatistics) GetPropertyCollection() (value CIM_CollectionOfMSEs, err error) {
	retValue, err := instance.GetProperty("Collection")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_CollectionOfMSEs)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStats sets the value of Stats for the instance
func (instance *Win32_CollectionStatistics) SetPropertyStats(value CIM_StatisticalInformation) (err error) {
	return instance.SetProperty("Stats", value)
}

// GetStats gets the value of Stats for the instance
func (instance *Win32_CollectionStatistics) GetPropertyStats() (value CIM_StatisticalInformation, err error) {
	retValue, err := instance.GetProperty("Stats")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_StatisticalInformation)
	if !ok {
		// TODO: Set an error
	}
	return
}
