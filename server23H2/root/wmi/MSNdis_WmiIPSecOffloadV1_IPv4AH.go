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

// MSNdis_WmiIPSecOffloadV1_IPv4AH struct
type MSNdis_WmiIPSecOffloadV1_IPv4AH struct {
	*MSNdis

	//
	Md5 uint32

	//
	Receive uint32

	//
	Send uint32

	//
	Sha_1 uint32

	//
	Transport uint32

	//
	Tunnel uint32
}

func NewMSNdis_WmiIPSecOffloadV1_IPv4AHEx1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiIPSecOffloadV1_IPv4AH, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiIPSecOffloadV1_IPv4AH{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiIPSecOffloadV1_IPv4AHEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiIPSecOffloadV1_IPv4AH, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiIPSecOffloadV1_IPv4AH{
		MSNdis: tmp,
	}
	return
}

// SetMd5 sets the value of Md5 for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4AH) SetPropertyMd5(value uint32) (err error) {
	return instance.SetProperty("Md5", (value))
}

// GetMd5 gets the value of Md5 for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4AH) GetPropertyMd5() (value uint32, err error) {
	retValue, err := instance.GetProperty("Md5")
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

// SetReceive sets the value of Receive for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4AH) SetPropertyReceive(value uint32) (err error) {
	return instance.SetProperty("Receive", (value))
}

// GetReceive gets the value of Receive for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4AH) GetPropertyReceive() (value uint32, err error) {
	retValue, err := instance.GetProperty("Receive")
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

// SetSend sets the value of Send for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4AH) SetPropertySend(value uint32) (err error) {
	return instance.SetProperty("Send", (value))
}

// GetSend gets the value of Send for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4AH) GetPropertySend() (value uint32, err error) {
	retValue, err := instance.GetProperty("Send")
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

// SetSha_1 sets the value of Sha_1 for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4AH) SetPropertySha_1(value uint32) (err error) {
	return instance.SetProperty("Sha_1", (value))
}

// GetSha_1 gets the value of Sha_1 for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4AH) GetPropertySha_1() (value uint32, err error) {
	retValue, err := instance.GetProperty("Sha_1")
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

// SetTransport sets the value of Transport for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4AH) SetPropertyTransport(value uint32) (err error) {
	return instance.SetProperty("Transport", (value))
}

// GetTransport gets the value of Transport for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4AH) GetPropertyTransport() (value uint32, err error) {
	retValue, err := instance.GetProperty("Transport")
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

// SetTunnel sets the value of Tunnel for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4AH) SetPropertyTunnel(value uint32) (err error) {
	return instance.SetProperty("Tunnel", (value))
}

// GetTunnel gets the value of Tunnel for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4AH) GetPropertyTunnel() (value uint32, err error) {
	retValue, err := instance.GetProperty("Tunnel")
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
