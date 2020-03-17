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

// MsftUal_Dns struct
type MsftUal_Dns struct {
	cim.WmiInstance

	// The host name of the client. This is associated with IPAddress.
	HostName string

	// The IP address of a DNS client record. This is associated with hostname.
	IPAddress string

	// The date and time when a DNS client record was last seen in DNS.
	LastSeen string

	// The name of the software parent product, or product line, that is providing User Access Logging data. This is also associated with a RoleName, and a RoleGuid.
	ProductName string

	// The UAL assigned, or registered, GUID representing the server role, or installed product. When a role or product reports its UAL data, this GUID accompanies the payload.
	RoleGuid string

	// The name of the role, component, or sub-product that is providing User Access Logging data. This is also associated with a ProductName, and a RoleGuid.
	RoleName string
}

// SetHostName sets the value of HostName for the instance
func (instance *MsftUal_Dns) SetPropertyHostName(value string) (err error) {
	return instance.SetProperty("HostName", value)
}

// GetHostName gets the value of HostName for the instance
func (instance *MsftUal_Dns) GetPropertyHostName() (value string, err error) {
	retValue, err := instance.GetProperty("HostName")
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
func (instance *MsftUal_Dns) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", value)
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *MsftUal_Dns) GetPropertyIPAddress() (value string, err error) {
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

// SetLastSeen sets the value of LastSeen for the instance
func (instance *MsftUal_Dns) SetPropertyLastSeen(value string) (err error) {
	return instance.SetProperty("LastSeen", value)
}

// GetLastSeen gets the value of LastSeen for the instance
func (instance *MsftUal_Dns) GetPropertyLastSeen() (value string, err error) {
	retValue, err := instance.GetProperty("LastSeen")
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
func (instance *MsftUal_Dns) SetPropertyProductName(value string) (err error) {
	return instance.SetProperty("ProductName", value)
}

// GetProductName gets the value of ProductName for the instance
func (instance *MsftUal_Dns) GetPropertyProductName() (value string, err error) {
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
func (instance *MsftUal_Dns) SetPropertyRoleGuid(value string) (err error) {
	return instance.SetProperty("RoleGuid", value)
}

// GetRoleGuid gets the value of RoleGuid for the instance
func (instance *MsftUal_Dns) GetPropertyRoleGuid() (value string, err error) {
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
func (instance *MsftUal_Dns) SetPropertyRoleName(value string) (err error) {
	return instance.SetProperty("RoleName", value)
}

// GetRoleName gets the value of RoleName for the instance
func (instance *MsftUal_Dns) GetPropertyRoleName() (value string, err error) {
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
