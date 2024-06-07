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

// Header_DbgIdRSDS_TypeGroup struct
type Header_DbgIdRSDS_TypeGroup struct { 
	*EventTraceEvent

	// 
	Age uint32

	// 
	Guid interface{}

	// 
	PdbName string
}

	func NewHeader_DbgIdRSDS_TypeGroupEx1(instance *cim.WmiInstance) (newInstance *Header_DbgIdRSDS_TypeGroup, err error) {tmp, err := NewEventTraceEventEx1(instance)
		
	if err != nil { return }
	newInstance = &Header_DbgIdRSDS_TypeGroup {
	EventTraceEvent: tmp,
	}
	return
	}
	

	func NewHeader_DbgIdRSDS_TypeGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Header_DbgIdRSDS_TypeGroup, err error) {tmp, err := NewEventTraceEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Header_DbgIdRSDS_TypeGroup {
	EventTraceEvent: tmp,
	}
	return
	}
	

// SetAge sets the value of Age for the instance
func (instance *Header_DbgIdRSDS_TypeGroup) SetPropertyAge(value uint32) (err error) { 
    return instance.SetProperty("Age", (value))
}

// GetAge gets the value of Age for the instance
func (instance *Header_DbgIdRSDS_TypeGroup) GetPropertyAge()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Age")
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

// SetGuid sets the value of Guid for the instance
func (instance *Header_DbgIdRSDS_TypeGroup) SetPropertyGuid(value interface{}) (err error) { 
    return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *Header_DbgIdRSDS_TypeGroup) GetPropertyGuid()(value interface{}, err error) { 
    retValue, err := instance.GetProperty("Guid")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(interface{}); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = interface{}(valuetmp)

    return
}

// SetPdbName sets the value of PdbName for the instance
func (instance *Header_DbgIdRSDS_TypeGroup) SetPropertyPdbName(value string) (err error) { 
    return instance.SetProperty("PdbName", (value))
}

// GetPdbName gets the value of PdbName for the instance
func (instance *Header_DbgIdRSDS_TypeGroup) GetPropertyPdbName()(value string, err error) { 
    retValue, err := instance.GetProperty("PdbName")
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

