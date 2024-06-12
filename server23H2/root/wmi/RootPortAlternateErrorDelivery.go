// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RootPortAlternateErrorDelivery struct
type RootPortAlternateErrorDelivery struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewRootPortAlternateErrorDeliveryEx1(instance *cim.WmiInstance) (newInstance *RootPortAlternateErrorDelivery, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RootPortAlternateErrorDelivery{
		WmiInstance: tmp,
	}
	return
}

func NewRootPortAlternateErrorDeliveryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RootPortAlternateErrorDelivery, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RootPortAlternateErrorDelivery{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *RootPortAlternateErrorDelivery) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *RootPortAlternateErrorDelivery) GetPropertyActive() (value bool, err error) {
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
func (instance *RootPortAlternateErrorDelivery) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *RootPortAlternateErrorDelivery) GetPropertyInstanceName() (value string, err error) {
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

// Change root port delivery from NMI to SCI

// <param name="Bus" type="uint8 ">The Bus number of the root port.</param>
// <param name="Device" type="uint8 ">The Device number of the root port.</param>
// <param name="Function" type="uint8 ">The Function number of the root port.</param>
// <param name="Segment" type="uint16 ">The segment number of the root port.</param>

// <param name="Status" type="uint8 ">Status of the method</param>
func (instance *RootPortAlternateErrorDelivery) EnableAlternateErrorDelivery( /* IN */ Segment uint16,
	/* IN */ Bus uint8,
	/* IN */ Device uint8,
	/* IN */ Function uint8,
	/* OUT */ Status uint8) (err error) {
	_, err = instance.InvokeMethod("EnableAlternateErrorDelivery", Segment, Bus, Device, Function)
	if err != nil {
		return
	}
	return

}

// Change root port delivery from SCI to NMI

// <param name="Bus" type="uint8 ">The Bus number of the root port.</param>
// <param name="Device" type="uint8 ">The Device number of the root port.</param>
// <param name="Function" type="uint8 ">The Function number of the root port.</param>
// <param name="Segment" type="uint16 ">The segment number of the root port.</param>

// <param name="Status" type="uint8 ">Status of the method</param>
func (instance *RootPortAlternateErrorDelivery) DisableAlternateErrorDelivery( /* IN */ Segment uint16,
	/* IN */ Bus uint8,
	/* IN */ Device uint8,
	/* IN */ Function uint8,
	/* OUT */ Status uint8) (err error) {
	_, err = instance.InvokeMethod("DisableAlternateErrorDelivery", Segment, Bus, Device, Function)
	if err != nil {
		return
	}
	return

}

// Reenable error delivery after an error occurs

// <param name="Bus" type="uint8 ">The Bus number of the root port.</param>
// <param name="Device" type="uint8 ">The Device number of the root port.</param>
// <param name="Function" type="uint8 ">The Function number of the root port.</param>
// <param name="Segment" type="uint16 ">The segment number of the root port.</param>

// <param name="Status" type="uint8 ">Status of the method</param>
func (instance *RootPortAlternateErrorDelivery) ReenableErrorDelivery( /* IN */ Segment uint16,
	/* IN */ Bus uint8,
	/* IN */ Device uint8,
	/* IN */ Function uint8,
	/* OUT */ Status uint8) (err error) {
	_, err = instance.InvokeMethod("ReenableErrorDelivery", Segment, Bus, Device, Function)
	if err != nil {
		return
	}
	return

}
