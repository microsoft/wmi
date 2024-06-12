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

// MSNdis_CoLinkSpeed struct
type MSNdis_CoLinkSpeed struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	NdisCoLinkSpeed MSNdis_NetworkLinkSpeed
}

func NewMSNdis_CoLinkSpeedEx1(instance *cim.WmiInstance) (newInstance *MSNdis_CoLinkSpeed, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_CoLinkSpeed{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_CoLinkSpeedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_CoLinkSpeed, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_CoLinkSpeed{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_CoLinkSpeed) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_CoLinkSpeed) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSNdis_CoLinkSpeed) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_CoLinkSpeed) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetNdisCoLinkSpeed sets the value of NdisCoLinkSpeed for the instance
func (instance *MSNdis_CoLinkSpeed) SetPropertyNdisCoLinkSpeed(value MSNdis_NetworkLinkSpeed) (err error) {
	return instance.SetProperty("NdisCoLinkSpeed", (value))
}

// GetNdisCoLinkSpeed gets the value of NdisCoLinkSpeed for the instance
func (instance *MSNdis_CoLinkSpeed) GetPropertyNdisCoLinkSpeed() (value MSNdis_NetworkLinkSpeed, err error) {
	retValue, err := instance.GetProperty("NdisCoLinkSpeed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_NetworkLinkSpeed)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_NetworkLinkSpeed is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_NetworkLinkSpeed(valuetmp)

	return
}
