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

// MsftUal_ServerDevice struct
type MsftUal_ServerDevice struct {
	cim.WmiInstance

	// The incremental counter of client device accesses for a particular client device.
	ActivityCount uint32

	// The unit identification for the local server.
	ChassisSerialNumber string

	// The date and time when a client IP address is first seen by a server.
	FirstSeen string

	// The IP address of the client that accompanies the UAL payload from installed roles and products.
	IPAddress string

	// The date and time when a client IP address is last seen by a server.
	LastSeen string

	// SMBIOS reported universally unique identifier for this server unit.
	UUID string
}

// SetActivityCount sets the value of ActivityCount for the instance
func (instance *MsftUal_ServerDevice) SetPropertyActivityCount(value uint32) (err error) {
	return instance.SetProperty("ActivityCount", value)
}

// GetActivityCount gets the value of ActivityCount for the instance
func (instance *MsftUal_ServerDevice) GetPropertyActivityCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActivityCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetChassisSerialNumber sets the value of ChassisSerialNumber for the instance
func (instance *MsftUal_ServerDevice) SetPropertyChassisSerialNumber(value string) (err error) {
	return instance.SetProperty("ChassisSerialNumber", value)
}

// GetChassisSerialNumber gets the value of ChassisSerialNumber for the instance
func (instance *MsftUal_ServerDevice) GetPropertyChassisSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("ChassisSerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFirstSeen sets the value of FirstSeen for the instance
func (instance *MsftUal_ServerDevice) SetPropertyFirstSeen(value string) (err error) {
	return instance.SetProperty("FirstSeen", value)
}

// GetFirstSeen gets the value of FirstSeen for the instance
func (instance *MsftUal_ServerDevice) GetPropertyFirstSeen() (value string, err error) {
	retValue, err := instance.GetProperty("FirstSeen")
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
func (instance *MsftUal_ServerDevice) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", value)
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *MsftUal_ServerDevice) GetPropertyIPAddress() (value string, err error) {
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
func (instance *MsftUal_ServerDevice) SetPropertyLastSeen(value string) (err error) {
	return instance.SetProperty("LastSeen", value)
}

// GetLastSeen gets the value of LastSeen for the instance
func (instance *MsftUal_ServerDevice) GetPropertyLastSeen() (value string, err error) {
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

// SetUUID sets the value of UUID for the instance
func (instance *MsftUal_ServerDevice) SetPropertyUUID(value string) (err error) {
	return instance.SetProperty("UUID", value)
}

// GetUUID gets the value of UUID for the instance
func (instance *MsftUal_ServerDevice) GetPropertyUUID() (value string, err error) {
	retValue, err := instance.GetProperty("UUID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
