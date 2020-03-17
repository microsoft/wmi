// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_EthernetSwitchPortOffloadData struct
type Msvm_EthernetSwitchPortOffloadData struct {
	Msvm_EthernetPortData

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

// SetIovOffloadUsage sets the value of IovOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyIovOffloadUsage(value uint32) (err error) {
	return instance.SetProperty("IovOffloadUsage", value)
}

// GetIovOffloadUsage gets the value of IovOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyIovOffloadUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("IovOffloadUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIovQueuePairUsage sets the value of IovQueuePairUsage for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyIovQueuePairUsage(value uint32) (err error) {
	return instance.SetProperty("IovQueuePairUsage", value)
}

// GetIovQueuePairUsage gets the value of IovQueuePairUsage for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyIovQueuePairUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("IovQueuePairUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIovVfDataPathActive sets the value of IovVfDataPathActive for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyIovVfDataPathActive(value bool) (err error) {
	return instance.SetProperty("IovVfDataPathActive", value)
}

// GetIovVfDataPathActive gets the value of IovVfDataPathActive for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyIovVfDataPathActive() (value bool, err error) {
	retValue, err := instance.GetProperty("IovVfDataPathActive")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIovVfId sets the value of IovVfId for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyIovVfId(value uint16) (err error) {
	return instance.SetProperty("IovVfId", value)
}

// GetIovVfId gets the value of IovVfId for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyIovVfId() (value uint16, err error) {
	retValue, err := instance.GetProperty("IovVfId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIpsecCurrentOffloadSaCount sets the value of IpsecCurrentOffloadSaCount for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyIpsecCurrentOffloadSaCount(value uint32) (err error) {
	return instance.SetProperty("IpsecCurrentOffloadSaCount", value)
}

// GetIpsecCurrentOffloadSaCount gets the value of IpsecCurrentOffloadSaCount for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyIpsecCurrentOffloadSaCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpsecCurrentOffloadSaCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmmqEnabled sets the value of VmmqEnabled for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVmmqEnabled(value bool) (err error) {
	return instance.SetProperty("VmmqEnabled", value)
}

// GetVmmqEnabled gets the value of VmmqEnabled for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVmmqEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("VmmqEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmmqQueuePairs sets the value of VmmqQueuePairs for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVmmqQueuePairs(value uint32) (err error) {
	return instance.SetProperty("VmmqQueuePairs", value)
}

// GetVmmqQueuePairs gets the value of VmmqQueuePairs for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVmmqQueuePairs() (value uint32, err error) {
	retValue, err := instance.GetProperty("VmmqQueuePairs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVMQId sets the value of VMQId for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVMQId(value uint32) (err error) {
	return instance.SetProperty("VMQId", value)
}

// GetVMQId gets the value of VMQId for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVMQId() (value uint32, err error) {
	retValue, err := instance.GetProperty("VMQId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVMQOffloadUsage sets the value of VMQOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVMQOffloadUsage(value uint32) (err error) {
	return instance.SetProperty("VMQOffloadUsage", value)
}

// GetVMQOffloadUsage gets the value of VMQOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVMQOffloadUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("VMQOffloadUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVrssEnabled sets the value of VrssEnabled for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVrssEnabled(value bool) (err error) {
	return instance.SetProperty("VrssEnabled", value)
}

// GetVrssEnabled gets the value of VrssEnabled for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVrssEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("VrssEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVrssExcludePrimaryProcessor sets the value of VrssExcludePrimaryProcessor for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVrssExcludePrimaryProcessor(value bool) (err error) {
	return instance.SetProperty("VrssExcludePrimaryProcessor", value)
}

// GetVrssExcludePrimaryProcessor gets the value of VrssExcludePrimaryProcessor for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVrssExcludePrimaryProcessor() (value bool, err error) {
	retValue, err := instance.GetProperty("VrssExcludePrimaryProcessor")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVrssIndependentHostSpreading sets the value of VrssIndependentHostSpreading for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVrssIndependentHostSpreading(value bool) (err error) {
	return instance.SetProperty("VrssIndependentHostSpreading", value)
}

// GetVrssIndependentHostSpreading gets the value of VrssIndependentHostSpreading for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVrssIndependentHostSpreading() (value bool, err error) {
	retValue, err := instance.GetProperty("VrssIndependentHostSpreading")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVrssMinQueuePairs sets the value of VrssMinQueuePairs for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVrssMinQueuePairs(value uint32) (err error) {
	return instance.SetProperty("VrssMinQueuePairs", value)
}

// GetVrssMinQueuePairs gets the value of VrssMinQueuePairs for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVrssMinQueuePairs() (value uint32, err error) {
	retValue, err := instance.GetProperty("VrssMinQueuePairs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVrssQueueSchedulingMode sets the value of VrssQueueSchedulingMode for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVrssQueueSchedulingMode(value uint32) (err error) {
	return instance.SetProperty("VrssQueueSchedulingMode", value)
}

// GetVrssQueueSchedulingMode gets the value of VrssQueueSchedulingMode for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVrssQueueSchedulingMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("VrssQueueSchedulingMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVrssVmbusChannelAffinityPolicy sets the value of VrssVmbusChannelAffinityPolicy for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) SetPropertyVrssVmbusChannelAffinityPolicy(value uint32) (err error) {
	return instance.SetProperty("VrssVmbusChannelAffinityPolicy", value)
}

// GetVrssVmbusChannelAffinityPolicy gets the value of VrssVmbusChannelAffinityPolicy for the instance
func (instance *Msvm_EthernetSwitchPortOffloadData) GetPropertyVrssVmbusChannelAffinityPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("VrssVmbusChannelAffinityPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchPortOffloadData) GetRelatedEthernetSwitchPort() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchPort")
}
