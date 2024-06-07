// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// AspLimits struct
type AspLimits struct { 
	*EmbeddedObject

	// 
	BufferingLimit uint32

	// 
	MaxRequestEntityAllowed uint32

	// 
	ProcessorThreadMax uint32

	// 
	QueueConnectionTestTime string

	// 
	QueueTimeout string

	// 
	RequestQueueMax uint32

	// 
	ScriptTimeout string
}

	func NewAspLimitsEx1(instance *cim.WmiInstance) (newInstance *AspLimits, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &AspLimits {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewAspLimitsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *AspLimits, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &AspLimits {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetBufferingLimit sets the value of BufferingLimit for the instance
func (instance *AspLimits) SetPropertyBufferingLimit(value uint32) (err error) { 
    return instance.SetProperty("BufferingLimit", (value))
}

// GetBufferingLimit gets the value of BufferingLimit for the instance
func (instance *AspLimits) GetPropertyBufferingLimit()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BufferingLimit")
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

// SetMaxRequestEntityAllowed sets the value of MaxRequestEntityAllowed for the instance
func (instance *AspLimits) SetPropertyMaxRequestEntityAllowed(value uint32) (err error) { 
    return instance.SetProperty("MaxRequestEntityAllowed", (value))
}

// GetMaxRequestEntityAllowed gets the value of MaxRequestEntityAllowed for the instance
func (instance *AspLimits) GetPropertyMaxRequestEntityAllowed()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxRequestEntityAllowed")
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

// SetProcessorThreadMax sets the value of ProcessorThreadMax for the instance
func (instance *AspLimits) SetPropertyProcessorThreadMax(value uint32) (err error) { 
    return instance.SetProperty("ProcessorThreadMax", (value))
}

// GetProcessorThreadMax gets the value of ProcessorThreadMax for the instance
func (instance *AspLimits) GetPropertyProcessorThreadMax()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ProcessorThreadMax")
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

// SetQueueConnectionTestTime sets the value of QueueConnectionTestTime for the instance
func (instance *AspLimits) SetPropertyQueueConnectionTestTime(value string) (err error) { 
    return instance.SetProperty("QueueConnectionTestTime", (value))
}

// GetQueueConnectionTestTime gets the value of QueueConnectionTestTime for the instance
func (instance *AspLimits) GetPropertyQueueConnectionTestTime()(value string, err error) { 
    retValue, err := instance.GetProperty("QueueConnectionTestTime")
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

// SetQueueTimeout sets the value of QueueTimeout for the instance
func (instance *AspLimits) SetPropertyQueueTimeout(value string) (err error) { 
    return instance.SetProperty("QueueTimeout", (value))
}

// GetQueueTimeout gets the value of QueueTimeout for the instance
func (instance *AspLimits) GetPropertyQueueTimeout()(value string, err error) { 
    retValue, err := instance.GetProperty("QueueTimeout")
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

// SetRequestQueueMax sets the value of RequestQueueMax for the instance
func (instance *AspLimits) SetPropertyRequestQueueMax(value uint32) (err error) { 
    return instance.SetProperty("RequestQueueMax", (value))
}

// GetRequestQueueMax gets the value of RequestQueueMax for the instance
func (instance *AspLimits) GetPropertyRequestQueueMax()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RequestQueueMax")
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

// SetScriptTimeout sets the value of ScriptTimeout for the instance
func (instance *AspLimits) SetPropertyScriptTimeout(value string) (err error) { 
    return instance.SetProperty("ScriptTimeout", (value))
}

// GetScriptTimeout gets the value of ScriptTimeout for the instance
func (instance *AspLimits) GetPropertyScriptTimeout()(value string, err error) { 
    retValue, err := instance.GetProperty("ScriptTimeout")
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

