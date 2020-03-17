// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapter_QosClassificationElement struct
type MSFT_NetAdapter_QosClassificationElement struct {
	cim.WmiInstance

	//
	Priority uint8

	//
	ProtocolSelector uint16

	//
	ProtocolSpecificValue uint16
}

// SetPriority sets the value of Priority for the instance
func (instance *MSFT_NetAdapter_QosClassificationElement) SetPropertyPriority(value uint8) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *MSFT_NetAdapter_QosClassificationElement) GetPropertyPriority() (value uint8, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocolSelector sets the value of ProtocolSelector for the instance
func (instance *MSFT_NetAdapter_QosClassificationElement) SetPropertyProtocolSelector(value uint16) (err error) {
	return instance.SetProperty("ProtocolSelector", value)
}

// GetProtocolSelector gets the value of ProtocolSelector for the instance
func (instance *MSFT_NetAdapter_QosClassificationElement) GetPropertyProtocolSelector() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProtocolSelector")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocolSpecificValue sets the value of ProtocolSpecificValue for the instance
func (instance *MSFT_NetAdapter_QosClassificationElement) SetPropertyProtocolSpecificValue(value uint16) (err error) {
	return instance.SetProperty("ProtocolSpecificValue", value)
}

// GetProtocolSpecificValue gets the value of ProtocolSpecificValue for the instance
func (instance *MSFT_NetAdapter_QosClassificationElement) GetPropertyProtocolSpecificValue() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProtocolSpecificValue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
