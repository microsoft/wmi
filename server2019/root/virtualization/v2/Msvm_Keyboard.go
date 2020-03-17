// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_Keyboard struct
type Msvm_Keyboard struct {
	CIM_UserDevice

	//
	Layout string

	//
	NumberOfFunctionKeys uint16

	//
	Password uint16

	//
	UnicodeSupported bool
}

// SetLayout sets the value of Layout for the instance
func (instance *Msvm_Keyboard) SetPropertyLayout(value string) (err error) {
	return instance.SetProperty("Layout", value)
}

// GetLayout gets the value of Layout for the instance
func (instance *Msvm_Keyboard) GetPropertyLayout() (value string, err error) {
	retValue, err := instance.GetProperty("Layout")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfFunctionKeys sets the value of NumberOfFunctionKeys for the instance
func (instance *Msvm_Keyboard) SetPropertyNumberOfFunctionKeys(value uint16) (err error) {
	return instance.SetProperty("NumberOfFunctionKeys", value)
}

// GetNumberOfFunctionKeys gets the value of NumberOfFunctionKeys for the instance
func (instance *Msvm_Keyboard) GetPropertyNumberOfFunctionKeys() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfFunctionKeys")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPassword sets the value of Password for the instance
func (instance *Msvm_Keyboard) SetPropertyPassword(value uint16) (err error) {
	return instance.SetProperty("Password", value)
}

// GetPassword gets the value of Password for the instance
func (instance *Msvm_Keyboard) GetPropertyPassword() (value uint16, err error) {
	retValue, err := instance.GetProperty("Password")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUnicodeSupported sets the value of UnicodeSupported for the instance
func (instance *Msvm_Keyboard) SetPropertyUnicodeSupported(value bool) (err error) {
	return instance.SetProperty("UnicodeSupported", value)
}

// GetUnicodeSupported gets the value of UnicodeSupported for the instance
func (instance *Msvm_Keyboard) GetPropertyUnicodeSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("UnicodeSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="KeyCode" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_Keyboard) PressKey( /* IN */ KeyCode uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("PressKey", KeyCode)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="KeyCode" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_Keyboard) ReleaseKey( /* IN */ KeyCode uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ReleaseKey", KeyCode)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="KeyCode" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_Keyboard) TypeKey( /* IN */ KeyCode uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("TypeKey", KeyCode)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="KeyCode" type="uint32 "></param>

// <param name="KeyState" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_Keyboard) IsKeyPressed( /* IN */ KeyCode uint32,
	/* OUT */ KeyState bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsKeyPressed", KeyCode)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AsciiText" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_Keyboard) TypeText( /* IN */ AsciiText string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("TypeText", AsciiText)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_Keyboard) TypeCtrlAltDel() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("TypeCtrlAltDel")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Scancodes" type="uint8 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_Keyboard) TypeScancodes( /* IN */ Scancodes []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("TypeScancodes", Scancodes)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
