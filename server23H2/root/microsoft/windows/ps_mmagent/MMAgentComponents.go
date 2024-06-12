// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.PS_MMAgent
//////////////////////////////////////////////
package ps_mmagent

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MMAgentComponents struct
type MMAgentComponents struct {
	*cim.WmiInstance

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

func NewMMAgentComponentsEx1(instance *cim.WmiInstance) (newInstance *MMAgentComponents, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MMAgentComponents{
		WmiInstance: tmp,
	}
	return
}

func NewMMAgentComponentsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MMAgentComponents, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MMAgentComponents{
		WmiInstance: tmp,
	}
	return
}

// SetApplicationLaunchPrefetching sets the value of ApplicationLaunchPrefetching for the instance
func (instance *MMAgentComponents) SetPropertyApplicationLaunchPrefetching(value bool) (err error) {
	return instance.SetProperty("ApplicationLaunchPrefetching", (value))
}

// GetApplicationLaunchPrefetching gets the value of ApplicationLaunchPrefetching for the instance
func (instance *MMAgentComponents) GetPropertyApplicationLaunchPrefetching() (value bool, err error) {
	retValue, err := instance.GetProperty("ApplicationLaunchPrefetching")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetApplicationPreLaunch sets the value of ApplicationPreLaunch for the instance
func (instance *MMAgentComponents) SetPropertyApplicationPreLaunch(value bool) (err error) {
	return instance.SetProperty("ApplicationPreLaunch", (value))
}

// GetApplicationPreLaunch gets the value of ApplicationPreLaunch for the instance
func (instance *MMAgentComponents) GetPropertyApplicationPreLaunch() (value bool, err error) {
	retValue, err := instance.GetProperty("ApplicationPreLaunch")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetMaxOperationAPIFiles sets the value of MaxOperationAPIFiles for the instance
func (instance *MMAgentComponents) SetPropertyMaxOperationAPIFiles(value uint32) (err error) {
	return instance.SetProperty("MaxOperationAPIFiles", (value))
}

// GetMaxOperationAPIFiles gets the value of MaxOperationAPIFiles for the instance
func (instance *MMAgentComponents) GetPropertyMaxOperationAPIFiles() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxOperationAPIFiles")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetMemoryCompression sets the value of MemoryCompression for the instance
func (instance *MMAgentComponents) SetPropertyMemoryCompression(value bool) (err error) {
	return instance.SetProperty("MemoryCompression", (value))
}

// GetMemoryCompression gets the value of MemoryCompression for the instance
func (instance *MMAgentComponents) GetPropertyMemoryCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("MemoryCompression")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetOperationAPI sets the value of OperationAPI for the instance
func (instance *MMAgentComponents) SetPropertyOperationAPI(value bool) (err error) {
	return instance.SetProperty("OperationAPI", (value))
}

// GetOperationAPI gets the value of OperationAPI for the instance
func (instance *MMAgentComponents) GetPropertyOperationAPI() (value bool, err error) {
	retValue, err := instance.GetProperty("OperationAPI")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetPageCombining sets the value of PageCombining for the instance
func (instance *MMAgentComponents) SetPropertyPageCombining(value bool) (err error) {
	return instance.SetProperty("PageCombining", (value))
}

// GetPageCombining gets the value of PageCombining for the instance
func (instance *MMAgentComponents) GetPropertyPageCombining() (value bool, err error) {
	retValue, err := instance.GetProperty("PageCombining")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
