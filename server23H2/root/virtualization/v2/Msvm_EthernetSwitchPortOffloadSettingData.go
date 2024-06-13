// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_EthernetSwitchPortOffloadSettingData struct
type Msvm_EthernetSwitchPortOffloadSettingData struct {
	*Msvm_EthernetSwitchPortFeatureSettingData

	//
	IOVInterruptModeration EthernetSwitchPortOffloadSettingData_IOVInterruptModeration

	//
	IOVOffloadWeight uint32

	//
	IOVQueuePairsRequested uint32

	//
	IPSecOffloadLimit uint32

	//
	PacketDirectModerationCount uint32

	//
	PacketDirectModerationInterval uint32

	//
	PacketDirectNumProcs uint32

	//
	RscEnabled bool

	//
	VmmqEnabled bool

	//
	VmmqQueuePairs uint32

	//
	VMQOffloadWeight uint32

	//
	VrssEnabled bool

	//
	VrssExcludePrimaryProcessor bool

	//
	VrssIndependentHostSpreading bool

	//
	VrssMinQueuePairs uint32

	//
	VrssQueueSchedulingMode uint32

	//
	VrssVmbusChannelAffinityPolicy uint32
}

func NewMsvm_EthernetSwitchPortOffloadSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortOffloadSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortOffloadSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchPortOffloadSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchPortOffloadSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortOffloadSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

// SetIOVInterruptModeration sets the value of IOVInterruptModeration for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyIOVInterruptModeration(value EthernetSwitchPortOffloadSettingData_IOVInterruptModeration) (err error) {
	return instance.SetProperty("IOVInterruptModeration", (value))
}

// GetIOVInterruptModeration gets the value of IOVInterruptModeration for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyIOVInterruptModeration() (value EthernetSwitchPortOffloadSettingData_IOVInterruptModeration, err error) {
	retValue, err := instance.GetProperty("IOVInterruptModeration")
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

	value = EthernetSwitchPortOffloadSettingData_IOVInterruptModeration(valuetmp)

	return
}

// SetIOVOffloadWeight sets the value of IOVOffloadWeight for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyIOVOffloadWeight(value uint32) (err error) {
	return instance.SetProperty("IOVOffloadWeight", (value))
}

// GetIOVOffloadWeight gets the value of IOVOffloadWeight for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyIOVOffloadWeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("IOVOffloadWeight")
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

// SetIOVQueuePairsRequested sets the value of IOVQueuePairsRequested for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyIOVQueuePairsRequested(value uint32) (err error) {
	return instance.SetProperty("IOVQueuePairsRequested", (value))
}

// GetIOVQueuePairsRequested gets the value of IOVQueuePairsRequested for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyIOVQueuePairsRequested() (value uint32, err error) {
	retValue, err := instance.GetProperty("IOVQueuePairsRequested")
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

// SetIPSecOffloadLimit sets the value of IPSecOffloadLimit for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyIPSecOffloadLimit(value uint32) (err error) {
	return instance.SetProperty("IPSecOffloadLimit", (value))
}

// GetIPSecOffloadLimit gets the value of IPSecOffloadLimit for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyIPSecOffloadLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPSecOffloadLimit")
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

// SetPacketDirectModerationCount sets the value of PacketDirectModerationCount for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyPacketDirectModerationCount(value uint32) (err error) {
	return instance.SetProperty("PacketDirectModerationCount", (value))
}

// GetPacketDirectModerationCount gets the value of PacketDirectModerationCount for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyPacketDirectModerationCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketDirectModerationCount")
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

// SetPacketDirectModerationInterval sets the value of PacketDirectModerationInterval for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyPacketDirectModerationInterval(value uint32) (err error) {
	return instance.SetProperty("PacketDirectModerationInterval", (value))
}

// GetPacketDirectModerationInterval gets the value of PacketDirectModerationInterval for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyPacketDirectModerationInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketDirectModerationInterval")
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

// SetPacketDirectNumProcs sets the value of PacketDirectNumProcs for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyPacketDirectNumProcs(value uint32) (err error) {
	return instance.SetProperty("PacketDirectNumProcs", (value))
}

