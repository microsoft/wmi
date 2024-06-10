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

// TcpIp_SendIPV4 struct
type TcpIp_SendIPV4 struct {
	*TcpIp

	//
	connid uint32

	//
	daddr interface{}

	//
	dport interface{}

	//
	endtime uint32

	//
	PID uint32

	//
	saddr interface{}

	//
	seqnum uint32

	//
	size uint32

	//
	sport interface{}

	//
	startime uint32
}

func NewTcpIp_SendIPV4Ex1(instance *cim.WmiInstance) (newInstance *TcpIp_SendIPV4, err error) {
	tmp, err := NewTcpIpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TcpIp_SendIPV4{
		TcpIp: tmp,
	}
	return
}

func NewTcpIp_SendIPV4Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TcpIp_SendIPV4, err error) {
	tmp, err := NewTcpIpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TcpIp_SendIPV4{
		TcpIp: tmp,
	}
	return
}

// Setconnid sets the value of connid for the instance
func (instance *TcpIp_SendIPV4) SetPropertyconnid(value uint32) (err error) {
	return instance.SetProperty("connid", (value))
}

// Getconnid gets the value of connid for the instance
func (instance *TcpIp_SendIPV4) GetPropertyconnid() (value uint32, err error) {
	retValue, err := instance.GetProperty("connid")
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
func (instance *TcpIp_SendIPV4) SetPropertydaddr(value interface{}) (err error) {
	return instance.SetProperty("daddr", (value))
}

// Getdaddr gets the value of daddr for the instance
func (instance *TcpIp_SendIPV4) GetPropertydaddr() (value interface{}, err error) {
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
func (instance *TcpIp_SendIPV4) SetPropertydport(value interface{}) (err error) {
	return instance.SetProperty("dport", (value))
}

// Getdport gets the value of dport for the instance
func (instance *TcpIp_SendIPV4) GetPropertydport() (value interface{}, err error) {
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

// Setendtime sets the value of endtime for the instance
func (instance *TcpIp_SendIPV4) SetPropertyendtime(value uint32) (err error) {
	return instance.SetProperty("endtime", (value))
}

// Getendtime gets the value of endtime for the instance
func (instance *TcpIp_SendIPV4) GetPropertyendtime() (value uint32, err error) {
	retValue, err := instance.GetProperty("endtime")
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

// SetPID sets the value of PID for the instance
func (instance *TcpIp_SendIPV4) SetPropertyPID(value uint32) (err error) {
	return instance.SetProperty("PID", (value))
}

// GetPID gets the value of PID for the instance
func (instance *TcpIp_SendIPV4) GetPropertyPID() (value uint32, err error) {
	retValue, err := instance.GetProperty("PID")
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

// Setsaddr sets the value of saddr for the instance
func (instance *TcpIp_SendIPV4) SetPropertysaddr(value interface{}) (err error) {
	return instance.SetProperty("saddr", (value))
}

// Getsaddr gets the value of saddr for the instance
func (instance *TcpIp_SendIPV4) GetPropertysaddr() (value interface{}, err error) {
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

// Setseqnum sets the value of seqnum for the instance
func (instance *TcpIp_SendIPV4) SetPropertyseqnum(value uint32) (err error) {
	return instance.SetProperty("seqnum", (value))
}

// Getseqnum gets the value of seqnum for the instance
func (instance *TcpIp_SendIPV4) GetPropertyseqnum() (value uint32, err error) {
	retValue, err := instance.GetProperty("seqnum")
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

// Setsize sets the value of size for the instance
func (instance *TcpIp_SendIPV4) SetPropertysize(value uint32) (err error) {
	return instance.SetProperty("size", (value))
}

// Getsize gets the value of size for the instance
func (instance *TcpIp_SendIPV4) GetPropertysize() (value uint32, err error) {
	retValue, err := instance.GetProperty("size")
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

// Setsport sets the value of sport for the instance
func (instance *TcpIp_SendIPV4) SetPropertysport(value interface{}) (err error) {
	return instance.SetProperty("sport", (value))
}

// Getsport gets the value of sport for the instance
func (instance *TcpIp_SendIPV4) GetPropertysport() (value interface{}, err error) {
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

// Setstartime sets the value of startime for the instance
func (instance *TcpIp_SendIPV4) SetPropertystartime(value uint32) (err error) {
	return instance.SetProperty("startime", (value))
}

// Getstartime gets the value of startime for the instance
func (instance *TcpIp_SendIPV4) GetPropertystartime() (value uint32, err error) {
	retValue, err := instance.GetProperty("startime")
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
