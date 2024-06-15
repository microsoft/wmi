// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_PciExpressSettingData struct
type Msvm_PciExpressSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	AllowDirectTranslatedP2P []bool

	//
	NumaAwarePlacement bool

	//
	TargetVtl uint8

	//
	VirtualFunctions []uint16

	//
	VirtualSystemIdentifiers []string
}

func NewMsvm_PciExpressSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_PciExpressSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_PciExpressSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_PciExpressSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_PciExpressSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_PciExpressSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetAllowDirectTranslatedP2P sets the value of AllowDirectTranslatedP2P for the instance
func (instance *Msvm_PciExpressSettingData) SetPropertyAllowDirectTranslatedP2P(value []bool) (err error) {
	return instance.SetProperty("AllowDirectTranslatedP2P", (value))
}

// GetAllowDirectTranslatedP2P gets the value of AllowDirectTranslatedP2P for the instance
func (instance *Msvm_PciExpressSettingData) GetPropertyAllowDirectTranslatedP2P() (value []bool, err error) {
	retValue, err := instance.GetProperty("AllowDirectTranslatedP2P")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(bool)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, bool(valuetmp))
	}

	return
}

// SetNumaAwarePlacement sets the value of NumaAwarePlacement for the instance
func (instance *Msvm_PciExpressSettingData) SetPropertyNumaAwarePlacement(value bool) (err error) {
	return instance.SetProperty("NumaAwarePlacement", (value))
}

// GetNumaAwarePlacement gets the value of NumaAwarePlacement for the instance
func (instance *Msvm_PciExpressSettingData) GetPropertyNumaAwarePlacement() (value bool, err error) {
	retValue, err := instance.GetProperty("NumaAwarePlacement")
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

// SetTargetVtl sets the value of TargetVtl for the instance
func (instance *Msvm_PciExpressSettingData) SetPropertyTargetVtl(value uint8) (err error) {
	return instance.SetProperty("TargetVtl", (value))
}

// GetTargetVtl gets the value of TargetVtl for the instance
func (instance *Msvm_PciExpressSettingData) GetPropertyTargetVtl() (value uint8, err error) {
	retValue, err := instance.GetProperty("TargetVtl")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetVirtualFunctions sets the value of VirtualFunctions for the instance
func (instance *Msvm_PciExpressSettingData) SetPropertyVirtualFunctions(value []uint16) (err error) {
	return instance.SetProperty("VirtualFunctions", (value))
}

// GetVirtualFunctions gets the value of VirtualFunctions for the instance
func (instance *Msvm_PciExpressSettingData) GetPropertyVirtualFunctions() (value []uint16, err error) {
	retValue, err := instance.GetProperty("VirtualFunctions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetVirtualSystemIdentifiers sets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_PciExpressSettingData) SetPropertyVirtualSystemIdentifiers(value []string) (err error) {
	return instance.SetProperty("VirtualSystemIdentifiers", (value))
}

// GetVirtualSystemIdentifiers gets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_PciExpressSettingData) GetPropertyVirtualSystemIdentifiers() (value []string, err error) {
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
func (instance *Msvm_PciExpressSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
