// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_FCPort struct
type CIM_FCPort struct {
	CIM_NetworkPort

	// An array of integers that indicates the Classes of Service that are active. The Active COS is indicated in ActiveCOS.
	ActiveCOS []FCPort_ActiveCOS

	// An array of integers that indicates the Fibre Channel FC-4 protocols currently running. A list of all supported protocols is indicated in the SupportedFC4Types property.
	ActiveFC4Types []FCPort_ActiveFC4Types

	// An array of integers that indicates the Fibre Channel Classes of Service that are supported. The active COS are indicated in ActiveCOS.
	SupportedCOS []FCPort_SupportedCOS

	// An array of integers that indicates the Fibre Channel FC-4 protocols supported. The protocols that are active and running are indicated in the ActiveFC4Types property.
	SupportedFC4Types []FCPort_SupportedFC4Types
}

// SetActiveCOS sets the value of ActiveCOS for the instance
func (instance *CIM_FCPort) SetPropertyActiveCOS(value []FCPort_ActiveCOS) (err error) {
	return instance.SetProperty("ActiveCOS", value)
}

// GetActiveCOS gets the value of ActiveCOS for the instance
func (instance *CIM_FCPort) GetPropertyActiveCOS() (value []FCPort_ActiveCOS, err error) {
	retValue, err := instance.GetProperty("ActiveCOS")
	if err != nil {
		return
	}
	value, ok := retValue.([]FCPort_ActiveCOS)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActiveFC4Types sets the value of ActiveFC4Types for the instance
func (instance *CIM_FCPort) SetPropertyActiveFC4Types(value []FCPort_ActiveFC4Types) (err error) {
	return instance.SetProperty("ActiveFC4Types", value)
}

// GetActiveFC4Types gets the value of ActiveFC4Types for the instance
func (instance *CIM_FCPort) GetPropertyActiveFC4Types() (value []FCPort_ActiveFC4Types, err error) {
	retValue, err := instance.GetProperty("ActiveFC4Types")
	if err != nil {
		return
	}
	value, ok := retValue.([]FCPort_ActiveFC4Types)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedCOS sets the value of SupportedCOS for the instance
func (instance *CIM_FCPort) SetPropertySupportedCOS(value []FCPort_SupportedCOS) (err error) {
	return instance.SetProperty("SupportedCOS", value)
}

// GetSupportedCOS gets the value of SupportedCOS for the instance
func (instance *CIM_FCPort) GetPropertySupportedCOS() (value []FCPort_SupportedCOS, err error) {
	retValue, err := instance.GetProperty("SupportedCOS")
	if err != nil {
		return
	}
	value, ok := retValue.([]FCPort_SupportedCOS)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedFC4Types sets the value of SupportedFC4Types for the instance
func (instance *CIM_FCPort) SetPropertySupportedFC4Types(value []FCPort_SupportedFC4Types) (err error) {
	return instance.SetProperty("SupportedFC4Types", value)
}

// GetSupportedFC4Types gets the value of SupportedFC4Types for the instance
func (instance *CIM_FCPort) GetPropertySupportedFC4Types() (value []FCPort_SupportedFC4Types, err error) {
	retValue, err := instance.GetProperty("SupportedFC4Types")
	if err != nil {
		return
	}
	value, ok := retValue.([]FCPort_SupportedFC4Types)
	if !ok {
		// TODO: Set an error
	}
	return
}
