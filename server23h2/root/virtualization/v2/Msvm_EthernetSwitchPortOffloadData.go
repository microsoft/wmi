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

// Msvm_EthernetSwitchPortOffloadData struct
type Msvm_EthernetSwitchPortOffloadData struct {
	*Msvm_EthernetPortData

	//
	IovOffloadUsage uint32

	//
	IovQueuePairUsage uint32

	//
	IovVfDataPathActive bool

	//
	IovVfId uint16

	//
	IpsecCurrentOffloadSaCount uint32

	//
	RscEnabled bool

	//
	VmmqEnabled bool

	//
	VmmqQueuePairs uint32

	//
	VMQId uint32

	//
	VMQOffloadUsage uint32

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

func NewMsvm_EthernetSwitchPortOffloadDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortOffloadData, err error) {
	tmp, err := NewMsvm_EthernetPortDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortOffloadData{
		Msvm_EthernetPortData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchPortOffloadDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchPortOffloadData, err error) {
	tmp, err := NewMsvm_EthernetPortDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortOffloadData{
		Msvm_EthernetPortData: tmp,
	}
	return
}

// SetIovOffloadUsage sets the value of IovOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyIovOffloadUsage(value uint32) (err error) {
	return instance.SetProperty("IovOffloadUsage", (value))
}

// GetIovOffloadUsage gets the value of IovOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyIovOffloadUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("IovOffloadUsage")
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

// SetIovQueuePairUsage sets the value of IovQueuePairUsage for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyIovQueuePairUsage(value uint32) (err error) {
	return instance.SetProperty("IovQueuePairUsage", (value))
}

// GetIovQueuePairUsage gets the value of IovQueuePairUsage for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyIovQueuePairUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("IovQueuePairUsage")
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

// SetIovVfDataPathActive sets the value of IovVfDataPathActive for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyIovVfDataPathActive(value bool) (err error) {
	return instance.SetProperty("IovVfDataPathActive", (value))
}

// GetIovVfDataPathActive gets the value of IovVfDataPathActive for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyIovVfDataPathActive() (value bool, err error) {
	retValue, err := instance.GetProperty("IovVfDataPathActive")
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

// SetIovVfId sets the value of IovVfId for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyIovVfId(value uint16) (err error) {
	return instance.SetProperty("IovVfId", (value))
}

// GetIovVfId gets the value of IovVfId for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyIovVfId() (value uint16, err error) {
	retValue, err := instance.GetProperty("IovVfId")
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

// SetIpsecCurrentOffloadSaCount sets the value of IpsecCurrentOffloadSaCount for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyIpsecCurrentOffloadSaCount(value uint32) (err error) {
	return instance.SetProperty("IpsecCurrentOffloadSaCount", (value))
}

// GetIpsecCurrentOffloadSaCount gets the value of IpsecCurrentOffloadSaCount for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyIpsecCurrentOffloadSaCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpsecCurrentOffloadSaCount")
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
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyRscEnabled(value bool) (err error) {
	return instance.SetProperty("RscEnabled", (value))
}

// GetRscEnabled gets the value of RscEnabled for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyRscEnabled() (value bool, err error) {
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
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVmmqEnabled(value bool) (err error) {
	return instance.SetProperty("VmmqEnabled", (value))
}

// GetVmmqEnabled gets the value of VmmqEnabled for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVmmqEnabled() (value bool, err error) {
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
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVmmqQueuePairs(value uint32) (err error) {
	return instance.SetProperty("VmmqQueuePairs", (value))
}

// GetVmmqQueuePairs gets the value of VmmqQueuePairs for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVmmqQueuePairs() (value uint32, err error) {
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

// SetVMQId sets the value of VMQId for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVMQId(value uint32) (err error) {
	return instance.SetProperty("VMQId", (value))
}

// GetVMQId gets the value of VMQId for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVMQId() (value uint32, err error) {
	retValue, err := instance.GetProperty("VMQId")
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

// SetVMQOffloadUsage sets the value of VMQOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVMQOffloadUsage(value uint32) (err error) {
	return instance.SetProperty("VMQOffloadUsage", (value))
}

// GetVMQOffloadUsage gets the value of VMQOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVMQOffloadUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("VMQOffloadUsage")
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
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVrssEnabled(value bool) (err error) {
	return instance.SetProperty("VrssEnabled", (value))
}

// GetVrssEnabled gets the value of VrssEnabled for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVrssEnabled() (value bool, err error) {
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
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVrssExcludePrimaryProcessor(value bool) (err error) {
	return instance.SetProperty("VrssExcludePrimaryProcessor", (value))
}

// GetVrssExcludePrimaryProcessor gets the value of VrssExcludePrimaryProcessor for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVrssExcludePrimaryProcessor() (value bool, err error) {
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
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVrssIndependentHostSpreading(value bool) (err error) {
	return instance.SetProperty("VrssIndependentHostSpreading", (value))
}

// GetVrssIndependentHostSpreading gets the value of VrssIndependentHostSpreading for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVrssIndependentHostSpreading() (value bool, err error) {
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
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVrssMinQueuePairs(value uint32) (err error) {
	return instance.SetProperty("VrssMinQueuePairs", (value))
}

// GetVrssMinQueuePairs gets the value of VrssMinQueuePairs for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVrssMinQueuePairs() (value uint32, err error) {
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
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVrssQueueSchedulingMode(value uint32) (err error) {
	return instance.SetProperty("VrssQueueSchedulingMode", (value))
}

// GetVrssQueueSchedulingMode gets the value of VrssQueueSchedulingMode for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVrssQueueSchedulingMode() (value uint32, err error) {
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
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVrssVmbusChannelAffinityPolicy(value uint32) (err error) {
	return instance.SetProperty("VrssVmbusChannelAffinityPolicy", (value))
}

// GetVrssVmbusChannelAffinityPolicy gets the value of VrssVmbusChannelAffinityPolicy for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVrssVmbusChannelAffinityPolicy() (value uint32, err error) {
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
func (instance *Msvm_EthernetSwitchPortOffloadData) GetRelatedEthernetSwitchPort() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchPort")
}
