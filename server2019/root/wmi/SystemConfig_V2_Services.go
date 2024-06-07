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

// SystemConfig_V2_Services struct
type SystemConfig_V2_Services struct { 
	*SystemConfig_V2

	// 
	DisplayName string

	// 
	ProcessId uint32

	// 
	ProcessName string

	// 
	ServiceName string

	// 
	ServiceState uint32

	// 
	SubProcessTag uint32
}

	func NewSystemConfig_V2_ServicesEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_Services, err error) {tmp, err := NewSystemConfig_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &SystemConfig_V2_Services {
	SystemConfig_V2: tmp,
	}
	return
	}
	

	func NewSystemConfig_V2_ServicesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SystemConfig_V2_Services, err error) {tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SystemConfig_V2_Services {
	SystemConfig_V2: tmp,
	}
	return
	}
	

// SetDisplayName sets the value of DisplayName for the instance
func (instance *SystemConfig_V2_Services) SetPropertyDisplayName(value string) (err error) { 
    return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *SystemConfig_V2_Services) GetPropertyDisplayName()(value string, err error) { 
    retValue, err := instance.GetProperty("DisplayName")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *SystemConfig_V2_Services) SetPropertyProcessId(value uint32) (err error) { 
    return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *SystemConfig_V2_Services) GetPropertyProcessId()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ProcessId")
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

// SetProcessName sets the value of ProcessName for the instance
func (instance *SystemConfig_V2_Services) SetPropertyProcessName(value string) (err error) { 
    return instance.SetProperty("ProcessName", (value))
}

// GetProcessName gets the value of ProcessName for the instance
func (instance *SystemConfig_V2_Services) GetPropertyProcessName()(value string, err error) { 
    retValue, err := instance.GetProperty("ProcessName")
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

// SetServiceName sets the value of ServiceName for the instance
func (instance *SystemConfig_V2_Services) SetPropertyServiceName(value string) (err error) { 
    return instance.SetProperty("ServiceName", (value))
}

// GetServiceName gets the value of ServiceName for the instance
func (instance *SystemConfig_V2_Services) GetPropertyServiceName()(value string, err error) { 
    retValue, err := instance.GetProperty("ServiceName")
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

// SetServiceState sets the value of ServiceState for the instance
func (instance *SystemConfig_V2_Services) SetPropertyServiceState(value uint32) (err error) { 
    return instance.SetProperty("ServiceState", (value))
}

// GetServiceState gets the value of ServiceState for the instance
func (instance *SystemConfig_V2_Services) GetPropertyServiceState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ServiceState")
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

// SetSubProcessTag sets the value of SubProcessTag for the instance
func (instance *SystemConfig_V2_Services) SetPropertySubProcessTag(value uint32) (err error) { 
    return instance.SetProperty("SubProcessTag", (value))
}

// GetSubProcessTag gets the value of SubProcessTag for the instance
func (instance *SystemConfig_V2_Services) GetPropertySubProcessTag()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SubProcessTag")
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

