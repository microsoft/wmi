// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// MLNX_PCIDevicePortTypeSettingData struct
type MLNX_PCIDevicePortTypeSettingData struct {
	MLNX_PCIDeviceSettingData

	//
	Configuration []string

	//
	CurrentSetting []string

	//
	DefaultSetting []string
}

// SetConfiguration sets the value of Configuration for the instance
func (instance *MLNX_PCIDevicePortTypeSettingData) SetPropertyConfiguration(value []string) (err error) {
	return instance.SetProperty("Configuration", value)
}

// GetConfiguration gets the value of Configuration for the instance
func (instance *MLNX_PCIDevicePortTypeSettingData) GetPropertyConfiguration() (value []string, err error) {
	retValue, err := instance.GetProperty("Configuration")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentSetting sets the value of CurrentSetting for the instance
func (instance *MLNX_PCIDevicePortTypeSettingData) SetPropertyCurrentSetting(value []string) (err error) {
	return instance.SetProperty("CurrentSetting", value)
}

// GetCurrentSetting gets the value of CurrentSetting for the instance
func (instance *MLNX_PCIDevicePortTypeSettingData) GetPropertyCurrentSetting() (value []string, err error) {
	retValue, err := instance.GetProperty("CurrentSetting")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultSetting sets the value of DefaultSetting for the instance
func (instance *MLNX_PCIDevicePortTypeSettingData) SetPropertyDefaultSetting(value []string) (err error) {
	return instance.SetProperty("DefaultSetting", value)
}

// GetDefaultSetting gets the value of DefaultSetting for the instance
func (instance *MLNX_PCIDevicePortTypeSettingData) GetPropertyDefaultSetting() (value []string, err error) {
	retValue, err := instance.GetProperty("DefaultSetting")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="port1" type="string "></param>
// <param name="port2" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_PCIDevicePortTypeSettingData) SetValue( /* IN */ port1 string,
	/* IN */ port2 string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetValue", port1, port2)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_PCIDevicePortTypeSettingData) SetDefault() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetDefault")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
