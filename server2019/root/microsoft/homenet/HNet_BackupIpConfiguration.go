// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// HNet_BackupIpConfiguration struct
type HNet_BackupIpConfiguration struct {
	cim.WmiInstance

	//
	Connection HNet_Connection

	//
	DefaultGateway string

	//
	EnableDHCP uint32

	//
	IPAddress string

	//
	SubnetMask string
}

// SetConnection sets the value of Connection for the instance
func (instance *HNet_BackupIpConfiguration) SetPropertyConnection(value HNet_Connection) (err error) {
	return instance.SetProperty("Connection", value)
}

// GetConnection gets the value of Connection for the instance
func (instance *HNet_BackupIpConfiguration) GetPropertyConnection() (value HNet_Connection, err error) {
	retValue, err := instance.GetProperty("Connection")
	if err != nil {
		return
	}
	value, ok := retValue.(HNet_Connection)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultGateway sets the value of DefaultGateway for the instance
func (instance *HNet_BackupIpConfiguration) SetPropertyDefaultGateway(value string) (err error) {
	return instance.SetProperty("DefaultGateway", value)
}

// GetDefaultGateway gets the value of DefaultGateway for the instance
func (instance *HNet_BackupIpConfiguration) GetPropertyDefaultGateway() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultGateway")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableDHCP sets the value of EnableDHCP for the instance
func (instance *HNet_BackupIpConfiguration) SetPropertyEnableDHCP(value uint32) (err error) {
	return instance.SetProperty("EnableDHCP", value)
}

// GetEnableDHCP gets the value of EnableDHCP for the instance
func (instance *HNet_BackupIpConfiguration) GetPropertyEnableDHCP() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnableDHCP")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPAddress sets the value of IPAddress for the instance
func (instance *HNet_BackupIpConfiguration) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", value)
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *HNet_BackupIpConfiguration) GetPropertyIPAddress() (value string, err error) {
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

// SetSubnetMask sets the value of SubnetMask for the instance
func (instance *HNet_BackupIpConfiguration) SetPropertySubnetMask(value string) (err error) {
	return instance.SetProperty("SubnetMask", value)
}

// GetSubnetMask gets the value of SubnetMask for the instance
func (instance *HNet_BackupIpConfiguration) GetPropertySubnetMask() (value string, err error) {
	retValue, err := instance.GetProperty("SubnetMask")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
