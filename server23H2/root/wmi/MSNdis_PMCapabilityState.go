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

// MSNdis_PMCapabilityState struct
type MSNdis_PMCapabilityState struct {
	*MSNdis

	//
	NdisPMCapabilityState PMCapabilityState_NdisPMCapabilityState
}

func NewMSNdis_PMCapabilityStateEx1(instance *cim.WmiInstance) (newInstance *MSNdis_PMCapabilityState, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PMCapabilityState{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_PMCapabilityStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_PMCapabilityState, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PMCapabilityState{
		MSNdis: tmp,
	}
	return
}

// SetNdisPMCapabilityState sets the value of NdisPMCapabilityState for the instance
func (instance *MSNdis_PMCapabilityState) SetPropertyNdisPMCapabilityState(value PMCapabilityState_NdisPMCapabilityState) (err error) {
	return instance.SetProperty("NdisPMCapabilityState", (value))
}

// GetNdisPMCapabilityState gets the value of NdisPMCapabilityState for the instance
func (instance *MSNdis_PMCapabilityState) GetPropertyNdisPMCapabilityState() (value PMCapabilityState_NdisPMCapabilityState, err error) {
	retValue, err := instance.GetProperty("NdisPMCapabilityState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = PMCapabilityState_NdisPMCapabilityState(valuetmp)

	return
}
