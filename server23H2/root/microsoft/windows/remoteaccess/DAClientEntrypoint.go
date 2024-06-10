// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DAClientEntrypoint struct
type DAClientEntrypoint struct {
	*cim.WmiInstance

	//
	DownlevelGpoName []string

	//
	DownlevelSecurityGroupNameList []string

	//
	EntrypointName string
}

func NewDAClientEntrypointEx1(instance *cim.WmiInstance) (newInstance *DAClientEntrypoint, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DAClientEntrypoint{
		WmiInstance: tmp,
	}
	return
}

func NewDAClientEntrypointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DAClientEntrypoint, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DAClientEntrypoint{
		WmiInstance: tmp,
	}
	return
}

// SetDownlevelGpoName sets the value of DownlevelGpoName for the instance
func (instance *DAClientEntrypoint) SetPropertyDownlevelGpoName(value []string) (err error) {
	return instance.SetProperty("DownlevelGpoName", (value))
}

// GetDownlevelGpoName gets the value of DownlevelGpoName for the instance
func (instance *DAClientEntrypoint) GetPropertyDownlevelGpoName() (value []string, err error) {
	retValue, err := instance.GetProperty("DownlevelGpoName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDownlevelSecurityGroupNameList sets the value of DownlevelSecurityGroupNameList for the instance
func (instance *DAClientEntrypoint) SetPropertyDownlevelSecurityGroupNameList(value []string) (err error) {
	return instance.SetProperty("DownlevelSecurityGroupNameList", (value))
}

// GetDownlevelSecurityGroupNameList gets the value of DownlevelSecurityGroupNameList for the instance
func (instance *DAClientEntrypoint) GetPropertyDownlevelSecurityGroupNameList() (value []string, err error) {
	retValue, err := instance.GetProperty("DownlevelSecurityGroupNameList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetEntrypointName sets the value of EntrypointName for the instance
func (instance *DAClientEntrypoint) SetPropertyEntrypointName(value string) (err error) {
	return instance.SetProperty("EntrypointName", (value))
}

// GetEntrypointName gets the value of EntrypointName for the instance
func (instance *DAClientEntrypoint) GetPropertyEntrypointName() (value string, err error) {
	retValue, err := instance.GetProperty("EntrypointName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
