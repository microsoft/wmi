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

// TcpIp_V1_Receive struct
type TcpIp_V1_Receive struct {
	*TcpIp_V1

	//
	connid uint32

	//
	daddr interface{}

	//
	dport interface{}

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
}

func NewTcpIp_V1_ReceiveEx1(instance *cim.WmiInstance) (newInstance *TcpIp_V1_Receive, err error) {
	tmp, err := NewTcpIp_V1Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &TcpIp_V1_Receive{
		TcpIp_V1: tmp,
	}
	return
}

func NewTcpIp_V1_ReceiveEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TcpIp_V1_Receive, err error) {
	tmp, err := NewTcpIp_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TcpIp_V1_Receive{
		TcpIp_V1: tmp,
	}
	return
}

// Setconnid sets the value of connid for the instance
func (instance *TcpIp_V1_Receive) SetPropertyconnid(value uint32) (err error) {
	return instance.SetProperty("connid", (value))
}

// Getconnid gets the value of connid for the instance
func (instance *TcpIp_V1_Receive) GetPropertyconnid() (value uint32, err error) {
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
func (instance *TcpIp_V1_Receive) SetPropertydaddr(value interface{}) (err error) {
	return instance.SetProperty("daddr", (value))
}

// Getdaddr gets the value of daddr for the instance
func (instance *TcpIp_V1_Receive) GetPropertydaddr() (value interface{}, err error) {
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
func (instance *TcpIp_V1_Receive) SetPropertydport(value interface{}) (err error) {
	return instance.SetProperty("dport", (value))
}

// Getdport gets the value of dport for the instance
func (instance *TcpIp_V1_Receive) GetPropertydport() (value interface{}, err error) {
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

// SetPID sets the value of PID for the instance
func (instance *TcpIp_V1_Receive) SetPropertyPID(value uint32) (err error) {
	return instance.SetProperty("PID", (value))
}

// GetPID gets the value of PID for the instance
func (instance *TcpIp_V1_Receive) GetPropertyPID() (value uint32, err error) {
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
func (instance *TcpIp_V1_Receive) SetPropertysaddr(value interface{}) (err error) {
	return instance.SetProperty("saddr", (value))
}

// Getsaddr gets the value of saddr for the instance
func (instance *TcpIp_V1_Receive) GetPropertysaddr() (value interface{}, err error) {
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
func (instance *TcpIp_V1_Receive) SetPropertyseqnum(value uint32) (err error) {
	return instance.SetProperty("seqnum", (value))
}

// Getseqnum gets the value of seqnum for the instance
func (instance *TcpIp_V1_Receive) GetPropertyseqnum() (value uint32, err error) {
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
func (instance *TcpIp_V1_Receive) SetPropertysize(value uint32) (err error) {
	return instance.SetProperty("size", (value))
}

// Getsize gets the value of size for the instance
func (instance *TcpIp_V1_Receive) GetPropertysize() (value uint32, err error) {
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
func (instance *TcpIp_V1_Receive) SetPropertysport(value interface{}) (err error) {
	return instance.SetProperty("sport", (value))
}

// Getsport gets the value of sport for the instance
func (instance *TcpIp_V1_Receive) GetPropertysport() (value interface{}, err error) {
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
