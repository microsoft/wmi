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

// WmiMonitorDescriptorMethods struct
type WmiMonitorDescriptorMethods struct {
	*MSMonitorClass

	//
	Active bool

	//
	InstanceName string
}

func NewWmiMonitorDescriptorMethodsEx1(instance *cim.WmiInstance) (newInstance *WmiMonitorDescriptorMethods, err error) {
	tmp, err := NewMSMonitorClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorDescriptorMethods{
		MSMonitorClass: tmp,
	}
	return
}

func NewWmiMonitorDescriptorMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WmiMonitorDescriptorMethods, err error) {
	tmp, err := NewMSMonitorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorDescriptorMethods{
		MSMonitorClass: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *WmiMonitorDescriptorMethods) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WmiMonitorDescriptorMethods) GetPropertyActive() (value bool, err error) {
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
func (instance *WmiMonitorDescriptorMethods) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WmiMonitorDescriptorMethods) GetPropertyInstanceName() (value string, err error) {
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

//

// <param name="BlockId" type="uint8 "></param>

// <param name="BlockContent" type="uint8 []"></param>
// <param name="BlockType" type="uint8 "></param>
func (instance *WmiMonitorDescriptorMethods) WmiGetMonitorRawEEdidV1Block( /* IN */ BlockId uint8,
	/* OUT */ BlockType uint8,
	/* OUT */ BlockContent []uint8) (err error) {
	_, err = instance.InvokeMethod("WmiGetMonitorRawEEdidV1Block", BlockId)
	if err != nil {
		return
	}
	return

}
