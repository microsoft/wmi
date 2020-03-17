// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.PS_MMAgent
//////////////////////////////////////////////
package ps_mmagent

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MMAgentComponents struct
type MMAgentComponents struct {
	cim.WmiInstance

	//
	ApplicationLaunchPrefetching bool

	//
	ApplicationPreLaunch bool

	//
	MaxOperationAPIFiles uint32

	//
	MemoryCompression bool

	//
	OperationAPI bool

	//
	PageCombining bool
}

// SetApplicationLaunchPrefetching sets the value of ApplicationLaunchPrefetching for the instance
func (instance *MMAgentComponents) SetPropertyApplicationLaunchPrefetching(value bool) (err error) {
	return instance.SetProperty("ApplicationLaunchPrefetching", value)
}

// GetApplicationLaunchPrefetching gets the value of ApplicationLaunchPrefetching for the instance
func (instance *MMAgentComponents) GetPropertyApplicationLaunchPrefetching() (value bool, err error) {
	retValue, err := instance.GetProperty("ApplicationLaunchPrefetching")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetApplicationPreLaunch sets the value of ApplicationPreLaunch for the instance
func (instance *MMAgentComponents) SetPropertyApplicationPreLaunch(value bool) (err error) {
	return instance.SetProperty("ApplicationPreLaunch", value)
}

// GetApplicationPreLaunch gets the value of ApplicationPreLaunch for the instance
func (instance *MMAgentComponents) GetPropertyApplicationPreLaunch() (value bool, err error) {
	retValue, err := instance.GetProperty("ApplicationPreLaunch")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxOperationAPIFiles sets the value of MaxOperationAPIFiles for the instance
func (instance *MMAgentComponents) SetPropertyMaxOperationAPIFiles(value uint32) (err error) {
	return instance.SetProperty("MaxOperationAPIFiles", value)
}

// GetMaxOperationAPIFiles gets the value of MaxOperationAPIFiles for the instance
func (instance *MMAgentComponents) GetPropertyMaxOperationAPIFiles() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxOperationAPIFiles")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMemoryCompression sets the value of MemoryCompression for the instance
func (instance *MMAgentComponents) SetPropertyMemoryCompression(value bool) (err error) {
	return instance.SetProperty("MemoryCompression", value)
}

// GetMemoryCompression gets the value of MemoryCompression for the instance
func (instance *MMAgentComponents) GetPropertyMemoryCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("MemoryCompression")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationAPI sets the value of OperationAPI for the instance
func (instance *MMAgentComponents) SetPropertyOperationAPI(value bool) (err error) {
	return instance.SetProperty("OperationAPI", value)
}

// GetOperationAPI gets the value of OperationAPI for the instance
func (instance *MMAgentComponents) GetPropertyOperationAPI() (value bool, err error) {
	retValue, err := instance.GetProperty("OperationAPI")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPageCombining sets the value of PageCombining for the instance
func (instance *MMAgentComponents) SetPropertyPageCombining(value bool) (err error) {
	return instance.SetProperty("PageCombining", value)
}

// GetPageCombining gets the value of PageCombining for the instance
func (instance *MMAgentComponents) GetPropertyPageCombining() (value bool, err error) {
	retValue, err := instance.GetProperty("PageCombining")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
