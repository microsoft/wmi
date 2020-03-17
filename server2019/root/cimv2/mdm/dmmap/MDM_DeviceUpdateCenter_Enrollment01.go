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

// MDM_DeviceUpdateCenter_Enrollment01 struct
type MDM_DeviceUpdateCenter_Enrollment01 struct {
	cim.WmiInstance

	//
	CustomPackageId string

	//
	DeviceModelId string

	//
	InstanceID string

	//
	OemPartnerRing string

	//
	ParentID string

	//
	PublisherId string
}

// SetCustomPackageId sets the value of CustomPackageId for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) SetPropertyCustomPackageId(value string) (err error) {
	return instance.SetProperty("CustomPackageId", value)
}

// GetCustomPackageId gets the value of CustomPackageId for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) GetPropertyCustomPackageId() (value string, err error) {
	retValue, err := instance.GetProperty("CustomPackageId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceModelId sets the value of DeviceModelId for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) SetPropertyDeviceModelId(value string) (err error) {
	return instance.SetProperty("DeviceModelId", value)
}

// GetDeviceModelId gets the value of DeviceModelId for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) GetPropertyDeviceModelId() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceModelId")
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
func (instance *MDM_DeviceUpdateCenter_Enrollment01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) GetPropertyInstanceID() (value string, err error) {
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

// SetOemPartnerRing sets the value of OemPartnerRing for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) SetPropertyOemPartnerRing(value string) (err error) {
	return instance.SetProperty("OemPartnerRing", value)
}

// GetOemPartnerRing gets the value of OemPartnerRing for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) GetPropertyOemPartnerRing() (value string, err error) {
	retValue, err := instance.GetProperty("OemPartnerRing")
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
func (instance *MDM_DeviceUpdateCenter_Enrollment01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) GetPropertyParentID() (value string, err error) {
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

// SetPublisherId sets the value of PublisherId for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) SetPropertyPublisherId(value string) (err error) {
	return instance.SetProperty("PublisherId", value)
}

// GetPublisherId gets the value of PublisherId for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) GetPropertyPublisherId() (value string, err error) {
	retValue, err := instance.GetProperty("PublisherId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
