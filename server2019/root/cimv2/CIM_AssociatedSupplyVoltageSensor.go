// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_AssociatedSupplyVoltageSensor struct
type CIM_AssociatedSupplyVoltageSensor struct {
	CIM_AssociatedSensor

	//
	MonitoringRange uint16
}

// SetMonitoringRange sets the value of MonitoringRange for the instance
func (instance *CIM_AssociatedSupplyVoltageSensor) SetPropertyMonitoringRange(value uint16) (err error) {
	return instance.SetProperty("MonitoringRange", value)
}

// GetMonitoringRange gets the value of MonitoringRange for the instance
func (instance *CIM_AssociatedSupplyVoltageSensor) GetPropertyMonitoringRange() (value uint16, err error) {
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
