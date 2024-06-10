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
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DAEntryPoint struct
type DAEntryPoint struct {
	*DAEntryPointBase

	//
	GslbIp string
}

func NewDAEntryPointEx1(instance *cim.WmiInstance) (newInstance *DAEntryPoint, err error) {
	tmp, err := NewDAEntryPointBaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DAEntryPoint{
		DAEntryPointBase: tmp,
	}
	return
}

func NewDAEntryPointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DAEntryPoint, err error) {
	tmp, err := NewDAEntryPointBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DAEntryPoint{
		DAEntryPointBase: tmp,
	}
	return
}

// SetGslbIp sets the value of GslbIp for the instance
func (instance *DAEntryPoint) SetPropertyGslbIp(value string) (err error) {
	return instance.SetProperty("GslbIp", (value))
}

// GetGslbIp gets the value of GslbIp for the instance
func (instance *DAEntryPoint) GetPropertyGslbIp() (value string, err error) {
	retValue, err := instance.GetProperty("GslbIp")
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
