// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Msft_WmiProvider_ComServerLoadOperationEvent struct
type Msft_WmiProvider_ComServerLoadOperationEvent struct {
	Msft_WmiProvider_OperationEvent

	//
	Clsid string

	//
	InProcServer bool

	//
	InProcServerPath string

	//
	LocalServer bool

	//
	LocalServerPath string

	//
	ServerName string
}

// SetClsid sets the value of Clsid for the instance
func (instance *Msft_WmiProvider_ComServerLoadOperationEvent) SetPropertyClsid(value string) (err error) {
	return instance.SetProperty("Clsid", value)
}

// GetClsid gets the value of Clsid for the instance
func (instance *Msft_WmiProvider_ComServerLoadOperationEvent) GetPropertyClsid() (value string, err error) {
	retValue, err := instance.GetProperty("Clsid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInProcServer sets the value of InProcServer for the instance
func (instance *Msft_WmiProvider_ComServerLoadOperationEvent) SetPropertyInProcServer(value bool) (err error) {
	return instance.SetProperty("InProcServer", value)
}

// GetInProcServer gets the value of InProcServer for the instance
func (instance *Msft_WmiProvider_ComServerLoadOperationEvent) GetPropertyInProcServer() (value bool, err error) {
	retValue, err := instance.GetProperty("InProcServer")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInProcServerPath sets the value of InProcServerPath for the instance
func (instance *Msft_WmiProvider_ComServerLoadOperationEvent) SetPropertyInProcServerPath(value string) (err error) {
	return instance.SetProperty("InProcServerPath", value)
}

// GetInProcServerPath gets the value of InProcServerPath for the instance
func (instance *Msft_WmiProvider_ComServerLoadOperationEvent) GetPropertyInProcServerPath() (value string, err error) {
	retValue, err := instance.GetProperty("InProcServerPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalServer sets the value of LocalServer for the instance
func (instance *Msft_WmiProvider_ComServerLoadOperationEvent) SetPropertyLocalServer(value bool) (err error) {
	return instance.SetProperty("LocalServer", value)
}

// GetLocalServer gets the value of LocalServer for the instance
func (instance *Msft_WmiProvider_ComServerLoadOperationEvent) GetPropertyLocalServer() (value bool, err error) {
	retValue, err := instance.GetProperty("LocalServer")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalServerPath sets the value of LocalServerPath for the instance
func (instance *Msft_WmiProvider_ComServerLoadOperationEvent) SetPropertyLocalServerPath(value string) (err error) {
	return instance.SetProperty("LocalServerPath", value)
}

// GetLocalServerPath gets the value of LocalServerPath for the instance
func (instance *Msft_WmiProvider_ComServerLoadOperationEvent) GetPropertyLocalServerPath() (value string, err error) {
	retValue, err := instance.GetProperty("LocalServerPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerName sets the value of ServerName for the instance
func (instance *Msft_WmiProvider_ComServerLoadOperationEvent) SetPropertyServerName(value string) (err error) {
	return instance.SetProperty("ServerName", value)
}

// GetServerName gets the value of ServerName for the instance
func (instance *Msft_WmiProvider_ComServerLoadOperationEvent) GetPropertyServerName() (value string, err error) {
	retValue, err := instance.GetProperty("ServerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
