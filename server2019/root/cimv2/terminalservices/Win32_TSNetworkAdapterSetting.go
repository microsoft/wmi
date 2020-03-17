// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TSNetworkAdapterSetting struct
type Win32_TSNetworkAdapterSetting struct {
	Win32_TerminalSetting

	//
	DeviceIDList []string

	//
	MaximumConnections uint32

	//
	NetworkAdapterID string

	//
	NetworkAdapterLanaID uint32

	//
	NetworkAdapterList []string

	//
	NetworkAdapterName string

	//
	PolicySourceMaximumConnections uint32
}

// SetDeviceIDList sets the value of DeviceIDList for the instance
func (instance *Win32_TSNetworkAdapterSetting) SetPropertyDeviceIDList(value []string) (err error) {
	return instance.SetProperty("DeviceIDList", value)
}

// GetDeviceIDList gets the value of DeviceIDList for the instance
func (instance *Win32_TSNetworkAdapterSetting) GetPropertyDeviceIDList() (value []string, err error) {
	retValue, err := instance.GetProperty("DeviceIDList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumConnections sets the value of MaximumConnections for the instance
func (instance *Win32_TSNetworkAdapterSetting) SetPropertyMaximumConnections(value uint32) (err error) {
	return instance.SetProperty("MaximumConnections", value)
}

// GetMaximumConnections gets the value of MaximumConnections for the instance
func (instance *Win32_TSNetworkAdapterSetting) GetPropertyMaximumConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumConnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkAdapterID sets the value of NetworkAdapterID for the instance
func (instance *Win32_TSNetworkAdapterSetting) SetPropertyNetworkAdapterID(value string) (err error) {
	return instance.SetProperty("NetworkAdapterID", value)
}

// GetNetworkAdapterID gets the value of NetworkAdapterID for the instance
func (instance *Win32_TSNetworkAdapterSetting) GetPropertyNetworkAdapterID() (value string, err error) {
	retValue, err := instance.GetProperty("NetworkAdapterID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkAdapterLanaID sets the value of NetworkAdapterLanaID for the instance
func (instance *Win32_TSNetworkAdapterSetting) SetPropertyNetworkAdapterLanaID(value uint32) (err error) {
	return instance.SetProperty("NetworkAdapterLanaID", value)
}

// GetNetworkAdapterLanaID gets the value of NetworkAdapterLanaID for the instance
func (instance *Win32_TSNetworkAdapterSetting) GetPropertyNetworkAdapterLanaID() (value uint32, err error) {
	retValue, err := instance.GetProperty("NetworkAdapterLanaID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkAdapterList sets the value of NetworkAdapterList for the instance
func (instance *Win32_TSNetworkAdapterSetting) SetPropertyNetworkAdapterList(value []string) (err error) {
	return instance.SetProperty("NetworkAdapterList", value)
}

// GetNetworkAdapterList gets the value of NetworkAdapterList for the instance
func (instance *Win32_TSNetworkAdapterSetting) GetPropertyNetworkAdapterList() (value []string, err error) {
	retValue, err := instance.GetProperty("NetworkAdapterList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkAdapterName sets the value of NetworkAdapterName for the instance
func (instance *Win32_TSNetworkAdapterSetting) SetPropertyNetworkAdapterName(value string) (err error) {
	return instance.SetProperty("NetworkAdapterName", value)
}

// GetNetworkAdapterName gets the value of NetworkAdapterName for the instance
func (instance *Win32_TSNetworkAdapterSetting) GetPropertyNetworkAdapterName() (value string, err error) {
	retValue, err := instance.GetProperty("NetworkAdapterName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceMaximumConnections sets the value of PolicySourceMaximumConnections for the instance
func (instance *Win32_TSNetworkAdapterSetting) SetPropertyPolicySourceMaximumConnections(value uint32) (err error) {
	return instance.SetProperty("PolicySourceMaximumConnections", value)
}

// GetPolicySourceMaximumConnections gets the value of PolicySourceMaximumConnections for the instance
func (instance *Win32_TSNetworkAdapterSetting) GetPropertyPolicySourceMaximumConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceMaximumConnections")
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

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSNetworkAdapterSetting) SelectAllNetworkAdapters() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SelectAllNetworkAdapters")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NetworkAdapterLanaID" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSNetworkAdapterSetting) SetNetworkAdapterLanaID( /* IN */ NetworkAdapterLanaID uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetNetworkAdapterLanaID", NetworkAdapterLanaID)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NetworkAdapterIP" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSNetworkAdapterSetting) SelectNetworkAdapterIP( /* IN */ NetworkAdapterIP string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SelectNetworkAdapterIP", NetworkAdapterIP)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
