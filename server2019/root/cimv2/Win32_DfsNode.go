// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_DfsNode struct
type Win32_DfsNode struct {
	CIM_LogicalElement

	//
	Root bool

	//
	State uint32

	//
	Timeout uint32
}

// SetRoot sets the value of Root for the instance
func (instance *Win32_DfsNode) SetPropertyRoot(value bool) (err error) {
	return instance.SetProperty("Root", value)
}

// GetRoot gets the value of Root for the instance
func (instance *Win32_DfsNode) GetPropertyRoot() (value bool, err error) {
	retValue, err := instance.GetProperty("Root")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *Win32_DfsNode) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *Win32_DfsNode) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeout sets the value of Timeout for the instance
func (instance *Win32_DfsNode) SetPropertyTimeout(value uint32) (err error) {
	return instance.SetProperty("Timeout", value)
}

// GetTimeout gets the value of Timeout for the instance
func (instance *Win32_DfsNode) GetPropertyTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("Timeout")
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

// <param name="Description" type="string "></param>
// <param name="DfsEntryPath" type="string "></param>
// <param name="ServerName" type="string "></param>
// <param name="ShareName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_DfsNode) Create( /* IN */ DfsEntryPath string,
	/* IN */ ServerName string,
	/* IN */ ShareName string,
	/* OPTIONAL IN */ Description string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Create", DfsEntryPath, ServerName, ShareName, Description)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
