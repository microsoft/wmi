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
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSSerial_HardwareConfiguration struct
type MSSerial_HardwareConfiguration struct { 
	*MSSerial

	// 
	Active bool

	// 
	BaseIOAddress uint64

	// 
	InstanceName string

	// 
	InterruptType uint32

	// 
	IrqAffinityMask uint64

	// 
	IrqLevel uint32

	// 
	IrqNumber uint32

	// 
	IrqVector uint32
}

	func NewMSSerial_HardwareConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSSerial_HardwareConfiguration, err error) {tmp, err := NewMSSerialEx1(instance)
		
	if err != nil { return }
	newInstance = &MSSerial_HardwareConfiguration {
	MSSerial: tmp,
	}
	return
	}
	

	func NewMSSerial_HardwareConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSSerial_HardwareConfiguration, err error) {tmp, err := NewMSSerialEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSSerial_HardwareConfiguration {
	MSSerial: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSSerial_HardwareConfiguration) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSSerial_HardwareConfiguration) GetPropertyActive()(value bool, err error) { 
    retValue, err := instance.GetProperty("Active")
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

// SetBaseIOAddress sets the value of BaseIOAddress for the instance
func (instance *MSSerial_HardwareConfiguration) SetPropertyBaseIOAddress(value uint64) (err error) { 
    return instance.SetProperty("BaseIOAddress", (value))
}

// GetBaseIOAddress gets the value of BaseIOAddress for the instance
func (instance *MSSerial_HardwareConfiguration) GetPropertyBaseIOAddress()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BaseIOAddress")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSSerial_HardwareConfiguration) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSSerial_HardwareConfiguration) GetPropertyInstanceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceName")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetInterruptType sets the value of InterruptType for the instance
func (instance *MSSerial_HardwareConfiguration) SetPropertyInterruptType(value uint32) (err error) { 
    return instance.SetProperty("InterruptType", (value))
}

// GetInterruptType gets the value of InterruptType for the instance
func (instance *MSSerial_HardwareConfiguration) GetPropertyInterruptType()(value uint32, err error) { 
    retValue, err := instance.GetProperty("InterruptType")
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

// SetIrqAffinityMask sets the value of IrqAffinityMask for the instance
func (instance *MSSerial_HardwareConfiguration) SetPropertyIrqAffinityMask(value uint64) (err error) { 
    return instance.SetProperty("IrqAffinityMask", (value))
}

// GetIrqAffinityMask gets the value of IrqAffinityMask for the instance
func (instance *MSSerial_HardwareConfiguration) GetPropertyIrqAffinityMask()(value uint64, err error) { 
    retValue, err := instance.GetProperty("IrqAffinityMask")
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

// SetIrqLevel sets the value of IrqLevel for the instance
func (instance *MSSerial_HardwareConfiguration) SetPropertyIrqLevel(value uint32) (err error) { 
    return instance.SetProperty("IrqLevel", (value))
}

// GetIrqLevel gets the value of IrqLevel for the instance
func (instance *MSSerial_HardwareConfiguration) GetPropertyIrqLevel()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IrqLevel")
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

// SetIrqNumber sets the value of IrqNumber for the instance
func (instance *MSSerial_HardwareConfiguration) SetPropertyIrqNumber(value uint32) (err error) { 
    return instance.SetProperty("IrqNumber", (value))
}

// GetIrqNumber gets the value of IrqNumber for the instance
func (instance *MSSerial_HardwareConfiguration) GetPropertyIrqNumber()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IrqNumber")
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

// SetIrqVector sets the value of IrqVector for the instance
func (instance *MSSerial_HardwareConfiguration) SetPropertyIrqVector(value uint32) (err error) { 
    return instance.SetProperty("IrqVector", (value))
}

// GetIrqVector gets the value of IrqVector for the instance
func (instance *MSSerial_HardwareConfiguration) GetPropertyIrqVector()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IrqVector")
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

