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

// MSNdis_NetworkDirectCapabilities struct
type MSNdis_NetworkDirectCapabilities struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	NdCapabilities MSNdis_NetworkDirectAdapterCapabilities
}

func NewMSNdis_NetworkDirectCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSNdis_NetworkDirectCapabilities, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_NetworkDirectCapabilities{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_NetworkDirectCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_NetworkDirectCapabilities, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_NetworkDirectCapabilities{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_NetworkDirectCapabilities) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_NetworkDirectCapabilities) GetPropertyActive() (value bool, err error) {
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
func (instance *MSNdis_NetworkDirectCapabilities) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_NetworkDirectCapabilities) GetPropertyInstanceName() (value string, err error) {
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

// SetNdCapabilities sets the value of NdCapabilities for the instance
func (instance *MSNdis_NetworkDirectCapabilities) SetPropertyNdCapabilities(value MSNdis_NetworkDirectAdapterCapabilities) (err error) {
	return instance.SetProperty("NdCapabilities", (value))
}

// GetNdCapabilities gets the value of NdCapabilities for the instance
func (instance *MSNdis_NetworkDirectCapabilities) GetPropertyNdCapabilities() (value MSNdis_NetworkDirectAdapterCapabilities, err error) {
	retValue, err := instance.GetProperty("NdCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_NetworkDirectAdapterCapabilities)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_NetworkDirectAdapterCapabilities is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_NetworkDirectAdapterCapabilities(valuetmp)

	return
}
