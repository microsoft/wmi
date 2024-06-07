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

// HandlerAction struct
type HandlerAction struct { 
	*CollectionElement

	// 
	AllowPathInfo bool

	// 
	Modules string

	// 
	Name string

	// 
	Path string

	// 
	PreCondition string

	// 
	RequireAccess int32

	// 
	ResourceType int32

	// 
	ResponseBufferLimit uint32

	// 
	ScriptProcessor string

	// 
	Type string

	// 
	Verb string
}

	func NewHandlerActionEx1(instance *cim.WmiInstance) (newInstance *HandlerAction, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &HandlerAction {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewHandlerActionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HandlerAction, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HandlerAction {
	CollectionElement: tmp,
	}
	return
	}
	

// SetAllowPathInfo sets the value of AllowPathInfo for the instance
func (instance *HandlerAction) SetPropertyAllowPathInfo(value bool) (err error) { 
    return instance.SetProperty("AllowPathInfo", (value))
}

// GetAllowPathInfo gets the value of AllowPathInfo for the instance
func (instance *HandlerAction) GetPropertyAllowPathInfo()(value bool, err error) { 
    retValue, err := instance.GetProperty("AllowPathInfo")
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

// SetModules sets the value of Modules for the instance
func (instance *HandlerAction) SetPropertyModules(value string) (err error) { 
    return instance.SetProperty("Modules", (value))
}

// GetModules gets the value of Modules for the instance
func (instance *HandlerAction) GetPropertyModules()(value string, err error) { 
    retValue, err := instance.GetProperty("Modules")
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

// SetName sets the value of Name for the instance
func (instance *HandlerAction) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *HandlerAction) GetPropertyName()(value string, err error) { 
    retValue, err := instance.GetProperty("Name")
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

// SetPath sets the value of Path for the instance
func (instance *HandlerAction) SetPropertyPath(value string) (err error) { 
    return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *HandlerAction) GetPropertyPath()(value string, err error) { 
    retValue, err := instance.GetProperty("Path")
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

// SetPreCondition sets the value of PreCondition for the instance
func (instance *HandlerAction) SetPropertyPreCondition(value string) (err error) { 
    return instance.SetProperty("PreCondition", (value))
}

// GetPreCondition gets the value of PreCondition for the instance
func (instance *HandlerAction) GetPropertyPreCondition()(value string, err error) { 
    retValue, err := instance.GetProperty("PreCondition")
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

// SetRequireAccess sets the value of RequireAccess for the instance
func (instance *HandlerAction) SetPropertyRequireAccess(value int32) (err error) { 
    return instance.SetProperty("RequireAccess", (value))
}

// GetRequireAccess gets the value of RequireAccess for the instance
func (instance *HandlerAction) GetPropertyRequireAccess()(value int32, err error) { 
    retValue, err := instance.GetProperty("RequireAccess")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetResourceType sets the value of ResourceType for the instance
func (instance *HandlerAction) SetPropertyResourceType(value int32) (err error) { 
    return instance.SetProperty("ResourceType", (value))
}

// GetResourceType gets the value of ResourceType for the instance
func (instance *HandlerAction) GetPropertyResourceType()(value int32, err error) { 
    retValue, err := instance.GetProperty("ResourceType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetResponseBufferLimit sets the value of ResponseBufferLimit for the instance
func (instance *HandlerAction) SetPropertyResponseBufferLimit(value uint32) (err error) { 
    return instance.SetProperty("ResponseBufferLimit", (value))
}

// GetResponseBufferLimit gets the value of ResponseBufferLimit for the instance
func (instance *HandlerAction) GetPropertyResponseBufferLimit()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ResponseBufferLimit")
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

// SetScriptProcessor sets the value of ScriptProcessor for the instance
func (instance *HandlerAction) SetPropertyScriptProcessor(value string) (err error) { 
    return instance.SetProperty("ScriptProcessor", (value))
}

// GetScriptProcessor gets the value of ScriptProcessor for the instance
func (instance *HandlerAction) GetPropertyScriptProcessor()(value string, err error) { 
    retValue, err := instance.GetProperty("ScriptProcessor")
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

// SetType sets the value of Type for the instance
func (instance *HandlerAction) SetPropertyType(value string) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *HandlerAction) GetPropertyType()(value string, err error) { 
    retValue, err := instance.GetProperty("Type")
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

// SetVerb sets the value of Verb for the instance
func (instance *HandlerAction) SetPropertyVerb(value string) (err error) { 
    return instance.SetProperty("Verb", (value))
}

// GetVerb gets the value of Verb for the instance
func (instance *HandlerAction) GetPropertyVerb()(value string, err error) { 
    retValue, err := instance.GetProperty("Verb")
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

