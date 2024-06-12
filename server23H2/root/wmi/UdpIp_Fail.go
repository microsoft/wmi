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

// UdpIp_Fail struct
type UdpIp_Fail struct {
	*UdpIp

	//
	FailureCode uint16

	//
	Proto uint16
}

func NewUdpIp_FailEx1(instance *cim.WmiInstance) (newInstance *UdpIp_Fail, err error) {
	tmp, err := NewUdpIpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &UdpIp_Fail{
		UdpIp: tmp,
	}
	return
}

func NewUdpIp_FailEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *UdpIp_Fail, err error) {
	tmp, err := NewUdpIpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &UdpIp_Fail{
		UdpIp: tmp,
	}
	return
}

// SetFailureCode sets the value of FailureCode for the instance
func (instance *UdpIp_Fail) SetPropertyFailureCode(value uint16) (err error) {
	return instance.SetProperty("FailureCode", (value))
}

// GetFailureCode gets the value of FailureCode for the instance
func (instance *UdpIp_Fail) GetPropertyFailureCode() (value uint16, err error) {
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
func (instance *UdpIp_Fail) SetPropertyProto(value uint16) (err error) {
	return instance.SetProperty("Proto", (value))
}

// GetProto gets the value of Proto for the instance
func (instance *UdpIp_Fail) GetPropertyProto() (value uint16, err error) {
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
