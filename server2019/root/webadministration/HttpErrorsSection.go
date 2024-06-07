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

// HttpErrorsSection struct
type HttpErrorsSection struct { 
	*ConfigurationSectionWithCollection

	// 
	AllowAbsolutePathsWhenDelegated bool

	// 
	DefaultPath string

	// 
	DefaultResponseMode int32

	// 
	DetailedMoreInformationLink string

	// 
	ErrorMode int32

	// 
	ExistingResponse int32

	// 
	HttpErrors []HttpErrorElement
}

	func NewHttpErrorsSectionEx1(instance *cim.WmiInstance) (newInstance *HttpErrorsSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &HttpErrorsSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewHttpErrorsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HttpErrorsSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HttpErrorsSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetAllowAbsolutePathsWhenDelegated sets the value of AllowAbsolutePathsWhenDelegated for the instance
func (instance *HttpErrorsSection) SetPropertyAllowAbsolutePathsWhenDelegated(value bool) (err error) { 
    return instance.SetProperty("AllowAbsolutePathsWhenDelegated", (value))
}

// GetAllowAbsolutePathsWhenDelegated gets the value of AllowAbsolutePathsWhenDelegated for the instance
func (instance *HttpErrorsSection) GetPropertyAllowAbsolutePathsWhenDelegated()(value bool, err error) { 
    retValue, err := instance.GetProperty("AllowAbsolutePathsWhenDelegated")
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

// SetDefaultPath sets the value of DefaultPath for the instance
func (instance *HttpErrorsSection) SetPropertyDefaultPath(value string) (err error) { 
    return instance.SetProperty("DefaultPath", (value))
}

// GetDefaultPath gets the value of DefaultPath for the instance
func (instance *HttpErrorsSection) GetPropertyDefaultPath()(value string, err error) { 
    retValue, err := instance.GetProperty("DefaultPath")
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

// SetDefaultResponseMode sets the value of DefaultResponseMode for the instance
func (instance *HttpErrorsSection) SetPropertyDefaultResponseMode(value int32) (err error) { 
    return instance.SetProperty("DefaultResponseMode", (value))
}

// GetDefaultResponseMode gets the value of DefaultResponseMode for the instance
func (instance *HttpErrorsSection) GetPropertyDefaultResponseMode()(value int32, err error) { 
    retValue, err := instance.GetProperty("DefaultResponseMode")
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

// SetDetailedMoreInformationLink sets the value of DetailedMoreInformationLink for the instance
func (instance *HttpErrorsSection) SetPropertyDetailedMoreInformationLink(value string) (err error) { 
    return instance.SetProperty("DetailedMoreInformationLink", (value))
}

// GetDetailedMoreInformationLink gets the value of DetailedMoreInformationLink for the instance
func (instance *HttpErrorsSection) GetPropertyDetailedMoreInformationLink()(value string, err error) { 
    retValue, err := instance.GetProperty("DetailedMoreInformationLink")
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

// SetErrorMode sets the value of ErrorMode for the instance
func (instance *HttpErrorsSection) SetPropertyErrorMode(value int32) (err error) { 
    return instance.SetProperty("ErrorMode", (value))
}

// GetErrorMode gets the value of ErrorMode for the instance
func (instance *HttpErrorsSection) GetPropertyErrorMode()(value int32, err error) { 
    retValue, err := instance.GetProperty("ErrorMode")
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

// SetExistingResponse sets the value of ExistingResponse for the instance
func (instance *HttpErrorsSection) SetPropertyExistingResponse(value int32) (err error) { 
    return instance.SetProperty("ExistingResponse", (value))
}

// GetExistingResponse gets the value of ExistingResponse for the instance
func (instance *HttpErrorsSection) GetPropertyExistingResponse()(value int32, err error) { 
    retValue, err := instance.GetProperty("ExistingResponse")
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

// SetHttpErrors sets the value of HttpErrors for the instance
func (instance *HttpErrorsSection) SetPropertyHttpErrors(value []HttpErrorElement) (err error) { 
    return instance.SetProperty("HttpErrors", (value))
}

// GetHttpErrors gets the value of HttpErrors for the instance
func (instance *HttpErrorsSection) GetPropertyHttpErrors()(value []HttpErrorElement, err error) { 
    retValue, err := instance.GetProperty("HttpErrors")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(HttpErrorElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " HttpErrorElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, HttpErrorElement(valuetmp))
    }

    return
}

