// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DAEntryPointBase struct
type DAEntryPointBase struct {
	*cim.WmiInstance

	//
	EntryPointName string

	//
	Servers []string
}

func NewDAEntryPointBaseEx1(instance *cim.WmiInstance) (newInstance *DAEntryPointBase, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DAEntryPointBase{
		WmiInstance: tmp,
	}
	return
}

func NewDAEntryPointBaseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DAEntryPointBase, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DAEntryPointBase{
		WmiInstance: tmp,
	}
	return
}

// SetEntryPointName sets the value of EntryPointName for the instance
func (instance *DAEntryPointBase) SetPropertyEntryPointName(value string) (err error) {
	return instance.SetProperty("EntryPointName", (value))
}

// GetEntryPointName gets the value of EntryPointName for the instance
func (instance *DAEntryPointBase) GetPropertyEntryPointName() (value string, err error) {
	retValue, err := instance.GetProperty("EntryPointName")
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

// SetServers sets the value of Servers for the instance
func (instance *DAEntryPointBase) SetPropertyServers(value []string) (err error) {
	return instance.SetProperty("Servers", (value))
}

// GetServers gets the value of Servers for the instance
func (instance *DAEntryPointBase) GetPropertyServers() (value []string, err error) {
	retValue, err := instance.GetProperty("Servers")
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
