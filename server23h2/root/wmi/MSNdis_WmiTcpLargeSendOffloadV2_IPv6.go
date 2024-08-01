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

// MSNdis_WmiTcpLargeSendOffloadV2_IPv6 struct
type MSNdis_WmiTcpLargeSendOffloadV2_IPv6 struct {
	*MSNdis

	//
	Encapsulation uint32

	//
	IpExtensionHeadersSupported uint32

	//
	MaxOffLoadSize uint32

	//
	MinSegmentCount uint32

	//
	TcpOptionsSupported uint32
}

func NewMSNdis_WmiTcpLargeSendOffloadV2_IPv6Ex1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiTcpLargeSendOffloadV2_IPv6, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiTcpLargeSendOffloadV2_IPv6{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiTcpLargeSendOffloadV2_IPv6Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiTcpLargeSendOffloadV2_IPv6, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiTcpLargeSendOffloadV2_IPv6{
		MSNdis: tmp,
	}
	return
}

// SetEncapsulation sets the value of Encapsulation for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV2_IPv6) SetPropertyEncapsulation(value uint32) (err error) {
	return instance.SetProperty("Encapsulation", (value))
}

// GetEncapsulation gets the value of Encapsulation for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV2_IPv6) GetPropertyEncapsulation() (value uint32, err error) {
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

// SetIpExtensionHeadersSupported sets the value of IpExtensionHeadersSupported for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV2_IPv6) SetPropertyIpExtensionHeadersSupported(value uint32) (err error) {
	return instance.SetProperty("IpExtensionHeadersSupported", (value))
}

// GetIpExtensionHeadersSupported gets the value of IpExtensionHeadersSupported for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV2_IPv6) GetPropertyIpExtensionHeadersSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpExtensionHeadersSupported")
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

// SetMaxOffLoadSize sets the value of MaxOffLoadSize for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV2_IPv6) SetPropertyMaxOffLoadSize(value uint32) (err error) {
	return instance.SetProperty("MaxOffLoadSize", (value))
}

// GetMaxOffLoadSize gets the value of MaxOffLoadSize for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV2_IPv6) GetPropertyMaxOffLoadSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxOffLoadSize")
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

// SetMinSegmentCount sets the value of MinSegmentCount for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV2_IPv6) SetPropertyMinSegmentCount(value uint32) (err error) {
	return instance.SetProperty("MinSegmentCount", (value))
}

// GetMinSegmentCount gets the value of MinSegmentCount for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV2_IPv6) GetPropertyMinSegmentCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinSegmentCount")
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
func (instance *MSNdis_WmiTcpLargeSendOffloadV2_IPv6) SetPropertyTcpOptionsSupported(value uint32) (err error) {
	return instance.SetProperty("TcpOptionsSupported", (value))
}

// GetTcpOptionsSupported gets the value of TcpOptionsSupported for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV2_IPv6) GetPropertyTcpOptionsSupported() (value uint32, err error) {
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
