// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RemoteAccessHealthMonitor struct
type RemoteAccessHealthMonitor struct { 
	*cim.WmiInstance

	// 
	Component string

	// 
	HealthState string

	// 
	Heuristics []RemoteAccessHealthHeuristic

	// 
	RemoteAccessServer string

	// 
	TimeStamp string
}

	func NewRemoteAccessHealthMonitorEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessHealthMonitor, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RemoteAccessHealthMonitor {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRemoteAccessHealthMonitorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RemoteAccessHealthMonitor, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RemoteAccessHealthMonitor {
	WmiInstance: tmp,
	}
	return
	}
	

// SetComponent sets the value of Component for the instance
func (instance *RemoteAccessHealthMonitor) SetPropertyComponent(value string) (err error) { 
    return instance.SetProperty("Component", (value))
}

// GetComponent gets the value of Component for the instance
func (instance *RemoteAccessHealthMonitor) GetPropertyComponent()(value string, err error) { 
    retValue, err := instance.GetProperty("Component")
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

// SetHealthState sets the value of HealthState for the instance
func (instance *RemoteAccessHealthMonitor) SetPropertyHealthState(value string) (err error) { 
    return instance.SetProperty("HealthState", (value))
}

// GetHealthState gets the value of HealthState for the instance
func (instance *RemoteAccessHealthMonitor) GetPropertyHealthState()(value string, err error) { 
    retValue, err := instance.GetProperty("HealthState")
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

// SetHeuristics sets the value of Heuristics for the instance
func (instance *RemoteAccessHealthMonitor) SetPropertyHeuristics(value []RemoteAccessHealthHeuristic) (err error) { 
    return instance.SetProperty("Heuristics", (value))
}

// GetHeuristics gets the value of Heuristics for the instance
func (instance *RemoteAccessHealthMonitor) GetPropertyHeuristics()(value []RemoteAccessHealthHeuristic, err error) { 
    retValue, err := instance.GetProperty("Heuristics")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(RemoteAccessHealthHeuristic); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " RemoteAccessHealthHeuristic is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, RemoteAccessHealthHeuristic(valuetmp))
    }

    return
}

// SetRemoteAccessServer sets the value of RemoteAccessServer for the instance
func (instance *RemoteAccessHealthMonitor) SetPropertyRemoteAccessServer(value string) (err error) { 
    return instance.SetProperty("RemoteAccessServer", (value))
}

// GetRemoteAccessServer gets the value of RemoteAccessServer for the instance
func (instance *RemoteAccessHealthMonitor) GetPropertyRemoteAccessServer()(value string, err error) { 
    retValue, err := instance.GetProperty("RemoteAccessServer")
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

// SetTimeStamp sets the value of TimeStamp for the instance
func (instance *RemoteAccessHealthMonitor) SetPropertyTimeStamp(value string) (err error) { 
    return instance.SetProperty("TimeStamp", (value))
}

// GetTimeStamp gets the value of TimeStamp for the instance
func (instance *RemoteAccessHealthMonitor) GetPropertyTimeStamp()(value string, err error) { 
    retValue, err := instance.GetProperty("TimeStamp")
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

