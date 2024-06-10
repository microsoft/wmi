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

// MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive struct
type MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive struct {
	*MSNdis

	//
	Encapsulation uint32

	//
	IpChecksum uint32

	//
	IpOptionsSupported uint32

	//
	TcpChecksum uint32

	//
	TcpOptionsSupported uint32

	//
	UdpChecksum uint32
}

func NewMSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceiveEx1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceiveEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive{
		MSNdis: tmp,
	}
	return
}

// SetEncapsulation sets the value of Encapsulation for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) SetPropertyEncapsulation(value uint32) (err error) {
	return instance.SetProperty("Encapsulation", (value))
}

// GetEncapsulation gets the value of Encapsulation for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) GetPropertyEncapsulation() (value uint32, err error) {
	retValue, err := instance.GetProperty("Encapsulation")
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

// SetIpChecksum sets the value of IpChecksum for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) SetPropertyIpChecksum(value uint32) (err error) {
	return instance.SetProperty("IpChecksum", (value))
}

// GetIpChecksum gets the value of IpChecksum for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) GetPropertyIpChecksum() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpChecksum")
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

// SetIpOptionsSupported sets the value of IpOptionsSupported for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) SetPropertyIpOptionsSupported(value uint32) (err error) {
	return instance.SetProperty("IpOptionsSupported", (value))
}

// GetIpOptionsSupported gets the value of IpOptionsSupported for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) GetPropertyIpOptionsSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpOptionsSupported")
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

// SetTcpChecksum sets the value of TcpChecksum for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) SetPropertyTcpChecksum(value uint32) (err error) {
	return instance.SetProperty("TcpChecksum", (value))
}

// GetTcpChecksum gets the value of TcpChecksum for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) GetPropertyTcpChecksum() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpChecksum")
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

// SetTcpOptionsSupported sets the value of TcpOptionsSupported for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) SetPropertyTcpOptionsSupported(value uint32) (err error) {
	return instance.SetProperty("TcpOptionsSupported", (value))
}

// GetTcpOptionsSupported gets the value of TcpOptionsSupported for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) GetPropertyTcpOptionsSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpOptionsSupported")
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

// SetUdpChecksum sets the value of UdpChecksum for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) SetPropertyUdpChecksum(value uint32) (err error) {
	return instance.SetProperty("UdpChecksum", (value))
}

// GetUdpChecksum gets the value of UdpChecksum for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) GetPropertyUdpChecksum() (value uint32, err error) {
	retValue, err := instance.GetProperty("UdpChecksum")
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
