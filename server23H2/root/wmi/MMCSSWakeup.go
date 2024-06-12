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

// MMCSSWakeup struct
type MMCSSWakeup struct {
	*MMCSSTrace

	//
	Reason uint32
}

func NewMMCSSWakeupEx1(instance *cim.WmiInstance) (newInstance *MMCSSWakeup, err error) {
	tmp, err := NewMMCSSTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MMCSSWakeup{
		MMCSSTrace: tmp,
	}
	return
}

func NewMMCSSWakeupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MMCSSWakeup, err error) {
	tmp, err := NewMMCSSTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MMCSSWakeup{
		MMCSSTrace: tmp,
	}
	return
}

// SetReason sets the value of Reason for the instance
func (instance *MMCSSWakeup) SetPropertyReason(value uint32) (err error) {
	return instance.SetProperty("Reason", (value))
}

// GetReason gets the value of Reason for the instance
func (instance *MMCSSWakeup) GetPropertyReason() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reason")
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
