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

// Msvm_EthernetSwitchHardwareOffloadData struct
type Msvm_EthernetSwitchHardwareOffloadData struct {
	Msvm_EthernetSwitchData

	//
	DefaultQueueVmmqEnabled bool

	//
	DefaultQueueVmmqQueuePairs uint32

	//
	DefaultQueueVrssEnabled bool

	//
	DefaultQueueVrssExcludePrimaryProcessor bool

	//
	DefaultQueueVrssIndependentHostSpreading bool

	//
	DefaultQueueVrssMinQueuePairs uint32

	//
	DefaultQueueVrssQueueSchedulingMode uint32

	//
	IovQueuePairCapacity uint32

	//
	IovQueuePairUsage uint32

	//
	IovVfCapacity uint32

	//
	IovVfUsage uint32

	//
	IPsecSACapacity uint32

	//
	IPsecSAUsage uint32

	//
	PacketDirectInUse bool

	//
	VmqCapacity uint32

	//
	VmqUsage uint32
}

// SetDefaultQueueVmmqEnabled sets the value of DefaultQueueVmmqEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVmmqEnabled(value bool) (err error) {
	return instance.SetProperty("DefaultQueueVmmqEnabled", value)
}

// GetDefaultQueueVmmqEnabled gets the value of DefaultQueueVmmqEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVmmqEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVmmqEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultQueueVmmqQueuePairs sets the value of DefaultQueueVmmqQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVmmqQueuePairs(value uint32) (err error) {
	return instance.SetProperty("DefaultQueueVmmqQueuePairs", value)
}

// GetDefaultQueueVmmqQueuePairs gets the value of DefaultQueueVmmqQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVmmqQueuePairs() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVmmqQueuePairs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultQueueVrssEnabled sets the value of DefaultQueueVrssEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVrssEnabled(value bool) (err error) {
	return instance.SetProperty("DefaultQueueVrssEnabled", value)
}

// GetDefaultQueueVrssEnabled gets the value of DefaultQueueVrssEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVrssEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultQueueVrssExcludePrimaryProcessor sets the value of DefaultQueueVrssExcludePrimaryProcessor for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVrssExcludePrimaryProcessor(value bool) (err error) {
	return instance.SetProperty("DefaultQueueVrssExcludePrimaryProcessor", value)
}

// GetDefaultQueueVrssExcludePrimaryProcessor gets the value of DefaultQueueVrssExcludePrimaryProcessor for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVrssExcludePrimaryProcessor() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssExcludePrimaryProcessor")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultQueueVrssIndependentHostSpreading sets the value of DefaultQueueVrssIndependentHostSpreading for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVrssIndependentHostSpreading(value bool) (err error) {
	return instance.SetProperty("DefaultQueueVrssIndependentHostSpreading", value)
}

// GetDefaultQueueVrssIndependentHostSpreading gets the value of DefaultQueueVrssIndependentHostSpreading for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVrssIndependentHostSpreading() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssIndependentHostSpreading")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultQueueVrssMinQueuePairs sets the value of DefaultQueueVrssMinQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVrssMinQueuePairs(value uint32) (err error) {
	return instance.SetProperty("DefaultQueueVrssMinQueuePairs", value)
}

// GetDefaultQueueVrssMinQueuePairs gets the value of DefaultQueueVrssMinQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVrssMinQueuePairs() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssMinQueuePairs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultQueueVrssQueueSchedulingMode sets the value of DefaultQueueVrssQueueSchedulingMode for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVrssQueueSchedulingMode(value uint32) (err error) {
	return instance.SetProperty("DefaultQueueVrssQueueSchedulingMode", value)
}

// GetDefaultQueueVrssQueueSchedulingMode gets the value of DefaultQueueVrssQueueSchedulingMode for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVrssQueueSchedulingMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssQueueSchedulingMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIovQueuePairCapacity sets the value of IovQueuePairCapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyIovQueuePairCapacity(value uint32) (err error) {
	return instance.SetProperty("IovQueuePairCapacity", value)
}

// GetIovQueuePairCapacity gets the value of IovQueuePairCapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyIovQueuePairCapacity() (value uint32, err error) {
	retValue, err := instance.GetProperty("IovQueuePairCapacity")
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
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyIovQueuePairUsage(value uint32) (err error) {
	return instance.SetProperty("IovQueuePairUsage", value)
}

// GetIovQueuePairUsage gets the value of IovQueuePairUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyIovQueuePairUsage() (value uint32, err error) {
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

// SetIovVfCapacity sets the value of IovVfCapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyIovVfCapacity(value uint32) (err error) {
	return instance.SetProperty("IovVfCapacity", value)
}

// GetIovVfCapacity gets the value of IovVfCapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyIovVfCapacity() (value uint32, err error) {
	retValue, err := instance.GetProperty("IovVfCapacity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIovVfUsage sets the value of IovVfUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyIovVfUsage(value uint32) (err error) {
	return instance.SetProperty("IovVfUsage", value)
}

// GetIovVfUsage gets the value of IovVfUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyIovVfUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("IovVfUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPsecSACapacity sets the value of IPsecSACapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyIPsecSACapacity(value uint32) (err error) {
	return instance.SetProperty("IPsecSACapacity", value)
}

// GetIPsecSACapacity gets the value of IPsecSACapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyIPsecSACapacity() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPsecSACapacity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPsecSAUsage sets the value of IPsecSAUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyIPsecSAUsage(value uint32) (err error) {
	return instance.SetProperty("IPsecSAUsage", value)
}

// GetIPsecSAUsage gets the value of IPsecSAUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyIPsecSAUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPsecSAUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketDirectInUse sets the value of PacketDirectInUse for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyPacketDirectInUse(value bool) (err error) {
	return instance.SetProperty("PacketDirectInUse", value)
}

// GetPacketDirectInUse gets the value of PacketDirectInUse for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyPacketDirectInUse() (value bool, err error) {
	retValue, err := instance.GetProperty("PacketDirectInUse")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmqCapacity sets the value of VmqCapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyVmqCapacity(value uint32) (err error) {
	return instance.SetProperty("VmqCapacity", value)
}

// GetVmqCapacity gets the value of VmqCapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyVmqCapacity() (value uint32, err error) {
	retValue, err := instance.GetProperty("VmqCapacity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmqUsage sets the value of VmqUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyVmqUsage(value uint32) (err error) {
	return instance.SetProperty("VmqUsage", value)
}

// GetVmqUsage gets the value of VmqUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyVmqUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("VmqUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetRelatedVirtualEthernetSwitch() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitch")
}
