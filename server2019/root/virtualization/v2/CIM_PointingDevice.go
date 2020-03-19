// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_PointingDevice struct
type CIM_PointingDevice struct {
	*CIM_UserDevice

	// Integer indicating whether the PointingDevice is configured for right (value=2) or left handed operation (value=3). Also, the values, "Unknown" (0) and "Not Applicable" (1), can be defined.
	Handedness PointingDevice_Handedness

	// Number of buttons. If the PointingDevice has no buttons, enter 0.
	NumberOfButtons uint8

	// The type of the pointing device.
	PointingType PointingDevice_PointingType

	// Tracking resolution of the PointingDevice in Counts per Inch.
	Resolution uint32
}

func NewCIM_PointingDeviceEx1(instance *cim.WmiInstance) (newInstance *CIM_PointingDevice, err error) {
	tmp, err := NewCIM_UserDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PointingDevice{
		CIM_UserDevice: tmp,
	}
	return
}

func NewCIM_PointingDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PointingDevice, err error) {
	tmp, err := NewCIM_UserDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PointingDevice{
		CIM_UserDevice: tmp,
	}
	return
}

// SetHandedness sets the value of Handedness for the instance
func (instance *CIM_PointingDevice) SetPropertyHandedness(value PointingDevice_Handedness) (err error) {
	return instance.SetProperty("Handedness", value)
}

// GetHandedness gets the value of Handedness for the instance
func (instance *CIM_PointingDevice) GetPropertyHandedness() (value PointingDevice_Handedness, err error) {
	retValue, err := instance.GetProperty("Handedness")
	if err != nil {
		return
	}
	value, ok := retValue.(PointingDevice_Handedness)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfButtons sets the value of NumberOfButtons for the instance
func (instance *CIM_PointingDevice) SetPropertyNumberOfButtons(value uint8) (err error) {
	return instance.SetProperty("NumberOfButtons", value)
}

// GetNumberOfButtons gets the value of NumberOfButtons for the instance
func (instance *CIM_PointingDevice) GetPropertyNumberOfButtons() (value uint8, err error) {
	retValue, err := instance.GetProperty("NumberOfButtons")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPointingType sets the value of PointingType for the instance
func (instance *CIM_PointingDevice) SetPropertyPointingType(value PointingDevice_PointingType) (err error) {
	return instance.SetProperty("PointingType", value)
}

// GetPointingType gets the value of PointingType for the instance
func (instance *CIM_PointingDevice) GetPropertyPointingType() (value PointingDevice_PointingType, err error) {
	retValue, err := instance.GetProperty("PointingType")
	if err != nil {
		return
	}
	value, ok := retValue.(PointingDevice_PointingType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResolution sets the value of Resolution for the instance
func (instance *CIM_PointingDevice) SetPropertyResolution(value uint32) (err error) {
	return instance.SetProperty("Resolution", value)
}

// GetResolution gets the value of Resolution for the instance
func (instance *CIM_PointingDevice) GetPropertyResolution() (value uint32, err error) {
	retValue, err := instance.GetProperty("Resolution")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
