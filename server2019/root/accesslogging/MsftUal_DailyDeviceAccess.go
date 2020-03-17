// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.AccessLogging
//////////////////////////////////////////////
package accesslogging

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MsftUal_DailyDeviceAccess struct
type MsftUal_DailyDeviceAccess struct {
	cim.WmiInstance

	// The number of accesses of a role, or installed product, on the local server from a unique client device.
	AccessCount uint32

	// The date that a device accessed a role, or installed product, on the local server.
	AccessDate string

	// The IP address of a client device that accompanies the UAL payload from installed roles and products.
	IPAddress string

	// The name of the software parent product, or product line, that is providing User Access Logging data. This is also associated with a RoleName, and a RoleGuid.
	ProductName string

	// The UAL assigned, or registered, GUID representing the server role, or installed product. When a role or product reports its UAL data, this GUID accompanies the payload.
	RoleGuid string

	// The name of the role, component, or sub-product that is providing User Access Logging data. This is also associated with a ProductName, and a RoleGuid.
	RoleName string
}

// SetAccessCount sets the value of AccessCount for the instance
func (instance *MsftUal_DailyDeviceAccess) SetPropertyAccessCount(value uint32) (err error) {
	return instance.SetProperty("AccessCount", value)
}

// GetAccessCount gets the value of AccessCount for the instance
func (instance *MsftUal_DailyDeviceAccess) GetPropertyAccessCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("AccessCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAccessDate sets the value of AccessDate for the instance
func (instance *MsftUal_DailyDeviceAccess) SetPropertyAccessDate(value string) (err error) {
	return instance.SetProperty("AccessDate", value)
}

// GetAccessDate gets the value of AccessDate for the instance
func (instance *MsftUal_DailyDeviceAccess) GetPropertyAccessDate() (value string, err error) {
	retValue, err := instance.GetProperty("AccessDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPAddress sets the value of IPAddress for the instance
func (instance *MsftUal_DailyDeviceAccess) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", value)
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *MsftUal_DailyDeviceAccess) GetPropertyIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProductName sets the value of ProductName for the instance
func (instance *MsftUal_DailyDeviceAccess) SetPropertyProductName(value string) (err error) {
	return instance.SetProperty("ProductName", value)
}

// GetProductName gets the value of ProductName for the instance
func (instance *MsftUal_DailyDeviceAccess) GetPropertyProductName() (value string, err error) {
	retValue, err := instance.GetProperty("ProductName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoleGuid sets the value of RoleGuid for the instance
func (instance *MsftUal_DailyDeviceAccess) SetPropertyRoleGuid(value string) (err error) {
	return instance.SetProperty("RoleGuid", value)
}

// GetRoleGuid gets the value of RoleGuid for the instance
func (instance *MsftUal_DailyDeviceAccess) GetPropertyRoleGuid() (value string, err error) {
	retValue, err := instance.GetProperty("RoleGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoleName sets the value of RoleName for the instance
func (instance *MsftUal_DailyDeviceAccess) SetPropertyRoleName(value string) (err error) {
	return instance.SetProperty("RoleName", value)
}

// GetRoleName gets the value of RoleName for the instance
func (instance *MsftUal_DailyDeviceAccess) GetPropertyRoleName() (value string, err error) {
	retValue, err := instance.GetProperty("RoleName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
