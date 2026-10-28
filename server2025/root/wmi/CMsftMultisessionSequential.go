// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CMsftMultisessionSequential struct
type CMsftMultisessionSequential struct {
	*EventTrace

	// Enable Flags
	Flags CMsftMultisessionSequential_Flags

	// Levels
	Level CMsftMultisessionSequential_Level
}

func NewCMsftMultisessionSequentialEx1(instance *cim.WmiInstance) (newInstance *CMsftMultisessionSequential, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CMsftMultisessionSequential{
		EventTrace: tmp,
	}
	return
}

func NewCMsftMultisessionSequentialEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CMsftMultisessionSequential, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CMsftMultisessionSequential{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *CMsftMultisessionSequential) SetPropertyFlags(value CMsftMultisessionSequential_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *CMsftMultisessionSequential) GetPropertyFlags() (value CMsftMultisessionSequential_Flags, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CMsftMultisessionSequential_Flags(valuetmp)

	return
}

// SetLevel sets the value of Level for the instance
func (instance *CMsftMultisessionSequential) SetPropertyLevel(value CMsftMultisessionSequential_Level) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *CMsftMultisessionSequential) GetPropertyLevel() (value CMsftMultisessionSequential_Level, err error) {
	retValue, err := instance.GetProperty("Level")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CMsftMultisessionSequential_Level(valuetmp)

	return
}
