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

// TSRdpEndPointTrace struct
type TSRdpEndPointTrace struct {
	*EventTrace

	// Enable Flags
	Flags TSRdpEndPointTrace_Flags
}

func NewTSRdpEndPointTraceEx1(instance *cim.WmiInstance) (newInstance *TSRdpEndPointTrace, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TSRdpEndPointTrace{
		EventTrace: tmp,
	}
	return
}

func NewTSRdpEndPointTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TSRdpEndPointTrace, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TSRdpEndPointTrace{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *TSRdpEndPointTrace) SetPropertyFlags(value TSRdpEndPointTrace_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *TSRdpEndPointTrace) GetPropertyFlags() (value TSRdpEndPointTrace_Flags, err error) {
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

	value = TSRdpEndPointTrace_Flags(valuetmp)

	return
}
