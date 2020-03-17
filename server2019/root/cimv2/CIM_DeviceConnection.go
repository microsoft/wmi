// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_DeviceConnection struct
type CIM_DeviceConnection struct {
	CIM_Dependency

	//
	NegotiatedDataWidth uint32

	//
	NegotiatedSpeed uint64
}

// SetNegotiatedDataWidth sets the value of NegotiatedDataWidth for the instance
func (instance *CIM_DeviceConnection) SetPropertyNegotiatedDataWidth(value uint32) (err error) {
	return instance.SetProperty("NegotiatedDataWidth", value)
}

// GetNegotiatedDataWidth gets the value of NegotiatedDataWidth for the instance
func (instance *CIM_DeviceConnection) GetPropertyNegotiatedDataWidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("NegotiatedDataWidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNegotiatedSpeed sets the value of NegotiatedSpeed for the instance
func (instance *CIM_DeviceConnection) SetPropertyNegotiatedSpeed(value uint64) (err error) {
	return instance.SetProperty("NegotiatedSpeed", value)
}

// GetNegotiatedSpeed gets the value of NegotiatedSpeed for the instance
func (instance *CIM_DeviceConnection) GetPropertyNegotiatedSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("NegotiatedSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
