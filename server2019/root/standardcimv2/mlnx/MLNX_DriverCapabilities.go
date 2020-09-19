// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MLNX_DriverCapabilities struct
type MLNX_DriverCapabilities struct {
	*CIM_Capabilities

	//
	Name string

	//
	SystemName string
}

func NewMLNX_DriverCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MLNX_DriverCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

func NewMLNX_DriverCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DriverCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MLNX_DriverCapabilities) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MLNX_DriverCapabilities) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetSystemName sets the value of SystemName for the instance
func (instance *MLNX_DriverCapabilities) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", (value))
}

// GetSystemName gets the value of SystemName for the instance
func (instance *MLNX_DriverCapabilities) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
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
