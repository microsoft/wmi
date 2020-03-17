// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_SecuritySetting struct
type Win32_SecuritySetting struct {
	CIM_Setting

	//
	ControlFlags uint32
}

// SetControlFlags sets the value of ControlFlags for the instance
func (instance *Win32_SecuritySetting) SetPropertyControlFlags(value uint32) (err error) {
	return instance.SetProperty("ControlFlags", value)
}

// GetControlFlags gets the value of ControlFlags for the instance
func (instance *Win32_SecuritySetting) GetPropertyControlFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("ControlFlags")
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

// <param name="Descriptor" type="Win32_SecurityDescriptor "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_SecuritySetting) GetSecurityDescriptor( /* OUT */ Descriptor Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSecurityDescriptor")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Descriptor" type="Win32_SecurityDescriptor "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_SecuritySetting) SetSecurityDescriptor( /* IN */ Descriptor Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetSecurityDescriptor", Descriptor)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
