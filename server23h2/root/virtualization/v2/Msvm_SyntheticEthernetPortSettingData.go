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

// Msvm_SyntheticEthernetPortSettingData struct
type Msvm_SyntheticEthernetPortSettingData struct {
	*CIM_EthernetPortAllocationSettingData

	//
	AllowDirectTranslatedP2P bool

	//
	AllowPacketDirect bool

	//
	ClusterMonitored bool

	//
	DeviceNamingEnabled bool

	//
	InterruptModeration bool

	//
	MediaType uint32

	//
	NumaAwarePlacement bool

	//
	StaticMacAddress bool

	//
	VirtualSystemIdentifiers []string
}

func NewMsvm_SyntheticEthernetPortSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_SyntheticEthernetPortSettingData, err error) {
	tmp, err := NewCIM_EthernetPortAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticEthernetPortSettingData{
		CIM_EthernetPortAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_SyntheticEthernetPortSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SyntheticEthernetPortSettingData, err error) {
	tmp, err := NewCIM_EthernetPortAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticEthernetPortSettingData{
		CIM_EthernetPortAllocationSettingData: tmp,
	}
	return
}

// SetAllowDirectTranslatedP2P sets the value of AllowDirectTranslatedP2P for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyAllowDirectTranslatedP2P(value bool) (err error) {
	return instance.SetProperty("AllowDirectTranslatedP2P", (value))
}

// GetAllowDirectTranslatedP2P gets the value of AllowDirectTranslatedP2P for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyAllowDirectTranslatedP2P() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowDirectTranslatedP2P")
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

// SetAllowPacketDirect sets the value of AllowPacketDirect for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyAllowPacketDirect(value bool) (err error) {
	return instance.SetProperty("AllowPacketDirect", (value))
}

// GetAllowPacketDirect gets the value of AllowPacketDirect for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyAllowPacketDirect() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowPacketDirect")
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

// SetClusterMonitored sets the value of ClusterMonitored for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyClusterMonitored(value bool) (err error) {
	return instance.SetProperty("ClusterMonitored", (value))
}

// GetClusterMonitored gets the value of ClusterMonitored for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyClusterMonitored() (value bool, err error) {
	retValue, err := instance.GetProperty("ClusterMonitored")
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

// SetDeviceNamingEnabled sets the value of DeviceNamingEnabled for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyDeviceNamingEnabled(value bool) (err error) {
	return instance.SetProperty("DeviceNamingEnabled", (value))
}

// GetDeviceNamingEnabled gets the value of DeviceNamingEnabled for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyDeviceNamingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DeviceNamingEnabled")
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

// SetInterruptModeration sets the value of InterruptModeration for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyInterruptModeration(value bool) (err error) {
	return instance.SetProperty("InterruptModeration", (value))
}

// GetInterruptModeration gets the value of InterruptModeration for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyInterruptModeration() (value bool, err error) {
	retValue, err := instance.GetProperty("InterruptModeration")
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

// SetMediaType sets the value of MediaType for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyMediaType(value uint32) (err error) {
	return instance.SetProperty("MediaType", (value))
}

// GetMediaType gets the value of MediaType for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyMediaType() (value uint32, err error) {
	retValue, err := instance.GetProperty("MediaType")
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

// SetNumaAwarePlacement sets the value of NumaAwarePlacement for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyNumaAwarePlacement(value bool) (err error) {
	return instance.SetProperty("NumaAwarePlacement", (value))
}

// GetNumaAwarePlacement gets the value of NumaAwarePlacement for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyNumaAwarePlacement() (value bool, err error) {
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

// SetStaticMacAddress sets the value of StaticMacAddress for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyStaticMacAddress(value bool) (err error) {
	return instance.SetProperty("StaticMacAddress", (value))
}

// GetStaticMacAddress gets the value of StaticMacAddress for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyStaticMacAddress() (value bool, err error) {
	retValue, err := instance.GetProperty("StaticMacAddress")
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

// SetVirtualSystemIdentifiers sets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyVirtualSystemIdentifiers(value []string) (err error) {
	return instance.SetProperty("VirtualSystemIdentifiers", (value))
}

// GetVirtualSystemIdentifiers gets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyVirtualSystemIdentifiers() (value []string, err error) {
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
func (instance *Msvm_SyntheticEthernetPortSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
