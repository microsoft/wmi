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

// TcpIp_Fail struct
type TcpIp_Fail struct {
	*TcpIp

	//
	FailureCode uint16

	//
	Proto uint16
}

func NewTcpIp_FailEx1(instance *cim.WmiInstance) (newInstance *TcpIp_Fail, err error) {
	tmp, err := NewTcpIpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TcpIp_Fail{
		TcpIp: tmp,
	}
	return
}

func NewTcpIp_FailEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TcpIp_Fail, err error) {
	tmp, err := NewTcpIpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TcpIp_Fail{
		TcpIp: tmp,
	}
	return
}

// SetFailureCode sets the value of FailureCode for the instance
func (instance *TcpIp_Fail) SetPropertyFailureCode(value uint16) (err error) {
	return instance.SetProperty("FailureCode", (value))
}

// GetFailureCode gets the value of FailureCode for the instance
func (instance *TcpIp_Fail) GetPropertyFailureCode() (value uint16, err error) {
	retValue, err := instance.GetProperty("FailureCode")
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

// SetProto sets the value of Proto for the instance
func (instance *TcpIp_Fail) SetPropertyProto(value uint16) (err error) {
	return instance.SetProperty("Proto", (value))
}

// GetProto gets the value of Proto for the instance
func (instance *TcpIp_Fail) GetPropertyProto() (value uint16, err error) {
	retValue, err := instance.GetProperty("Proto")
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
