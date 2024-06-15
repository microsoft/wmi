// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// TraceListener struct
type TraceListener struct {
	*cim.WmiInstance

	// The name of the trace listener.
	Name string

	// The arguments of the trace listener.
	TraceListenerArguments []TraceListenerArgument
}

func NewTraceListenerEx1(instance *cim.WmiInstance) (newInstance *TraceListener, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &TraceListener{
		WmiInstance: tmp,
	}
	return
}

func NewTraceListenerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TraceListener, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TraceListener{
		WmiInstance: tmp,
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *TraceListener) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *TraceListener) GetPropertyName() (value string, err error) {
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

// SetTraceListenerArguments sets the value of TraceListenerArguments for the instance
func (instance *TraceListener) SetPropertyTraceListenerArguments(value []TraceListenerArgument) (err error) {
	return instance.SetProperty("TraceListenerArguments", (value))
}

// GetTraceListenerArguments gets the value of TraceListenerArguments for the instance
func (instance *TraceListener) GetPropertyTraceListenerArguments() (value []TraceListenerArgument, err error) {
	retValue, err := instance.GetProperty("TraceListenerArguments")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(TraceListenerArgument)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " TraceListenerArgument is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, TraceListenerArgument(valuetmp))
	}

	return
}
