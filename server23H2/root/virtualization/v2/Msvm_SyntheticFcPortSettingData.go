// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_SyntheticFcPortSettingData struct
type Msvm_SyntheticFcPortSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	ChapEnabled bool

	//
	SecondaryWWNN string

	//
	SecondaryWWPN string

	//
	VirtualPortWWNN string

	//
	VirtualPortWWPN string

	//
	VirtualSystemIdentifiers []string
}

func NewMsvm_SyntheticFcPortSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_SyntheticFcPortSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticFcPortSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_SyntheticFcPortSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SyntheticFcPortSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticFcPortSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetChapEnabled sets the value of ChapEnabled for the instance
func (instance *Msvm_SyntheticFcPortSettingData) SetPropertyChapEnabled(value bool) (err error) {
	return instance.SetProperty("ChapEnabled", (value))
}

// GetChapEnabled gets the value of ChapEnabled for the instance
func (instance *Msvm_SyntheticFcPortSettingData) GetPropertyChapEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("ChapEnabled")
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

// SetSecondaryWWNN sets the value of SecondaryWWNN for the instance
func (instance *Msvm_SyntheticFcPortSettingData) SetPropertySecondaryWWNN(value string) (err error) {
	return instance.SetProperty("SecondaryWWNN", (value))
}

// GetSecondaryWWNN gets the value of SecondaryWWNN for the instance
func (instance *Msvm_SyntheticFcPortSettingData) GetPropertySecondaryWWNN() (value string, err error) {
	retValue, err := instance.GetProperty("SecondaryWWNN")
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

// SetSecondaryWWPN sets the value of SecondaryWWPN for the instance
func (instance *Msvm_SyntheticFcPortSettingData) SetPropertySecondaryWWPN(value string) (err error) {
	return instance.SetProperty("SecondaryWWPN", (value))
}

// GetSecondaryWWPN gets the value of SecondaryWWPN for the instance
func (instance *Msvm_SyntheticFcPortSettingData) GetPropertySecondaryWWPN() (value string, err error) {
	retValue, err := instance.GetProperty("SecondaryWWPN")
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

// SetVirtualPortWWNN sets the value of VirtualPortWWNN for the instance
func (instance *Msvm_SyntheticFcPortSettingData) SetPropertyVirtualPortWWNN(value string) (err error) {
	return instance.SetProperty("VirtualPortWWNN", (value))
}

// GetVirtualPortWWNN gets the value of VirtualPortWWNN for the instance
func (instance *Msvm_SyntheticFcPortSettingData) GetPropertyVirtualPortWWNN() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualPortWWNN")
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

// SetVirtualPortWWPN sets the value of VirtualPortWWPN for the instance
func (instance *Msvm_SyntheticFcPortSettingData) SetPropertyVirtualPortWWPN(value string) (err error) {
	return instance.SetProperty("VirtualPortWWPN", (value))
}

// GetVirtualPortWWPN gets the value of VirtualPortWWPN for the instance
func (instance *Msvm_SyntheticFcPortSettingData) GetPropertyVirtualPortWWPN() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualPortWWPN")
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

// SetVirtualSystemIdentifiers sets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_SyntheticFcPortSettingData) SetPropertyVirtualSystemIdentifiers(value []string) (err error) {
	return instance.SetProperty("VirtualSystemIdentifiers", (value))
}

// GetVirtualSystemIdentifiers gets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_SyntheticFcPortSettingData) GetPropertyVirtualSystemIdentifiers() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemIdentifiers")
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
func (instance *Msvm_SyntheticFcPortSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
