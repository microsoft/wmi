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

// Msvm_EthernetSwitchPortVlanSettingData struct
type Msvm_EthernetSwitchPortVlanSettingData struct {
	*Msvm_EthernetSwitchPortFeatureSettingData

	//
	AccessVlanId uint16

	//
	NativeVlanId uint16

	//
	OperationMode uint32

	//
	PrimaryVlanId uint16

	//
	PruneVlanIdArray []uint16

	//
	PvlanMode uint32

	//
	SecondaryVlanId uint16

	//
	SecondaryVlanIdArray []uint16

	//
	TrunkVlanIdArray []uint16
}

func NewMsvm_EthernetSwitchPortVlanSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortVlanSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortVlanSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchPortVlanSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchPortVlanSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortVlanSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

// SetAccessVlanId sets the value of AccessVlanId for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) SetPropertyAccessVlanId(value uint16) (err error) {
	return instance.SetProperty("AccessVlanId", (value))
}

// GetAccessVlanId gets the value of AccessVlanId for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) GetPropertyAccessVlanId() (value uint16, err error) {
	retValue, err := instance.GetProperty("AccessVlanId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetNativeVlanId sets the value of NativeVlanId for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) SetPropertyNativeVlanId(value uint16) (err error) {
	return instance.SetProperty("NativeVlanId", (value))
}

// GetNativeVlanId gets the value of NativeVlanId for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) GetPropertyNativeVlanId() (value uint16, err error) {
	retValue, err := instance.GetProperty("NativeVlanId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetOperationMode sets the value of OperationMode for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) SetPropertyOperationMode(value uint32) (err error) {
	return instance.SetProperty("OperationMode", (value))
}

// GetOperationMode gets the value of OperationMode for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) GetPropertyOperationMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("OperationMode")
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

// SetPrimaryVlanId sets the value of PrimaryVlanId for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) SetPropertyPrimaryVlanId(value uint16) (err error) {
	return instance.SetProperty("PrimaryVlanId", (value))
}

// GetPrimaryVlanId gets the value of PrimaryVlanId for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) GetPropertyPrimaryVlanId() (value uint16, err error) {
	retValue, err := instance.GetProperty("PrimaryVlanId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetPruneVlanIdArray sets the value of PruneVlanIdArray for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) SetPropertyPruneVlanIdArray(value []uint16) (err error) {
	return instance.SetProperty("PruneVlanIdArray", (value))
}

// GetPruneVlanIdArray gets the value of PruneVlanIdArray for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) GetPropertyPruneVlanIdArray() (value []uint16, err error) {
	retValue, err := instance.GetProperty("PruneVlanIdArray")
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

// SetPvlanMode sets the value of PvlanMode for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) SetPropertyPvlanMode(value uint32) (err error) {
	return instance.SetProperty("PvlanMode", (value))
}

// GetPvlanMode gets the value of PvlanMode for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) GetPropertyPvlanMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("PvlanMode")
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

// SetSecondaryVlanId sets the value of SecondaryVlanId for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) SetPropertySecondaryVlanId(value uint16) (err error) {
	return instance.SetProperty("SecondaryVlanId", (value))
}

// GetSecondaryVlanId gets the value of SecondaryVlanId for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) GetPropertySecondaryVlanId() (value uint16, err error) {
	retValue, err := instance.GetProperty("SecondaryVlanId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetSecondaryVlanIdArray sets the value of SecondaryVlanIdArray for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) SetPropertySecondaryVlanIdArray(value []uint16) (err error) {
	return instance.SetProperty("SecondaryVlanIdArray", (value))
}

// GetSecondaryVlanIdArray gets the value of SecondaryVlanIdArray for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) GetPropertySecondaryVlanIdArray() (value []uint16, err error) {
	retValue, err := instance.GetProperty("SecondaryVlanIdArray")
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

// SetTrunkVlanIdArray sets the value of TrunkVlanIdArray for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) SetPropertyTrunkVlanIdArray(value []uint16) (err error) {
	return instance.SetProperty("TrunkVlanIdArray", (value))
}

// GetTrunkVlanIdArray gets the value of TrunkVlanIdArray for the instance
func (instance *Msvm_EthernetSwitchPortVlanSettingData) GetPropertyTrunkVlanIdArray() (value []uint16, err error) {
	retValue, err := instance.GetProperty("TrunkVlanIdArray")
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
func (instance *Msvm_EthernetSwitchPortVlanSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
