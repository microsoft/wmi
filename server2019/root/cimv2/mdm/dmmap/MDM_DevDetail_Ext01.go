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

// MDM_DevDetail_Ext01 struct
type MDM_DevDetail_Ext01 struct {
	cim.WmiInstance

	//
	DeviceHardwareData string

	//
	InstanceID string

	//
	ParentID string

	//
	WLANMACAddress string
}

// SetDeviceHardwareData sets the value of DeviceHardwareData for the instance
func (instance *MDM_DevDetail_Ext01) SetPropertyDeviceHardwareData(value string) (err error) {
	return instance.SetProperty("DeviceHardwareData", value)
}

// GetDeviceHardwareData gets the value of DeviceHardwareData for the instance
func (instance *MDM_DevDetail_Ext01) GetPropertyDeviceHardwareData() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceHardwareData")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_DevDetail_Ext01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DevDetail_Ext01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_DevDetail_Ext01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DevDetail_Ext01) GetPropertyParentID() (value string, err error) {
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

// SetWLANMACAddress sets the value of WLANMACAddress for the instance
func (instance *MDM_DevDetail_Ext01) SetPropertyWLANMACAddress(value string) (err error) {
	return instance.SetProperty("WLANMACAddress", value)
}

// GetWLANMACAddress gets the value of WLANMACAddress for the instance
func (instance *MDM_DevDetail_Ext01) GetPropertyWLANMACAddress() (value string, err error) {
	retValue, err := instance.GetProperty("WLANMACAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
