// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TSVirtualIP struct
type Win32_TSVirtualIP struct {
	CIM_Setting

	//
	NetworkAdapterDescription string

	//
	NetworkAdapterDescriptionList []string

	//
	NetworkAdapterMacAddress string

	//
	NetworkAdapterMacList []string

	//
	PolicySourceNetworkAdapter uint32

	//
	PolicySourceProgramList uint32

	//
	PolicySourceVirtualIPActive uint32

	//
	PolicySourceVirtualIPMode uint32

	//
	ProgramList []string

	//
	VirtualIPActive uint32

	//
	VirtualIPMode uint32

	//
	VirtualizeLoopbackAddressesEnabled uint32
}

// SetNetworkAdapterDescription sets the value of NetworkAdapterDescription for the instance
func (instance *Win32_TSVirtualIP) SetPropertyNetworkAdapterDescription(value string) (err error) {
	return instance.SetProperty("NetworkAdapterDescription", value)
}

// GetNetworkAdapterDescription gets the value of NetworkAdapterDescription for the instance
func (instance *Win32_TSVirtualIP) GetPropertyNetworkAdapterDescription() (value string, err error) {
	retValue, err := instance.GetProperty("NetworkAdapterDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkAdapterDescriptionList sets the value of NetworkAdapterDescriptionList for the instance
func (instance *Win32_TSVirtualIP) SetPropertyNetworkAdapterDescriptionList(value []string) (err error) {
	return instance.SetProperty("NetworkAdapterDescriptionList", value)
}

// GetNetworkAdapterDescriptionList gets the value of NetworkAdapterDescriptionList for the instance
func (instance *Win32_TSVirtualIP) GetPropertyNetworkAdapterDescriptionList() (value []string, err error) {
	retValue, err := instance.GetProperty("NetworkAdapterDescriptionList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkAdapterMacAddress sets the value of NetworkAdapterMacAddress for the instance
func (instance *Win32_TSVirtualIP) SetPropertyNetworkAdapterMacAddress(value string) (err error) {
	return instance.SetProperty("NetworkAdapterMacAddress", value)
}

// GetNetworkAdapterMacAddress gets the value of NetworkAdapterMacAddress for the instance
func (instance *Win32_TSVirtualIP) GetPropertyNetworkAdapterMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("NetworkAdapterMacAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkAdapterMacList sets the value of NetworkAdapterMacList for the instance
func (instance *Win32_TSVirtualIP) SetPropertyNetworkAdapterMacList(value []string) (err error) {
	return instance.SetProperty("NetworkAdapterMacList", value)
}

// GetNetworkAdapterMacList gets the value of NetworkAdapterMacList for the instance
func (instance *Win32_TSVirtualIP) GetPropertyNetworkAdapterMacList() (value []string, err error) {
	retValue, err := instance.GetProperty("NetworkAdapterMacList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceNetworkAdapter sets the value of PolicySourceNetworkAdapter for the instance
func (instance *Win32_TSVirtualIP) SetPropertyPolicySourceNetworkAdapter(value uint32) (err error) {
	return instance.SetProperty("PolicySourceNetworkAdapter", value)
}

// GetPolicySourceNetworkAdapter gets the value of PolicySourceNetworkAdapter for the instance
func (instance *Win32_TSVirtualIP) GetPropertyPolicySourceNetworkAdapter() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceNetworkAdapter")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceProgramList sets the value of PolicySourceProgramList for the instance
func (instance *Win32_TSVirtualIP) SetPropertyPolicySourceProgramList(value uint32) (err error) {
	return instance.SetProperty("PolicySourceProgramList", value)
}

// GetPolicySourceProgramList gets the value of PolicySourceProgramList for the instance
func (instance *Win32_TSVirtualIP) GetPropertyPolicySourceProgramList() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceProgramList")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceVirtualIPActive sets the value of PolicySourceVirtualIPActive for the instance
func (instance *Win32_TSVirtualIP) SetPropertyPolicySourceVirtualIPActive(value uint32) (err error) {
	return instance.SetProperty("PolicySourceVirtualIPActive", value)
}

// GetPolicySourceVirtualIPActive gets the value of PolicySourceVirtualIPActive for the instance
func (instance *Win32_TSVirtualIP) GetPropertyPolicySourceVirtualIPActive() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceVirtualIPActive")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceVirtualIPMode sets the value of PolicySourceVirtualIPMode for the instance
func (instance *Win32_TSVirtualIP) SetPropertyPolicySourceVirtualIPMode(value uint32) (err error) {
	return instance.SetProperty("PolicySourceVirtualIPMode", value)
}

// GetPolicySourceVirtualIPMode gets the value of PolicySourceVirtualIPMode for the instance
func (instance *Win32_TSVirtualIP) GetPropertyPolicySourceVirtualIPMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceVirtualIPMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProgramList sets the value of ProgramList for the instance
func (instance *Win32_TSVirtualIP) SetPropertyProgramList(value []string) (err error) {
	return instance.SetProperty("ProgramList", value)
}

// GetProgramList gets the value of ProgramList for the instance
func (instance *Win32_TSVirtualIP) GetPropertyProgramList() (value []string, err error) {
	retValue, err := instance.GetProperty("ProgramList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualIPActive sets the value of VirtualIPActive for the instance
func (instance *Win32_TSVirtualIP) SetPropertyVirtualIPActive(value uint32) (err error) {
	return instance.SetProperty("VirtualIPActive", value)
}

// GetVirtualIPActive gets the value of VirtualIPActive for the instance
func (instance *Win32_TSVirtualIP) GetPropertyVirtualIPActive() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualIPActive")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualIPMode sets the value of VirtualIPMode for the instance
func (instance *Win32_TSVirtualIP) SetPropertyVirtualIPMode(value uint32) (err error) {
	return instance.SetProperty("VirtualIPMode", value)
}

// GetVirtualIPMode gets the value of VirtualIPMode for the instance
func (instance *Win32_TSVirtualIP) GetPropertyVirtualIPMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualIPMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualizeLoopbackAddressesEnabled sets the value of VirtualizeLoopbackAddressesEnabled for the instance
func (instance *Win32_TSVirtualIP) SetPropertyVirtualizeLoopbackAddressesEnabled(value uint32) (err error) {
	return instance.SetProperty("VirtualizeLoopbackAddressesEnabled", value)
}

// GetVirtualizeLoopbackAddressesEnabled gets the value of VirtualizeLoopbackAddressesEnabled for the instance
func (instance *Win32_TSVirtualIP) GetPropertyVirtualizeLoopbackAddressesEnabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualizeLoopbackAddressesEnabled")
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

// <param name="VirtualIPActive" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSVirtualIP) SetVirtualIPActive( /* IN */ VirtualIPActive uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetVirtualIPActive", VirtualIPActive)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="VirtualIPMode" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSVirtualIP) SetVirtualIPMode( /* IN */ VirtualIPMode uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetVirtualIPMode", VirtualIPMode)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ProgramPath" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSVirtualIP) AddProgram( /* IN */ ProgramPath string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddProgram", ProgramPath)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ProgramPath" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSVirtualIP) RemoveProgram( /* IN */ ProgramPath string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveProgram", ProgramPath)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ProgramList" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSVirtualIP) SetProgramList( /* IN */ ProgramList []string) (result interface{}, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetProgramList", ProgramList)
	if err != nil {
		return
	}
	result = interface{}(retVal)
	return

}

//

// <param name="NetworkAdapterMacAddress" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSVirtualIP) SelectNetworkAdapter( /* IN */ NetworkAdapterMacAddress string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SelectNetworkAdapter", NetworkAdapterMacAddress)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="VirtualizeLoopbackAddressesEnabled" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSVirtualIP) SetVirtualizeLoopbackAddressesEnabled( /* IN */ VirtualizeLoopbackAddressesEnabled uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetVirtualizeLoopbackAddressesEnabled", VirtualizeLoopbackAddressesEnabled)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
