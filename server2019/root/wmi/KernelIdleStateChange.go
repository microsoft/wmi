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

// KernelIdleStateChange struct
type KernelIdleStateChange struct { 
	*WMIEvent

	// 
	Active bool

	// 
	InstanceName string

	// 
	NewState uint32

	// 
	OldState uint32

	// 
	Processors uint64
}

	func NewKernelIdleStateChangeEx1(instance *cim.WmiInstance) (newInstance *KernelIdleStateChange, err error) {tmp, err := NewWMIEventEx1(instance)
		
	if err != nil { return }
	newInstance = &KernelIdleStateChange {
	WMIEvent: tmp,
	}
	return
	}
	

	func NewKernelIdleStateChangeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *KernelIdleStateChange, err error) {tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &KernelIdleStateChange {
	WMIEvent: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *KernelIdleStateChange) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *KernelIdleStateChange) GetPropertyActive()(value bool, err error) { 
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *KernelIdleStateChange) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *KernelIdleStateChange) GetPropertyInstanceName()(value string, err error) { 
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

// SetNewState sets the value of NewState for the instance
func (instance *KernelIdleStateChange) SetPropertyNewState(value uint32) (err error) { 
    return instance.SetProperty("NewState", (value))
}

// GetNewState gets the value of NewState for the instance
func (instance *KernelIdleStateChange) GetPropertyNewState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NewState")
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

// SetOldState sets the value of OldState for the instance
func (instance *KernelIdleStateChange) SetPropertyOldState(value uint32) (err error) { 
    return instance.SetProperty("OldState", (value))
}

// GetOldState gets the value of OldState for the instance
func (instance *KernelIdleStateChange) GetPropertyOldState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OldState")
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

// SetProcessors sets the value of Processors for the instance
func (instance *KernelIdleStateChange) SetPropertyProcessors(value uint64) (err error) { 
    return instance.SetProperty("Processors", (value))
}

// GetProcessors gets the value of Processors for the instance
func (instance *KernelIdleStateChange) GetPropertyProcessors()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Processors")
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

