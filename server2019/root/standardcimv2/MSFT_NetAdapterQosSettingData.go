// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetAdapterQosSettingData struct
type MSFT_NetAdapterQosSettingData struct {
	MSFT_NetAdapterSettingData

	//
	CurrentCapabilities MSFT_NetAdapter_QosCapabilities

	//
	Enabled bool

	//
	HardwareCapabilities MSFT_NetAdapter_QosCapabilities

	//
	OperationalSettings MSFT_NetAdapter_QosSettings

	//
	RemoteSettings MSFT_NetAdapter_QosSettings
}

// SetCurrentCapabilities sets the value of CurrentCapabilities for the instance
func (instance *MSFT_NetAdapterQosSettingData) SetPropertyCurrentCapabilities(value MSFT_NetAdapter_QosCapabilities) (err error) {
	return instance.SetProperty("CurrentCapabilities", value)
}

// GetCurrentCapabilities gets the value of CurrentCapabilities for the instance
func (instance *MSFT_NetAdapterQosSettingData) GetPropertyCurrentCapabilities() (value MSFT_NetAdapter_QosCapabilities, err error) {
	retValue, err := instance.GetProperty("CurrentCapabilities")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetAdapter_QosCapabilities)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_NetAdapterQosSettingData) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_NetAdapterQosSettingData) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHardwareCapabilities sets the value of HardwareCapabilities for the instance
func (instance *MSFT_NetAdapterQosSettingData) SetPropertyHardwareCapabilities(value MSFT_NetAdapter_QosCapabilities) (err error) {
	return instance.SetProperty("HardwareCapabilities", value)
}

// GetHardwareCapabilities gets the value of HardwareCapabilities for the instance
func (instance *MSFT_NetAdapterQosSettingData) GetPropertyHardwareCapabilities() (value MSFT_NetAdapter_QosCapabilities, err error) {
	retValue, err := instance.GetProperty("HardwareCapabilities")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetAdapter_QosCapabilities)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalSettings sets the value of OperationalSettings for the instance
func (instance *MSFT_NetAdapterQosSettingData) SetPropertyOperationalSettings(value MSFT_NetAdapter_QosSettings) (err error) {
	return instance.SetProperty("OperationalSettings", value)
}

// GetOperationalSettings gets the value of OperationalSettings for the instance
func (instance *MSFT_NetAdapterQosSettingData) GetPropertyOperationalSettings() (value MSFT_NetAdapter_QosSettings, err error) {
	retValue, err := instance.GetProperty("OperationalSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetAdapter_QosSettings)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteSettings sets the value of RemoteSettings for the instance
func (instance *MSFT_NetAdapterQosSettingData) SetPropertyRemoteSettings(value MSFT_NetAdapter_QosSettings) (err error) {
	return instance.SetProperty("RemoteSettings", value)
}

// GetRemoteSettings gets the value of RemoteSettings for the instance
func (instance *MSFT_NetAdapterQosSettingData) GetPropertyRemoteSettings() (value MSFT_NetAdapter_QosSettings, err error) {
	retValue, err := instance.GetProperty("RemoteSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetAdapter_QosSettings)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="CmdletOutput" type="MSFT_NetAdapterQosSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterQosSettingData) Enable( /* OUT */ CmdletOutput MSFT_NetAdapterQosSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CmdletOutput" type="MSFT_NetAdapterQosSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterQosSettingData) Disable( /* OUT */ CmdletOutput MSFT_NetAdapterQosSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
