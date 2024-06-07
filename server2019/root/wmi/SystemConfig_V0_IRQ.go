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

// SystemConfig_V0_IRQ struct
type SystemConfig_V0_IRQ struct { 
	*SystemConfig_V0

	// 
	DeviceDescription string

	// 
	DeviceDescriptionLen uint32

	// 
	IRQAffinity uint64

	// 
	IRQNum uint32
}

	func NewSystemConfig_V0_IRQEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V0_IRQ, err error) {tmp, err := NewSystemConfig_V0Ex1(instance)
		
	if err != nil { return }
	newInstance = &SystemConfig_V0_IRQ {
	SystemConfig_V0: tmp,
	}
	return
	}
	

	func NewSystemConfig_V0_IRQEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SystemConfig_V0_IRQ, err error) {tmp, err := NewSystemConfig_V0Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SystemConfig_V0_IRQ {
	SystemConfig_V0: tmp,
	}
	return
	}
	

// SetDeviceDescription sets the value of DeviceDescription for the instance
func (instance *SystemConfig_V0_IRQ) SetPropertyDeviceDescription(value string) (err error) { 
    return instance.SetProperty("DeviceDescription", (value))
}

// GetDeviceDescription gets the value of DeviceDescription for the instance
func (instance *SystemConfig_V0_IRQ) GetPropertyDeviceDescription()(value string, err error) { 
    retValue, err := instance.GetProperty("DeviceDescription")
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

// SetDeviceDescriptionLen sets the value of DeviceDescriptionLen for the instance
func (instance *SystemConfig_V0_IRQ) SetPropertyDeviceDescriptionLen(value uint32) (err error) { 
    return instance.SetProperty("DeviceDescriptionLen", (value))
}

// GetDeviceDescriptionLen gets the value of DeviceDescriptionLen for the instance
func (instance *SystemConfig_V0_IRQ) GetPropertyDeviceDescriptionLen()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DeviceDescriptionLen")
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

// SetIRQAffinity sets the value of IRQAffinity for the instance
func (instance *SystemConfig_V0_IRQ) SetPropertyIRQAffinity(value uint64) (err error) { 
    return instance.SetProperty("IRQAffinity", (value))
}

// GetIRQAffinity gets the value of IRQAffinity for the instance
func (instance *SystemConfig_V0_IRQ) GetPropertyIRQAffinity()(value uint64, err error) { 
    retValue, err := instance.GetProperty("IRQAffinity")
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

// SetIRQNum sets the value of IRQNum for the instance
func (instance *SystemConfig_V0_IRQ) SetPropertyIRQNum(value uint32) (err error) { 
    return instance.SetProperty("IRQNum", (value))
}

// GetIRQNum gets the value of IRQNum for the instance
func (instance *SystemConfig_V0_IRQ) GetPropertyIRQNum()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IRQNum")
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

