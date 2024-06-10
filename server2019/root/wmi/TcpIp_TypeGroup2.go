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

// TcpIp_TypeGroup2 struct
type TcpIp_TypeGroup2 struct {
	*TcpIp

	//
	connid uint32

	//
	daddr interface{}

	//
	dport interface{}

	//
	mss uint16

	//
	PID uint32

	//
	rcvwin uint32

	//
	rcvwinscale int16

	//
	sackopt uint16

	//
	saddr interface{}

	//
	seqnum uint32

	//
	size uint32

	//
	sndwinscale int16

	//
	sport interface{}

	//
	tsopt uint16

	//
	wsopt uint16
}

func NewTcpIp_TypeGroup2Ex1(instance *cim.WmiInstance) (newInstance *TcpIp_TypeGroup2, err error) {
	tmp, err := NewTcpIpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TcpIp_TypeGroup2{
		TcpIp: tmp,
	}
	return
}

func NewTcpIp_TypeGroup2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TcpIp_TypeGroup2, err error) {
	tmp, err := NewTcpIpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TcpIp_TypeGroup2{
		TcpIp: tmp,
	}
	return
}

// Setconnid sets the value of connid for the instance
func (instance *TcpIp_TypeGroup2) SetPropertyconnid(value uint32) (err error) {
	return instance.SetProperty("connid", (value))
}

// Getconnid gets the value of connid for the instance
func (instance *TcpIp_TypeGroup2) GetPropertyconnid() (value uint32, err error) {
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
func (instance *TcpIp_TypeGroup2) SetPropertydaddr(value interface{}) (err error) {
	return instance.SetProperty("daddr", (value))
}

// Getdaddr gets the value of daddr for the instance
func (instance *TcpIp_TypeGroup2) GetPropertydaddr() (value interface{}, err error) {
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
func (instance *TcpIp_TypeGroup2) SetPropertydport(value interface{}) (err error) {
	return instance.SetProperty("dport", (value))
}

// Getdport gets the value of dport for the instance
func (instance *TcpIp_TypeGroup2) GetPropertydport() (value interface{}, err error) {
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

// Setmss sets the value of mss for the instance
func (instance *TcpIp_TypeGroup2) SetPropertymss(value uint16) (err error) {
	return instance.SetProperty("mss", (value))
}

// Getmss gets the value of mss for the instance
func (instance *TcpIp_TypeGroup2) GetPropertymss() (value uint16, err error) {
	retValue, err := instance.GetProperty("mss")
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

// SetPID sets the value of PID for the instance
func (instance *TcpIp_TypeGroup2) SetPropertyPID(value uint32) (err error) {
	return instance.SetProperty("PID", (value))
}

// GetPID gets the value of PID for the instance
func (instance *TcpIp_TypeGroup2) GetPropertyPID() (value uint32, err error) {
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

// Setrcvwin sets the value of rcvwin for the instance
func (instance *TcpIp_TypeGroup2) SetPropertyrcvwin(value uint32) (err error) {
	return instance.SetProperty("rcvwin", (value))
}

// Getrcvwin gets the value of rcvwin for the instance
func (instance *TcpIp_TypeGroup2) GetPropertyrcvwin() (value uint32, err error) {
	retValue, err := instance.GetProperty("rcvwin")
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

// Setrcvwinscale sets the value of rcvwinscale for the instance
func (instance *TcpIp_TypeGroup2) SetPropertyrcvwinscale(value int16) (err error) {
	return instance.SetProperty("rcvwinscale", (value))
}

// Getrcvwinscale gets the value of rcvwinscale for the instance
func (instance *TcpIp_TypeGroup2) GetPropertyrcvwinscale() (value int16, err error) {
	retValue, err := instance.GetProperty("rcvwinscale")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int16(valuetmp)

	return
}

// Setsackopt sets the value of sackopt for the instance
func (instance *TcpIp_TypeGroup2) SetPropertysackopt(value uint16) (err error) {
	return instance.SetProperty("sackopt", (value))
}

// Getsackopt gets the value of sackopt for the instance
func (instance *TcpIp_TypeGroup2) GetPropertysackopt() (value uint16, err error) {
	retValue, err := instance.GetProperty("sackopt")
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
func (instance *TcpIp_TypeGroup2) SetPropertysaddr(value interface{}) (err error) {
	return instance.SetProperty("saddr", (value))
}

// Getsaddr gets the value of saddr for the instance
func (instance *TcpIp_TypeGroup2) GetPropertysaddr() (value interface{}, err error) {
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
func (instance *TcpIp_TypeGroup2) SetPropertyseqnum(value uint32) (err error) {
	return instance.SetProperty("seqnum", (value))
}

// Getseqnum gets the value of seqnum for the instance
func (instance *TcpIp_TypeGroup2) GetPropertyseqnum() (value uint32, err error) {
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
func (instance *TcpIp_TypeGroup2) SetPropertysize(value uint32) (err error) {
	return instance.SetProperty("size", (value))
}

// Getsize gets the value of size for the instance
func (instance *TcpIp_TypeGroup2) GetPropertysize() (value uint32, err error) {
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

// Setsndwinscale sets the value of sndwinscale for the instance
func (instance *TcpIp_TypeGroup2) SetPropertysndwinscale(value int16) (err error) {
	return instance.SetProperty("sndwinscale", (value))
}

// Getsndwinscale gets the value of sndwinscale for the instance
func (instance *TcpIp_TypeGroup2) GetPropertysndwinscale() (value int16, err error) {
	retValue, err := instance.GetProperty("sndwinscale")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int16(valuetmp)

	return
}

// Setsport sets the value of sport for the instance
func (instance *TcpIp_TypeGroup2) SetPropertysport(value interface{}) (err error) {
	return instance.SetProperty("sport", (value))
}

// Getsport gets the value of sport for the instance
func (instance *TcpIp_TypeGroup2) GetPropertysport() (value interface{}, err error) {
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

// Settsopt sets the value of tsopt for the instance
func (instance *TcpIp_TypeGroup2) SetPropertytsopt(value uint16) (err error) {
	return instance.SetProperty("tsopt", (value))
}

// Gettsopt gets the value of tsopt for the instance
func (instance *TcpIp_TypeGroup2) GetPropertytsopt() (value uint16, err error) {
	retValue, err := instance.GetProperty("tsopt")
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

// Setwsopt sets the value of wsopt for the instance
func (instance *TcpIp_TypeGroup2) SetPropertywsopt(value uint16) (err error) {
	return instance.SetProperty("wsopt", (value))
}

// Getwsopt gets the value of wsopt for the instance
func (instance *TcpIp_TypeGroup2) GetPropertywsopt() (value uint16, err error) {
	retValue, err := instance.GetProperty("wsopt")
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
