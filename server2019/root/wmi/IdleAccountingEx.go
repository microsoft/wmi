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

// IdleAccountingEx struct
type IdleAccountingEx struct { 
	*WMIEvent

	// 
	Active bool

	// 
	InstanceName string

	// 
	ResetCount uint32

	// 
	StartTime uint64

	// 
	State []IdleStateAccountingEx

	// 
	StateCount uint32

	// 
	TotalTransitions uint32
}

	func NewIdleAccountingExEx1(instance *cim.WmiInstance) (newInstance *IdleAccountingEx, err error) {tmp, err := NewWMIEventEx1(instance)
		
	if err != nil { return }
	newInstance = &IdleAccountingEx {
	WMIEvent: tmp,
	}
	return
	}
	

	func NewIdleAccountingExEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *IdleAccountingEx, err error) {tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &IdleAccountingEx {
	WMIEvent: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *IdleAccountingEx) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *IdleAccountingEx) GetPropertyActive()(value bool, err error) { 
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
func (instance *IdleAccountingEx) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *IdleAccountingEx) GetPropertyInstanceName()(value string, err error) { 
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

// SetResetCount sets the value of ResetCount for the instance
func (instance *IdleAccountingEx) SetPropertyResetCount(value uint32) (err error) { 
    return instance.SetProperty("ResetCount", (value))
}

// GetResetCount gets the value of ResetCount for the instance
func (instance *IdleAccountingEx) GetPropertyResetCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ResetCount")
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

// SetStartTime sets the value of StartTime for the instance
func (instance *IdleAccountingEx) SetPropertyStartTime(value uint64) (err error) { 
    return instance.SetProperty("StartTime", (value))
}

// GetStartTime gets the value of StartTime for the instance
func (instance *IdleAccountingEx) GetPropertyStartTime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("StartTime")
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

// SetState sets the value of State for the instance
func (instance *IdleAccountingEx) SetPropertyState(value []IdleStateAccountingEx) (err error) { 
    return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *IdleAccountingEx) GetPropertyState()(value []IdleStateAccountingEx, err error) { 
    retValue, err := instance.GetProperty("State")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(IdleStateAccountingEx); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " IdleStateAccountingEx is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, IdleStateAccountingEx(valuetmp))
    }

    return
}

// SetStateCount sets the value of StateCount for the instance
func (instance *IdleAccountingEx) SetPropertyStateCount(value uint32) (err error) { 
    return instance.SetProperty("StateCount", (value))
}

// GetStateCount gets the value of StateCount for the instance
func (instance *IdleAccountingEx) GetPropertyStateCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("StateCount")
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

// SetTotalTransitions sets the value of TotalTransitions for the instance
func (instance *IdleAccountingEx) SetPropertyTotalTransitions(value uint32) (err error) { 
    return instance.SetProperty("TotalTransitions", (value))
}

// GetTotalTransitions gets the value of TotalTransitions for the instance
func (instance *IdleAccountingEx) GetPropertyTotalTransitions()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalTransitions")
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

