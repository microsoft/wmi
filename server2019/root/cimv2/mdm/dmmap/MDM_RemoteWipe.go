// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_RemoteWipe struct
type MDM_RemoteWipe struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_RemoteWipe) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_RemoteWipe) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_RemoteWipe) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_RemoteWipe) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_RemoteWipe) doWipeMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("doWipeMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_RemoteWipe) doWipePersistProvisionedDataMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("doWipePersistProvisionedDataMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_RemoteWipe) doWipeProtectedMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("doWipeProtectedMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_RemoteWipe) doWipePersistUserDataMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("doWipePersistUserDataMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
