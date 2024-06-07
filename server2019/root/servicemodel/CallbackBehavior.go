// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// CallbackBehavior struct
type CallbackBehavior struct { 
	*Behavior

	// When true the session is automatically closed when a service closes a duplex session.
	AutomaticSessionShutdown bool

	// Specifies whether the service supports one thread, multiple threads, or reentrant calls.
	ConcurrencyMode string

	// A value that specifies whether to send unknown serialization data onto the wire.
	IgnoreExtensionDataObject bool

	// When enabled details about exceptions on the callback are attached to the faults returned to the service.
	IncludeExceptionDetailInFaults bool

	// The maximum number of items allowed in a serialized object.
	MaxItemsInObjectGraph bool

	// Specifies whether to use the current synchronization context to choose the thread of execution.
	UseSynchronizationContext bool

	// Specifies whether the system or the application enforces SOAP MustUnderstand header processing.
	ValidateMustUnderstand bool
}

	func NewCallbackBehaviorEx1(instance *cim.WmiInstance) (newInstance *CallbackBehavior, err error) {tmp, err := NewBehaviorEx1(instance)
		
	if err != nil { return }
	newInstance = &CallbackBehavior {
	Behavior: tmp,
	}
	return
	}
	

	func NewCallbackBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CallbackBehavior, err error) {tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CallbackBehavior {
	Behavior: tmp,
	}
	return
	}
	

// SetAutomaticSessionShutdown sets the value of AutomaticSessionShutdown for the instance
func (instance *CallbackBehavior) SetPropertyAutomaticSessionShutdown(value bool) (err error) { 
    return instance.SetProperty("AutomaticSessionShutdown", (value))
}

// GetAutomaticSessionShutdown gets the value of AutomaticSessionShutdown for the instance
func (instance *CallbackBehavior) GetPropertyAutomaticSessionShutdown()(value bool, err error) { 
    retValue, err := instance.GetProperty("AutomaticSessionShutdown")
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

// SetConcurrencyMode sets the value of ConcurrencyMode for the instance
func (instance *CallbackBehavior) SetPropertyConcurrencyMode(value string) (err error) { 
    return instance.SetProperty("ConcurrencyMode", (value))
}

// GetConcurrencyMode gets the value of ConcurrencyMode for the instance
func (instance *CallbackBehavior) GetPropertyConcurrencyMode()(value string, err error) { 
    retValue, err := instance.GetProperty("ConcurrencyMode")
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

// SetIgnoreExtensionDataObject sets the value of IgnoreExtensionDataObject for the instance
func (instance *CallbackBehavior) SetPropertyIgnoreExtensionDataObject(value bool) (err error) { 
    return instance.SetProperty("IgnoreExtensionDataObject", (value))
}

// GetIgnoreExtensionDataObject gets the value of IgnoreExtensionDataObject for the instance
func (instance *CallbackBehavior) GetPropertyIgnoreExtensionDataObject()(value bool, err error) { 
    retValue, err := instance.GetProperty("IgnoreExtensionDataObject")
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

// SetIncludeExceptionDetailInFaults sets the value of IncludeExceptionDetailInFaults for the instance
func (instance *CallbackBehavior) SetPropertyIncludeExceptionDetailInFaults(value bool) (err error) { 
    return instance.SetProperty("IncludeExceptionDetailInFaults", (value))
}

// GetIncludeExceptionDetailInFaults gets the value of IncludeExceptionDetailInFaults for the instance
func (instance *CallbackBehavior) GetPropertyIncludeExceptionDetailInFaults()(value bool, err error) { 
    retValue, err := instance.GetProperty("IncludeExceptionDetailInFaults")
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

// SetMaxItemsInObjectGraph sets the value of MaxItemsInObjectGraph for the instance
func (instance *CallbackBehavior) SetPropertyMaxItemsInObjectGraph(value bool) (err error) { 
    return instance.SetProperty("MaxItemsInObjectGraph", (value))
}

// GetMaxItemsInObjectGraph gets the value of MaxItemsInObjectGraph for the instance
func (instance *CallbackBehavior) GetPropertyMaxItemsInObjectGraph()(value bool, err error) { 
    retValue, err := instance.GetProperty("MaxItemsInObjectGraph")
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

// SetUseSynchronizationContext sets the value of UseSynchronizationContext for the instance
func (instance *CallbackBehavior) SetPropertyUseSynchronizationContext(value bool) (err error) { 
    return instance.SetProperty("UseSynchronizationContext", (value))
}

// GetUseSynchronizationContext gets the value of UseSynchronizationContext for the instance
func (instance *CallbackBehavior) GetPropertyUseSynchronizationContext()(value bool, err error) { 
    retValue, err := instance.GetProperty("UseSynchronizationContext")
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

// SetValidateMustUnderstand sets the value of ValidateMustUnderstand for the instance
func (instance *CallbackBehavior) SetPropertyValidateMustUnderstand(value bool) (err error) { 
    return instance.SetProperty("ValidateMustUnderstand", (value))
}

// GetValidateMustUnderstand gets the value of ValidateMustUnderstand for the instance
func (instance *CallbackBehavior) GetPropertyValidateMustUnderstand()(value bool, err error) { 
    retValue, err := instance.GetProperty("ValidateMustUnderstand")
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

