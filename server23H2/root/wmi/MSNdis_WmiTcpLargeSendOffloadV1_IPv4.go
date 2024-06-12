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

// MSNdis_WmiTcpLargeSendOffloadV1_IPv4 struct
type MSNdis_WmiTcpLargeSendOffloadV1_IPv4 struct {
	*MSNdis

	//
	Encapsulation uint32

	//
	IpOptions uint32

	//
	MaxOffLoadSize uint32

	//
	MinSegmentCount uint32

	//
	TcpOptions uint32
}

func NewMSNdis_WmiTcpLargeSendOffloadV1_IPv4Ex1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiTcpLargeSendOffloadV1_IPv4, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiTcpLargeSendOffloadV1_IPv4{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiTcpLargeSendOffloadV1_IPv4Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiTcpLargeSendOffloadV1_IPv4, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiTcpLargeSendOffloadV1_IPv4{
		MSNdis: tmp,
	}
	return
}

// SetEncapsulation sets the value of Encapsulation for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV1_IPv4) SetPropertyEncapsulation(value uint32) (err error) {
	return instance.SetProperty("Encapsulation", (value))
}

// GetEncapsulation gets the value of Encapsulation for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV1_IPv4) GetPropertyEncapsulation() (value uint32, err error) {
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

// SetIpOptions sets the value of IpOptions for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV1_IPv4) SetPropertyIpOptions(value uint32) (err error) {
	return instance.SetProperty("IpOptions", (value))
}

// GetIpOptions gets the value of IpOptions for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV1_IPv4) GetPropertyIpOptions() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpOptions")
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
func (instance *MSNdis_WmiTcpLargeSendOffloadV1_IPv4) SetPropertyMaxOffLoadSize(value uint32) (err error) {
	return instance.SetProperty("MaxOffLoadSize", (value))
}

// GetMaxOffLoadSize gets the value of MaxOffLoadSize for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV1_IPv4) GetPropertyMaxOffLoadSize() (value uint32, err error) {
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
func (instance *MSNdis_WmiTcpLargeSendOffloadV1_IPv4) SetPropertyMinSegmentCount(value uint32) (err error) {
	return instance.SetProperty("MinSegmentCount", (value))
}

// GetMinSegmentCount gets the value of MinSegmentCount for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV1_IPv4) GetPropertyMinSegmentCount() (value uint32, err error) {
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

// SetTcpOptions sets the value of TcpOptions for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV1_IPv4) SetPropertyTcpOptions(value uint32) (err error) {
	return instance.SetProperty("TcpOptions", (value))
}

// GetTcpOptions gets the value of TcpOptions for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV1_IPv4) GetPropertyTcpOptions() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpOptions")
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
