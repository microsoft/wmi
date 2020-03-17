// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetEventNetworkAdapter struct
type MSFT_NetEventNetworkAdapter struct {
	MSFT_NetEventPacketCaptureTarget

	//
	InterfaceDescription string

	//
	PromiscuousMode bool
}

// SetInterfaceDescription sets the value of InterfaceDescription for the instance
func (instance *MSFT_NetEventNetworkAdapter) SetPropertyInterfaceDescription(value string) (err error) {
	return instance.SetProperty("InterfaceDescription", value)
}

// GetInterfaceDescription gets the value of InterfaceDescription for the instance
func (instance *MSFT_NetEventNetworkAdapter) GetPropertyInterfaceDescription() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPromiscuousMode sets the value of PromiscuousMode for the instance
func (instance *MSFT_NetEventNetworkAdapter) SetPropertyPromiscuousMode(value bool) (err error) {
	return instance.SetProperty("PromiscuousMode", value)
}

// GetPromiscuousMode gets the value of PromiscuousMode for the instance
func (instance *MSFT_NetEventNetworkAdapter) GetPropertyPromiscuousMode() (value bool, err error) {
	retValue, err := instance.GetProperty("PromiscuousMode")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
