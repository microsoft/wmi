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

// TP_V2_CBCancel struct
type TP_V2_CBCancel struct {
	*ThreadPoolTrace_V2

	//
	CancelCount uint32

	//
	TaskId uint32
}

func NewTP_V2_CBCancelEx1(instance *cim.WmiInstance) (newInstance *TP_V2_CBCancel, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &TP_V2_CBCancel{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

func NewTP_V2_CBCancelEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TP_V2_CBCancel, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TP_V2_CBCancel{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

// SetCancelCount sets the value of CancelCount for the instance
func (instance *TP_V2_CBCancel) SetPropertyCancelCount(value uint32) (err error) {
	return instance.SetProperty("CancelCount", (value))
}

// GetCancelCount gets the value of CancelCount for the instance
func (instance *TP_V2_CBCancel) GetPropertyCancelCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("CancelCount")
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

// SetTaskId sets the value of TaskId for the instance
func (instance *TP_V2_CBCancel) SetPropertyTaskId(value uint32) (err error) {
	return instance.SetProperty("TaskId", (value))
}

// GetTaskId gets the value of TaskId for the instance
func (instance *TP_V2_CBCancel) GetPropertyTaskId() (value uint32, err error) {
	retValue, err := instance.GetProperty("TaskId")
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
