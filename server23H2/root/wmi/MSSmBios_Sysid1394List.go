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

// MSSmBios_Sysid1394List struct
type MSSmBios_Sysid1394List struct {
	*MS_SmBios

	//
	Active bool

	//
	Count uint32

	//
	InstanceName string

	//
	List []MSSmBios_Sysid1394
}

func NewMSSmBios_Sysid1394ListEx1(instance *cim.WmiInstance) (newInstance *MSSmBios_Sysid1394List, err error) {
	tmp, err := NewMS_SmBiosEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSSmBios_Sysid1394List{
		MS_SmBios: tmp,
	}
	return
}

func NewMSSmBios_Sysid1394ListEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSSmBios_Sysid1394List, err error) {
	tmp, err := NewMS_SmBiosEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSSmBios_Sysid1394List{
		MS_SmBios: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSSmBios_Sysid1394List) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSSmBios_Sysid1394List) GetPropertyActive() (value bool, err error) {
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

// SetCount sets the value of Count for the instance
func (instance *MSSmBios_Sysid1394List) SetPropertyCount(value uint32) (err error) {
	return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *MSSmBios_Sysid1394List) GetPropertyCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("Count")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSSmBios_Sysid1394List) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSSmBios_Sysid1394List) GetPropertyInstanceName() (value string, err error) {
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

// SetList sets the value of List for the instance
func (instance *MSSmBios_Sysid1394List) SetPropertyList(value []MSSmBios_Sysid1394) (err error) {
	return instance.SetProperty("List", (value))
}

// GetList gets the value of List for the instance
func (instance *MSSmBios_Sysid1394List) GetPropertyList() (value []MSSmBios_Sysid1394, err error) {
	retValue, err := instance.GetProperty("List")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSSmBios_Sysid1394)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSSmBios_Sysid1394 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSSmBios_Sysid1394(valuetmp))
	}

	return
}
