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

// ISR_Unexpected struct
type ISR_Unexpected struct {
	*PerfInfo_V2

	//
	Vector uint16
}

func NewISR_UnexpectedEx1(instance *cim.WmiInstance) (newInstance *ISR_Unexpected, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &ISR_Unexpected{
		PerfInfo_V2: tmp,
	}
	return
}

func NewISR_UnexpectedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ISR_Unexpected, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ISR_Unexpected{
		PerfInfo_V2: tmp,
	}
	return
}

// SetVector sets the value of Vector for the instance
func (instance *ISR_Unexpected) SetPropertyVector(value uint16) (err error) {
	return instance.SetProperty("Vector", (value))
}

// GetVector gets the value of Vector for the instance
func (instance *ISR_Unexpected) GetPropertyVector() (value uint16, err error) {
	retValue, err := instance.GetProperty("Vector")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
