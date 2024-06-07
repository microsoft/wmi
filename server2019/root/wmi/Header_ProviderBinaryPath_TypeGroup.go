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

// Header_ProviderBinaryPath_TypeGroup struct
type Header_ProviderBinaryPath_TypeGroup struct { 
	*EventTraceEvent

	// 
	BinaryPath string

	// 
	Guid []interface{}

	// 
	GuidCount uint32
}

	func NewHeader_ProviderBinaryPath_TypeGroupEx1(instance *cim.WmiInstance) (newInstance *Header_ProviderBinaryPath_TypeGroup, err error) {tmp, err := NewEventTraceEventEx1(instance)
		
	if err != nil { return }
	newInstance = &Header_ProviderBinaryPath_TypeGroup {
	EventTraceEvent: tmp,
	}
	return
	}
	

	func NewHeader_ProviderBinaryPath_TypeGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Header_ProviderBinaryPath_TypeGroup, err error) {tmp, err := NewEventTraceEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Header_ProviderBinaryPath_TypeGroup {
	EventTraceEvent: tmp,
	}
	return
	}
	

// SetBinaryPath sets the value of BinaryPath for the instance
func (instance *Header_ProviderBinaryPath_TypeGroup) SetPropertyBinaryPath(value string) (err error) { 
    return instance.SetProperty("BinaryPath", (value))
}

// GetBinaryPath gets the value of BinaryPath for the instance
func (instance *Header_ProviderBinaryPath_TypeGroup) GetPropertyBinaryPath()(value string, err error) { 
    retValue, err := instance.GetProperty("BinaryPath")
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

// SetGuid sets the value of Guid for the instance
func (instance *Header_ProviderBinaryPath_TypeGroup) SetPropertyGuid(value []interface{}) (err error) { 
    return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *Header_ProviderBinaryPath_TypeGroup) GetPropertyGuid()(value []interface{}, err error) { 
    retValue, err := instance.GetProperty("Guid")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(interface{}); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, interface{}(valuetmp))
    }

    return
}

// SetGuidCount sets the value of GuidCount for the instance
func (instance *Header_ProviderBinaryPath_TypeGroup) SetPropertyGuidCount(value uint32) (err error) { 
    return instance.SetProperty("GuidCount", (value))
}

// GetGuidCount gets the value of GuidCount for the instance
func (instance *Header_ProviderBinaryPath_TypeGroup) GetPropertyGuidCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("GuidCount")
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

