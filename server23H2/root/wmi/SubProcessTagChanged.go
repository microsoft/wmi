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

// SubProcessTagChanged struct
type SubProcessTagChanged struct {
	*Thread_V2

	//
	NewTag uint32

	//
	OldTag uint32
}

func NewSubProcessTagChangedEx1(instance *cim.WmiInstance) (newInstance *SubProcessTagChanged, err error) {
	tmp, err := NewThread_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SubProcessTagChanged{
		Thread_V2: tmp,
	}
	return
}

func NewSubProcessTagChangedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SubProcessTagChanged, err error) {
	tmp, err := NewThread_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SubProcessTagChanged{
		Thread_V2: tmp,
	}
	return
}

// SetNewTag sets the value of NewTag for the instance
func (instance *SubProcessTagChanged) SetPropertyNewTag(value uint32) (err error) {
	return instance.SetProperty("NewTag", (value))
}

// GetNewTag gets the value of NewTag for the instance
func (instance *SubProcessTagChanged) GetPropertyNewTag() (value uint32, err error) {
	retValue, err := instance.GetProperty("NewTag")
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

// SetOldTag sets the value of OldTag for the instance
func (instance *SubProcessTagChanged) SetPropertyOldTag(value uint32) (err error) {
	return instance.SetProperty("OldTag", (value))
}

// GetOldTag gets the value of OldTag for the instance
func (instance *SubProcessTagChanged) GetPropertyOldTag() (value uint32, err error) {
	retValue, err := instance.GetProperty("OldTag")
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
