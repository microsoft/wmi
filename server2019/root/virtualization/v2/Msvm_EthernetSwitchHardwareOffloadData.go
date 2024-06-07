// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Msvm_EthernetSwitchHardwareOffloadData struct
type Msvm_EthernetSwitchHardwareOffloadData struct { 
	*Msvm_EthernetSwitchData

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

	func NewMsvm_EthernetSwitchHardwareOffloadDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchHardwareOffloadData, err error) {tmp, err := NewMsvm_EthernetSwitchDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetSwitchHardwareOffloadData {
	Msvm_EthernetSwitchData: tmp,
	}
	return
	}
	

	func NewMsvm_EthernetSwitchHardwareOffloadDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_EthernetSwitchHardwareOffloadData, err error) {tmp, err := NewMsvm_EthernetSwitchDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetSwitchHardwareOffloadData {
	Msvm_EthernetSwitchData: tmp,
	}
	return
	}
	

// SetDefaultQueueVmmqEnabled sets the value of DefaultQueueVmmqEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVmmqEnabled(value bool) (err error) { 
    return instance.SetProperty("DefaultQueueVmmqEnabled", (value))
}

// GetDefaultQueueVmmqEnabled gets the value of DefaultQueueVmmqEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVmmqEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("DefaultQueueVmmqEnabled")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetDefaultQueueVmmqQueuePairs sets the value of DefaultQueueVmmqQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVmmqQueuePairs(value uint32) (err error) { 
    return instance.SetProperty("DefaultQueueVmmqQueuePairs", (value))
}

// GetDefaultQueueVmmqQueuePairs gets the value of DefaultQueueVmmqQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVmmqQueuePairs()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DefaultQueueVmmqQueuePairs")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetDefaultQueueVrssEnabled sets the value of DefaultQueueVrssEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVrssEnabled(value bool) (err error) { 
    return instance.SetProperty("DefaultQueueVrssEnabled", (value))
}

// GetDefaultQueueVrssEnabled gets the value of DefaultQueueVrssEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVrssEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("DefaultQueueVrssEnabled")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetDefaultQueueVrssExcludePrimaryProcessor sets the value of DefaultQueueVrssExcludePrimaryProcessor for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVrssExcludePrimaryProcessor(value bool) (err error) { 
    return instance.SetProperty("DefaultQueueVrssExcludePrimaryProcessor", (value))
}

// GetDefaultQueueVrssExcludePrimaryProcessor gets the value of DefaultQueueVrssExcludePrimaryProcessor for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVrssExcludePrimaryProcessor()(value bool, err error) { 
    retValue, err := instance.GetProperty("DefaultQueueVrssExcludePrimaryProcessor")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetDefaultQueueVrssIndependentHostSpreading sets the value of DefaultQueueVrssIndependentHostSpreading for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVrssIndependentHostSpreading(value bool) (err error) { 
    return instance.SetProperty("DefaultQueueVrssIndependentHostSpreading", (value))
}

// GetDefaultQueueVrssIndependentHostSpreading gets the value of DefaultQueueVrssIndependentHostSpreading for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVrssIndependentHostSpreading()(value bool, err error) { 
    retValue, err := instance.GetProperty("DefaultQueueVrssIndependentHostSpreading")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetDefaultQueueVrssMinQueuePairs sets the value of DefaultQueueVrssMinQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVrssMinQueuePairs(value uint32) (err error) { 
    return instance.SetProperty("DefaultQueueVrssMinQueuePairs", (value))
}

// GetDefaultQueueVrssMinQueuePairs gets the value of DefaultQueueVrssMinQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVrssMinQueuePairs()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DefaultQueueVrssMinQueuePairs")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetDefaultQueueVrssQueueSchedulingMode sets the value of DefaultQueueVrssQueueSchedulingMode for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyDefaultQueueVrssQueueSchedulingMode(value uint32) (err error) { 
    return instance.SetProperty("DefaultQueueVrssQueueSchedulingMode", (value))
}

// GetDefaultQueueVrssQueueSchedulingMode gets the value of DefaultQueueVrssQueueSchedulingMode for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyDefaultQueueVrssQueueSchedulingMode()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DefaultQueueVrssQueueSchedulingMode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetIovQueuePairCapacity sets the value of IovQueuePairCapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyIovQueuePairCapacity(value uint32) (err error) { 
    return instance.SetProperty("IovQueuePairCapacity", (value))
}

// GetIovQueuePairCapacity gets the value of IovQueuePairCapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyIovQueuePairCapacity()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IovQueuePairCapacity")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetIovQueuePairUsage sets the value of IovQueuePairUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyIovQueuePairUsage(value uint32) (err error) { 
    return instance.SetProperty("IovQueuePairUsage", (value))
}

// GetIovQueuePairUsage gets the value of IovQueuePairUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyIovQueuePairUsage()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IovQueuePairUsage")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetIovVfCapacity sets the value of IovVfCapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyIovVfCapacity(value uint32) (err error) { 
    return instance.SetProperty("IovVfCapacity", (value))
}

// GetIovVfCapacity gets the value of IovVfCapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyIovVfCapacity()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IovVfCapacity")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetIovVfUsage sets the value of IovVfUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyIovVfUsage(value uint32) (err error) { 
    return instance.SetProperty("IovVfUsage", (value))
}

// GetIovVfUsage gets the value of IovVfUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyIovVfUsage()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IovVfUsage")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetIPsecSACapacity sets the value of IPsecSACapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyIPsecSACapacity(value uint32) (err error) { 
    return instance.SetProperty("IPsecSACapacity", (value))
}

// GetIPsecSACapacity gets the value of IPsecSACapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyIPsecSACapacity()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IPsecSACapacity")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetIPsecSAUsage sets the value of IPsecSAUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyIPsecSAUsage(value uint32) (err error) { 
    return instance.SetProperty("IPsecSAUsage", (value))
}

// GetIPsecSAUsage gets the value of IPsecSAUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyIPsecSAUsage()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IPsecSAUsage")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetPacketDirectInUse sets the value of PacketDirectInUse for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyPacketDirectInUse(value bool) (err error) { 
    return instance.SetProperty("PacketDirectInUse", (value))
}

// GetPacketDirectInUse gets the value of PacketDirectInUse for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyPacketDirectInUse()(value bool, err error) { 
    retValue, err := instance.GetProperty("PacketDirectInUse")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetVmqCapacity sets the value of VmqCapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyVmqCapacity(value uint32) (err error) { 
    return instance.SetProperty("VmqCapacity", (value))
}

// GetVmqCapacity gets the value of VmqCapacity for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyVmqCapacity()(value uint32, err error) { 
    retValue, err := instance.GetProperty("VmqCapacity")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetVmqUsage sets the value of VmqUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) SetPropertyVmqUsage(value uint32) (err error) { 
    return instance.SetProperty("VmqUsage", (value))
}

// GetVmqUsage gets the value of VmqUsage for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadData) GetPropertyVmqUsage()(value uint32, err error) { 
    retValue, err := instance.GetProperty("VmqUsage")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}
func  (instance* Msvm_EthernetSwitchHardwareOffloadData) GetRelatedVirtualEthernetSwitch() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_VirtualEthernetSwitch"); 
	}
	

