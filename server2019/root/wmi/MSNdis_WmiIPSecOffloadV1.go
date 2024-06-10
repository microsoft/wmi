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

// MSNdis_WmiIPSecOffloadV1 struct
type MSNdis_WmiIPSecOffloadV1 struct {
	*MSNdis

	//
	WmiIPv4AH MSNdis_WmiIPSecOffloadV1_IPv4AH

	//
	WmiIPv4ESP MSNdis_WmiIPSecOffloadV1_IPv4ESP

	//
	WmiSupported MSNdis_WmiIPSecOffloadV1_Supported
}

func NewMSNdis_WmiIPSecOffloadV1Ex1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiIPSecOffloadV1, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiIPSecOffloadV1{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiIPSecOffloadV1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiIPSecOffloadV1, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiIPSecOffloadV1{
		MSNdis: tmp,
	}
	return
}

// SetWmiIPv4AH sets the value of WmiIPv4AH for the instance
func (instance *MSNdis_WmiIPSecOffloadV1) SetPropertyWmiIPv4AH(value MSNdis_WmiIPSecOffloadV1_IPv4AH) (err error) {
	return instance.SetProperty("WmiIPv4AH", (value))
}

// GetWmiIPv4AH gets the value of WmiIPv4AH for the instance
func (instance *MSNdis_WmiIPSecOffloadV1) GetPropertyWmiIPv4AH() (value MSNdis_WmiIPSecOffloadV1_IPv4AH, err error) {
	retValue, err := instance.GetProperty("WmiIPv4AH")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_WmiIPSecOffloadV1_IPv4AH)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiIPSecOffloadV1_IPv4AH is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_WmiIPSecOffloadV1_IPv4AH(valuetmp)

	return
}

// SetWmiIPv4ESP sets the value of WmiIPv4ESP for the instance
func (instance *MSNdis_WmiIPSecOffloadV1) SetPropertyWmiIPv4ESP(value MSNdis_WmiIPSecOffloadV1_IPv4ESP) (err error) {
	return instance.SetProperty("WmiIPv4ESP", (value))
}

// GetWmiIPv4ESP gets the value of WmiIPv4ESP for the instance
func (instance *MSNdis_WmiIPSecOffloadV1) GetPropertyWmiIPv4ESP() (value MSNdis_WmiIPSecOffloadV1_IPv4ESP, err error) {
	retValue, err := instance.GetProperty("WmiIPv4ESP")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_WmiIPSecOffloadV1_IPv4ESP)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiIPSecOffloadV1_IPv4ESP is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_WmiIPSecOffloadV1_IPv4ESP(valuetmp)

	return
}

// SetWmiSupported sets the value of WmiSupported for the instance
func (instance *MSNdis_WmiIPSecOffloadV1) SetPropertyWmiSupported(value MSNdis_WmiIPSecOffloadV1_Supported) (err error) {
	return instance.SetProperty("WmiSupported", (value))
}

// GetWmiSupported gets the value of WmiSupported for the instance
func (instance *MSNdis_WmiIPSecOffloadV1) GetPropertyWmiSupported() (value MSNdis_WmiIPSecOffloadV1_Supported, err error) {
	retValue, err := instance.GetProperty("WmiSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_WmiIPSecOffloadV1_Supported)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiIPSecOffloadV1_Supported is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_WmiIPSecOffloadV1_Supported(valuetmp)

	return
}
