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

// MDM_DeviceStatus_NetworkIdentifiers01_01 struct
type MDM_DeviceStatus_NetworkIdentifiers01_01 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	IPAddressV4 string

	//
	IPAddressV6 string

	//
	IsConnected bool

	//
	ParentID string

	//
	Type int32
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) GetPropertyInstanceID() (value string, err error) {
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

// SetIPAddressV4 sets the value of IPAddressV4 for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) SetPropertyIPAddressV4(value string) (err error) {
	return instance.SetProperty("IPAddressV4", value)
}

// GetIPAddressV4 gets the value of IPAddressV4 for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) GetPropertyIPAddressV4() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddressV4")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPAddressV6 sets the value of IPAddressV6 for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) SetPropertyIPAddressV6(value string) (err error) {
	return instance.SetProperty("IPAddressV6", value)
}

// GetIPAddressV6 gets the value of IPAddressV6 for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) GetPropertyIPAddressV6() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddressV6")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsConnected sets the value of IsConnected for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) SetPropertyIsConnected(value bool) (err error) {
	return instance.SetProperty("IsConnected", value)
}

// GetIsConnected gets the value of IsConnected for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) GetPropertyIsConnected() (value bool, err error) {
	retValue, err := instance.GetProperty("IsConnected")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) GetPropertyParentID() (value string, err error) {
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

// SetType sets the value of Type for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) SetPropertyType(value int32) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) GetPropertyType() (value int32, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
