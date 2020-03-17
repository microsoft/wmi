// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_WsdPrinterPort struct
type MSFT_WsdPrinterPort struct {
	MSFT_PrinterPort

	//
	DeviceURL string

	//
	DeviceUUID string

	//
	DiscoveryMethod uint32
}

// SetDeviceURL sets the value of DeviceURL for the instance
func (instance *MSFT_WsdPrinterPort) SetPropertyDeviceURL(value string) (err error) {
	return instance.SetProperty("DeviceURL", value)
}

// GetDeviceURL gets the value of DeviceURL for the instance
func (instance *MSFT_WsdPrinterPort) GetPropertyDeviceURL() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceURL")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceUUID sets the value of DeviceUUID for the instance
func (instance *MSFT_WsdPrinterPort) SetPropertyDeviceUUID(value string) (err error) {
	return instance.SetProperty("DeviceUUID", value)
}

// GetDeviceUUID gets the value of DeviceUUID for the instance
func (instance *MSFT_WsdPrinterPort) GetPropertyDeviceUUID() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceUUID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiscoveryMethod sets the value of DiscoveryMethod for the instance
func (instance *MSFT_WsdPrinterPort) SetPropertyDiscoveryMethod(value uint32) (err error) {
	return instance.SetProperty("DiscoveryMethod", value)
}

// GetDiscoveryMethod gets the value of DiscoveryMethod for the instance
func (instance *MSFT_WsdPrinterPort) GetPropertyDiscoveryMethod() (value uint32, err error) {
	retValue, err := instance.GetProperty("DiscoveryMethod")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