// GetPacketDirectNumProcs gets the value of PacketDirectNumProcs for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyPacketDirectNumProcs() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketDirectNumProcs")
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

// SetRscEnabled sets the value of RscEnabled for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyRscEnabled(value bool) (err error) {
	return instance.SetProperty("RscEnabled", (value))
}

// GetRscEnabled gets the value of RscEnabled for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyRscEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("RscEnabled")
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

// SetVmmqEnabled sets the value of VmmqEnabled for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyVmmqEnabled(value bool) (err error) {
	return instance.SetProperty("VmmqEnabled", (value))
}

// GetVmmqEnabled gets the value of VmmqEnabled for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyVmmqEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("VmmqEnabled")
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

// SetVmmqQueuePairs sets the value of VmmqQueuePairs for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyVmmqQueuePairs(value uint32) (err error) {
	return instance.SetProperty("VmmqQueuePairs", (value))
}

// GetVmmqQueuePairs gets the value of VmmqQueuePairs for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyVmmqQueuePairs() (value uint32, err error) {
	retValue, err := instance.GetProperty("VmmqQueuePairs")
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

// SetVMQOffloadWeight sets the value of VMQOffloadWeight for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyVMQOffloadWeight(value uint32) (err error) {
	return instance.SetProperty("VMQOffloadWeight", (value))
}

// GetVMQOffloadWeight gets the value of VMQOffloadWeight for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyVMQOffloadWeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("VMQOffloadWeight")
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

// SetVrssEnabled sets the value of VrssEnabled for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyVrssEnabled(value bool) (err error) {
	return instance.SetProperty("VrssEnabled", (value))
}

// GetVrssEnabled gets the value of VrssEnabled for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyVrssEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("VrssEnabled")
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

// SetVrssExcludePrimaryProcessor sets the value of VrssExcludePrimaryProcessor for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyVrssExcludePrimaryProcessor(value bool) (err error) {
	return instance.SetProperty("VrssExcludePrimaryProcessor", (value))
}

// GetVrssExcludePrimaryProcessor gets the value of VrssExcludePrimaryProcessor for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyVrssExcludePrimaryProcessor() (value bool, err error) {
	retValue, err := instance.GetProperty("VrssExcludePrimaryProcessor")
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

// SetVrssIndependentHostSpreading sets the value of VrssIndependentHostSpreading for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyVrssIndependentHostSpreading(value bool) (err error) {
	return instance.SetProperty("VrssIndependentHostSpreading", (value))
}

// GetVrssIndependentHostSpreading gets the value of VrssIndependentHostSpreading for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyVrssIndependentHostSpreading() (value bool, err error) {
	retValue, err := instance.GetProperty("VrssIndependentHostSpreading")
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

// SetVrssMinQueuePairs sets the value of VrssMinQueuePairs for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyVrssMinQueuePairs(value uint32) (err error) {
	return instance.SetProperty("VrssMinQueuePairs", (value))
}

// GetVrssMinQueuePairs gets the value of VrssMinQueuePairs for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyVrssMinQueuePairs() (value uint32, err error) {
	retValue, err := instance.GetProperty("VrssMinQueuePairs")
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

// SetVrssQueueSchedulingMode sets the value of VrssQueueSchedulingMode for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyVrssQueueSchedulingMode(value uint32) (err error) {
	return instance.SetProperty("VrssQueueSchedulingMode", (value))
}

// GetVrssQueueSchedulingMode gets the value of VrssQueueSchedulingMode for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyVrssQueueSchedulingMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("VrssQueueSchedulingMode")
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

// SetVrssVmbusChannelAffinityPolicy sets the value of VrssVmbusChannelAffinityPolicy for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) SetPropertyVrssVmbusChannelAffinityPolicy(value uint32) (err error) {
	return instance.SetProperty("VrssVmbusChannelAffinityPolicy", (value))
}

// GetVrssVmbusChannelAffinityPolicy gets the value of VrssVmbusChannelAffinityPolicy for the instance
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetPropertyVrssVmbusChannelAffinityPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("VrssVmbusChannelAffinityPolicy")
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
func (instance *Msvm_EthernetSwitchPortOffloadSettingData) GetRelatedEthernetPortAllocationSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetPortAllocationSettingData")
}
