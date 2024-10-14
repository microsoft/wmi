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

// TcpIp_V1_Fail struct
type TcpIp_V1_Fail struct {
	*TcpIp_V1

	//
	Proto uint32
}

func NewTcpIp_V1_FailEx1(instance *cim.WmiInstance) (newInstance *TcpIp_V1_Fail, err error) {
	tmp, err := NewTcpIp_V1Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &TcpIp_V1_Fail{
		TcpIp_V1: tmp,
	}
	return
}

func NewTcpIp_V1_FailEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TcpIp_V1_Fail, err error) {
	tmp, err := NewTcpIp_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TcpIp_V1_Fail{
		TcpIp_V1: tmp,
	}
	return
}

// SetProto sets the value of Proto for the instance
func (instance *TcpIp_V1_Fail) SetPropertyProto(value uint32) (err error) {
	return instance.SetProperty("Proto", (value))
}

// GetProto gets the value of Proto for the instance
func (instance *TcpIp_V1_Fail) GetPropertyProto() (value uint32, err error) {
	retValue, err := instance.GetProperty("Proto")
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
