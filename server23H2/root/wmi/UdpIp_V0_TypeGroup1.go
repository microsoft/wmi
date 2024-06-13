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

// UdpIp_V0_TypeGroup1 struct
type UdpIp_V0_TypeGroup1 struct {
	*UdpIp_V0

	//
	context uint32

	//
	daddr interface{}

	//
	dport interface{}

	//
	dsize uint16

	//
	saddr interface{}

	//
	size uint16

	//
	sport interface{}
}

func NewUdpIp_V0_TypeGroup1Ex1(instance *cim.WmiInstance) (newInstance *UdpIp_V0_TypeGroup1, err error) {
	tmp, err := NewUdpIp_V0Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &UdpIp_V0_TypeGroup1{
		UdpIp_V0: tmp,
	}
	return
}

func NewUdpIp_V0_TypeGroup1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *UdpIp_V0_TypeGroup1, err error) {
	tmp, err := NewUdpIp_V0Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &UdpIp_V0_TypeGroup1{
		UdpIp_V0: tmp,
	}
	return
}

// Setcontext sets the value of context for the instance
func (instance *UdpIp_V0_TypeGroup1) SetPropertycontext(value uint32) (err error) {
	return instance.SetProperty("context", (value))
}

// Getcontext gets the value of context for the instance
func (instance *UdpIp_V0_TypeGroup1) GetPropertycontext() (value uint32, err error) {
	retValue, err := instance.GetProperty("context")
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

// Setdaddr sets the value of daddr for the instance
func (instance *UdpIp_V0_TypeGroup1) SetPropertydaddr(value interface{}) (err error) {
	return instance.SetProperty("daddr", (value))
}

// Getdaddr gets the value of daddr for the instance
func (instance *UdpIp_V0_TypeGroup1) GetPropertydaddr() (value interface{}, err error) {
	retValue, err := instance.GetProperty("daddr")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// Setdport sets the value of dport for the instance
func (instance *UdpIp_V0_TypeGroup1) SetPropertydport(value interface{}) (err error) {
	return instance.SetProperty("dport", (value))
}

// Getdport gets the value of dport for the instance
func (instance *UdpIp_V0_TypeGroup1) GetPropertydport() (value interface{}, err error) {
	retValue, err := instance.GetProperty("dport")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// Setdsize sets the value of dsize for the instance
func (instance *UdpIp_V0_TypeGroup1) SetPropertydsize(value uint16) (err error) {
	return instance.SetProperty("dsize", (value))
}

// Getdsize gets the value of dsize for the instance
func (instance *UdpIp_V0_TypeGroup1) GetPropertydsize() (value uint16, err error) {
	retValue, err := instance.GetProperty("dsize")
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

// Setsaddr sets the value of saddr for the instance
func (instance *UdpIp_V0_TypeGroup1) SetPropertysaddr(value interface{}) (err error) {
	return instance.SetProperty("saddr", (value))
}

// Getsaddr gets the value of saddr for the instance
func (instance *UdpIp_V0_TypeGroup1) GetPropertysaddr() (value interface{}, err error) {
	retValue, err := instance.GetProperty("saddr")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// Setsize sets the value of size for the instance
func (instance *UdpIp_V0_TypeGroup1) SetPropertysize(value uint16) (err error) {
	return instance.SetProperty("size", (value))
}

// Getsize gets the value of size for the instance
func (instance *UdpIp_V0_TypeGroup1) GetPropertysize() (value uint16, err error) {
	retValue, err := instance.GetProperty("size")
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

// Setsport sets the value of sport for the instance
func (instance *UdpIp_V0_TypeGroup1) SetPropertysport(value interface{}) (err error) {
	return instance.SetProperty("sport", (value))
}

// Getsport gets the value of sport for the instance
func (instance *UdpIp_V0_TypeGroup1) GetPropertysport() (value interface{}, err error) {
	retValue, err := instance.GetProperty("sport")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
