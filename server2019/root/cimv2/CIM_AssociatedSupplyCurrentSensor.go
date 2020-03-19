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

// CIM_AssociatedSupplyCurrentSensor struct
type CIM_AssociatedSupplyCurrentSensor struct {
	*CIM_AssociatedSensor

	//
	MonitoringRange uint16
}

func NewCIM_AssociatedSupplyCurrentSensorEx1(instance *cim.WmiInstance) (newInstance *CIM_AssociatedSupplyCurrentSensor, err error) {
	tmp, err := NewCIM_AssociatedSensorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_AssociatedSupplyCurrentSensor{
		CIM_AssociatedSensor: tmp,
	}
	return
}

func NewCIM_AssociatedSupplyCurrentSensorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_AssociatedSupplyCurrentSensor, err error) {
	tmp, err := NewCIM_AssociatedSensorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_AssociatedSupplyCurrentSensor{
		CIM_AssociatedSensor: tmp,
	}
	return
}

// SetMonitoringRange sets the value of MonitoringRange for the instance
func (instance *CIM_AssociatedSupplyCurrentSensor) SetPropertyMonitoringRange(value uint16) (err error) {
	return instance.SetProperty("MonitoringRange", value)
}

// GetMonitoringRange gets the value of MonitoringRange for the instance
func (instance *CIM_AssociatedSupplyCurrentSensor) GetPropertyMonitoringRange() (value uint16, err error) {
	retValue, err := instance.GetProperty("MonitoringRange")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
