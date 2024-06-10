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

// MSNdis_WmiTcpIpChecksumOffload struct
type MSNdis_WmiTcpIpChecksumOffload struct {
	*MSNdis

	//
	IPv4Receive MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive

	//
	IPv4Transmit MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive

	//
	IPv6Receive MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive

	//
	IPv6Transmit MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive
}

func NewMSNdis_WmiTcpIpChecksumOffloadEx1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiTcpIpChecksumOffload, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiTcpIpChecksumOffload{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiTcpIpChecksumOffloadEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiTcpIpChecksumOffload, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiTcpIpChecksumOffload{
		MSNdis: tmp,
	}
	return
}

// SetIPv4Receive sets the value of IPv4Receive for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload) SetPropertyIPv4Receive(value MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) (err error) {
	return instance.SetProperty("IPv4Receive", (value))
}

// GetIPv4Receive gets the value of IPv4Receive for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload) GetPropertyIPv4Receive() (value MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive, err error) {
	retValue, err := instance.GetProperty("IPv4Receive")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive(valuetmp)

	return
}

// SetIPv4Transmit sets the value of IPv4Transmit for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload) SetPropertyIPv4Transmit(value MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive) (err error) {
	return instance.SetProperty("IPv4Transmit", (value))
}

// GetIPv4Transmit gets the value of IPv4Transmit for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload) GetPropertyIPv4Transmit() (value MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive, err error) {
	retValue, err := instance.GetProperty("IPv4Transmit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_WmiTcpIpChecksumOffload_IPv4TransmitReceive(valuetmp)

	return
}

// SetIPv6Receive sets the value of IPv6Receive for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload) SetPropertyIPv6Receive(value MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive) (err error) {
	return instance.SetProperty("IPv6Receive", (value))
}

// GetIPv6Receive gets the value of IPv6Receive for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload) GetPropertyIPv6Receive() (value MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive, err error) {
	retValue, err := instance.GetProperty("IPv6Receive")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive(valuetmp)

	return
}

// SetIPv6Transmit sets the value of IPv6Transmit for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload) SetPropertyIPv6Transmit(value MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive) (err error) {
	return instance.SetProperty("IPv6Transmit", (value))
}

// GetIPv6Transmit gets the value of IPv6Transmit for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload) GetPropertyIPv6Transmit() (value MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive, err error) {
	retValue, err := instance.GetProperty("IPv6Transmit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive(valuetmp)

	return
}
