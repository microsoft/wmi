// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Msft_WmiProvider_LoadOperationEvent struct
type Msft_WmiProvider_LoadOperationEvent struct {
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

	//
	Synchronisation uint32

	//
	ThreadingModel uint32
}

// SetClsid sets the value of Clsid for the instance
func (instance *Msft_WmiProvider_LoadOperationEvent) SetPropertyClsid(value string) (err error) {
	return instance.SetProperty("Clsid", value)
}

// GetClsid gets the value of Clsid for the instance
func (instance *Msft_WmiProvider_LoadOperationEvent) GetPropertyClsid() (value string, err error) {
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
func (instance *Msft_WmiProvider_LoadOperationEvent) SetPropertyInProcServer(value bool) (err error) {
	return instance.SetProperty("InProcServer", value)
}

// GetInProcServer gets the value of InProcServer for the instance
func (instance *Msft_WmiProvider_LoadOperationEvent) GetPropertyInProcServer() (value bool, err error) {
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
func (instance *Msft_WmiProvider_LoadOperationEvent) SetPropertyInProcServerPath(value string) (err error) {
	return instance.SetProperty("InProcServerPath", value)
}

// GetInProcServerPath gets the value of InProcServerPath for the instance
func (instance *Msft_WmiProvider_LoadOperationEvent) GetPropertyInProcServerPath() (value string, err error) {
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
func (instance *Msft_WmiProvider_LoadOperationEvent) SetPropertyLocalServer(value bool) (err error) {
	return instance.SetProperty("LocalServer", value)
}

// GetLocalServer gets the value of LocalServer for the instance
func (instance *Msft_WmiProvider_LoadOperationEvent) GetPropertyLocalServer() (value bool, err error) {
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
func (instance *Msft_WmiProvider_LoadOperationEvent) SetPropertyLocalServerPath(value string) (err error) {
	return instance.SetProperty("LocalServerPath", value)
}

// GetLocalServerPath gets the value of LocalServerPath for the instance
func (instance *Msft_WmiProvider_LoadOperationEvent) GetPropertyLocalServerPath() (value string, err error) {
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
func (instance *Msft_WmiProvider_LoadOperationEvent) SetPropertyServerName(value string) (err error) {
	return instance.SetProperty("ServerName", value)
}

// GetServerName gets the value of ServerName for the instance
func (instance *Msft_WmiProvider_LoadOperationEvent) GetPropertyServerName() (value string, err error) {
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

// SetSynchronisation sets the value of Synchronisation for the instance
func (instance *Msft_WmiProvider_LoadOperationEvent) SetPropertySynchronisation(value uint32) (err error) {
	return instance.SetProperty("Synchronisation", value)
}

// GetSynchronisation gets the value of Synchronisation for the instance
func (instance *Msft_WmiProvider_LoadOperationEvent) GetPropertySynchronisation() (value uint32, err error) {
	retValue, err := instance.GetProperty("Synchronisation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreadingModel sets the value of ThreadingModel for the instance
func (instance *Msft_WmiProvider_LoadOperationEvent) SetPropertyThreadingModel(value uint32) (err error) {
	return instance.SetProperty("ThreadingModel", value)
}

// GetThreadingModel gets the value of ThreadingModel for the instance
func (instance *Msft_WmiProvider_LoadOperationEvent) GetPropertyThreadingModel() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadingModel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
