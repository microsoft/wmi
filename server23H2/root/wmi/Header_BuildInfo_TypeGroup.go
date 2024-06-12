// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Header_BuildInfo_TypeGroup struct
type Header_BuildInfo_TypeGroup struct {
	*EventTraceEvent

	//
	BuildString string
}

func NewHeader_BuildInfo_TypeGroupEx1(instance *cim.WmiInstance) (newInstance *Header_BuildInfo_TypeGroup, err error) {
	tmp, err := NewEventTraceEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Header_BuildInfo_TypeGroup{
		EventTraceEvent: tmp,
	}
	return
}

func NewHeader_BuildInfo_TypeGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Header_BuildInfo_TypeGroup, err error) {
	tmp, err := NewEventTraceEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Header_BuildInfo_TypeGroup{
		EventTraceEvent: tmp,
	}
	return
}

// SetBuildString sets the value of BuildString for the instance
func (instance *Header_BuildInfo_TypeGroup) SetPropertyBuildString(value string) (err error) {
	return instance.SetProperty("BuildString", (value))
}

// GetBuildString gets the value of BuildString for the instance
func (instance *Header_BuildInfo_TypeGroup) GetPropertyBuildString() (value string, err error) {
	retValue, err := instance.GetProperty("BuildString")
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
