// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_IBPortCounters struct
type MLNX_IBPortCounters struct {
	*CIM_IBPortStatistics

	//
	SystemName string
}

func NewMLNX_IBPortCountersEx1(instance *cim.WmiInstance) (newInstance *MLNX_IBPortCounters, err error) {
	tmp, err := NewCIM_IBPortStatisticsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_IBPortCounters{
		CIM_IBPortStatistics: tmp,
	}
	return
}

func NewMLNX_IBPortCountersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_IBPortCounters, err error) {
	tmp, err := NewCIM_IBPortStatisticsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_IBPortCounters{
		CIM_IBPortStatistics: tmp,
	}
	return
}

// SetSystemName sets the value of SystemName for the instance
func (instance *MLNX_IBPortCounters) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", value)
}

// GetSystemName gets the value of SystemName for the instance
func (instance *MLNX_IBPortCounters) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
