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

// MSNdis_80211_InfrastructureMode struct
type MSNdis_80211_InfrastructureMode struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	Ndis80211InfrastructureMode MSNdis_80211_NetworkInfrastructure
}

func NewMSNdis_80211_InfrastructureModeEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_InfrastructureMode, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_InfrastructureMode{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_80211_InfrastructureModeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_80211_InfrastructureMode, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_InfrastructureMode{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_80211_InfrastructureMode) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_80211_InfrastructureMode) GetPropertyActive() (value bool, err error) {
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
func (instance *MSNdis_80211_InfrastructureMode) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_80211_InfrastructureMode) GetPropertyInstanceName() (value string, err error) {
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

// SetNdis80211InfrastructureMode sets the value of Ndis80211InfrastructureMode for the instance
func (instance *MSNdis_80211_InfrastructureMode) SetPropertyNdis80211InfrastructureMode(value MSNdis_80211_NetworkInfrastructure) (err error) {
	return instance.SetProperty("Ndis80211InfrastructureMode", (value))
}

// GetNdis80211InfrastructureMode gets the value of Ndis80211InfrastructureMode for the instance
func (instance *MSNdis_80211_InfrastructureMode) GetPropertyNdis80211InfrastructureMode() (value MSNdis_80211_NetworkInfrastructure, err error) {
	retValue, err := instance.GetProperty("Ndis80211InfrastructureMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_80211_NetworkInfrastructure)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_80211_NetworkInfrastructure is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_80211_NetworkInfrastructure(valuetmp)

	return
}
