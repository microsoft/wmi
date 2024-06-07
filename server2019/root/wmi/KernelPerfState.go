// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// KernelPerfState struct
type KernelPerfState struct { 
	*cim.WmiInstance

	// 
	Control uint64

	// 
	DecreaseLevel uint8

	// 
	DecreaseTime uint32

	// 
	Frequency uint32

	// 
	HitCount uint32

	// 
	IncreaseLevel uint8

	// 
	IncreaseTime uint32

	// 
	PercentFrequency uint8

	// 
	Power uint32

	// 
	Reserved1 uint32

	// 
	Reserved2 uint64

	// 
	Reserved3 uint64

	// 
	Status uint64

	// 
	Type uint8
}

	func NewKernelPerfStateEx1(instance *cim.WmiInstance) (newInstance *KernelPerfState, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &KernelPerfState {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewKernelPerfStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *KernelPerfState, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &KernelPerfState {
	WmiInstance: tmp,
	}
	return
	}
	

// SetControl sets the value of Control for the instance
func (instance *KernelPerfState) SetPropertyControl(value uint64) (err error) { 
    return instance.SetProperty("Control", (value))
}

// GetControl gets the value of Control for the instance
func (instance *KernelPerfState) GetPropertyControl()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Control")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetDecreaseLevel sets the value of DecreaseLevel for the instance
func (instance *KernelPerfState) SetPropertyDecreaseLevel(value uint8) (err error) { 
    return instance.SetProperty("DecreaseLevel", (value))
}

// GetDecreaseLevel gets the value of DecreaseLevel for the instance
func (instance *KernelPerfState) GetPropertyDecreaseLevel()(value uint8, err error) { 
    retValue, err := instance.GetProperty("DecreaseLevel")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetDecreaseTime sets the value of DecreaseTime for the instance
func (instance *KernelPerfState) SetPropertyDecreaseTime(value uint32) (err error) { 
    return instance.SetProperty("DecreaseTime", (value))
}

// GetDecreaseTime gets the value of DecreaseTime for the instance
func (instance *KernelPerfState) GetPropertyDecreaseTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DecreaseTime")
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

// SetFrequency sets the value of Frequency for the instance
func (instance *KernelPerfState) SetPropertyFrequency(value uint32) (err error) { 
    return instance.SetProperty("Frequency", (value))
}

// GetFrequency gets the value of Frequency for the instance
func (instance *KernelPerfState) GetPropertyFrequency()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Frequency")
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

// SetHitCount sets the value of HitCount for the instance
func (instance *KernelPerfState) SetPropertyHitCount(value uint32) (err error) { 
    return instance.SetProperty("HitCount", (value))
}

// GetHitCount gets the value of HitCount for the instance
func (instance *KernelPerfState) GetPropertyHitCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HitCount")
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

// SetIncreaseLevel sets the value of IncreaseLevel for the instance
func (instance *KernelPerfState) SetPropertyIncreaseLevel(value uint8) (err error) { 
    return instance.SetProperty("IncreaseLevel", (value))
}

// GetIncreaseLevel gets the value of IncreaseLevel for the instance
func (instance *KernelPerfState) GetPropertyIncreaseLevel()(value uint8, err error) { 
    retValue, err := instance.GetProperty("IncreaseLevel")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetIncreaseTime sets the value of IncreaseTime for the instance
func (instance *KernelPerfState) SetPropertyIncreaseTime(value uint32) (err error) { 
    return instance.SetProperty("IncreaseTime", (value))
}

// GetIncreaseTime gets the value of IncreaseTime for the instance
func (instance *KernelPerfState) GetPropertyIncreaseTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IncreaseTime")
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

// SetPercentFrequency sets the value of PercentFrequency for the instance
func (instance *KernelPerfState) SetPropertyPercentFrequency(value uint8) (err error) { 
    return instance.SetProperty("PercentFrequency", (value))
}

// GetPercentFrequency gets the value of PercentFrequency for the instance
func (instance *KernelPerfState) GetPropertyPercentFrequency()(value uint8, err error) { 
    retValue, err := instance.GetProperty("PercentFrequency")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetPower sets the value of Power for the instance
func (instance *KernelPerfState) SetPropertyPower(value uint32) (err error) { 
    return instance.SetProperty("Power", (value))
}

// GetPower gets the value of Power for the instance
func (instance *KernelPerfState) GetPropertyPower()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Power")
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

// SetReserved1 sets the value of Reserved1 for the instance
func (instance *KernelPerfState) SetPropertyReserved1(value uint32) (err error) { 
    return instance.SetProperty("Reserved1", (value))
}

// GetReserved1 gets the value of Reserved1 for the instance
func (instance *KernelPerfState) GetPropertyReserved1()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Reserved1")
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

// SetReserved2 sets the value of Reserved2 for the instance
func (instance *KernelPerfState) SetPropertyReserved2(value uint64) (err error) { 
    return instance.SetProperty("Reserved2", (value))
}

// GetReserved2 gets the value of Reserved2 for the instance
func (instance *KernelPerfState) GetPropertyReserved2()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Reserved2")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetReserved3 sets the value of Reserved3 for the instance
func (instance *KernelPerfState) SetPropertyReserved3(value uint64) (err error) { 
    return instance.SetProperty("Reserved3", (value))
}

// GetReserved3 gets the value of Reserved3 for the instance
func (instance *KernelPerfState) GetPropertyReserved3()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Reserved3")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetStatus sets the value of Status for the instance
func (instance *KernelPerfState) SetPropertyStatus(value uint64) (err error) { 
    return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *KernelPerfState) GetPropertyStatus()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Status")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetType sets the value of Type for the instance
func (instance *KernelPerfState) SetPropertyType(value uint8) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *KernelPerfState) GetPropertyType()(value uint8, err error) { 
    retValue, err := instance.GetProperty("Type")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

