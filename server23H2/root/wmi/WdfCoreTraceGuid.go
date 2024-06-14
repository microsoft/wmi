// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WdfCoreTraceGuid struct
type WdfCoreTraceGuid struct {
	*EventTrace

	//
	Flags uint32

	//
	Level uint32
}

func NewWdfCoreTraceGuidEx1(instance *cim.WmiInstance) (newInstance *WdfCoreTraceGuid, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WdfCoreTraceGuid{
		EventTrace: tmp,
	}
	return
}

func NewWdfCoreTraceGuidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WdfCoreTraceGuid, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WdfCoreTraceGuid{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *WdfCoreTraceGuid) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *WdfCoreTraceGuid) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetLevel sets the value of Level for the instance
func (instance *WdfCoreTraceGuid) SetPropertyLevel(value uint32) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *WdfCoreTraceGuid) GetPropertyLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("Level")
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
