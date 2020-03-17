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

// CIM_RelatedStatistics struct
type CIM_RelatedStatistics struct {
	cim.WmiInstance

	//
	RelatedStats CIM_StatisticalInformation

	//
	Stats CIM_StatisticalInformation
}

// SetRelatedStats sets the value of RelatedStats for the instance
func (instance *CIM_RelatedStatistics) SetPropertyRelatedStats(value CIM_StatisticalInformation) (err error) {
	return instance.SetProperty("RelatedStats", value)
}

// GetRelatedStats gets the value of RelatedStats for the instance
func (instance *CIM_RelatedStatistics) GetPropertyRelatedStats() (value CIM_StatisticalInformation, err error) {
	retValue, err := instance.GetProperty("RelatedStats")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_StatisticalInformation)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStats sets the value of Stats for the instance
func (instance *CIM_RelatedStatistics) SetPropertyStats(value CIM_StatisticalInformation) (err error) {
	return instance.SetProperty("Stats", value)
}

// GetStats gets the value of Stats for the instance
func (instance *CIM_RelatedStatistics) GetPropertyStats() (value CIM_StatisticalInformation, err error) {
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
