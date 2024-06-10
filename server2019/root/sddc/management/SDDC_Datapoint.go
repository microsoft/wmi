// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.SDDC.Management
//
// ////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SDDC_Datapoint struct
type SDDC_Datapoint struct {
	*cim.WmiInstance

	//
	Timestamp string

	//
	Value float64
}

func NewSDDC_DatapointEx1(instance *cim.WmiInstance) (newInstance *SDDC_Datapoint, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_Datapoint{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_DatapointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_Datapoint, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_Datapoint{
		WmiInstance: tmp,
	}
	return
}

// SetTimestamp sets the value of Timestamp for the instance
func (instance *SDDC_Datapoint) SetPropertyTimestamp(value string) (err error) {
	return instance.SetProperty("Timestamp", (value))
}

// GetTimestamp gets the value of Timestamp for the instance
func (instance *SDDC_Datapoint) GetPropertyTimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("Timestamp")
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

// SetValue sets the value of Value for the instance
func (instance *SDDC_Datapoint) SetPropertyValue(value float64) (err error) {
	return instance.SetProperty("Value", (value))
}

// GetValue gets the value of Value for the instance
func (instance *SDDC_Datapoint) GetPropertyValue() (value float64, err error) {
	retValue, err := instance.GetProperty("Value")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float64(valuetmp)

	return
}
