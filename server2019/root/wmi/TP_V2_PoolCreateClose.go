// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// TP_V2_PoolCreateClose struct
type TP_V2_PoolCreateClose struct {
	*ThreadPoolTrace_V2

	//
	PoolId uint32
}

func NewTP_V2_PoolCreateCloseEx1(instance *cim.WmiInstance) (newInstance *TP_V2_PoolCreateClose, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &TP_V2_PoolCreateClose{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

func NewTP_V2_PoolCreateCloseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TP_V2_PoolCreateClose, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TP_V2_PoolCreateClose{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

// SetPoolId sets the value of PoolId for the instance
func (instance *TP_V2_PoolCreateClose) SetPropertyPoolId(value uint32) (err error) {
	return instance.SetProperty("PoolId", (value))
}

// GetPoolId gets the value of PoolId for the instance
func (instance *TP_V2_PoolCreateClose) GetPropertyPoolId() (value uint32, err error) {
	retValue, err := instance.GetProperty("PoolId")
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
