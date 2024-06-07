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

// ProcessorAcpiTsd struct
type ProcessorAcpiTsd struct { 
	*MSProcessorClass

	// 
	Active bool

	// 
	Count uint32

	// 
	Dependency []ProcessorAcpiTsdDependency

	// 
	InstanceName string
}

	func NewProcessorAcpiTsdEx1(instance *cim.WmiInstance) (newInstance *ProcessorAcpiTsd, err error) {tmp, err := NewMSProcessorClassEx1(instance)
		
	if err != nil { return }
	newInstance = &ProcessorAcpiTsd {
	MSProcessorClass: tmp,
	}
	return
	}
	

	func NewProcessorAcpiTsdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ProcessorAcpiTsd, err error) {tmp, err := NewMSProcessorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ProcessorAcpiTsd {
	MSProcessorClass: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *ProcessorAcpiTsd) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *ProcessorAcpiTsd) GetPropertyActive()(value bool, err error) { 
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

// SetCount sets the value of Count for the instance
func (instance *ProcessorAcpiTsd) SetPropertyCount(value uint32) (err error) { 
    return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *ProcessorAcpiTsd) GetPropertyCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Count")
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

// SetDependency sets the value of Dependency for the instance
func (instance *ProcessorAcpiTsd) SetPropertyDependency(value []ProcessorAcpiTsdDependency) (err error) { 
    return instance.SetProperty("Dependency", (value))
}

// GetDependency gets the value of Dependency for the instance
func (instance *ProcessorAcpiTsd) GetPropertyDependency()(value []ProcessorAcpiTsdDependency, err error) { 
    retValue, err := instance.GetProperty("Dependency")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(ProcessorAcpiTsdDependency); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " ProcessorAcpiTsdDependency is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, ProcessorAcpiTsdDependency(valuetmp))
    }

    return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *ProcessorAcpiTsd) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ProcessorAcpiTsd) GetPropertyInstanceName()(value string, err error) { 
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

