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

// CritSecInit struct
type CritSecInit struct {
	*CritSecTrace

	//
	CritSecAddr uint32

	//
	SpinCount uint32
}

func NewCritSecInitEx1(instance *cim.WmiInstance) (newInstance *CritSecInit, err error) {
	tmp, err := NewCritSecTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CritSecInit{
		CritSecTrace: tmp,
	}
	return
}

func NewCritSecInitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CritSecInit, err error) {
	tmp, err := NewCritSecTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CritSecInit{
		CritSecTrace: tmp,
	}
	return
}

// SetCritSecAddr sets the value of CritSecAddr for the instance
func (instance *CritSecInit) SetPropertyCritSecAddr(value uint32) (err error) {
	return instance.SetProperty("CritSecAddr", (value))
}

// GetCritSecAddr gets the value of CritSecAddr for the instance
func (instance *CritSecInit) GetPropertyCritSecAddr() (value uint32, err error) {
	retValue, err := instance.GetProperty("CritSecAddr")
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

// SetSpinCount sets the value of SpinCount for the instance
func (instance *CritSecInit) SetPropertySpinCount(value uint32) (err error) {
	return instance.SetProperty("SpinCount", (value))
}

// GetSpinCount gets the value of SpinCount for the instance
func (instance *CritSecInit) GetPropertySpinCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("SpinCount")
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
