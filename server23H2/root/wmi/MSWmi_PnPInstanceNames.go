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

// MSWmi_PnPInstanceNames struct
type MSWmi_PnPInstanceNames struct {
	*MS_WmiInternal

	//
	Active bool

	//
	Count uint32

	//
	InstanceName string

	//
	InstanceNameList []string
}

func NewMSWmi_PnPInstanceNamesEx1(instance *cim.WmiInstance) (newInstance *MSWmi_PnPInstanceNames, err error) {
	tmp, err := NewMS_WmiInternalEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSWmi_PnPInstanceNames{
		MS_WmiInternal: tmp,
	}
	return
}

func NewMSWmi_PnPInstanceNamesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSWmi_PnPInstanceNames, err error) {
	tmp, err := NewMS_WmiInternalEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSWmi_PnPInstanceNames{
		MS_WmiInternal: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSWmi_PnPInstanceNames) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSWmi_PnPInstanceNames) GetPropertyActive() (value bool, err error) {
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
func (instance *MSWmi_PnPInstanceNames) SetPropertyCount(value uint32) (err error) {
	return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *MSWmi_PnPInstanceNames) GetPropertyCount() (value uint32, err error) {
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
func (instance *MSWmi_PnPInstanceNames) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSWmi_PnPInstanceNames) GetPropertyInstanceName() (value string, err error) {
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

// SetInstanceNameList sets the value of InstanceNameList for the instance
func (instance *MSWmi_PnPInstanceNames) SetPropertyInstanceNameList(value []string) (err error) {
	return instance.SetProperty("InstanceNameList", (value))
}

// GetInstanceNameList gets the value of InstanceNameList for the instance
func (instance *MSWmi_PnPInstanceNames) GetPropertyInstanceNameList() (value []string, err error) {
	retValue, err := instance.GetProperty("InstanceNameList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
