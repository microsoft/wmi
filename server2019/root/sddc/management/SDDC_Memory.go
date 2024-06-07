// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// SDDC_Memory struct
type SDDC_Memory struct { 
	*cim.WmiInstance

	// 
	Manufacturer string

	// 
	Model string

	// 
	SerialNumber string

	// 
	SizeInBytes uint64

	// 
	SpeedInMHz uint32
}

	func NewSDDC_MemoryEx1(instance *cim.WmiInstance) (newInstance *SDDC_Memory, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &SDDC_Memory {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewSDDC_MemoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SDDC_Memory, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SDDC_Memory {
	WmiInstance: tmp,
	}
	return
	}
	

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *SDDC_Memory) SetPropertyManufacturer(value string) (err error) { 
    return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *SDDC_Memory) GetPropertyManufacturer()(value string, err error) { 
    retValue, err := instance.GetProperty("Manufacturer")
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

// SetModel sets the value of Model for the instance
func (instance *SDDC_Memory) SetPropertyModel(value string) (err error) { 
    return instance.SetProperty("Model", (value))
}

// GetModel gets the value of Model for the instance
func (instance *SDDC_Memory) GetPropertyModel()(value string, err error) { 
    retValue, err := instance.GetProperty("Model")
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

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *SDDC_Memory) SetPropertySerialNumber(value string) (err error) { 
    return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *SDDC_Memory) GetPropertySerialNumber()(value string, err error) { 
    retValue, err := instance.GetProperty("SerialNumber")
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

// SetSizeInBytes sets the value of SizeInBytes for the instance
func (instance *SDDC_Memory) SetPropertySizeInBytes(value uint64) (err error) { 
    return instance.SetProperty("SizeInBytes", (value))
}

// GetSizeInBytes gets the value of SizeInBytes for the instance
func (instance *SDDC_Memory) GetPropertySizeInBytes()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SizeInBytes")
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

// SetSpeedInMHz sets the value of SpeedInMHz for the instance
func (instance *SDDC_Memory) SetPropertySpeedInMHz(value uint32) (err error) { 
    return instance.SetProperty("SpeedInMHz", (value))
}

// GetSpeedInMHz gets the value of SpeedInMHz for the instance
func (instance *SDDC_Memory) GetPropertySpeedInMHz()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SpeedInMHz")
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

