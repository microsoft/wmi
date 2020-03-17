// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

// ServerComponent_UpdateServices_Database struct
type ServerComponent_UpdateServices_Database struct {
	MSFT_ServerManagerServerComponentDescriptor

	//
	InstanceName string
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *ServerComponent_UpdateServices_Database) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", value)
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ServerComponent_UpdateServices_Database) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
