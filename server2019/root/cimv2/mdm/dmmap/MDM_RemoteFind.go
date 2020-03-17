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

// MDM_RemoteFind struct
type MDM_RemoteFind struct {
	cim.WmiInstance

	//
	DesiredAccuracy int32

	//
	InstanceID string

	//
	MaximumAge int32

	//
	ParentID string

	//
	Timeout int32
}

// SetDesiredAccuracy sets the value of DesiredAccuracy for the instance
func (instance *MDM_RemoteFind) SetPropertyDesiredAccuracy(value int32) (err error) {
	return instance.SetProperty("DesiredAccuracy", value)
}

// GetDesiredAccuracy gets the value of DesiredAccuracy for the instance
func (instance *MDM_RemoteFind) GetPropertyDesiredAccuracy() (value int32, err error) {
	retValue, err := instance.GetProperty("DesiredAccuracy")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_RemoteFind) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_RemoteFind) GetPropertyInstanceID() (value string, err error) {
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

// SetMaximumAge sets the value of MaximumAge for the instance
func (instance *MDM_RemoteFind) SetPropertyMaximumAge(value int32) (err error) {
	return instance.SetProperty("MaximumAge", value)
}

// GetMaximumAge gets the value of MaximumAge for the instance
func (instance *MDM_RemoteFind) GetPropertyMaximumAge() (value int32, err error) {
	retValue, err := instance.GetProperty("MaximumAge")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_RemoteFind) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_RemoteFind) GetPropertyParentID() (value string, err error) {
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

// SetTimeout sets the value of Timeout for the instance
func (instance *MDM_RemoteFind) SetPropertyTimeout(value int32) (err error) {
	return instance.SetProperty("Timeout", value)
}

// GetTimeout gets the value of Timeout for the instance
func (instance *MDM_RemoteFind) GetPropertyTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("Timeout")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
