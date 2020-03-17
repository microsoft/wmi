// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

// ServerComponent_HyperV struct
type ServerComponent_HyperV struct {
	MSFT_ServerManagerServerComponentDescriptor

	//
	DefaultVirtualHardDiskPath string

	//
	DefaultVirtualMachinePath string

	//
	EnableVirtualMachineMigration bool

	//
	VirtualMachineMigrationAuthenticationType string

	//
	VirtualSwitchNetworkAdapters []string
}

// SetDefaultVirtualHardDiskPath sets the value of DefaultVirtualHardDiskPath for the instance
func (instance *ServerComponent_HyperV) SetPropertyDefaultVirtualHardDiskPath(value string) (err error) {
	return instance.SetProperty("DefaultVirtualHardDiskPath", value)
}

// GetDefaultVirtualHardDiskPath gets the value of DefaultVirtualHardDiskPath for the instance
func (instance *ServerComponent_HyperV) GetPropertyDefaultVirtualHardDiskPath() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultVirtualHardDiskPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultVirtualMachinePath sets the value of DefaultVirtualMachinePath for the instance
func (instance *ServerComponent_HyperV) SetPropertyDefaultVirtualMachinePath(value string) (err error) {
	return instance.SetProperty("DefaultVirtualMachinePath", value)
}

// GetDefaultVirtualMachinePath gets the value of DefaultVirtualMachinePath for the instance
func (instance *ServerComponent_HyperV) GetPropertyDefaultVirtualMachinePath() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultVirtualMachinePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableVirtualMachineMigration sets the value of EnableVirtualMachineMigration for the instance
func (instance *ServerComponent_HyperV) SetPropertyEnableVirtualMachineMigration(value bool) (err error) {
	return instance.SetProperty("EnableVirtualMachineMigration", value)
}

// GetEnableVirtualMachineMigration gets the value of EnableVirtualMachineMigration for the instance
func (instance *ServerComponent_HyperV) GetPropertyEnableVirtualMachineMigration() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableVirtualMachineMigration")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualMachineMigrationAuthenticationType sets the value of VirtualMachineMigrationAuthenticationType for the instance
func (instance *ServerComponent_HyperV) SetPropertyVirtualMachineMigrationAuthenticationType(value string) (err error) {
	return instance.SetProperty("VirtualMachineMigrationAuthenticationType", value)
}

// GetVirtualMachineMigrationAuthenticationType gets the value of VirtualMachineMigrationAuthenticationType for the instance
func (instance *ServerComponent_HyperV) GetPropertyVirtualMachineMigrationAuthenticationType() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualMachineMigrationAuthenticationType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSwitchNetworkAdapters sets the value of VirtualSwitchNetworkAdapters for the instance
func (instance *ServerComponent_HyperV) SetPropertyVirtualSwitchNetworkAdapters(value []string) (err error) {
	return instance.SetProperty("VirtualSwitchNetworkAdapters", value)
}

// GetVirtualSwitchNetworkAdapters gets the value of VirtualSwitchNetworkAdapters for the instance
func (instance *ServerComponent_HyperV) GetPropertyVirtualSwitchNetworkAdapters() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualSwitchNetworkAdapters")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
