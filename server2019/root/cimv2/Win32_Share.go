// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_Share struct
type Win32_Share struct {
	CIM_LogicalElement

	//
	AccessMask uint32

	//
	AllowMaximum bool

	//
	MaximumAllowed uint32

	//
	Path string

	//
	Type uint32
}

// SetAccessMask sets the value of AccessMask for the instance
func (instance *Win32_Share) SetPropertyAccessMask(value uint32) (err error) {
	return instance.SetProperty("AccessMask", value)
}

// GetAccessMask gets the value of AccessMask for the instance
func (instance *Win32_Share) GetPropertyAccessMask() (value uint32, err error) {
	retValue, err := instance.GetProperty("AccessMask")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowMaximum sets the value of AllowMaximum for the instance
func (instance *Win32_Share) SetPropertyAllowMaximum(value bool) (err error) {
	return instance.SetProperty("AllowMaximum", value)
}

// GetAllowMaximum gets the value of AllowMaximum for the instance
func (instance *Win32_Share) GetPropertyAllowMaximum() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowMaximum")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumAllowed sets the value of MaximumAllowed for the instance
func (instance *Win32_Share) SetPropertyMaximumAllowed(value uint32) (err error) {
	return instance.SetProperty("MaximumAllowed", value)
}

// GetMaximumAllowed gets the value of MaximumAllowed for the instance
func (instance *Win32_Share) GetPropertyMaximumAllowed() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumAllowed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPath sets the value of Path for the instance
func (instance *Win32_Share) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", value)
}

// GetPath gets the value of Path for the instance
func (instance *Win32_Share) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *Win32_Share) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *Win32_Share) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
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

// <param name="Access" type="Win32_SecurityDescriptor "></param>
// <param name="Description" type="string "></param>
// <param name="MaximumAllowed" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="Password" type="string "></param>
// <param name="Path" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Share) Create( /* IN */ Path string,
	/* IN */ Name string,
	/* IN */ Type uint32,
	/* OPTIONAL IN */ MaximumAllowed uint32,
	/* OPTIONAL IN */ Description string,
	/* OPTIONAL IN */ Password string,
	/* OPTIONAL IN */ Access Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Create", Path, Name, Type, MaximumAllowed, Description, Password, Access)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Access" type="Win32_SecurityDescriptor "></param>
// <param name="Description" type="string "></param>
// <param name="MaximumAllowed" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Share) SetShareInfo( /* OPTIONAL IN */ MaximumAllowed uint32,
	/* OPTIONAL IN */ Description string,
	/* OPTIONAL IN */ Access Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetShareInfo", MaximumAllowed, Description, Access)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Share) GetAccessMask() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetAccessMask")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Share) Delete() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Delete")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
