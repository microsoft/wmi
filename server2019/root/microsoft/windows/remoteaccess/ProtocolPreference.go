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

// ProtocolPreference struct
type ProtocolPreference struct {
	*cim.WmiInstance

	//
	ProtocolMetric []ProtocolMetric
}

func NewProtocolPreferenceEx1(instance *cim.WmiInstance) (newInstance *ProtocolPreference, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ProtocolPreference{
		WmiInstance: tmp,
	}
	return
}

func NewProtocolPreferenceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProtocolPreference, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProtocolPreference{
		WmiInstance: tmp,
	}
	return
}

// SetProtocolMetric sets the value of ProtocolMetric for the instance
func (instance *ProtocolPreference) SetPropertyProtocolMetric(value []ProtocolMetric) (err error) {
	return instance.SetProperty("ProtocolMetric", (value))
}

// GetProtocolMetric gets the value of ProtocolMetric for the instance
func (instance *ProtocolPreference) GetPropertyProtocolMetric() (value []ProtocolMetric, err error) {
	retValue, err := instance.GetProperty("ProtocolMetric")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ProtocolMetric)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ProtocolMetric is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ProtocolMetric(valuetmp))
	}

	return
}
