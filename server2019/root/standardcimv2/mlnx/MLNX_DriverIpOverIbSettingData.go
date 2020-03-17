// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// MLNX_DriverIpOverIbSettingData struct
type MLNX_DriverIpOverIbSettingData struct {
	MLNX_DriverSettingData

	//
	ModeFlags uint32

	//
	NdkDebugFlags uint32

	//
	NdkDebugLevel uint32
}

// SetModeFlags sets the value of ModeFlags for the instance
func (instance *MLNX_DriverIpOverIbSettingData) SetPropertyModeFlags(value uint32) (err error) {
	return instance.SetProperty("ModeFlags", value)
}

// GetModeFlags gets the value of ModeFlags for the instance
func (instance *MLNX_DriverIpOverIbSettingData) GetPropertyModeFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("ModeFlags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdkDebugFlags sets the value of NdkDebugFlags for the instance
func (instance *MLNX_DriverIpOverIbSettingData) SetPropertyNdkDebugFlags(value uint32) (err error) {
	return instance.SetProperty("NdkDebugFlags", value)
}

// GetNdkDebugFlags gets the value of NdkDebugFlags for the instance
func (instance *MLNX_DriverIpOverIbSettingData) GetPropertyNdkDebugFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdkDebugFlags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdkDebugLevel sets the value of NdkDebugLevel for the instance
func (instance *MLNX_DriverIpOverIbSettingData) SetPropertyNdkDebugLevel(value uint32) (err error) {
	return instance.SetProperty("NdkDebugLevel", value)
}

// GetNdkDebugLevel gets the value of NdkDebugLevel for the instance
func (instance *MLNX_DriverIpOverIbSettingData) GetPropertyNdkDebugLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdkDebugLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ModeFlags" type="uint32 "></param>
// <param name="NdkDebugFlags" type="uint32 "></param>
// <param name="NdkDebugLevel" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_DriverIpOverIbSettingData) SetValue( /* IN */ ModeFlags uint32,
	/* IN */ NdkDebugFlags uint32,
	/* IN */ NdkDebugLevel uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetValue", ModeFlags, NdkDebugFlags, NdkDebugLevel)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_DriverIpOverIbSettingData) SetDefault() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetDefault")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
