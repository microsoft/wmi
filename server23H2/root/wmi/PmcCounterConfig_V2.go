// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PmcCounterConfig_V2 struct
type PmcCounterConfig_V2 struct {
	*PerfInfo_V2

	//
	CounterCount uint32

	//
	CounterName []string
}

func NewPmcCounterConfig_V2Ex1(instance *cim.WmiInstance) (newInstance *PmcCounterConfig_V2, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &PmcCounterConfig_V2{
		PerfInfo_V2: tmp,
	}
	return
}

func NewPmcCounterConfig_V2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PmcCounterConfig_V2, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PmcCounterConfig_V2{
		PerfInfo_V2: tmp,
	}
	return
}

// SetCounterCount sets the value of CounterCount for the instance
func (instance *PmcCounterConfig_V2) SetPropertyCounterCount(value uint32) (err error) {
	return instance.SetProperty("CounterCount", (value))
}

// GetCounterCount gets the value of CounterCount for the instance
func (instance *PmcCounterConfig_V2) GetPropertyCounterCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("CounterCount")
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

// SetCounterName sets the value of CounterName for the instance
func (instance *PmcCounterConfig_V2) SetPropertyCounterName(value []string) (err error) {
	return instance.SetProperty("CounterName", (value))
}

// GetCounterName gets the value of CounterName for the instance
func (instance *PmcCounterConfig_V2) GetPropertyCounterName() (value []string, err error) {
	retValue, err := instance.GetProperty("CounterName")
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
