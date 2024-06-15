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

// MSNdis_WmiTcpLargeSendOffloadV2 struct
type MSNdis_WmiTcpLargeSendOffloadV2 struct {
	*MSNdis

	//
	WmiIPv4 MSNdis_WmiTcpLargeSendOffloadV2_IPv4

	//
	WmiIPv6 MSNdis_WmiTcpLargeSendOffloadV2_IPv6
}

func NewMSNdis_WmiTcpLargeSendOffloadV2Ex1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiTcpLargeSendOffloadV2, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiTcpLargeSendOffloadV2{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiTcpLargeSendOffloadV2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiTcpLargeSendOffloadV2, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiTcpLargeSendOffloadV2{
		MSNdis: tmp,
	}
	return
}

// SetWmiIPv4 sets the value of WmiIPv4 for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV2) SetPropertyWmiIPv4(value MSNdis_WmiTcpLargeSendOffloadV2_IPv4) (err error) {
	return instance.SetProperty("WmiIPv4", (value))
}

// GetWmiIPv4 gets the value of WmiIPv4 for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV2) GetPropertyWmiIPv4() (value MSNdis_WmiTcpLargeSendOffloadV2_IPv4, err error) {
	retValue, err := instance.GetProperty("WmiIPv4")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_WmiTcpLargeSendOffloadV2_IPv4)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiTcpLargeSendOffloadV2_IPv4 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_WmiTcpLargeSendOffloadV2_IPv4(valuetmp)

	return
}

// SetWmiIPv6 sets the value of WmiIPv6 for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV2) SetPropertyWmiIPv6(value MSNdis_WmiTcpLargeSendOffloadV2_IPv6) (err error) {
	return instance.SetProperty("WmiIPv6", (value))
}

// GetWmiIPv6 gets the value of WmiIPv6 for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV2) GetPropertyWmiIPv6() (value MSNdis_WmiTcpLargeSendOffloadV2_IPv6, err error) {
	retValue, err := instance.GetProperty("WmiIPv6")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_WmiTcpLargeSendOffloadV2_IPv6)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiTcpLargeSendOffloadV2_IPv6 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_WmiTcpLargeSendOffloadV2_IPv6(valuetmp)

	return
}
