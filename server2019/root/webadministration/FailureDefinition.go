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

// FailureDefinition struct
type FailureDefinition struct { 
	*EmbeddedObject

	// 
	StatusCodes string

	// 
	TimeTaken string

	// 
	Verbosity int32
}

	func NewFailureDefinitionEx1(instance *cim.WmiInstance) (newInstance *FailureDefinition, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &FailureDefinition {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewFailureDefinitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FailureDefinition, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FailureDefinition {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetStatusCodes sets the value of StatusCodes for the instance
func (instance *FailureDefinition) SetPropertyStatusCodes(value string) (err error) { 
    return instance.SetProperty("StatusCodes", (value))
}

// GetStatusCodes gets the value of StatusCodes for the instance
func (instance *FailureDefinition) GetPropertyStatusCodes()(value string, err error) { 
    retValue, err := instance.GetProperty("StatusCodes")
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

// SetTimeTaken sets the value of TimeTaken for the instance
func (instance *FailureDefinition) SetPropertyTimeTaken(value string) (err error) { 
    return instance.SetProperty("TimeTaken", (value))
}

// GetTimeTaken gets the value of TimeTaken for the instance
func (instance *FailureDefinition) GetPropertyTimeTaken()(value string, err error) { 
    retValue, err := instance.GetProperty("TimeTaken")
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

// SetVerbosity sets the value of Verbosity for the instance
func (instance *FailureDefinition) SetPropertyVerbosity(value int32) (err error) { 
    return instance.SetProperty("Verbosity", (value))
}

// GetVerbosity gets the value of Verbosity for the instance
func (instance *FailureDefinition) GetPropertyVerbosity()(value int32, err error) { 
    retValue, err := instance.GetProperty("Verbosity")
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

