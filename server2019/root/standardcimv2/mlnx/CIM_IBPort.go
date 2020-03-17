// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_IBPort struct
type CIM_IBPort struct {
	CIM_NetworkPort

	//
	LIDMask uint8

	//
	LinkSpeedActive uint8

	//
	LinkWidthActive uint16
}

// SetLIDMask sets the value of LIDMask for the instance
func (instance *CIM_IBPort) SetPropertyLIDMask(value uint8) (err error) {
	return instance.SetProperty("LIDMask", value)
}

// GetLIDMask gets the value of LIDMask for the instance
func (instance *CIM_IBPort) GetPropertyLIDMask() (value uint8, err error) {
	retValue, err := instance.GetProperty("LIDMask")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLinkSpeedActive sets the value of LinkSpeedActive for the instance
func (instance *CIM_IBPort) SetPropertyLinkSpeedActive(value uint8) (err error) {
	return instance.SetProperty("LinkSpeedActive", value)
}

// GetLinkSpeedActive gets the value of LinkSpeedActive for the instance
func (instance *CIM_IBPort) GetPropertyLinkSpeedActive() (value uint8, err error) {
	retValue, err := instance.GetProperty("LinkSpeedActive")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLinkWidthActive sets the value of LinkWidthActive for the instance
func (instance *CIM_IBPort) SetPropertyLinkWidthActive(value uint16) (err error) {
	return instance.SetProperty("LinkWidthActive", value)
}

// GetLinkWidthActive gets the value of LinkWidthActive for the instance
func (instance *CIM_IBPort) GetPropertyLinkWidthActive() (value uint16, err error) {
	retValue, err := instance.GetProperty("LinkWidthActive")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
