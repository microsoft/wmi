// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_ServicingSettings struct
type Msvm_ServicingSettings struct {
	*cim.WmiInstance

	//
	Version string
}

func NewMsvm_ServicingSettingsEx1(instance *cim.WmiInstance) (newInstance *Msvm_ServicingSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Msvm_ServicingSettings{
		WmiInstance: tmp,
	}
	return
}

func NewMsvm_ServicingSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ServicingSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ServicingSettings{
		WmiInstance: tmp,
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *Msvm_ServicingSettings) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *Msvm_ServicingSettings) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
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
