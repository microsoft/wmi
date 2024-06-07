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

// ProcessorAcpiTssState struct
type ProcessorAcpiTssState struct { 
	*cim.WmiInstance

	// 
	Control uint32

	// 
	FreqPercentage uint32

	// 
	Power uint32

	// 
	Reserved1 uint64

	// 
	Reserved2 uint64

	// 
	Status uint32

	// 
	TransitionLatency uint32
}

	func NewProcessorAcpiTssStateEx1(instance *cim.WmiInstance) (newInstance *ProcessorAcpiTssState, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &ProcessorAcpiTssState {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewProcessorAcpiTssStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ProcessorAcpiTssState, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ProcessorAcpiTssState {
	WmiInstance: tmp,
	}
	return
	}
	

// SetControl sets the value of Control for the instance
func (instance *ProcessorAcpiTssState) SetPropertyControl(value uint32) (err error) { 
    return instance.SetProperty("Control", (value))
}

// GetControl gets the value of Control for the instance
func (instance *ProcessorAcpiTssState) GetPropertyControl()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Control")
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

// SetFreqPercentage sets the value of FreqPercentage for the instance
func (instance *ProcessorAcpiTssState) SetPropertyFreqPercentage(value uint32) (err error) { 
    return instance.SetProperty("FreqPercentage", (value))
}

// GetFreqPercentage gets the value of FreqPercentage for the instance
func (instance *ProcessorAcpiTssState) GetPropertyFreqPercentage()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FreqPercentage")
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

// SetPower sets the value of Power for the instance
func (instance *ProcessorAcpiTssState) SetPropertyPower(value uint32) (err error) { 
    return instance.SetProperty("Power", (value))
}

// GetPower gets the value of Power for the instance
func (instance *ProcessorAcpiTssState) GetPropertyPower()(value uint32, err error) { 
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
func (instance *ProcessorAcpiTssState) SetPropertyReserved1(value uint64) (err error) { 
    return instance.SetProperty("Reserved1", (value))
}

// GetReserved1 gets the value of Reserved1 for the instance
func (instance *ProcessorAcpiTssState) GetPropertyReserved1()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Reserved1")
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

// SetReserved2 sets the value of Reserved2 for the instance
func (instance *ProcessorAcpiTssState) SetPropertyReserved2(value uint64) (err error) { 
    return instance.SetProperty("Reserved2", (value))
}

// GetReserved2 gets the value of Reserved2 for the instance
func (instance *ProcessorAcpiTssState) GetPropertyReserved2()(value uint64, err error) { 
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

// SetStatus sets the value of Status for the instance
func (instance *ProcessorAcpiTssState) SetPropertyStatus(value uint32) (err error) { 
    return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *ProcessorAcpiTssState) GetPropertyStatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Status")
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

// SetTransitionLatency sets the value of TransitionLatency for the instance
func (instance *ProcessorAcpiTssState) SetPropertyTransitionLatency(value uint32) (err error) { 
    return instance.SetProperty("TransitionLatency", (value))
}

// GetTransitionLatency gets the value of TransitionLatency for the instance
func (instance *ProcessorAcpiTssState) GetPropertyTransitionLatency()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TransitionLatency")
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

