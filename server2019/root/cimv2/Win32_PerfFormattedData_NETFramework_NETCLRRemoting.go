// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_NETFramework_NETCLRRemoting struct
type Win32_PerfFormattedData_NETFramework_NETCLRRemoting struct {
	Win32_PerfFormattedData

	//
	Channels uint32

	//
	ContextBoundClassesLoaded uint32

	//
	ContextBoundObjectsAllocPersec uint32

	//
	ContextProxies uint32

	//
	Contexts uint32

	//
	RemoteCallsPersec uint32

	//
	TotalRemoteCalls uint32
}

// SetChannels sets the value of Channels for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) SetPropertyChannels(value uint32) (err error) {
	return instance.SetProperty("Channels", value)
}

// GetChannels gets the value of Channels for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) GetPropertyChannels() (value uint32, err error) {
	retValue, err := instance.GetProperty("Channels")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContextBoundClassesLoaded sets the value of ContextBoundClassesLoaded for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) SetPropertyContextBoundClassesLoaded(value uint32) (err error) {
	return instance.SetProperty("ContextBoundClassesLoaded", value)
}

// GetContextBoundClassesLoaded gets the value of ContextBoundClassesLoaded for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) GetPropertyContextBoundClassesLoaded() (value uint32, err error) {
	retValue, err := instance.GetProperty("ContextBoundClassesLoaded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContextBoundObjectsAllocPersec sets the value of ContextBoundObjectsAllocPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) SetPropertyContextBoundObjectsAllocPersec(value uint32) (err error) {
	return instance.SetProperty("ContextBoundObjectsAllocPersec", value)
}

// GetContextBoundObjectsAllocPersec gets the value of ContextBoundObjectsAllocPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) GetPropertyContextBoundObjectsAllocPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ContextBoundObjectsAllocPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContextProxies sets the value of ContextProxies for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) SetPropertyContextProxies(value uint32) (err error) {
	return instance.SetProperty("ContextProxies", value)
}

// GetContextProxies gets the value of ContextProxies for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) GetPropertyContextProxies() (value uint32, err error) {
	retValue, err := instance.GetProperty("ContextProxies")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContexts sets the value of Contexts for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) SetPropertyContexts(value uint32) (err error) {
	return instance.SetProperty("Contexts", value)
}

// GetContexts gets the value of Contexts for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) GetPropertyContexts() (value uint32, err error) {
	retValue, err := instance.GetProperty("Contexts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteCallsPersec sets the value of RemoteCallsPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) SetPropertyRemoteCallsPersec(value uint32) (err error) {
	return instance.SetProperty("RemoteCallsPersec", value)
}

// GetRemoteCallsPersec gets the value of RemoteCallsPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) GetPropertyRemoteCallsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("RemoteCallsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalRemoteCalls sets the value of TotalRemoteCalls for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) SetPropertyTotalRemoteCalls(value uint32) (err error) {
	return instance.SetProperty("TotalRemoteCalls", value)
}

// GetTotalRemoteCalls gets the value of TotalRemoteCalls for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRRemoting) GetPropertyTotalRemoteCalls() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalRemoteCalls")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
