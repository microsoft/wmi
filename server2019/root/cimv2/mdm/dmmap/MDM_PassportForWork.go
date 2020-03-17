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

// MDM_PassportForWork struct
type MDM_PassportForWork struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	UseBiometrics bool
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_PassportForWork) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_PassportForWork) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_PassportForWork) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_PassportForWork) GetPropertyParentID() (value string, err error) {
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

// SetUseBiometrics sets the value of UseBiometrics for the instance
func (instance *MDM_PassportForWork) SetPropertyUseBiometrics(value bool) (err error) {
	return instance.SetProperty("UseBiometrics", value)
}

// GetUseBiometrics gets the value of UseBiometrics for the instance
func (instance *MDM_PassportForWork) GetPropertyUseBiometrics() (value bool, err error) {
	retValue, err := instance.GetProperty("UseBiometrics")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
